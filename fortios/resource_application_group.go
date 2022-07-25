// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure firewall application groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceApplicationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationGroupCreate,
		Read:   resourceApplicationGroupRead,
		Update: resourceApplicationGroupUpdate,
		Delete: resourceApplicationGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"category": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"risk": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"protocols": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"technology": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"behavior": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"popularity": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceApplicationGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectApplicationGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationGroup resource while getting object: %v", err)
	}

	o, err := c.CreateApplicationGroup(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ApplicationGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ApplicationGroup")
	}

	return resourceApplicationGroupRead(d, m)
}

func resourceApplicationGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectApplicationGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateApplicationGroup(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ApplicationGroup")
	}

	return resourceApplicationGroupRead(d, m)
}

func resourceApplicationGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteApplicationGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadApplicationGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationGroup resource from API: %v", err)
	}
	return nil
}

func flattenApplicationGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenApplicationGroupApplicationId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationGroupApplicationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenApplicationGroupCategoryId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationGroupCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupRisk(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {

			tmp["level"] = flattenApplicationGroupRiskLevel(i["level"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "level", d)
	return result
}

func flattenApplicationGroupRiskLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupProtocols(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupTechnology(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupBehavior(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationGroupPopularity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectApplicationGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenApplicationGroupName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenApplicationGroupComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("type", flattenApplicationGroupType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("application", flattenApplicationGroupApplication(o["application"], d, "application", sv)); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenApplicationGroupApplication(o["application"], d, "application", sv)); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("category", flattenApplicationGroupCategory(o["category"], d, "category", sv)); err != nil {
			if !fortiAPIPatch(o["category"]) {
				return fmt.Errorf("Error reading category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("category"); ok {
			if err = d.Set("category", flattenApplicationGroupCategory(o["category"], d, "category", sv)); err != nil {
				if !fortiAPIPatch(o["category"]) {
					return fmt.Errorf("Error reading category: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("risk", flattenApplicationGroupRisk(o["risk"], d, "risk", sv)); err != nil {
			if !fortiAPIPatch(o["risk"]) {
				return fmt.Errorf("Error reading risk: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("risk"); ok {
			if err = d.Set("risk", flattenApplicationGroupRisk(o["risk"], d, "risk", sv)); err != nil {
				if !fortiAPIPatch(o["risk"]) {
					return fmt.Errorf("Error reading risk: %v", err)
				}
			}
		}
	}

	if err = d.Set("protocols", flattenApplicationGroupProtocols(o["protocols"], d, "protocols", sv)); err != nil {
		if !fortiAPIPatch(o["protocols"]) {
			return fmt.Errorf("Error reading protocols: %v", err)
		}
	}

	if err = d.Set("vendor", flattenApplicationGroupVendor(o["vendor"], d, "vendor", sv)); err != nil {
		if !fortiAPIPatch(o["vendor"]) {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	if err = d.Set("technology", flattenApplicationGroupTechnology(o["technology"], d, "technology", sv)); err != nil {
		if !fortiAPIPatch(o["technology"]) {
			return fmt.Errorf("Error reading technology: %v", err)
		}
	}

	if err = d.Set("behavior", flattenApplicationGroupBehavior(o["behavior"], d, "behavior", sv)); err != nil {
		if !fortiAPIPatch(o["behavior"]) {
			return fmt.Errorf("Error reading behavior: %v", err)
		}
	}

	if err = d.Set("popularity", flattenApplicationGroupPopularity(o["popularity"], d, "popularity", sv)); err != nil {
		if !fortiAPIPatch(o["popularity"]) {
			return fmt.Errorf("Error reading popularity: %v", err)
		}
	}

	return nil
}

func flattenApplicationGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandApplicationGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandApplicationGroupApplicationId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationGroupApplicationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandApplicationGroupCategoryId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationGroupCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupRisk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["level"], _ = expandApplicationGroupRiskLevel(d, i["level"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationGroupRiskLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupProtocols(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupTechnology(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupBehavior(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationGroupPopularity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandApplicationGroupName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandApplicationGroupComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandApplicationGroupType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {

		t, err := expandApplicationGroupApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {

		t, err := expandApplicationGroupCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("risk"); ok || d.HasChange("risk") {

		t, err := expandApplicationGroupRisk(d, v, "risk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risk"] = t
		}
	}

	if v, ok := d.GetOk("protocols"); ok {

		t, err := expandApplicationGroupProtocols(d, v, "protocols", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocols"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok {

		t, err := expandApplicationGroupVendor(d, v, "vendor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	if v, ok := d.GetOk("technology"); ok {

		t, err := expandApplicationGroupTechnology(d, v, "technology", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["technology"] = t
		}
	}

	if v, ok := d.GetOk("behavior"); ok {

		t, err := expandApplicationGroupBehavior(d, v, "behavior", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["behavior"] = t
		}
	}

	if v, ok := d.GetOk("popularity"); ok {

		t, err := expandApplicationGroupPopularity(d, v, "popularity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["popularity"] = t
		}
	}

	return &obj, nil
}
