// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Global settings for remote syslog server.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogSyslogd3Setting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd3SettingUpdate,
		Read:   resourceLogSyslogd3SettingRead,
		Update: resourceLogSyslogd3SettingUpdate,
		Delete: resourceLogSyslogd3SettingDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"facility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"source_ip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_log_rate": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100000),
				Optional:     true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"custom_field_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"custom": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
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
			"syslog_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceLogSyslogd3SettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogSyslogd3Setting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3Setting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogSyslogd3Setting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3Setting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogSyslogd3Setting")
	}

	return resourceLogSyslogd3SettingRead(d, m)
}

func resourceLogSyslogd3SettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogSyslogd3Setting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3Setting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogd3Setting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogSyslogd3Setting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd3SettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogSyslogd3Setting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd3Setting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd3Setting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd3Setting resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd3SettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogSyslogd3SettingFacility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingSourceIpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogSyslogd3SettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingCustomFieldName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenLogSyslogd3SettingCustomFieldNameId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenLogSyslogd3SettingCustomFieldNameName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom"
		if cur_v, ok := i["custom"]; ok {
			tmp["custom"] = flattenLogSyslogd3SettingCustomFieldNameCustom(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogSyslogd3SettingCustomFieldNameId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogSyslogd3SettingCustomFieldNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingCustomFieldNameCustom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3SettingVrfSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogSyslogd3SettingSyslogType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectLogSyslogd3Setting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("status", flattenLogSyslogd3SettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("server", flattenLogSyslogd3SettingServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("mode", flattenLogSyslogd3SettingMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("port", flattenLogSyslogd3SettingPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("facility", flattenLogSyslogd3SettingFacility(o["facility"], d, "facility", sv)); err != nil {
		if !fortiAPIPatch(o["facility"]) {
			return fmt.Errorf("Error reading facility: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenLogSyslogd3SettingSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip-interface"]) {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogSyslogd3SettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("format", flattenLogSyslogd3SettingFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("priority", flattenLogSyslogd3SettingPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenLogSyslogd3SettingMaxLogRate(o["max-log-rate"], d, "max_log_rate", sv)); err != nil {
		if !fortiAPIPatch(o["max-log-rate"]) {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogSyslogd3SettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogSyslogd3SettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("certificate", flattenLogSyslogd3SettingCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("custom_field_name", flattenLogSyslogd3SettingCustomFieldName(o["custom-field-name"], d, "custom_field_name", sv)); err != nil {
			if !fortiAPIPatch(o["custom-field-name"]) {
				return fmt.Errorf("Error reading custom_field_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_field_name"); ok {
			if err = d.Set("custom_field_name", flattenLogSyslogd3SettingCustomFieldName(o["custom-field-name"], d, "custom_field_name", sv)); err != nil {
				if !fortiAPIPatch(o["custom-field-name"]) {
					return fmt.Errorf("Error reading custom_field_name: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface_select_method", flattenLogSyslogd3SettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogSyslogd3SettingInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenLogSyslogd3SettingVrfSelect(o["vrf-select"], d, "vrf_select", sv)); err != nil {
		if !fortiAPIPatch(o["vrf-select"]) {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	if err = d.Set("syslog_type", flattenLogSyslogd3SettingSyslogType(o["syslog-type"], d, "syslog_type", sv)); err != nil {
		if !fortiAPIPatch(o["syslog-type"]) {
			return fmt.Errorf("Error reading syslog_type: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd3SettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogSyslogd3SettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingFacility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingSourceIpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingCustomFieldName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandLogSyslogd3SettingCustomFieldNameId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandLogSyslogd3SettingCustomFieldNameName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["custom"], _ = expandLogSyslogd3SettingCustomFieldNameCustom(d, i["custom"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["custom"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogSyslogd3SettingCustomFieldNameId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingCustomFieldNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingCustomFieldNameCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingVrfSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3SettingSyslogType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd3Setting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandLogSyslogd3SettingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("server"); ok {
		if setArgNil {
			obj["server"] = nil
		} else {
			t, err := expandLogSyslogd3SettingServer(d, v, "server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server"] = t
			}
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {
			t, err := expandLogSyslogd3SettingMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("port"); ok {
		if setArgNil {
			obj["port"] = nil
		} else {
			t, err := expandLogSyslogd3SettingPort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	if v, ok := d.GetOk("facility"); ok {
		if setArgNil {
			obj["facility"] = nil
		} else {
			t, err := expandLogSyslogd3SettingFacility(d, v, "facility", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["facility"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok {
		if setArgNil {
			obj["source-ip-interface"] = nil
		} else {
			t, err := expandLogSyslogd3SettingSourceIpInterface(d, v, "source_ip_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip-interface"] = t
			}
		}
	} else if d.HasChange("source_ip_interface") {
		obj["source-ip-interface"] = nil
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {
			t, err := expandLogSyslogd3SettingSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	} else if d.HasChange("source_ip") {
		obj["source-ip"] = nil
	}

	if v, ok := d.GetOk("format"); ok {
		if setArgNil {
			obj["format"] = nil
		} else {
			t, err := expandLogSyslogd3SettingFormat(d, v, "format", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["format"] = t
			}
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		if setArgNil {
			obj["priority"] = nil
		} else {
			t, err := expandLogSyslogd3SettingPriority(d, v, "priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["priority"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_log_rate"); ok {
		if setArgNil {
			obj["max-log-rate"] = nil
		} else {
			t, err := expandLogSyslogd3SettingMaxLogRate(d, v, "max_log_rate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-log-rate"] = t
			}
		}
	} else if d.HasChange("max_log_rate") {
		obj["max-log-rate"] = nil
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		if setArgNil {
			obj["enc-algorithm"] = nil
		} else {
			t, err := expandLogSyslogd3SettingEncAlgorithm(d, v, "enc_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["enc-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		if setArgNil {
			obj["ssl-min-proto-version"] = nil
		} else {
			t, err := expandLogSyslogd3SettingSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-min-proto-version"] = t
			}
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		if setArgNil {
			obj["certificate"] = nil
		} else {
			t, err := expandLogSyslogd3SettingCertificate(d, v, "certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certificate"] = t
			}
		}
	} else if d.HasChange("certificate") {
		obj["certificate"] = nil
	}

	if v, ok := d.GetOk("custom_field_name"); ok || d.HasChange("custom_field_name") {
		if setArgNil {
			obj["custom-field-name"] = make([]struct{}, 0)
		} else {
			t, err := expandLogSyslogd3SettingCustomFieldName(d, v, "custom_field_name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["custom-field-name"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		if setArgNil {
			obj["interface-select-method"] = nil
		} else {
			t, err := expandLogSyslogd3SettingInterfaceSelectMethod(d, v, "interface_select_method", sv)
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
			t, err := expandLogSyslogd3SettingInterface(d, v, "interface", sv)
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
			t, err := expandLogSyslogd3SettingVrfSelect(d, v, "vrf_select", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vrf-select"] = t
			}
		}
	} else if d.HasChange("vrf_select") {
		obj["vrf-select"] = nil
	}

	if v, ok := d.GetOkExists("syslog_type"); ok {
		if setArgNil {
			obj["syslog-type"] = nil
		} else {
			t, err := expandLogSyslogd3SettingSyslogType(d, v, "syslog_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["syslog-type"] = t
			}
		}
	} else if d.HasChange("syslog_type") {
		obj["syslog-type"] = nil
	}

	return &obj, nil
}
