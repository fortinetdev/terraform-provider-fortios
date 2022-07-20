// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Override settings for remote syslog server.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogSyslogd4OverrideSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd4OverrideSettingUpdate,
		Read:   resourceLogSyslogd4OverrideSettingRead,
		Update: resourceLogSyslogd4OverrideSettingUpdate,
		Delete: resourceLogSyslogd4OverrideSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				Computed:     true,
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
			"source_ip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
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
				Computed:     true,
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
				Computed:     true,
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
							Computed:     true,
						},
						"custom": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
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
				Computed:     true,
			},
			"syslog_type": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceLogSyslogd4OverrideSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogSyslogd4OverrideSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd4OverrideSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogSyslogd4OverrideSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd4OverrideSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogSyslogd4OverrideSetting")
	}

	return resourceLogSyslogd4OverrideSettingRead(d, m)
}

func resourceLogSyslogd4OverrideSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogSyslogd4OverrideSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd4OverrideSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogSyslogd4OverrideSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogSyslogd4OverrideSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd4OverrideSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogSyslogd4OverrideSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd4OverrideSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd4OverrideSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd4OverrideSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd4OverrideSettingOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingFacility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingCustomFieldName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenLogSyslogd4OverrideSettingCustomFieldNameId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenLogSyslogd4OverrideSettingCustomFieldNameName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom"
		if _, ok := i["custom"]; ok {

			tmp["custom"] = flattenLogSyslogd4OverrideSettingCustomFieldNameCustom(i["custom"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogSyslogd4OverrideSettingCustomFieldNameId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingCustomFieldNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingCustomFieldNameCustom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd4OverrideSettingSyslogType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogSyslogd4OverrideSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("override", flattenLogSyslogd4OverrideSettingOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("status", flattenLogSyslogd4OverrideSettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("server", flattenLogSyslogd4OverrideSettingServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("mode", flattenLogSyslogd4OverrideSettingMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("port", flattenLogSyslogd4OverrideSettingPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("facility", flattenLogSyslogd4OverrideSettingFacility(o["facility"], d, "facility", sv)); err != nil {
		if !fortiAPIPatch(o["facility"]) {
			return fmt.Errorf("Error reading facility: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogSyslogd4OverrideSettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("format", flattenLogSyslogd4OverrideSettingFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("priority", flattenLogSyslogd4OverrideSettingPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenLogSyslogd4OverrideSettingMaxLogRate(o["max-log-rate"], d, "max_log_rate", sv)); err != nil {
		if !fortiAPIPatch(o["max-log-rate"]) {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogSyslogd4OverrideSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogSyslogd4OverrideSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("certificate", flattenLogSyslogd4OverrideSettingCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_field_name", flattenLogSyslogd4OverrideSettingCustomFieldName(o["custom-field-name"], d, "custom_field_name", sv)); err != nil {
			if !fortiAPIPatch(o["custom-field-name"]) {
				return fmt.Errorf("Error reading custom_field_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_field_name"); ok {
			if err = d.Set("custom_field_name", flattenLogSyslogd4OverrideSettingCustomFieldName(o["custom-field-name"], d, "custom_field_name", sv)); err != nil {
				if !fortiAPIPatch(o["custom-field-name"]) {
					return fmt.Errorf("Error reading custom_field_name: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface_select_method", flattenLogSyslogd4OverrideSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogSyslogd4OverrideSettingInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("syslog_type", flattenLogSyslogd4OverrideSettingSyslogType(o["syslog-type"], d, "syslog_type", sv)); err != nil {
		if !fortiAPIPatch(o["syslog-type"]) {
			return fmt.Errorf("Error reading syslog_type: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd4OverrideSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogSyslogd4OverrideSettingOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingFacility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingCustomFieldName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandLogSyslogd4OverrideSettingCustomFieldNameId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandLogSyslogd4OverrideSettingCustomFieldNameName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["custom"], _ = expandLogSyslogd4OverrideSettingCustomFieldNameCustom(d, i["custom"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogSyslogd4OverrideSettingCustomFieldNameId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingCustomFieldNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingCustomFieldNameCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd4OverrideSettingSyslogType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd4OverrideSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("override"); ok {
		if setArgNil {
			obj["override"] = nil
		} else {

			t, err := expandLogSyslogd4OverrideSettingOverride(d, v, "override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override"] = t
			}
		}
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandLogSyslogd4OverrideSettingStatus(d, v, "status", sv)
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

			t, err := expandLogSyslogd4OverrideSettingServer(d, v, "server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server"] = t
			}
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {

			t, err := expandLogSyslogd4OverrideSettingMode(d, v, "mode", sv)
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

			t, err := expandLogSyslogd4OverrideSettingPort(d, v, "port", sv)
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

			t, err := expandLogSyslogd4OverrideSettingFacility(d, v, "facility", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["facility"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {

			t, err := expandLogSyslogd4OverrideSettingSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("format"); ok {
		if setArgNil {
			obj["format"] = nil
		} else {

			t, err := expandLogSyslogd4OverrideSettingFormat(d, v, "format", sv)
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

			t, err := expandLogSyslogd4OverrideSettingPriority(d, v, "priority", sv)
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

			t, err := expandLogSyslogd4OverrideSettingMaxLogRate(d, v, "max_log_rate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-log-rate"] = t
			}
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		if setArgNil {
			obj["enc-algorithm"] = nil
		} else {

			t, err := expandLogSyslogd4OverrideSettingEncAlgorithm(d, v, "enc_algorithm", sv)
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

			t, err := expandLogSyslogd4OverrideSettingSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
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

			t, err := expandLogSyslogd4OverrideSettingCertificate(d, v, "certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certificate"] = t
			}
		}
	}

	if v, ok := d.GetOk("custom_field_name"); ok {
		if setArgNil {
			obj["custom-field-name"] = make([]struct{}, 0)
		} else {

			t, err := expandLogSyslogd4OverrideSettingCustomFieldName(d, v, "custom_field_name", sv)
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

			t, err := expandLogSyslogd4OverrideSettingInterfaceSelectMethod(d, v, "interface_select_method", sv)
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

			t, err := expandLogSyslogd4OverrideSettingInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("syslog_type"); ok {
		if setArgNil {
			obj["syslog-type"] = nil
		} else {

			t, err := expandLogSyslogd4OverrideSettingSyslogType(d, v, "syslog_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["syslog-type"] = t
			}
		}
	}

	return &obj, nil
}
