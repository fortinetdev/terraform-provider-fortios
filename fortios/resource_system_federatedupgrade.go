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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceSystemFederatedUpgradeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFederatedUpgrade(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFederatedUpgrade resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFederatedUpgrade(obj, mkey)
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

	err := c.DeleteSystemFederatedUpgrade(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFederatedUpgrade resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFederatedUpgradeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFederatedUpgrade(mkey)
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

func flattenSystemFederatedUpgradeNodeList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {

			tmp["serial"] = flattenSystemFederatedUpgradeNodeListSerial(i["serial"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timing"
		if _, ok := i["timing"]; ok {

			tmp["timing"] = flattenSystemFederatedUpgradeNodeListTiming(i["timing"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time"
		if _, ok := i["time"]; ok {

			tmp["time"] = flattenSystemFederatedUpgradeNodeListTime(i["time"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "setup_time"
		if _, ok := i["setup-time"]; ok {

			tmp["setup_time"] = flattenSystemFederatedUpgradeNodeListSetupTime(i["setup-time"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upgrade_path"
		if _, ok := i["upgrade-path"]; ok {

			tmp["upgrade_path"] = flattenSystemFederatedUpgradeNodeListUpgradePath(i["upgrade-path"], d, pre_append, sv)
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

func flattenSystemFederatedUpgradeNodeListTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListSetupTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFederatedUpgradeNodeListUpgradePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFederatedUpgrade(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemFederatedUpgradeStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if isImportTable() {
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

func expandSystemFederatedUpgradeNodeList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["serial"], _ = expandSystemFederatedUpgradeNodeListSerial(d, i["serial"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timing"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["timing"], _ = expandSystemFederatedUpgradeNodeListTiming(d, i["timing"], pre_append, sv)
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

func expandSystemFederatedUpgradeNodeListTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListSetupTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFederatedUpgradeNodeListUpgradePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFederatedUpgrade(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSystemFederatedUpgradeStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("node_list"); ok {

		t, err := expandSystemFederatedUpgradeNodeList(d, v, "node_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["node-list"] = t
		}
	}

	return &obj, nil
}
