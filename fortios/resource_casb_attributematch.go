// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure CASB SaaS application.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbAttributeMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbAttributeMatchCreate,
		Read:   resourceCasbAttributeMatchRead,
		Update: resourceCasbAttributeMatchUpdate,
		Delete: resourceCasbAttributeMatchDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"application": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"match_strategy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"match_pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),
							Optional:     true,
						},
						"case_sensitive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"negate": &schema.Schema{
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceCasbAttributeMatchCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbAttributeMatch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CasbAttributeMatch resource while getting object: %v", err)
	}

	o, err := c.CreateCasbAttributeMatch(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating CasbAttributeMatch resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CasbAttributeMatch")
	}

	return resourceCasbAttributeMatchRead(d, m)
}

func resourceCasbAttributeMatchUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbAttributeMatch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatch resource while getting object: %v", err)
	}

	o, err := c.UpdateCasbAttributeMatch(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating CasbAttributeMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CasbAttributeMatch")
	}

	return resourceCasbAttributeMatchRead(d, m)
}

func resourceCasbAttributeMatchDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCasbAttributeMatch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting CasbAttributeMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbAttributeMatchRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbAttributeMatch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CasbAttributeMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbAttributeMatch(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CasbAttributeMatch resource from API: %v", err)
	}
	return nil
}

func flattenCasbAttributeMatchName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbAttributeMatchApplication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbAttributeMatchMatchStrategy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenCasbAttributeMatchAttributeName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if cur_v, ok := i["match-pattern"]; ok {
			tmp["match_pattern"] = flattenCasbAttributeMatchAttributeMatchPattern(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if cur_v, ok := i["match-value"]; ok {
			tmp["match_value"] = flattenCasbAttributeMatchAttributeMatchValue(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if cur_v, ok := i["case-sensitive"]; ok {
			tmp["case_sensitive"] = flattenCasbAttributeMatchAttributeCaseSensitive(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if cur_v, ok := i["negate"]; ok {
			tmp["negate"] = flattenCasbAttributeMatchAttributeNegate(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbAttributeMatchAttributeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeMatchPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeMatchValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeCaseSensitive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbAttributeMatchAttributeNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCasbAttributeMatch(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenCasbAttributeMatchName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("application", flattenCasbAttributeMatchApplication(o["application"], d, "application", sv)); err != nil {
		if !fortiAPIPatch(o["application"]) {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("match_strategy", flattenCasbAttributeMatchMatchStrategy(o["match-strategy"], d, "match_strategy", sv)); err != nil {
		if !fortiAPIPatch(o["match-strategy"]) {
			return fmt.Errorf("Error reading match_strategy: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("attribute", flattenCasbAttributeMatchAttribute(o["attribute"], d, "attribute", sv)); err != nil {
			if !fortiAPIPatch(o["attribute"]) {
				return fmt.Errorf("Error reading attribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("attribute"); ok {
			if err = d.Set("attribute", flattenCasbAttributeMatchAttribute(o["attribute"], d, "attribute", sv)); err != nil {
				if !fortiAPIPatch(o["attribute"]) {
					return fmt.Errorf("Error reading attribute: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenCasbAttributeMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCasbAttributeMatchName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchMatchStrategy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbAttributeMatchAttributeName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-pattern"], _ = expandCasbAttributeMatchAttributeMatchPattern(d, i["match_pattern"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-value"], _ = expandCasbAttributeMatchAttributeMatchValue(d, i["match_value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-value"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["case-sensitive"], _ = expandCasbAttributeMatchAttributeCaseSensitive(d, i["case_sensitive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["negate"], _ = expandCasbAttributeMatchAttributeNegate(d, i["negate"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbAttributeMatchAttributeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeMatchPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeMatchValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeCaseSensitive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbAttributeMatchAttributeNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCasbAttributeMatch(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandCasbAttributeMatchName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {
		t, err := expandCasbAttributeMatchApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	} else if d.HasChange("application") {
		obj["application"] = nil
	}

	if v, ok := d.GetOk("match_strategy"); ok {
		t, err := expandCasbAttributeMatchMatchStrategy(d, v, "match_strategy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-strategy"] = t
		}
	}

	if v, ok := d.GetOk("attribute"); ok || d.HasChange("attribute") {
		t, err := expandCasbAttributeMatchAttribute(d, v, "attribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute"] = t
		}
	}

	return &obj, nil
}
