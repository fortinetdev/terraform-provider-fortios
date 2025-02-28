// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Automation stitches.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationStitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationStitchCreate,
		Read:   resourceSystemAutomationStitchRead,
		Update: resourceSystemAutomationStitchUpdate,
		Delete: resourceSystemAutomationStitchDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"trigger": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"condition": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"condition_logic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
						},
						"delay": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),
							Optional:     true,
						},
						"required": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"destination": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
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

func resourceSystemAutomationStitchCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAutomationStitch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationStitch resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAutomationStitch(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationStitch resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationStitch")
	}

	return resourceSystemAutomationStitchRead(d, m)
}

func resourceSystemAutomationStitchUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAutomationStitch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationStitch resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutomationStitch(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationStitch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationStitch")
	}

	return resourceSystemAutomationStitchRead(d, m)
}

func resourceSystemAutomationStitchDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemAutomationStitch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationStitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationStitchRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAutomationStitch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationStitch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationStitch(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationStitch resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationStitchName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationStitchDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationStitchStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationStitchTrigger(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationStitchCondition(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemAutomationStitchConditionName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemAutomationStitchConditionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationStitchConditionLogic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationStitchActions(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenSystemAutomationStitchActionsId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenSystemAutomationStitchActionsAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delay"
		if cur_v, ok := i["delay"]; ok {
			tmp["delay"] = flattenSystemAutomationStitchActionsDelay(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if cur_v, ok := i["required"]; ok {
			tmp["required"] = flattenSystemAutomationStitchActionsRequired(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemAutomationStitchActionsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationStitchActionsAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationStitchActionsDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationStitchActionsRequired(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationStitchAction(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemAutomationStitchActionName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemAutomationStitchActionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationStitchDestination(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemAutomationStitchDestinationName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemAutomationStitchDestinationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAutomationStitch(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemAutomationStitchName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemAutomationStitchDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemAutomationStitchStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trigger", flattenSystemAutomationStitchTrigger(o["trigger"], d, "trigger", sv)); err != nil {
		if !fortiAPIPatch(o["trigger"]) {
			return fmt.Errorf("Error reading trigger: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("condition", flattenSystemAutomationStitchCondition(o["condition"], d, "condition", sv)); err != nil {
			if !fortiAPIPatch(o["condition"]) {
				return fmt.Errorf("Error reading condition: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("condition"); ok {
			if err = d.Set("condition", flattenSystemAutomationStitchCondition(o["condition"], d, "condition", sv)); err != nil {
				if !fortiAPIPatch(o["condition"]) {
					return fmt.Errorf("Error reading condition: %v", err)
				}
			}
		}
	}

	if err = d.Set("condition_logic", flattenSystemAutomationStitchConditionLogic(o["condition-logic"], d, "condition_logic", sv)); err != nil {
		if !fortiAPIPatch(o["condition-logic"]) {
			return fmt.Errorf("Error reading condition_logic: %v", err)
		}
	}

	if _, ok := o["actions"].([]interface{}); ok {
		if b_get_all_tables {
			if err = d.Set("actions", flattenSystemAutomationStitchActions(o["actions"], d, "actions", sv)); err != nil {
				if !fortiAPIPatch(o["actions"]) {
					return fmt.Errorf("Error reading actions: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("actions"); ok {
				if err = d.Set("actions", flattenSystemAutomationStitchActions(o["actions"], d, "actions", sv)); err != nil {
					if !fortiAPIPatch(o["actions"]) {
						return fmt.Errorf("Error reading actions: %v", err)
					}
				}
			}
		}
	}

	if _, ok := o["action"].([]interface{}); ok {
		if b_get_all_tables {
			if err = d.Set("action", flattenSystemAutomationStitchAction(o["action"], d, "action", sv)); err != nil {
				if !fortiAPIPatch(o["action"]) {
					return fmt.Errorf("Error reading action: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("action"); ok {
				if err = d.Set("action", flattenSystemAutomationStitchAction(o["action"], d, "action", sv)); err != nil {
					if !fortiAPIPatch(o["action"]) {
						return fmt.Errorf("Error reading action: %v", err)
					}
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("destination", flattenSystemAutomationStitchDestination(o["destination"], d, "destination", sv)); err != nil {
			if !fortiAPIPatch(o["destination"]) {
				return fmt.Errorf("Error reading destination: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("destination"); ok {
			if err = d.Set("destination", flattenSystemAutomationStitchDestination(o["destination"], d, "destination", sv)); err != nil {
				if !fortiAPIPatch(o["destination"]) {
					return fmt.Errorf("Error reading destination: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemAutomationStitchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAutomationStitchName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchTrigger(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchCondition(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemAutomationStitchConditionName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationStitchConditionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchConditionLogic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchActions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAutomationStitchActionsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandSystemAutomationStitchActionsAction(d, i["action"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["action"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["delay"], _ = expandSystemAutomationStitchActionsDelay(d, i["delay"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["delay"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "required"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["required"], _ = expandSystemAutomationStitchActionsRequired(d, i["required"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationStitchActionsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchActionsAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchActionsDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchActionsRequired(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemAutomationStitchActionName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationStitchActionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationStitchDestination(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemAutomationStitchDestinationName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationStitchDestinationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationStitch(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAutomationStitchName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemAutomationStitchDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemAutomationStitchStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trigger"); ok {
		t, err := expandSystemAutomationStitchTrigger(d, v, "trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger"] = t
		}
	} else if d.HasChange("trigger") {
		obj["trigger"] = nil
	}

	if v, ok := d.GetOk("condition"); ok || d.HasChange("condition") {
		t, err := expandSystemAutomationStitchCondition(d, v, "condition", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["condition"] = t
		}
	}

	if v, ok := d.GetOk("condition_logic"); ok {
		t, err := expandSystemAutomationStitchConditionLogic(d, v, "condition_logic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["condition-logic"] = t
		}
	}

	if v, ok := d.GetOk("actions"); ok || d.HasChange("actions") {
		new_version_map := map[string][]string{
			">=": []string{"7.0.1"},
		}
		if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
			if _, ok := d.GetOk("action"); !ok && !d.HasChange("action") {
				err := fmt.Errorf("Argument 'actions' %s.", err)
				return nil, err
			}
		} else {
			t, err := expandSystemAutomationStitchActions(d, v, "actions", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["actions"] = t
			}
		}
	}

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		new_version_map := map[string][]string{
			"<=": []string{"7.0.0"},
		}
		if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
			if _, ok := d.GetOk("actions"); !ok && !d.HasChange("actions") {
				err := fmt.Errorf("Argument 'action' %s.", err)
				return nil, err
			}
		} else {
			t, err := expandSystemAutomationStitchAction(d, v, "action", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["action"] = t
			}
		}
	}

	if v, ok := d.GetOk("destination"); ok || d.HasChange("destination") {
		t, err := expandSystemAutomationStitchDestination(d, v, "destination", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination"] = t
		}
	}

	return &obj, nil
}
