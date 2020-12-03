// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure shaping profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallShapingProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShapingProfileCreate,
		Read:   resourceFirewallShapingProfileRead,
		Update: resourceFirewallShapingProfileUpdate,
		Delete: resourceFirewallShapingProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"profile_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"default_class_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 31),
				Required:     true,
			},
			"shaping_entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"class_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 31),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"guaranteed_bandwidth_percentage": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"maximum_bandwidth_percentage": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceFirewallShapingProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallShapingProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallShapingProfile resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallShapingProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallShapingProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallShapingProfile")
	}

	return resourceFirewallShapingProfileRead(d, m)
}

func resourceFirewallShapingProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallShapingProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallShapingProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallShapingProfile")
	}

	return resourceFirewallShapingProfileRead(d, m)
}

func resourceFirewallShapingProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallShapingProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallShapingProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShapingProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallShapingProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShapingProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallShapingProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShapingProfile resource from API: %v", err)
	}
	return nil
}

func flattenFirewallShapingProfileProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileDefaultClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallShapingProfileShapingEntriesId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {
			tmp["class_id"] = flattenFirewallShapingProfileShapingEntriesClassId(i["class-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = flattenFirewallShapingProfileShapingEntriesPriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth_percentage"
		if _, ok := i["guaranteed-bandwidth-percentage"]; ok {
			tmp["guaranteed_bandwidth_percentage"] = flattenFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(i["guaranteed-bandwidth-percentage"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth_percentage"
		if _, ok := i["maximum-bandwidth-percentage"]; ok {
			tmp["maximum_bandwidth_percentage"] = flattenFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(i["maximum-bandwidth-percentage"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallShapingProfileShapingEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallShapingProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("profile_name", flattenFirewallShapingProfileProfileName(o["profile-name"], d, "profile_name")); err != nil {
		if !fortiAPIPatch(o["profile-name"]) {
			return fmt.Errorf("Error reading profile_name: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallShapingProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("default_class_id", flattenFirewallShapingProfileDefaultClassId(o["default-class-id"], d, "default_class_id")); err != nil {
		if !fortiAPIPatch(o["default-class-id"]) {
			return fmt.Errorf("Error reading default_class_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("shaping_entries", flattenFirewallShapingProfileShapingEntries(o["shaping-entries"], d, "shaping_entries")); err != nil {
			if !fortiAPIPatch(o["shaping-entries"]) {
				return fmt.Errorf("Error reading shaping_entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("shaping_entries"); ok {
			if err = d.Set("shaping_entries", flattenFirewallShapingProfileShapingEntries(o["shaping-entries"], d, "shaping_entries")); err != nil {
				if !fortiAPIPatch(o["shaping-entries"]) {
					return fmt.Errorf("Error reading shaping_entries: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallShapingProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallShapingProfileProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileDefaultClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallShapingProfileShapingEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["class-id"], _ = expandFirewallShapingProfileShapingEntriesClassId(d, i["class_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandFirewallShapingProfileShapingEntriesPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth_percentage"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["guaranteed-bandwidth-percentage"], _ = expandFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(d, i["guaranteed_bandwidth_percentage"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth_percentage"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-bandwidth-percentage"], _ = expandFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(d, i["maximum_bandwidth_percentage"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallShapingProfileShapingEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShapingProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("profile_name"); ok {
		t, err := expandFirewallShapingProfileProfileName(d, v, "profile_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallShapingProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("default_class_id"); ok {
		t, err := expandFirewallShapingProfileDefaultClassId(d, v, "default_class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-class-id"] = t
		}
	}

	if v, ok := d.GetOk("shaping_entries"); ok {
		t, err := expandFirewallShapingProfileShapingEntries(d, v, "shaping_entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shaping-entries"] = t
		}
	}

	return &obj, nil
}
