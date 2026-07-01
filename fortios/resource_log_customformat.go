// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure custom log format.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogCustomFormat() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogCustomFormatCreate,
		Read:   resourceLogCustomFormatRead,
		Update: resourceLogCustomFormatUpdate,
		Delete: resourceLogCustomFormatDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"field_exclusion_list": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"empty_value_indicator": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1),
				Optional:     true,
			},
			"log_templates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subtypes": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subtype": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"template": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 2047),
							Optional:     true,
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

func resourceLogCustomFormatCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectLogCustomFormat(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LogCustomFormat resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLogCustomFormat(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLogCustomFormat(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating LogCustomFormat resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateLogCustomFormat(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating LogCustomFormat resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogCustomFormat")
	}

	return resourceLogCustomFormatRead(d, m)
}

func resourceLogCustomFormatUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectLogCustomFormat(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogCustomFormat resource while getting object: %v", err)
	}

	o, err := c.UpdateLogCustomFormat(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogCustomFormat resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogCustomFormat")
	}

	return resourceLogCustomFormatRead(d, m)
}

func resourceLogCustomFormatDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteLogCustomFormat(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting LogCustomFormat resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogCustomFormatRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadLogCustomFormat(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogCustomFormat resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogCustomFormat(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogCustomFormat resource from API: %v", err)
	}
	return nil
}

func flattenLogCustomFormatName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogCustomFormatFieldExclusionList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "field", "field")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["field"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "field"
			}
			tmp["field"] = flattenLogCustomFormatFieldExclusionListField(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "field", d)
	return result
}

func flattenLogCustomFormatFieldExclusionListField(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogCustomFormatEmptyValueIndicator(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogCustomFormatLogTemplates(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenLogCustomFormatLogTemplatesName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["category"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
			}
			tmp["category"] = flattenLogCustomFormatLogTemplatesCategory(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["subtypes"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "subtypes"
			}
			tmp["subtypes"] = flattenLogCustomFormatLogTemplatesSubtypes(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["template"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "template"
			}
			tmp["template"] = flattenLogCustomFormatLogTemplatesTemplate(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenLogCustomFormatLogTemplatesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogCustomFormatLogTemplatesCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogCustomFormatLogTemplatesSubtypes(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "subtype", "subtype")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["subtype"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
			}
			tmp["subtype"] = flattenLogCustomFormatLogTemplatesSubtypesSubtype(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "subtype", d)
	return result
}

func flattenLogCustomFormatLogTemplatesSubtypesSubtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogCustomFormatLogTemplatesTemplate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogCustomFormat(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenLogCustomFormatName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("field_exclusion_list", flattenLogCustomFormatFieldExclusionList(o["field-exclusion-list"], d, "field_exclusion_list", sv)); err != nil {
			if !fortiAPIPatch(o["field-exclusion-list"]) {
				return fmt.Errorf("Error reading field_exclusion_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("field_exclusion_list"); ok {
			if err = d.Set("field_exclusion_list", flattenLogCustomFormatFieldExclusionList(o["field-exclusion-list"], d, "field_exclusion_list", sv)); err != nil {
				if !fortiAPIPatch(o["field-exclusion-list"]) {
					return fmt.Errorf("Error reading field_exclusion_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("empty_value_indicator", flattenLogCustomFormatEmptyValueIndicator(o["empty-value-indicator"], d, "empty_value_indicator", sv)); err != nil {
		if !fortiAPIPatch(o["empty-value-indicator"]) {
			return fmt.Errorf("Error reading empty_value_indicator: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("log_templates", flattenLogCustomFormatLogTemplates(o["log-templates"], d, "log_templates", sv)); err != nil {
			if !fortiAPIPatch(o["log-templates"]) {
				return fmt.Errorf("Error reading log_templates: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("log_templates"); ok {
			if err = d.Set("log_templates", flattenLogCustomFormatLogTemplates(o["log-templates"], d, "log_templates", sv)); err != nil {
				if !fortiAPIPatch(o["log-templates"]) {
					return fmt.Errorf("Error reading log_templates: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenLogCustomFormatFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogCustomFormatName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatFieldExclusionList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["field"], _ = expandLogCustomFormatFieldExclusionListField(d, i["field"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogCustomFormatFieldExclusionListField(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatEmptyValueIndicator(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatLogTemplates(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandLogCustomFormatLogTemplatesName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandLogCustomFormatLogTemplatesCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtypes"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["subtypes"], _ = expandLogCustomFormatLogTemplatesSubtypes(d, i["subtypes"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["subtypes"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "template"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["template"], _ = expandLogCustomFormatLogTemplatesTemplate(d, i["template"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["template"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogCustomFormatLogTemplatesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatLogTemplatesCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatLogTemplatesSubtypes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["subtype"], _ = expandLogCustomFormatLogTemplatesSubtypesSubtype(d, i["subtype"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogCustomFormatLogTemplatesSubtypesSubtype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogCustomFormatLogTemplatesTemplate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogCustomFormat(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandLogCustomFormatName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("field_exclusion_list"); ok || d.HasChange("field_exclusion_list") {
		t, err := expandLogCustomFormatFieldExclusionList(d, v, "field_exclusion_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["field-exclusion-list"] = t
		}
	}

	if v, ok := d.GetOk("empty_value_indicator"); ok {
		t, err := expandLogCustomFormatEmptyValueIndicator(d, v, "empty_value_indicator", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-value-indicator"] = t
		}
	} else if d.HasChange("empty_value_indicator") {
		obj["empty-value-indicator"] = nil
	}

	if v, ok := d.GetOk("log_templates"); ok || d.HasChange("log_templates") {
		t, err := expandLogCustomFormatLogTemplates(d, v, "log_templates", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-templates"] = t
		}
	}

	return &obj, nil
}
