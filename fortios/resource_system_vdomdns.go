// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DNS servers for a non-management VDOM.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVdomDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomDnsUpdate,
		Read:   resourceSystemVdomDnsRead,
		Update: resourceSystemVdomDnsUpdate,
		Delete: resourceSystemVdomDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vdom_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemVdomDnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVdomDns(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomDns resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdomDns(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomDns")
	}

	return resourceSystemVdomDnsRead(d, m)
}

func resourceSystemVdomDnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVdomDns(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemVdomDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomDns(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemVdomDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemVdomDns(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomDns(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomDnsVdomDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsSecondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsDnsOverTls(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsSslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsServerHostname(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {

			tmp["hostname"] = flattenSystemVdomDnsServerHostnameHostname(i["hostname"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "hostname", d)
	return result
}

func flattenSystemVdomDnsServerHostnameHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsServerSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsAltPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomDnsAltSecondary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVdomDns(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("vdom_dns", flattenSystemVdomDnsVdomDns(o["vdom-dns"], d, "vdom_dns", sv)); err != nil {
		if !fortiAPIPatch(o["vdom-dns"]) {
			return fmt.Errorf("Error reading vdom_dns: %v", err)
		}
	}

	if err = d.Set("primary", flattenSystemVdomDnsPrimary(o["primary"], d, "primary", sv)); err != nil {
		if !fortiAPIPatch(o["primary"]) {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystemVdomDnsSecondary(o["secondary"], d, "secondary", sv)); err != nil {
		if !fortiAPIPatch(o["secondary"]) {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemVdomDnsProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("dns_over_tls", flattenSystemVdomDnsDnsOverTls(o["dns-over-tls"], d, "dns_over_tls", sv)); err != nil {
		if !fortiAPIPatch(o["dns-over-tls"]) {
			return fmt.Errorf("Error reading dns_over_tls: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystemVdomDnsSslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_hostname", flattenSystemVdomDnsServerHostname(o["server-hostname"], d, "server_hostname", sv)); err != nil {
			if !fortiAPIPatch(o["server-hostname"]) {
				return fmt.Errorf("Error reading server_hostname: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_hostname"); ok {
			if err = d.Set("server_hostname", flattenSystemVdomDnsServerHostname(o["server-hostname"], d, "server_hostname", sv)); err != nil {
				if !fortiAPIPatch(o["server-hostname"]) {
					return fmt.Errorf("Error reading server_hostname: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_primary", flattenSystemVdomDnsIp6Primary(o["ip6-primary"], d, "ip6_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-primary"]) {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", flattenSystemVdomDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-secondary"]) {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemVdomDnsSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemVdomDnsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVdomDnsInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("server_select_method", flattenSystemVdomDnsServerSelectMethod(o["server-select-method"], d, "server_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["server-select-method"]) {
			return fmt.Errorf("Error reading server_select_method: %v", err)
		}
	}

	if err = d.Set("alt_primary", flattenSystemVdomDnsAltPrimary(o["alt-primary"], d, "alt_primary", sv)); err != nil {
		if !fortiAPIPatch(o["alt-primary"]) {
			return fmt.Errorf("Error reading alt_primary: %v", err)
		}
	}

	if err = d.Set("alt_secondary", flattenSystemVdomDnsAltSecondary(o["alt-secondary"], d, "alt_secondary", sv)); err != nil {
		if !fortiAPIPatch(o["alt-secondary"]) {
			return fmt.Errorf("Error reading alt_secondary: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemVdomDnsVdomDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsSecondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsDnsOverTls(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsSslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsServerHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hostname"], _ = expandSystemVdomDnsServerHostnameHostname(d, i["hostname"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVdomDnsServerHostnameHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsIp6Primary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsIp6Secondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsServerSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsAltPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomDnsAltSecondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomDns(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vdom_dns"); ok {
		if setArgNil {
			obj["vdom-dns"] = nil
		} else {

			t, err := expandSystemVdomDnsVdomDns(d, v, "vdom_dns", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vdom-dns"] = t
			}
		}
	}

	if v, ok := d.GetOk("primary"); ok {
		if setArgNil {
			obj["primary"] = nil
		} else {

			t, err := expandSystemVdomDnsPrimary(d, v, "primary", sv)
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

			t, err := expandSystemVdomDnsSecondary(d, v, "secondary", sv)
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

			t, err := expandSystemVdomDnsProtocol(d, v, "protocol", sv)
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

			t, err := expandSystemVdomDnsDnsOverTls(d, v, "dns_over_tls", sv)
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

			t, err := expandSystemVdomDnsSslCertificate(d, v, "ssl_certificate", sv)
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

			t, err := expandSystemVdomDnsServerHostname(d, v, "server_hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-hostname"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip6_primary"); ok {
		if setArgNil {
			obj["ip6-primary"] = nil
		} else {

			t, err := expandSystemVdomDnsIp6Primary(d, v, "ip6_primary", sv)
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

			t, err := expandSystemVdomDnsIp6Secondary(d, v, "ip6_secondary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip6-secondary"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandSystemVdomDnsSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		if setArgNil {
			obj["interface-select-method"] = nil
		} else {

			t, err := expandSystemVdomDnsInterfaceSelectMethod(d, v, "interface_select_method", sv)
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

			t, err := expandSystemVdomDnsInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_select_method"); ok {
		if setArgNil {
			obj["server-select-method"] = nil
		} else {

			t, err := expandSystemVdomDnsServerSelectMethod(d, v, "server_select_method", sv)
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

			t, err := expandSystemVdomDnsAltPrimary(d, v, "alt_primary", sv)
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

			t, err := expandSystemVdomDnsAltSecondary(d, v, "alt_secondary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["alt-secondary"] = t
			}
		}
	}

	return &obj, nil
}
