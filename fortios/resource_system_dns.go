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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"domain": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
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
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemDns(d, c.Fv)
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

	err := c.DeleteSystemDns(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {

			tmp["hostname"] = flattenSystemDnsServerHostnameHostname(i["hostname"], d, pre_append, sv)
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

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {

			tmp["domain"] = flattenSystemDnsDomainDomain(i["domain"], d, pre_append, sv)
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
	return v
}

func flattenSystemDnsRetry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDnsCacheLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDnsCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsCacheNotfoundResponses(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

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

	if isImportTable() {
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

	if isImportTable() {
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

func expandSystemDnsDnsOverTls(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsServerHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hostname"], _ = expandSystemDnsServerHostnameHostname(d, i["hostname"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDnsServerHostnameHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["domain"], _ = expandSystemDnsDomainDomain(d, i["domain"], pre_append, sv)
		}

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

func expandSystemDnsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("primary"); ok {

		t, err := expandSystemDnsPrimary(d, v, "primary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary"] = t
		}
	}

	if v, ok := d.GetOk("secondary"); ok {

		t, err := expandSystemDnsSecondary(d, v, "secondary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary"] = t
		}
	}

	if v, ok := d.GetOk("dns_over_tls"); ok {

		t, err := expandSystemDnsDnsOverTls(d, v, "dns_over_tls", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-over-tls"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok {

		t, err := expandSystemDnsSslCertificate(d, v, "ssl_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("server_hostname"); ok {

		t, err := expandSystemDnsServerHostname(d, v, "server_hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-hostname"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {

		t, err := expandSystemDnsDomain(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("ip6_primary"); ok {

		t, err := expandSystemDnsIp6Primary(d, v, "ip6_primary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-primary"] = t
		}
	}

	if v, ok := d.GetOk("ip6_secondary"); ok {

		t, err := expandSystemDnsIp6Secondary(d, v, "ip6_secondary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-secondary"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {

		t, err := expandSystemDnsTimeout(d, v, "timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("retry"); ok {

		t, err := expandSystemDnsRetry(d, v, "retry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry"] = t
		}
	}

	if v, ok := d.GetOkExists("dns_cache_limit"); ok {

		t, err := expandSystemDnsDnsCacheLimit(d, v, "dns_cache_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-cache-limit"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_ttl"); ok {

		t, err := expandSystemDnsDnsCacheTtl(d, v, "dns_cache_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("cache_notfound_responses"); ok {

		t, err := expandSystemDnsCacheNotfoundResponses(d, v, "cache_notfound_responses", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-notfound-responses"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandSystemDnsSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {

		t, err := expandSystemDnsInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandSystemDnsInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
