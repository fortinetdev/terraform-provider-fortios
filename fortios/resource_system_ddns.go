// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DDNS.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDdns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDdnsCreate,
		Read:   resourceSystemDdnsRead,
		Update: resourceSystemDdnsUpdate,
		Delete: resourceSystemDdnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ddnsid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ddns_server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 256),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_zone": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"ddns_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
			},
			"ddns_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_keyname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"ddns_key": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ddns_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"ddns_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"ddns_sn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"ddns_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"use_public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 2592000),
				Optional:     true,
				Computed:     true,
			},
			"clear_text": &schema.Schema{
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
			"bound_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
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

func resourceSystemDdnsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemDdns(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemDdns resource while getting object: %v", err)
	}

	o, err := c.CreateSystemDdns(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemDdns resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDdns")
	}

	return resourceSystemDdnsRead(d, m)
}

func resourceSystemDdnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemDdns(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDdns resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDdns(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemDdns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDdns")
	}

	return resourceSystemDdnsRead(d, m)
}

func resourceSystemDdnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemDdns(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDdns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDdnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemDdns(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemDdns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDdns(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDdns resource from API: %v", err)
	}
	return nil
}

func flattenSystemDdnsDdnsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsServerAddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr"
		if cur_v, ok := i["addr"]; ok {
			tmp["addr"] = flattenSystemDdnsDdnsServerAddrAddr(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "addr", d)
	return result
}

func flattenSystemDdnsDdnsServerAddrAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsZone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsKeyname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsSn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsDdnsPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsUsePublicIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsClearText(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsSslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsBoundIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDdnsMonitorInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if cur_v, ok := i["interface-name"]; ok {
			tmp["interface_name"] = flattenSystemDdnsMonitorInterfaceInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "interface_name", d)
	return result
}

func flattenSystemDdnsMonitorInterfaceInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDdns(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("ddnsid", flattenSystemDdnsDdnsid(o["ddnsid"], d, "ddnsid", sv)); err != nil {
		if !fortiAPIPatch(o["ddnsid"]) {
			return fmt.Errorf("Error reading ddnsid: %v", err)
		}
	}

	if err = d.Set("ddns_server", flattenSystemDdnsDdnsServer(o["ddns-server"], d, "ddns_server", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-server"]) {
			return fmt.Errorf("Error reading ddns_server: %v", err)
		}
	}

	if err = d.Set("server_type", flattenSystemDdnsServerType(o["server-type"], d, "server_type", sv)); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ddns_server_addr", flattenSystemDdnsDdnsServerAddr(o["ddns-server-addr"], d, "ddns_server_addr", sv)); err != nil {
			if !fortiAPIPatch(o["ddns-server-addr"]) {
				return fmt.Errorf("Error reading ddns_server_addr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ddns_server_addr"); ok {
			if err = d.Set("ddns_server_addr", flattenSystemDdnsDdnsServerAddr(o["ddns-server-addr"], d, "ddns_server_addr", sv)); err != nil {
				if !fortiAPIPatch(o["ddns-server-addr"]) {
					return fmt.Errorf("Error reading ddns_server_addr: %v", err)
				}
			}
		}
	}

	if err = d.Set("ddns_server_ip", flattenSystemDdnsDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip"]) {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenSystemDdnsDdnsZone(o["ddns-zone"], d, "ddns_zone", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-zone"]) {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenSystemDdnsDdnsTtl(o["ddns-ttl"], d, "ddns_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-ttl"]) {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenSystemDdnsDdnsAuth(o["ddns-auth"], d, "ddns_auth", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-auth"]) {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenSystemDdnsDdnsKeyname(o["ddns-keyname"], d, "ddns_keyname", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-keyname"]) {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_domain", flattenSystemDdnsDdnsDomain(o["ddns-domain"], d, "ddns_domain", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-domain"]) {
			return fmt.Errorf("Error reading ddns_domain: %v", err)
		}
	}

	if err = d.Set("ddns_username", flattenSystemDdnsDdnsUsername(o["ddns-username"], d, "ddns_username", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-username"]) {
			return fmt.Errorf("Error reading ddns_username: %v", err)
		}
	}

	if err = d.Set("ddns_sn", flattenSystemDdnsDdnsSn(o["ddns-sn"], d, "ddns_sn", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-sn"]) {
			return fmt.Errorf("Error reading ddns_sn: %v", err)
		}
	}

	if err = d.Set("use_public_ip", flattenSystemDdnsUsePublicIp(o["use-public-ip"], d, "use_public_ip", sv)); err != nil {
		if !fortiAPIPatch(o["use-public-ip"]) {
			return fmt.Errorf("Error reading use_public_ip: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenSystemDdnsAddrType(o["addr-type"], d, "addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["addr-type"]) {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenSystemDdnsUpdateInterval(o["update-interval"], d, "update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["update-interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("clear_text", flattenSystemDdnsClearText(o["clear-text"], d, "clear_text", sv)); err != nil {
		if !fortiAPIPatch(o["clear-text"]) {
			return fmt.Errorf("Error reading clear_text: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenSystemDdnsSslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("bound_ip", flattenSystemDdnsBoundIp(o["bound-ip"], d, "bound_ip", sv)); err != nil {
		if !fortiAPIPatch(o["bound-ip"]) {
			return fmt.Errorf("Error reading bound_ip: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("monitor_interface", flattenSystemDdnsMonitorInterface(o["monitor-interface"], d, "monitor_interface", sv)); err != nil {
			if !fortiAPIPatch(o["monitor-interface"]) {
				return fmt.Errorf("Error reading monitor_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("monitor_interface"); ok {
			if err = d.Set("monitor_interface", flattenSystemDdnsMonitorInterface(o["monitor-interface"], d, "monitor_interface", sv)); err != nil {
				if !fortiAPIPatch(o["monitor-interface"]) {
					return fmt.Errorf("Error reading monitor_interface: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemDdnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDdnsDdnsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsServerAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["addr"], _ = expandSystemDdnsDdnsServerAddrAddr(d, i["addr"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDdnsDdnsServerAddrAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsKeyname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsSn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsDdnsPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsUsePublicIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsClearText(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsSslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsBoundIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDdnsMonitorInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["interface-name"], _ = expandSystemDdnsMonitorInterfaceInterfaceName(d, i["interface_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDdnsMonitorInterfaceInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDdns(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("ddnsid"); ok {
		t, err := expandSystemDdnsDdnsid(d, v, "ddnsid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddnsid"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server"); ok {
		t, err := expandSystemDdnsDdnsServer(d, v, "ddns_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {
		t, err := expandSystemDdnsServerType(d, v, "server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_addr"); ok || d.HasChange("ddns_server_addr") {
		t, err := expandSystemDdnsDdnsServerAddr(d, v, "ddns_server_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-addr"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok {
		t, err := expandSystemDdnsDdnsServerIp(d, v, "ddns_server_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok {
		t, err := expandSystemDdnsDdnsZone(d, v, "ddns_zone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	}

	if v, ok := d.GetOk("ddns_ttl"); ok {
		t, err := expandSystemDdnsDdnsTtl(d, v, "ddns_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ddns_auth"); ok {
		t, err := expandSystemDdnsDdnsAuth(d, v, "ddns_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok {
		t, err := expandSystemDdnsDdnsKeyname(d, v, "ddns_keyname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	}

	if v, ok := d.GetOk("ddns_key"); ok {
		t, err := expandSystemDdnsDdnsKey(d, v, "ddns_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	}

	if v, ok := d.GetOk("ddns_domain"); ok {
		t, err := expandSystemDdnsDdnsDomain(d, v, "ddns_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-domain"] = t
		}
	}

	if v, ok := d.GetOk("ddns_username"); ok {
		t, err := expandSystemDdnsDdnsUsername(d, v, "ddns_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-username"] = t
		}
	}

	if v, ok := d.GetOk("ddns_sn"); ok {
		t, err := expandSystemDdnsDdnsSn(d, v, "ddns_sn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-sn"] = t
		}
	}

	if v, ok := d.GetOk("ddns_password"); ok {
		t, err := expandSystemDdnsDdnsPassword(d, v, "ddns_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-password"] = t
		}
	}

	if v, ok := d.GetOk("use_public_ip"); ok {
		t, err := expandSystemDdnsUsePublicIp(d, v, "use_public_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-public-ip"] = t
		}
	}

	if v, ok := d.GetOk("addr_type"); ok {
		t, err := expandSystemDdnsAddrType(d, v, "addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("update_interval"); ok {
		t, err := expandSystemDdnsUpdateInterval(d, v, "update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("clear_text"); ok {
		t, err := expandSystemDdnsClearText(d, v, "clear_text", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clear-text"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok {
		t, err := expandSystemDdnsSslCertificate(d, v, "ssl_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("bound_ip"); ok {
		t, err := expandSystemDdnsBoundIp(d, v, "bound_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bound-ip"] = t
		}
	}

	if v, ok := d.GetOk("monitor_interface"); ok || d.HasChange("monitor_interface") {
		t, err := expandSystemDdnsMonitorInterface(d, v, "monitor_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-interface"] = t
		}
	}

	return &obj, nil
}
