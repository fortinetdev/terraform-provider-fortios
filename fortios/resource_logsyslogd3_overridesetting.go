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

func resourceLogSyslogd3OverrideSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd3OverrideSettingUpdate,
		Read:   resourceLogSyslogd3OverrideSettingRead,
		Update: resourceLogSyslogd3OverrideSettingUpdate,
		Delete: resourceLogSyslogd3OverrideSettingDelete,

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

func resourceLogSyslogd3OverrideSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogSyslogd3OverrideSetting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3OverrideSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogSyslogd3OverrideSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3OverrideSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogSyslogd3OverrideSetting")
	}

	return resourceLogSyslogd3OverrideSettingRead(d, m)
}

func resourceLogSyslogd3OverrideSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteLogSyslogd3OverrideSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogd3OverrideSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd3OverrideSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogSyslogd3OverrideSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd3OverrideSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd3OverrideSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd3OverrideSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogSyslogd3OverrideSettingOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingFacility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingCustomFieldName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenLogSyslogd3OverrideSettingCustomFieldNameId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenLogSyslogd3OverrideSettingCustomFieldNameName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom"
		if _, ok := i["custom"]; ok {

			tmp["custom"] = flattenLogSyslogd3OverrideSettingCustomFieldNameCustom(i["custom"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogSyslogd3OverrideSettingCustomFieldNameId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingCustomFieldNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingCustomFieldNameCustom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSyslogd3OverrideSettingSyslogType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogSyslogd3OverrideSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("override", flattenLogSyslogd3OverrideSettingOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("status", flattenLogSyslogd3OverrideSettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("server", flattenLogSyslogd3OverrideSettingServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("mode", flattenLogSyslogd3OverrideSettingMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("port", flattenLogSyslogd3OverrideSettingPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("facility", flattenLogSyslogd3OverrideSettingFacility(o["facility"], d, "facility", sv)); err != nil {
		if !fortiAPIPatch(o["facility"]) {
			return fmt.Errorf("Error reading facility: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogSyslogd3OverrideSettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("format", flattenLogSyslogd3OverrideSettingFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("priority", flattenLogSyslogd3OverrideSettingPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenLogSyslogd3OverrideSettingMaxLogRate(o["max-log-rate"], d, "max_log_rate", sv)); err != nil {
		if !fortiAPIPatch(o["max-log-rate"]) {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogSyslogd3OverrideSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogSyslogd3OverrideSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("certificate", flattenLogSyslogd3OverrideSettingCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_field_name", flattenLogSyslogd3OverrideSettingCustomFieldName(o["custom-field-name"], d, "custom_field_name", sv)); err != nil {
			if !fortiAPIPatch(o["custom-field-name"]) {
				return fmt.Errorf("Error reading custom_field_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_field_name"); ok {
			if err = d.Set("custom_field_name", flattenLogSyslogd3OverrideSettingCustomFieldName(o["custom-field-name"], d, "custom_field_name", sv)); err != nil {
				if !fortiAPIPatch(o["custom-field-name"]) {
					return fmt.Errorf("Error reading custom_field_name: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface_select_method", flattenLogSyslogd3OverrideSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogSyslogd3OverrideSettingInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("syslog_type", flattenLogSyslogd3OverrideSettingSyslogType(o["syslog-type"], d, "syslog_type", sv)); err != nil {
		if !fortiAPIPatch(o["syslog-type"]) {
			return fmt.Errorf("Error reading syslog_type: %v", err)
		}
	}

	return nil
}

func flattenLogSyslogd3OverrideSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogSyslogd3OverrideSettingOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingFacility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingCustomFieldName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandLogSyslogd3OverrideSettingCustomFieldNameId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandLogSyslogd3OverrideSettingCustomFieldNameName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["custom"], _ = expandLogSyslogd3OverrideSettingCustomFieldNameCustom(d, i["custom"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogSyslogd3OverrideSettingCustomFieldNameId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingCustomFieldNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingCustomFieldNameCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3OverrideSettingSyslogType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogSyslogd3OverrideSetting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("override"); ok {

		t, err := expandLogSyslogd3OverrideSettingOverride(d, v, "override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandLogSyslogd3OverrideSettingStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandLogSyslogd3OverrideSettingServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {

		t, err := expandLogSyslogd3OverrideSettingMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOkExists("port"); ok {

		t, err := expandLogSyslogd3OverrideSettingPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("facility"); ok {

		t, err := expandLogSyslogd3OverrideSettingFacility(d, v, "facility", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["facility"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandLogSyslogd3OverrideSettingSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok {

		t, err := expandLogSyslogd3OverrideSettingFormat(d, v, "format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {

		t, err := expandLogSyslogd3OverrideSettingPriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOkExists("max_log_rate"); ok {

		t, err := expandLogSyslogd3OverrideSettingMaxLogRate(d, v, "max_log_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-rate"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {

		t, err := expandLogSyslogd3OverrideSettingEncAlgorithm(d, v, "enc_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {

		t, err := expandLogSyslogd3OverrideSettingSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {

		t, err := expandLogSyslogd3OverrideSettingCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("custom_field_name"); ok {

		t, err := expandLogSyslogd3OverrideSettingCustomFieldName(d, v, "custom_field_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-field-name"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {

		t, err := expandLogSyslogd3OverrideSettingInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandLogSyslogd3OverrideSettingInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOkExists("syslog_type"); ok {

		t, err := expandLogSyslogd3OverrideSettingSyslogType(d, v, "syslog_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-type"] = t
		}
	}

	return &obj, nil
}
