// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure shaping profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"profile_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_class_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
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
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
						"limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 10000),
							Optional:     true,
							Computed:     true,
						},
						"burst_in_msec": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2000),
							Optional:     true,
							Computed:     true,
						},
						"cburst_in_msec": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2000),
							Optional:     true,
							Computed:     true,
						},
						"red_probability": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 20),
							Optional:     true,
							Computed:     true,
						},
						"min": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 3000),
							Optional:     true,
							Computed:     true,
						},
						"max": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 3000),
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

func resourceFirewallShapingProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallShapingProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallShapingProfile resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallShapingProfile(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallShapingProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShapingProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallShapingProfile(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallShapingProfile(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallShapingProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShapingProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallShapingProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShapingProfile resource from API: %v", err)
	}
	return nil
}

func flattenFirewallShapingProfileProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileDefaultClassId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallShapingProfileShapingEntriesId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {

			tmp["class_id"] = flattenFirewallShapingProfileShapingEntriesClassId(i["class-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenFirewallShapingProfileShapingEntriesPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth_percentage"
		if _, ok := i["guaranteed-bandwidth-percentage"]; ok {

			tmp["guaranteed_bandwidth_percentage"] = flattenFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(i["guaranteed-bandwidth-percentage"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth_percentage"
		if _, ok := i["maximum-bandwidth-percentage"]; ok {

			tmp["maximum_bandwidth_percentage"] = flattenFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(i["maximum-bandwidth-percentage"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit"
		if _, ok := i["limit"]; ok {

			tmp["limit"] = flattenFirewallShapingProfileShapingEntriesLimit(i["limit"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "burst_in_msec"
		if _, ok := i["burst-in-msec"]; ok {

			tmp["burst_in_msec"] = flattenFirewallShapingProfileShapingEntriesBurstInMsec(i["burst-in-msec"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cburst_in_msec"
		if _, ok := i["cburst-in-msec"]; ok {

			tmp["cburst_in_msec"] = flattenFirewallShapingProfileShapingEntriesCburstInMsec(i["cburst-in-msec"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "red_probability"
		if _, ok := i["red-probability"]; ok {

			tmp["red_probability"] = flattenFirewallShapingProfileShapingEntriesRedProbability(i["red-probability"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min"
		if _, ok := i["min"]; ok {

			tmp["min"] = flattenFirewallShapingProfileShapingEntriesMin(i["min"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max"
		if _, ok := i["max"]; ok {

			tmp["max"] = flattenFirewallShapingProfileShapingEntriesMax(i["max"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallShapingProfileShapingEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesClassId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesBurstInMsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesCburstInMsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesRedProbability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallShapingProfileShapingEntriesMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallShapingProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("profile_name", flattenFirewallShapingProfileProfileName(o["profile-name"], d, "profile_name", sv)); err != nil {
		if !fortiAPIPatch(o["profile-name"]) {
			return fmt.Errorf("Error reading profile_name: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallShapingProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallShapingProfileType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("default_class_id", flattenFirewallShapingProfileDefaultClassId(o["default-class-id"], d, "default_class_id", sv)); err != nil {
		if !fortiAPIPatch(o["default-class-id"]) {
			return fmt.Errorf("Error reading default_class_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("shaping_entries", flattenFirewallShapingProfileShapingEntries(o["shaping-entries"], d, "shaping_entries", sv)); err != nil {
			if !fortiAPIPatch(o["shaping-entries"]) {
				return fmt.Errorf("Error reading shaping_entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("shaping_entries"); ok {
			if err = d.Set("shaping_entries", flattenFirewallShapingProfileShapingEntries(o["shaping-entries"], d, "shaping_entries", sv)); err != nil {
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
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallShapingProfileProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileDefaultClassId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallShapingProfileShapingEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["class-id"], _ = expandFirewallShapingProfileShapingEntriesClassId(d, i["class_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority"], _ = expandFirewallShapingProfileShapingEntriesPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "guaranteed_bandwidth_percentage"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["guaranteed-bandwidth-percentage"], _ = expandFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(d, i["guaranteed_bandwidth_percentage"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_bandwidth_percentage"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-bandwidth-percentage"], _ = expandFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(d, i["maximum_bandwidth_percentage"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["limit"], _ = expandFirewallShapingProfileShapingEntriesLimit(d, i["limit"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "burst_in_msec"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["burst-in-msec"], _ = expandFirewallShapingProfileShapingEntriesBurstInMsec(d, i["burst_in_msec"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cburst_in_msec"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cburst-in-msec"], _ = expandFirewallShapingProfileShapingEntriesCburstInMsec(d, i["cburst_in_msec"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "red_probability"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["red-probability"], _ = expandFirewallShapingProfileShapingEntriesRedProbability(d, i["red_probability"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["min"], _ = expandFirewallShapingProfileShapingEntriesMin(d, i["min"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["max"], _ = expandFirewallShapingProfileShapingEntriesMax(d, i["max"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallShapingProfileShapingEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesClassId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesGuaranteedBandwidthPercentage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMaximumBandwidthPercentage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesBurstInMsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesCburstInMsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesRedProbability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallShapingProfileShapingEntriesMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShapingProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("profile_name"); ok {

		t, err := expandFirewallShapingProfileProfileName(d, v, "profile_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandFirewallShapingProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandFirewallShapingProfileType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOkExists("default_class_id"); ok {

		t, err := expandFirewallShapingProfileDefaultClassId(d, v, "default_class_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-class-id"] = t
		}
	}

	if v, ok := d.GetOk("shaping_entries"); ok || d.HasChange("shaping_entries") {

		t, err := expandFirewallShapingProfileShapingEntries(d, v, "shaping_entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shaping-entries"] = t
		}
	}

	return &obj, nil
}
