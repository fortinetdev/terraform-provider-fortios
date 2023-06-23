// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure MS Exchange server entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserExchange() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserExchangeCreate,
		Read:   resourceUserExchangeRead,
		Update: resourceUserExchangeUpdate,
		Delete: resourceUserExchangeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"server_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"domain_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connect_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discover_kdc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kdc_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4": &schema.Schema{
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

func resourceUserExchangeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserExchange(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserExchange resource while getting object: %v", err)
	}

	o, err := c.CreateUserExchange(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserExchange resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserExchange")
	}

	return resourceUserExchangeRead(d, m)
}

func resourceUserExchangeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserExchange(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserExchange resource while getting object: %v", err)
	}

	o, err := c.UpdateUserExchange(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserExchange resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserExchange")
	}

	return resourceUserExchangeRead(d, m)
}

func resourceUserExchangeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserExchange(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserExchange resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserExchangeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserExchange(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserExchange resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserExchange(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserExchange resource from API: %v", err)
	}
	return nil
}

func flattenUserExchangeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeConnectProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeAuthLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeHttpAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeAutoDiscoverKdc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserExchangeKdcIp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4"
		if _, ok := i["ipv4"]; ok {
			tmp["ipv4"] = flattenUserExchangeKdcIpIpv4(i["ipv4"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ipv4", d)
	return result
}

func flattenUserExchangeKdcIpIpv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserExchange(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenUserExchangeName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_name", flattenUserExchangeServerName(o["server-name"], d, "server_name", sv)); err != nil {
		if !fortiAPIPatch(o["server-name"]) {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenUserExchangeDomainName(o["domain-name"], d, "domain_name", sv)); err != nil {
		if !fortiAPIPatch(o["domain-name"]) {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("username", flattenUserExchangeUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("ip", flattenUserExchangeIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("connect_protocol", flattenUserExchangeConnectProtocol(o["connect-protocol"], d, "connect_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["connect-protocol"]) {
			return fmt.Errorf("Error reading connect_protocol: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenUserExchangeAuthType(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("auth_level", flattenUserExchangeAuthLevel(o["auth-level"], d, "auth_level", sv)); err != nil {
		if !fortiAPIPatch(o["auth-level"]) {
			return fmt.Errorf("Error reading auth_level: %v", err)
		}
	}

	if err = d.Set("http_auth_type", flattenUserExchangeHttpAuthType(o["http-auth-type"], d, "http_auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["http-auth-type"]) {
			return fmt.Errorf("Error reading http_auth_type: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenUserExchangeSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("auto_discover_kdc", flattenUserExchangeAutoDiscoverKdc(o["auto-discover-kdc"], d, "auto_discover_kdc", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discover-kdc"]) {
			return fmt.Errorf("Error reading auto_discover_kdc: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("kdc_ip", flattenUserExchangeKdcIp(o["kdc-ip"], d, "kdc_ip", sv)); err != nil {
			if !fortiAPIPatch(o["kdc-ip"]) {
				return fmt.Errorf("Error reading kdc_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("kdc_ip"); ok {
			if err = d.Set("kdc_ip", flattenUserExchangeKdcIp(o["kdc-ip"], d, "kdc_ip", sv)); err != nil {
				if !fortiAPIPatch(o["kdc-ip"]) {
					return fmt.Errorf("Error reading kdc_ip: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenUserExchangeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserExchangeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeConnectProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeAuthLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeHttpAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeAutoDiscoverKdc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserExchangeKdcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipv4"], _ = expandUserExchangeKdcIpIpv4(d, i["ipv4"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserExchangeKdcIpIpv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserExchange(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserExchangeName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok {
		t, err := expandUserExchangeServerName(d, v, "server_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {
		t, err := expandUserExchangeDomainName(d, v, "domain_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandUserExchangeUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandUserExchangePassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandUserExchangeIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("connect_protocol"); ok {
		t, err := expandUserExchangeConnectProtocol(d, v, "connect_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connect-protocol"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandUserExchangeAuthType(d, v, "auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("auth_level"); ok {
		t, err := expandUserExchangeAuthLevel(d, v, "auth_level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-level"] = t
		}
	}

	if v, ok := d.GetOk("http_auth_type"); ok {
		t, err := expandUserExchangeHttpAuthType(d, v, "http_auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-auth-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		t, err := expandUserExchangeSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("auto_discover_kdc"); ok {
		t, err := expandUserExchangeAutoDiscoverKdc(d, v, "auto_discover_kdc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discover-kdc"] = t
		}
	}

	if v, ok := d.GetOk("kdc_ip"); ok || d.HasChange("kdc_ip") {
		t, err := expandUserExchangeKdcIp(d, v, "kdc_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kdc-ip"] = t
		}
	}

	return &obj, nil
}
