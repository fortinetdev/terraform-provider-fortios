// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure quarantine support.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserQuarantine() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserQuarantineUpdate,
		Read:   resourceUserQuarantineRead,
		Update: resourceUserQuarantineUpdate,
		Delete: resourceUserQuarantineDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"firewall_groups": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"targets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"macs": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mac": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"entry_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
									"drop": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"parent": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
								},
							},
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

func resourceUserQuarantineUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserQuarantine(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserQuarantine resource while getting object: %v", err)
	}

	o, err := c.UpdateUserQuarantine(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserQuarantine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserQuarantine")
	}

	return resourceUserQuarantineRead(d, m)
}

func resourceUserQuarantineDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserQuarantine(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating UserQuarantine resource while getting object: %v", err)
	}

	_, err = c.UpdateUserQuarantine(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing UserQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserQuarantineRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserQuarantine(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserQuarantine(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantine resource from API: %v", err)
	}
	return nil
}

func flattenUserQuarantineQuarantine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserQuarantineTrafficPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserQuarantineFirewallGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserQuarantineTargets(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry"
		if _, ok := i["entry"]; ok {

			tmp["entry"] = flattenUserQuarantineTargetsEntry(i["entry"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenUserQuarantineTargetsDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macs"
		if _, ok := i["macs"]; ok {

			tmp["macs"] = flattenUserQuarantineTargetsMacs(i["macs"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "entry", d)
	return result
}

func flattenUserQuarantineTargetsEntry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserQuarantineTargetsDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacs(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {

			tmp["mac"] = flattenUserQuarantineTargetsMacsMac(i["mac"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := i["entry-id"]; ok {

			tmp["entry_id"] = flattenUserQuarantineTargetsMacsEntryId(i["entry-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenUserQuarantineTargetsMacsDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop"
		if _, ok := i["drop"]; ok {

			tmp["drop"] = flattenUserQuarantineTargetsMacsDrop(i["drop"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parent"
		if _, ok := i["parent"]; ok {

			tmp["parent"] = flattenUserQuarantineTargetsMacsParent(i["parent"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenUserQuarantineTargetsMacsMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsEntryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsDrop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsParent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserQuarantine(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("quarantine", flattenUserQuarantineQuarantine(o["quarantine"], d, "quarantine", sv)); err != nil {
		if !fortiAPIPatch(o["quarantine"]) {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("traffic_policy", flattenUserQuarantineTrafficPolicy(o["traffic-policy"], d, "traffic_policy", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-policy"]) {
			return fmt.Errorf("Error reading traffic_policy: %v", err)
		}
	}

	if err = d.Set("firewall_groups", flattenUserQuarantineFirewallGroups(o["firewall-groups"], d, "firewall_groups", sv)); err != nil {
		if !fortiAPIPatch(o["firewall-groups"]) {
			return fmt.Errorf("Error reading firewall_groups: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("targets", flattenUserQuarantineTargets(o["targets"], d, "targets", sv)); err != nil {
			if !fortiAPIPatch(o["targets"]) {
				return fmt.Errorf("Error reading targets: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("targets"); ok {
			if err = d.Set("targets", flattenUserQuarantineTargets(o["targets"], d, "targets", sv)); err != nil {
				if !fortiAPIPatch(o["targets"]) {
					return fmt.Errorf("Error reading targets: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenUserQuarantineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserQuarantineQuarantine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTrafficPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineFirewallGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargets(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["entry"], _ = expandUserQuarantineTargetsEntry(d, i["entry"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandUserQuarantineTargetsDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macs"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["macs"], _ = expandUserQuarantineTargetsMacs(d, i["macs"], pre_append, sv)
		} else {
			tmp["macs"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserQuarantineTargetsEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mac"], _ = expandUserQuarantineTargetsMacsMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["entry-id"], _ = expandUserQuarantineTargetsMacsEntryId(d, i["entry_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandUserQuarantineTargetsMacsDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["drop"], _ = expandUserQuarantineTargetsMacsDrop(d, i["drop"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parent"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["parent"], _ = expandUserQuarantineTargetsMacsParent(d, i["parent"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserQuarantineTargetsMacsMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsEntryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsDrop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsParent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserQuarantine(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("quarantine"); ok {
		if setArgNil {
			obj["quarantine"] = nil
		} else {

			t, err := expandUserQuarantineQuarantine(d, v, "quarantine", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["quarantine"] = t
			}
		}
	}

	if v, ok := d.GetOk("traffic_policy"); ok {
		if setArgNil {
			obj["traffic-policy"] = nil
		} else {

			t, err := expandUserQuarantineTrafficPolicy(d, v, "traffic_policy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["traffic-policy"] = t
			}
		}
	}

	if v, ok := d.GetOk("firewall_groups"); ok {
		if setArgNil {
			obj["firewall-groups"] = nil
		} else {

			t, err := expandUserQuarantineFirewallGroups(d, v, "firewall_groups", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["firewall-groups"] = t
			}
		}
	}

	if v, ok := d.GetOk("targets"); ok {
		if setArgNil {
			obj["targets"] = make([]struct{}, 0)
		} else {

			t, err := expandUserQuarantineTargets(d, v, "targets", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["targets"] = t
			}
		}
	}

	return &obj, nil
}
