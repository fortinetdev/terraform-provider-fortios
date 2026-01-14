// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DNS.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDnsUpdate,
		Read:   resourceSystemDnsRead,
		Update: resourceSystemDnsUpdate,
		Delete: resourceSystemDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_over_tls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"server_hostname": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
						},
					},
				},
			},
			"domain": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
						},
					},
				},
			},
			"ip6_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"retry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 5),
				Optional:     true,
				Computed:     true,
			},
			"dns_cache_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_cache_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
			},
			"cache_notfound_responses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"root_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"vrf_select": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 511),
				Optional:     true,
			},
			"server_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alt_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alt_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn_cache_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),
				Optional:     true,
			},
			"fqdn_max_refresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3600, 86400),
				Optional:     true,
				Computed:     true,
			},
			"fqdn_min_refresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),
				Optional:     true,
				Computed:     true,
			},
			"hostname_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
			},
			"hostname_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 50000),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemDnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSystemDns(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDns(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDns")
	}

	return resourceSystemDnsRead(d, m)
}

func resourceSystemDnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemDns(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSystemDns(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDns(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemDnsPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsSecondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDnsOverTls(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsSslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsServerHostname(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "hostname", "hostname")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["hostname"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
			}
			tmp["hostname"] = flattenSystemDnsServerHostnameHostname(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "hostname", d)
	return result
}

func flattenSystemDnsServerHostnameHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDomain(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "domain", "domain")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["domain"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
			}
			tmp["domain"] = flattenSystemDnsDomainDomain(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "domain", d)
	return result
}

func flattenSystemDnsDomainDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsRetry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsDnsCacheLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsDnsCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsCacheNotfoundResponses(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsSourceIpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsRootServers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsVrfSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsServerSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsAltPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsAltSecondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsFqdnCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsFqdnMaxRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsFqdnMinRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsHostnameTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsHostnameLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("primary", flattenSystemDnsPrimary(o["primary"], d, "primary", sv)); err != nil {
		if !fortiAPIPatch(o["primary"]) {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystemDnsSecondary(o["secondary"], d, "secondary", sv)); err != nil {
		if !fortiAPIPatch(o["secondary"]) {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemDnsProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("dns_over_tls", flattenSystemDnsDnsOverTls(o["dns-over-tls"], d, "dns_over_tls", sv)); err != nil {
		if !fortiAPIPatch(o["dns-over-tls"]) {
			return fmt.Errorf("Error reading dns_over_tls: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystemDnsSslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("server_hostname", flattenSystemDnsServerHostname(o["server-hostname"], d, "server_hostname", sv)); err != nil {
			if !fortiAPIPatch(o["server-hostname"]) {
				return fmt.Errorf("Error reading server_hostname: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_hostname"); ok {
			if err = d.Set("server_hostname", flattenSystemDnsServerHostname(o["server-hostname"], d, "server_hostname", sv)); err != nil {
				if !fortiAPIPatch(o["server-hostname"]) {
					return fmt.Errorf("Error reading server_hostname: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("domain", flattenSystemDnsDomain(o["domain"], d, "domain", sv)); err != nil {
			if !fortiAPIPatch(o["domain"]) {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("domain"); ok {
			if err = d.Set("domain", flattenSystemDnsDomain(o["domain"], d, "domain", sv)); err != nil {
				if !fortiAPIPatch(o["domain"]) {
					return fmt.Errorf("Error reading domain: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_primary", flattenSystemDnsIp6Primary(o["ip6-primary"], d, "ip6_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-primary"]) {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", flattenSystemDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-secondary"]) {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemDnsTimeout(o["timeout"], d, "timeout", sv)); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("retry", flattenSystemDnsRetry(o["retry"], d, "retry", sv)); err != nil {
		if !fortiAPIPatch(o["retry"]) {
			return fmt.Errorf("Error reading retry: %v", err)
		}
	}

	if err = d.Set("dns_cache_limit", flattenSystemDnsDnsCacheLimit(o["dns-cache-limit"], d, "dns_cache_limit", sv)); err != nil {
		if !fortiAPIPatch(o["dns-cache-limit"]) {
			return fmt.Errorf("Error reading dns_cache_limit: %v", err)
		}
	}

	if err = d.Set("dns_cache_ttl", flattenSystemDnsDnsCacheTtl(o["dns-cache-ttl"], d, "dns_cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["dns-cache-ttl"]) {
			return fmt.Errorf("Error reading dns_cache_ttl: %v", err)
		}
	}

	if err = d.Set("cache_notfound_responses", flattenSystemDnsCacheNotfoundResponses(o["cache-notfound-responses"], d, "cache_notfound_responses", sv)); err != nil {
		if !fortiAPIPatch(o["cache-notfound-responses"]) {
			return fmt.Errorf("Error reading cache_notfound_responses: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemDnsSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenSystemDnsSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip-interface"]) {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("root_servers", flattenSystemDnsRootServers(o["root-servers"], d, "root_servers", sv)); err != nil {
		if !fortiAPIPatch(o["root-servers"]) {
			return fmt.Errorf("Error reading root_servers: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemDnsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDnsInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenSystemDnsVrfSelect(o["vrf-select"], d, "vrf_select", sv)); err != nil {
		if !fortiAPIPatch(o["vrf-select"]) {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	if err = d.Set("server_select_method", flattenSystemDnsServerSelectMethod(o["server-select-method"], d, "server_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["server-select-method"]) {
			return fmt.Errorf("Error reading server_select_method: %v", err)
		}
	}

	if err = d.Set("alt_primary", flattenSystemDnsAltPrimary(o["alt-primary"], d, "alt_primary", sv)); err != nil {
		if !fortiAPIPatch(o["alt-primary"]) {
			return fmt.Errorf("Error reading alt_primary: %v", err)
		}
	}

	if err = d.Set("alt_secondary", flattenSystemDnsAltSecondary(o["alt-secondary"], d, "alt_secondary", sv)); err != nil {
		if !fortiAPIPatch(o["alt-secondary"]) {
			return fmt.Errorf("Error reading alt_secondary: %v", err)
		}
	}

	if err = d.Set("log", flattenSystemDnsLog(o["log"], d, "log", sv)); err != nil {
		if !fortiAPIPatch(o["log"]) {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("fqdn_cache_ttl", flattenSystemDnsFqdnCacheTtl(o["fqdn-cache-ttl"], d, "fqdn_cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["fqdn-cache-ttl"]) {
			return fmt.Errorf("Error reading fqdn_cache_ttl: %v", err)
		}
	}

	if err = d.Set("fqdn_max_refresh", flattenSystemDnsFqdnMaxRefresh(o["fqdn-max-refresh"], d, "fqdn_max_refresh", sv)); err != nil {
		if !fortiAPIPatch(o["fqdn-max-refresh"]) {
			return fmt.Errorf("Error reading fqdn_max_refresh: %v", err)
		}
	}

	if err = d.Set("fqdn_min_refresh", flattenSystemDnsFqdnMinRefresh(o["fqdn-min-refresh"], d, "fqdn_min_refresh", sv)); err != nil {
		if !fortiAPIPatch(o["fqdn-min-refresh"]) {
			return fmt.Errorf("Error reading fqdn_min_refresh: %v", err)
		}
	}

	if err = d.Set("hostname_ttl", flattenSystemDnsHostnameTtl(o["hostname-ttl"], d, "hostname_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["hostname-ttl"]) {
			return fmt.Errorf("Error reading hostname_ttl: %v", err)
		}
	}

	if err = d.Set("hostname_limit", flattenSystemDnsHostnameLimit(o["hostname-limit"], d, "hostname_limit", sv)); err != nil {
		if !fortiAPIPatch(o["hostname-limit"]) {
			return fmt.Errorf("Error reading hostname_limit: %v", err)
		}
	}

	return nil
}

func flattenSystemDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDnsPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSecondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsOverTls(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsServerHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["hostname"], _ = expandSystemDnsServerHostnameHostname(d, i["hostname"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDnsServerHostnameHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["domain"], _ = expandSystemDnsDomainDomain(d, i["domain"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDnsDomainDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIp6Primary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIp6Secondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsRetry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsCacheLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsCacheNotfoundResponses(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSourceIpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsRootServers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsVrfSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsServerSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsAltPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsAltSecondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsFqdnCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsFqdnMaxRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsFqdnMinRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsHostnameTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsHostnameLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("primary"); ok {
		if setArgNil {
			obj["primary"] = nil
		} else {
			t, err := expandSystemDnsPrimary(d, v, "primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("secondary"); ok {
		if setArgNil {
			obj["secondary"] = nil
		} else {
			t, err := expandSystemDnsSecondary(d, v, "secondary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["secondary"] = t
			}
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		if setArgNil {
			obj["protocol"] = nil
		} else {
			t, err := expandSystemDnsProtocol(d, v, "protocol", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["protocol"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns_over_tls"); ok {
		if setArgNil {
			obj["dns-over-tls"] = nil
		} else {
			t, err := expandSystemDnsDnsOverTls(d, v, "dns_over_tls", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-over-tls"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok {
		if setArgNil {
			obj["ssl-certificate"] = nil
		} else {
			t, err := expandSystemDnsSslCertificate(d, v, "ssl_certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-certificate"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_hostname"); ok || d.HasChange("server_hostname") {
		if setArgNil {
			obj["server-hostname"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemDnsServerHostname(d, v, "server_hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-hostname"] = t
			}
		}
	}

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		if setArgNil {
			obj["domain"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemDnsDomain(d, v, "domain", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["domain"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip6_primary"); ok {
		if setArgNil {
			obj["ip6-primary"] = nil
		} else {
			t, err := expandSystemDnsIp6Primary(d, v, "ip6_primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip6-primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip6_secondary"); ok {
		if setArgNil {
			obj["ip6-secondary"] = nil
		} else {
			t, err := expandSystemDnsIp6Secondary(d, v, "ip6_secondary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip6-secondary"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout"); ok {
		if setArgNil {
			obj["timeout"] = nil
		} else {
			t, err := expandSystemDnsTimeout(d, v, "timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("retry"); ok {
		if setArgNil {
			obj["retry"] = nil
		} else {
			t, err := expandSystemDnsRetry(d, v, "retry", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["retry"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("dns_cache_limit"); ok {
		if setArgNil {
			obj["dns-cache-limit"] = nil
		} else {
			t, err := expandSystemDnsDnsCacheLimit(d, v, "dns_cache_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-cache-limit"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns_cache_ttl"); ok {
		if setArgNil {
			obj["dns-cache-ttl"] = nil
		} else {
			t, err := expandSystemDnsDnsCacheTtl(d, v, "dns_cache_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-cache-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("cache_notfound_responses"); ok {
		if setArgNil {
			obj["cache-notfound-responses"] = nil
		} else {
			t, err := expandSystemDnsCacheNotfoundResponses(d, v, "cache_notfound_responses", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cache-notfound-responses"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {
			t, err := expandSystemDnsSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok {
		if setArgNil {
			obj["source-ip-interface"] = nil
		} else {
			t, err := expandSystemDnsSourceIpInterface(d, v, "source_ip_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip-interface"] = t
			}
		}
	} else if d.HasChange("source_ip_interface") {
		obj["source-ip-interface"] = nil
	}

	if v, ok := d.GetOk("root_servers"); ok {
		if setArgNil {
			obj["root-servers"] = nil
		} else {
			t, err := expandSystemDnsRootServers(d, v, "root_servers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["root-servers"] = t
			}
		}
	} else if d.HasChange("root_servers") {
		obj["root-servers"] = nil
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		if setArgNil {
			obj["interface-select-method"] = nil
		} else {
			t, err := expandSystemDnsInterfaceSelectMethod(d, v, "interface_select_method", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface-select-method"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {
			t, err := expandSystemDnsInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOkExists("vrf_select"); ok {
		if setArgNil {
			obj["vrf-select"] = nil
		} else {
			t, err := expandSystemDnsVrfSelect(d, v, "vrf_select", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vrf-select"] = t
			}
		}
	} else if d.HasChange("vrf_select") {
		obj["vrf-select"] = nil
	}

	if v, ok := d.GetOk("server_select_method"); ok {
		if setArgNil {
			obj["server-select-method"] = nil
		} else {
			t, err := expandSystemDnsServerSelectMethod(d, v, "server_select_method", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-select-method"] = t
			}
		}
	}

	if v, ok := d.GetOk("alt_primary"); ok {
		if setArgNil {
			obj["alt-primary"] = nil
		} else {
			t, err := expandSystemDnsAltPrimary(d, v, "alt_primary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["alt-primary"] = t
			}
		}
	}

	if v, ok := d.GetOk("alt_secondary"); ok {
		if setArgNil {
			obj["alt-secondary"] = nil
		} else {
			t, err := expandSystemDnsAltSecondary(d, v, "alt_secondary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["alt-secondary"] = t
			}
		}
	}

	if v, ok := d.GetOk("log"); ok {
		if setArgNil {
			obj["log"] = nil
		} else {
			t, err := expandSystemDnsLog(d, v, "log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("fqdn_cache_ttl"); ok {
		if setArgNil {
			obj["fqdn-cache-ttl"] = nil
		} else {
			t, err := expandSystemDnsFqdnCacheTtl(d, v, "fqdn_cache_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fqdn-cache-ttl"] = t
			}
		}
	} else if d.HasChange("fqdn_cache_ttl") {
		obj["fqdn-cache-ttl"] = nil
	}

	if v, ok := d.GetOk("fqdn_max_refresh"); ok {
		if setArgNil {
			obj["fqdn-max-refresh"] = nil
		} else {
			t, err := expandSystemDnsFqdnMaxRefresh(d, v, "fqdn_max_refresh", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fqdn-max-refresh"] = t
			}
		}
	}

	if v, ok := d.GetOk("fqdn_min_refresh"); ok {
		if setArgNil {
			obj["fqdn-min-refresh"] = nil
		} else {
			t, err := expandSystemDnsFqdnMinRefresh(d, v, "fqdn_min_refresh", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fqdn-min-refresh"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostname_ttl"); ok {
		if setArgNil {
			obj["hostname-ttl"] = nil
		} else {
			t, err := expandSystemDnsHostnameTtl(d, v, "hostname_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("hostname_limit"); ok {
		if setArgNil {
			obj["hostname-limit"] = nil
		} else {
			t, err := expandSystemDnsHostnameLimit(d, v, "hostname_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname-limit"] = t
			}
		}
	}

	return &obj, nil
}
