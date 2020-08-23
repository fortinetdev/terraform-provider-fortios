// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure quarantine support.

package fortios

import (
	"fmt"
	"log"
	"strconv"

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
			"quarantine": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_policy": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"targets": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional: true,
							Computed: true,
						},
						"macs": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mac": &schema.Schema{
										Type: schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"entry_id": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Type: schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional: true,
										Computed: true,
									},
									"parent": &schema.Schema{
										Type: schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}


func resourceUserQuarantineUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserQuarantine(d)
	if err != nil {
		return fmt.Errorf("Error updating UserQuarantine resource while getting object: %v", err)
	}

	o, err := c.UpdateUserQuarantine(obj, mkey)
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

	err := c.DeleteUserQuarantine(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserQuarantineRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserQuarantine(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserQuarantine(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserQuarantine resource from API: %v", err)
	}
	return nil
}


func flattenUserQuarantineQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTrafficPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargets(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry"
		if _, ok := i["entry"]; ok {
			tmp["entry"] = flattenUserQuarantineTargetsEntry(i["entry"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = flattenUserQuarantineTargetsDescription(i["description"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macs"
		if _, ok := i["macs"]; ok {
			tmp["macs"] = flattenUserQuarantineTargetsMacs(i["macs"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenUserQuarantineTargetsEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			tmp["mac"] = flattenUserQuarantineTargetsMacsMac(i["mac"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := i["entry-id"]; ok {
			tmp["entry_id"] = flattenUserQuarantineTargetsMacsEntryId(i["entry-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = flattenUserQuarantineTargetsMacsDescription(i["description"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parent"
		if _, ok := i["parent"]; ok {
			tmp["parent"] = flattenUserQuarantineTargetsMacsParent(i["parent"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenUserQuarantineTargetsMacsMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserQuarantineTargetsMacsParent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectUserQuarantine(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("quarantine", flattenUserQuarantineQuarantine(o["quarantine"], d, "quarantine")); err != nil {
		if !fortiAPIPatch(o["quarantine"]) {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("traffic_policy", flattenUserQuarantineTrafficPolicy(o["traffic-policy"], d, "traffic_policy")); err != nil {
		if !fortiAPIPatch(o["traffic-policy"]) {
			return fmt.Errorf("Error reading traffic_policy: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("targets", flattenUserQuarantineTargets(o["targets"], d, "targets")); err != nil {
            if !fortiAPIPatch(o["targets"]) {
                return fmt.Errorf("Error reading targets: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("targets"); ok {
            if err = d.Set("targets", flattenUserQuarantineTargets(o["targets"], d, "targets")); err != nil {
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
	log.Printf("ER List: %v", e)
}


func expandUserQuarantineQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTrafficPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["entry"], _ = expandUserQuarantineTargetsEntry(d, i["entry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandUserQuarantineTargetsDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macs"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["macs"], _ = expandUserQuarantineTargetsMacs(d, i["macs"], pre_append)
		} else {
			tmp["macs"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserQuarantineTargetsEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandUserQuarantineTargetsMacsMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["entry-id"], _ = expandUserQuarantineTargetsMacsEntryId(d, i["entry_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandUserQuarantineTargetsMacsDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["parent"], _ = expandUserQuarantineTargetsMacsParent(d, i["parent"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserQuarantineTargetsMacsMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserQuarantineTargetsMacsParent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectUserQuarantine(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("quarantine"); ok {
		t, err := expandUserQuarantineQuarantine(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("traffic_policy"); ok {
		t, err := expandUserQuarantineTrafficPolicy(d, v, "traffic_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-policy"] = t
		}
	}

	if v, ok := d.GetOk("targets"); ok {
		t, err := expandUserQuarantineTargets(d, v, "targets")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["targets"] = t
		}
	}


	return &obj, nil
}

