// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: VDOM wireless controller configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSettingUpdate,
		Read:   resourceWirelessControllerSettingRead,
		Update: resourceWirelessControllerSettingUpdate,
		Delete: resourceWirelessControllerSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"account_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"duplicate_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fapc_compatibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wfa_compatibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phishing_ssid_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fake_ssid_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"offending_ssid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"ssid_pattern": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 33),
							Optional:     true,
							Computed:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"device_weight": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"device_holdoff": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 60),
				Optional:     true,
				Computed:     true,
			},
			"device_idle": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 14400),
				Optional:     true,
				Computed:     true,
			},
			"firmware_provision_on_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"darrp_optimize": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),
				Optional:     true,
				Computed:     true,
			},
			"darrp_optimize_schedules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
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
		},
	}
}

func resourceWirelessControllerSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerSetting")
	}

	return resourceWirelessControllerSettingRead(d, m)
}

func resourceWirelessControllerSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing WirelessControllerSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSetting resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSettingAccountId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingCountry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingDuplicateSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingFapcCompatibility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingWfaCompatibility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingPhishingSsidDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingFakeSsidAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingOffendingSsid(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenWirelessControllerSettingOffendingSsidId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid_pattern"
		if _, ok := i["ssid-pattern"]; ok {

			tmp["ssid_pattern"] = flattenWirelessControllerSettingOffendingSsidSsidPattern(i["ssid-pattern"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenWirelessControllerSettingOffendingSsidAction(i["action"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerSettingOffendingSsidId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingOffendingSsidSsidPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingOffendingSsidAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingDeviceWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingDeviceHoldoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingDeviceIdle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingFirmwareProvisionOnAuthorization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingDarrpOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSettingDarrpOptimizeSchedules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenWirelessControllerSettingDarrpOptimizeSchedulesName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerSettingDarrpOptimizeSchedulesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("account_id", flattenWirelessControllerSettingAccountId(o["account-id"], d, "account_id", sv)); err != nil {
		if !fortiAPIPatch(o["account-id"]) {
			return fmt.Errorf("Error reading account_id: %v", err)
		}
	}

	if err = d.Set("country", flattenWirelessControllerSettingCountry(o["country"], d, "country", sv)); err != nil {
		if !fortiAPIPatch(o["country"]) {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("duplicate_ssid", flattenWirelessControllerSettingDuplicateSsid(o["duplicate-ssid"], d, "duplicate_ssid", sv)); err != nil {
		if !fortiAPIPatch(o["duplicate-ssid"]) {
			return fmt.Errorf("Error reading duplicate_ssid: %v", err)
		}
	}

	if err = d.Set("fapc_compatibility", flattenWirelessControllerSettingFapcCompatibility(o["fapc-compatibility"], d, "fapc_compatibility", sv)); err != nil {
		if !fortiAPIPatch(o["fapc-compatibility"]) {
			return fmt.Errorf("Error reading fapc_compatibility: %v", err)
		}
	}

	if err = d.Set("wfa_compatibility", flattenWirelessControllerSettingWfaCompatibility(o["wfa-compatibility"], d, "wfa_compatibility", sv)); err != nil {
		if !fortiAPIPatch(o["wfa-compatibility"]) {
			return fmt.Errorf("Error reading wfa_compatibility: %v", err)
		}
	}

	if err = d.Set("phishing_ssid_detect", flattenWirelessControllerSettingPhishingSsidDetect(o["phishing-ssid-detect"], d, "phishing_ssid_detect", sv)); err != nil {
		if !fortiAPIPatch(o["phishing-ssid-detect"]) {
			return fmt.Errorf("Error reading phishing_ssid_detect: %v", err)
		}
	}

	if err = d.Set("fake_ssid_action", flattenWirelessControllerSettingFakeSsidAction(o["fake-ssid-action"], d, "fake_ssid_action", sv)); err != nil {
		if !fortiAPIPatch(o["fake-ssid-action"]) {
			return fmt.Errorf("Error reading fake_ssid_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("offending_ssid", flattenWirelessControllerSettingOffendingSsid(o["offending-ssid"], d, "offending_ssid", sv)); err != nil {
			if !fortiAPIPatch(o["offending-ssid"]) {
				return fmt.Errorf("Error reading offending_ssid: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("offending_ssid"); ok {
			if err = d.Set("offending_ssid", flattenWirelessControllerSettingOffendingSsid(o["offending-ssid"], d, "offending_ssid", sv)); err != nil {
				if !fortiAPIPatch(o["offending-ssid"]) {
					return fmt.Errorf("Error reading offending_ssid: %v", err)
				}
			}
		}
	}

	if err = d.Set("device_weight", flattenWirelessControllerSettingDeviceWeight(o["device-weight"], d, "device_weight", sv)); err != nil {
		if !fortiAPIPatch(o["device-weight"]) {
			return fmt.Errorf("Error reading device_weight: %v", err)
		}
	}

	if err = d.Set("device_holdoff", flattenWirelessControllerSettingDeviceHoldoff(o["device-holdoff"], d, "device_holdoff", sv)); err != nil {
		if !fortiAPIPatch(o["device-holdoff"]) {
			return fmt.Errorf("Error reading device_holdoff: %v", err)
		}
	}

	if err = d.Set("device_idle", flattenWirelessControllerSettingDeviceIdle(o["device-idle"], d, "device_idle", sv)); err != nil {
		if !fortiAPIPatch(o["device-idle"]) {
			return fmt.Errorf("Error reading device_idle: %v", err)
		}
	}

	if err = d.Set("firmware_provision_on_authorization", flattenWirelessControllerSettingFirmwareProvisionOnAuthorization(o["firmware-provision-on-authorization"], d, "firmware_provision_on_authorization", sv)); err != nil {
		if !fortiAPIPatch(o["firmware-provision-on-authorization"]) {
			return fmt.Errorf("Error reading firmware_provision_on_authorization: %v", err)
		}
	}

	if err = d.Set("darrp_optimize", flattenWirelessControllerSettingDarrpOptimize(o["darrp-optimize"], d, "darrp_optimize", sv)); err != nil {
		if !fortiAPIPatch(o["darrp-optimize"]) {
			return fmt.Errorf("Error reading darrp_optimize: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("darrp_optimize_schedules", flattenWirelessControllerSettingDarrpOptimizeSchedules(o["darrp-optimize-schedules"], d, "darrp_optimize_schedules", sv)); err != nil {
			if !fortiAPIPatch(o["darrp-optimize-schedules"]) {
				return fmt.Errorf("Error reading darrp_optimize_schedules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("darrp_optimize_schedules"); ok {
			if err = d.Set("darrp_optimize_schedules", flattenWirelessControllerSettingDarrpOptimizeSchedules(o["darrp-optimize-schedules"], d, "darrp_optimize_schedules", sv)); err != nil {
				if !fortiAPIPatch(o["darrp-optimize-schedules"]) {
					return fmt.Errorf("Error reading darrp_optimize_schedules: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerSettingAccountId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingCountry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDuplicateSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingFapcCompatibility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingWfaCompatibility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingPhishingSsidDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingFakeSsidAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingOffendingSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandWirelessControllerSettingOffendingSsidId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid_pattern"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssid-pattern"], _ = expandWirelessControllerSettingOffendingSsidSsidPattern(d, i["ssid_pattern"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandWirelessControllerSettingOffendingSsidAction(d, i["action"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSettingOffendingSsidId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingOffendingSsidSsidPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingOffendingSsidAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDeviceWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDeviceHoldoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDeviceIdle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingFirmwareProvisionOnAuthorization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDarrpOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDarrpOptimizeSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandWirelessControllerSettingDarrpOptimizeSchedulesName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSettingDarrpOptimizeSchedulesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("account_id"); ok {
		if setArgNil {
			obj["account-id"] = nil
		} else {

			t, err := expandWirelessControllerSettingAccountId(d, v, "account_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["account-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("country"); ok {
		if setArgNil {
			obj["country"] = nil
		} else {

			t, err := expandWirelessControllerSettingCountry(d, v, "country", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["country"] = t
			}
		}
	}

	if v, ok := d.GetOk("duplicate_ssid"); ok {
		if setArgNil {
			obj["duplicate-ssid"] = nil
		} else {

			t, err := expandWirelessControllerSettingDuplicateSsid(d, v, "duplicate_ssid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["duplicate-ssid"] = t
			}
		}
	}

	if v, ok := d.GetOk("fapc_compatibility"); ok {
		if setArgNil {
			obj["fapc-compatibility"] = nil
		} else {

			t, err := expandWirelessControllerSettingFapcCompatibility(d, v, "fapc_compatibility", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fapc-compatibility"] = t
			}
		}
	}

	if v, ok := d.GetOk("wfa_compatibility"); ok {
		if setArgNil {
			obj["wfa-compatibility"] = nil
		} else {

			t, err := expandWirelessControllerSettingWfaCompatibility(d, v, "wfa_compatibility", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wfa-compatibility"] = t
			}
		}
	}

	if v, ok := d.GetOk("phishing_ssid_detect"); ok {
		if setArgNil {
			obj["phishing-ssid-detect"] = nil
		} else {

			t, err := expandWirelessControllerSettingPhishingSsidDetect(d, v, "phishing_ssid_detect", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["phishing-ssid-detect"] = t
			}
		}
	}

	if v, ok := d.GetOk("fake_ssid_action"); ok {
		if setArgNil {
			obj["fake-ssid-action"] = nil
		} else {

			t, err := expandWirelessControllerSettingFakeSsidAction(d, v, "fake_ssid_action", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fake-ssid-action"] = t
			}
		}
	}

	if v, ok := d.GetOk("offending_ssid"); ok || d.HasChange("offending_ssid") {
		if setArgNil {
			obj["offending-ssid"] = make([]struct{}, 0)
		} else {

			t, err := expandWirelessControllerSettingOffendingSsid(d, v, "offending_ssid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["offending-ssid"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("device_weight"); ok {
		if setArgNil {
			obj["device-weight"] = nil
		} else {

			t, err := expandWirelessControllerSettingDeviceWeight(d, v, "device_weight", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["device-weight"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("device_holdoff"); ok {
		if setArgNil {
			obj["device-holdoff"] = nil
		} else {

			t, err := expandWirelessControllerSettingDeviceHoldoff(d, v, "device_holdoff", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["device-holdoff"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("device_idle"); ok {
		if setArgNil {
			obj["device-idle"] = nil
		} else {

			t, err := expandWirelessControllerSettingDeviceIdle(d, v, "device_idle", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["device-idle"] = t
			}
		}
	}

	if v, ok := d.GetOk("firmware_provision_on_authorization"); ok {
		if setArgNil {
			obj["firmware-provision-on-authorization"] = nil
		} else {

			t, err := expandWirelessControllerSettingFirmwareProvisionOnAuthorization(d, v, "firmware_provision_on_authorization", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["firmware-provision-on-authorization"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("darrp_optimize"); ok {
		if setArgNil {
			obj["darrp-optimize"] = nil
		} else {

			t, err := expandWirelessControllerSettingDarrpOptimize(d, v, "darrp_optimize", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["darrp-optimize"] = t
			}
		}
	}

	if v, ok := d.GetOk("darrp_optimize_schedules"); ok || d.HasChange("darrp_optimize_schedules") {
		if setArgNil {
			obj["darrp-optimize-schedules"] = make([]struct{}, 0)
		} else {

			t, err := expandWirelessControllerSettingDarrpOptimizeSchedules(d, v, "darrp_optimize_schedules", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["darrp-optimize-schedules"] = t
			}
		}
	}

	return &obj, nil
}
