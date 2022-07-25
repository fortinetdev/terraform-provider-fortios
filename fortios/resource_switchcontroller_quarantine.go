// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch quarantine support.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerQuarantine() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQuarantineUpdate,
		Read:   resourceSwitchControllerQuarantineRead,
		Update: resourceSwitchControllerQuarantineUpdate,
		Delete: resourceSwitchControllerQuarantineDelete,

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
			"targets": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tags": &schema.Schema{
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

func resourceSwitchControllerQuarantineUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerQuarantine(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQuarantine resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerQuarantine(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQuarantine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerQuarantine")
	}

	return resourceSwitchControllerQuarantineRead(d, m)
}

func resourceSwitchControllerQuarantineDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerQuarantine(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQuarantine resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerQuarantine(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQuarantineRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerQuarantine(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQuarantine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQuarantine(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQuarantine resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQuarantineQuarantine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQuarantineTargets(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["mac"] = flattenSwitchControllerQuarantineTargetsMac(i["mac"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := i["entry-id"]; ok {

			tmp["entry_id"] = flattenSwitchControllerQuarantineTargetsEntryId(i["entry-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchControllerQuarantineTargetsDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {

			tmp["tag"] = flattenSwitchControllerQuarantineTargetsTag(i["tag"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "mac", d)
	return result
}

func flattenSwitchControllerQuarantineTargetsMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQuarantineTargetsEntryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQuarantineTargetsDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQuarantineTargetsTag(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {

			tmp["tags"] = flattenSwitchControllerQuarantineTargetsTagTags(i["tags"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerQuarantineTargetsTagTags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerQuarantine(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("quarantine", flattenSwitchControllerQuarantineQuarantine(o["quarantine"], d, "quarantine", sv)); err != nil {
		if !fortiAPIPatch(o["quarantine"]) {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("targets", flattenSwitchControllerQuarantineTargets(o["targets"], d, "targets", sv)); err != nil {
			if !fortiAPIPatch(o["targets"]) {
				return fmt.Errorf("Error reading targets: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("targets"); ok {
			if err = d.Set("targets", flattenSwitchControllerQuarantineTargets(o["targets"], d, "targets", sv)); err != nil {
				if !fortiAPIPatch(o["targets"]) {
					return fmt.Errorf("Error reading targets: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerQuarantineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerQuarantineQuarantine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQuarantineTargets(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mac"], _ = expandSwitchControllerQuarantineTargetsMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["entry-id"], _ = expandSwitchControllerQuarantineTargetsEntryId(d, i["entry_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchControllerQuarantineTargetsDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["tag"], _ = expandSwitchControllerQuarantineTargetsTag(d, i["tag"], pre_append, sv)
		} else {
			tmp["tag"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerQuarantineTargetsMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQuarantineTargetsEntryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQuarantineTargetsDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQuarantineTargetsTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["tags"], _ = expandSwitchControllerQuarantineTargetsTagTags(d, i["tags"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerQuarantineTargetsTagTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerQuarantine(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("quarantine"); ok {
		if setArgNil {
			obj["quarantine"] = nil
		} else {

			t, err := expandSwitchControllerQuarantineQuarantine(d, v, "quarantine", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["quarantine"] = t
			}
		}
	}

	if v, ok := d.GetOk("targets"); ok || d.HasChange("targets") {
		if setArgNil {
			obj["targets"] = make([]struct{}, 0)
		} else {

			t, err := expandSwitchControllerQuarantineTargets(d, v, "targets", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["targets"] = t
			}
		}
	}

	return &obj, nil
}
