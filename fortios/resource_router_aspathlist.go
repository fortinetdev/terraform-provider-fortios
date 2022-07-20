// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Autonomous System (AS) path lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterAspathList() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterAspathListCreate,
		Read:   resourceRouterAspathListRead,
		Update: resourceRouterAspathListUpdate,
		Delete: resourceRouterAspathListDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"rule": &schema.Schema{
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
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"regexp": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
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

func resourceRouterAspathListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterAspathList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterAspathList resource while getting object: %v", err)
	}

	o, err := c.CreateRouterAspathList(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterAspathList resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterAspathList")
	}

	return resourceRouterAspathListRead(d, m)
}

func resourceRouterAspathListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterAspathList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterAspathList resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterAspathList(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterAspathList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterAspathList")
	}

	return resourceRouterAspathListRead(d, m)
}

func resourceRouterAspathListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterAspathList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterAspathList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterAspathListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadRouterAspathList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterAspathList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterAspathList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterAspathList resource from API: %v", err)
	}
	return nil
}

func flattenRouterAspathListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAspathListRule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterAspathListRuleId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenRouterAspathListRuleAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := i["regexp"]; ok {

			tmp["regexp"] = flattenRouterAspathListRuleRegexp(i["regexp"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterAspathListRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAspathListRuleAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAspathListRuleRegexp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterAspathList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenRouterAspathListName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenRouterAspathListRule(o["rule"], d, "rule", sv)); err != nil {
			if !fortiAPIPatch(o["rule"]) {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenRouterAspathListRule(o["rule"], d, "rule", sv)); err != nil {
				if !fortiAPIPatch(o["rule"]) {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterAspathListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterAspathListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAspathListRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandRouterAspathListRuleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandRouterAspathListRuleAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["regexp"], _ = expandRouterAspathListRuleRegexp(d, i["regexp"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterAspathListRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAspathListRuleAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAspathListRuleRegexp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterAspathList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterAspathListName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok {

		t, err := expandRouterAspathListRule(d, v, "rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
