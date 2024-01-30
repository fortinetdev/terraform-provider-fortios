// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Coordinate federated upgrades within the Security Fabric.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFederatedUpgrade() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFederatedUpgradeUpdate,
		Read:   resourceSystemFederatedUpgradeRead,
		Update: resourceSystemFederatedUpgradeUpdate,
		Delete: resourceSystemFederatedUpgradeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failure_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failure_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"upgrade_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"next_path_index": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10),
				Optional:     true,
				Computed:     true,
			},
			"ha_reboot_controller": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"known_ha_members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"serial": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"node_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"serial": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"timing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maximum_minutes": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 10080),
							Optional:     true,
							Computed:     true,
						},
						"time": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"setup_time": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"upgrade_path": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"coordinating_fortigate": &schema.Schema{
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

func resourceSystemFederatedUpgradeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFederatedUpgrade(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgrade resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFederatedUpgrade(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgrade resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFederatedUpgrade")
	}

	return resourceSystemFederatedUpgradeRead(d, m)
}

func resourceSystemFederatedUpgradeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFederatedUpgrade(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgrade resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFederatedUpgrade(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFederatedUpgrade resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFederatedUpgradeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemFederatedUpgrade(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemFederatedUpgrade resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFederatedUpgrade(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFederatedUpgrade resource from API: %v", err)
	}
	return nil
}

func flattenSystemFederatedUpgradeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeFailureReason(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeFailureDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeUpgradeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNextPathIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeHaRebootController(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeKnownHaMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if cur_v, ok := i["serial"]; ok {
			tmp["serial"] = flattenSystemFederatedUpgradeKnownHaMembersSerial(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "serial", d)
	return result
}

func flattenSystemFederatedUpgradeKnownHaMembersSerial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if cur_v, ok := i["serial"]; ok {
			tmp["serial"] = flattenSystemFederatedUpgradeNodeListSerial(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timing"
		if cur_v, ok := i["timing"]; ok {
			tmp["timing"] = flattenSystemFederatedUpgradeNodeListTiming(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_minutes"
		if cur_v, ok := i["maximum-minutes"]; ok {
			tmp["maximum_minutes"] = flattenSystemFederatedUpgradeNodeListMaximumMinutes(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time"
		if cur_v, ok := i["time"]; ok {
			tmp["time"] = flattenSystemFederatedUpgradeNodeListTime(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "setup_time"
		if cur_v, ok := i["setup-time"]; ok {
			tmp["setup_time"] = flattenSystemFederatedUpgradeNodeListSetupTime(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upgrade_path"
		if cur_v, ok := i["upgrade-path"]; ok {
			tmp["upgrade_path"] = flattenSystemFederatedUpgradeNodeListUpgradePath(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if cur_v, ok := i["device-type"]; ok {
			tmp["device_type"] = flattenSystemFederatedUpgradeNodeListDeviceType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "coordinating_fortigate"
		if cur_v, ok := i["coordinating-fortigate"]; ok {
			tmp["coordinating_fortigate"] = flattenSystemFederatedUpgradeNodeListCoordinatingFortigate(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "serial", d)
	return result
}

func flattenSystemFederatedUpgradeNodeListSerial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListTiming(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListMaximumMinutes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListSetupTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListUpgradePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListDeviceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListCoordinatingFortigate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFederatedUpgrade(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("status", flattenSystemFederatedUpgradeStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("failure_reason", flattenSystemFederatedUpgradeFailureReason(o["failure-reason"], d, "failure_reason", sv)); err != nil {
		if !fortiAPIPatch(o["failure-reason"]) {
			return fmt.Errorf("Error reading failure_reason: %v", err)
		}
	}

	if err = d.Set("failure_device", flattenSystemFederatedUpgradeFailureDevice(o["failure-device"], d, "failure_device", sv)); err != nil {
		if !fortiAPIPatch(o["failure-device"]) {
			return fmt.Errorf("Error reading failure_device: %v", err)
		}
	}

	if err = d.Set("upgrade_id", flattenSystemFederatedUpgradeUpgradeId(o["upgrade-id"], d, "upgrade_id", sv)); err != nil {
		if !fortiAPIPatch(o["upgrade-id"]) {
			return fmt.Errorf("Error reading upgrade_id: %v", err)
		}
	}

	if err = d.Set("next_path_index", flattenSystemFederatedUpgradeNextPathIndex(o["next-path-index"], d, "next_path_index", sv)); err != nil {
		if !fortiAPIPatch(o["next-path-index"]) {
			return fmt.Errorf("Error reading next_path_index: %v", err)
		}
	}

	if err = d.Set("ha_reboot_controller", flattenSystemFederatedUpgradeHaRebootController(o["ha-reboot-controller"], d, "ha_reboot_controller", sv)); err != nil {
		if !fortiAPIPatch(o["ha-reboot-controller"]) {
			return fmt.Errorf("Error reading ha_reboot_controller: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("known_ha_members", flattenSystemFederatedUpgradeKnownHaMembers(o["known-ha-members"], d, "known_ha_members", sv)); err != nil {
			if !fortiAPIPatch(o["known-ha-members"]) {
				return fmt.Errorf("Error reading known_ha_members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("known_ha_members"); ok {
			if err = d.Set("known_ha_members", flattenSystemFederatedUpgradeKnownHaMembers(o["known-ha-members"], d, "known_ha_members", sv)); err != nil {
				if !fortiAPIPatch(o["known-ha-members"]) {
					return fmt.Errorf("Error reading known_ha_members: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("node_list", flattenSystemFederatedUpgradeNodeList(o["node-list"], d, "node_list", sv)); err != nil {
			if !fortiAPIPatch(o["node-list"]) {
				return fmt.Errorf("Error reading node_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("node_list"); ok {
			if err = d.Set("node_list", flattenSystemFederatedUpgradeNodeList(o["node-list"], d, "node_list", sv)); err != nil {
				if !fortiAPIPatch(o["node-list"]) {
					return fmt.Errorf("Error reading node_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemFederatedUpgradeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemFederatedUpgradeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeFailureReason(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeFailureDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeUpgradeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNextPathIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeHaRebootController(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeKnownHaMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["serial"], _ = expandSystemFederatedUpgradeKnownHaMembersSerial(d, i["serial"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemFederatedUpgradeKnownHaMembersSerial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["serial"], _ = expandSystemFederatedUpgradeNodeListSerial(d, i["serial"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timing"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["timing"], _ = expandSystemFederatedUpgradeNodeListTiming(d, i["timing"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_minutes"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-minutes"], _ = expandSystemFederatedUpgradeNodeListMaximumMinutes(d, i["maximum_minutes"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["time"], _ = expandSystemFederatedUpgradeNodeListTime(d, i["time"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "setup_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["setup-time"], _ = expandSystemFederatedUpgradeNodeListSetupTime(d, i["setup_time"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upgrade_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["upgrade-path"], _ = expandSystemFederatedUpgradeNodeListUpgradePath(d, i["upgrade_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device-type"], _ = expandSystemFederatedUpgradeNodeListDeviceType(d, i["device_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "coordinating_fortigate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["coordinating-fortigate"], _ = expandSystemFederatedUpgradeNodeListCoordinatingFortigate(d, i["coordinating_fortigate"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemFederatedUpgradeNodeListSerial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListTiming(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListMaximumMinutes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListSetupTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListUpgradePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListDeviceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListCoordinatingFortigate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFederatedUpgrade(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemFederatedUpgradeStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("failure_reason"); ok {
		if setArgNil {
			obj["failure-reason"] = nil
		} else {
			t, err := expandSystemFederatedUpgradeFailureReason(d, v, "failure_reason", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["failure-reason"] = t
			}
		}
	}

	if v, ok := d.GetOk("failure_device"); ok {
		if setArgNil {
			obj["failure-device"] = nil
		} else {
			t, err := expandSystemFederatedUpgradeFailureDevice(d, v, "failure_device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["failure-device"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("upgrade_id"); ok {
		if setArgNil {
			obj["upgrade-id"] = nil
		} else {
			t, err := expandSystemFederatedUpgradeUpgradeId(d, v, "upgrade_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upgrade-id"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("next_path_index"); ok {
		if setArgNil {
			obj["next-path-index"] = nil
		} else {
			t, err := expandSystemFederatedUpgradeNextPathIndex(d, v, "next_path_index", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["next-path-index"] = t
			}
		}
	}

	if v, ok := d.GetOk("ha_reboot_controller"); ok {
		if setArgNil {
			obj["ha-reboot-controller"] = nil
		} else {
			t, err := expandSystemFederatedUpgradeHaRebootController(d, v, "ha_reboot_controller", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ha-reboot-controller"] = t
			}
		}
	}

	if v, ok := d.GetOk("known_ha_members"); ok || d.HasChange("known_ha_members") {
		if setArgNil {
			obj["known-ha-members"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemFederatedUpgradeKnownHaMembers(d, v, "known_ha_members", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["known-ha-members"] = t
			}
		}
	}

	if v, ok := d.GetOk("node_list"); ok || d.HasChange("node_list") {
		if setArgNil {
			obj["node-list"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemFederatedUpgradeNodeList(d, v, "node_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["node-list"] = t
			}
		}
	}

	return &obj, nil
}
