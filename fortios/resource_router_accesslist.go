// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure access lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterAccessList() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterAccessListCreate,
		Read:   resourceRouterAccessListRead,
		Update: resourceRouterAccessListUpdate,
		Delete: resourceRouterAccessListDelete,

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
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
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
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wildcard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exact_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
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
		},
	}
}

func resourceRouterAccessListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterAccessList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterAccessList resource while getting object: %v", err)
	}

	o, err := c.CreateRouterAccessList(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterAccessList resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterAccessList")
	}

	return resourceRouterAccessListRead(d, m)
}

func resourceRouterAccessListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterAccessList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterAccessList(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterAccessList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterAccessList")
	}

	return resourceRouterAccessListRead(d, m)
}

func resourceRouterAccessListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterAccessList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterAccessList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterAccessListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadRouterAccessList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterAccessList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterAccessList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterAccessList resource from API: %v", err)
	}
	return nil
}

func flattenRouterAccessListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessListComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessListRule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterAccessListRuleId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenRouterAccessListRuleAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenRouterAccessListRulePrefix(i["prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard"
		if _, ok := i["wildcard"]; ok {

			tmp["wildcard"] = flattenRouterAccessListRuleWildcard(i["wildcard"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exact_match"
		if _, ok := i["exact-match"]; ok {

			tmp["exact_match"] = flattenRouterAccessListRuleExactMatch(i["exact-match"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {

			tmp["flags"] = flattenRouterAccessListRuleFlags(i["flags"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterAccessListRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessListRuleAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessListRulePrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessListRuleWildcard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessListRuleExactMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterAccessListRuleFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterAccessList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenRouterAccessListName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterAccessListComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenRouterAccessListRule(o["rule"], d, "rule", sv)); err != nil {
			if !fortiAPIPatch(o["rule"]) {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenRouterAccessListRule(o["rule"], d, "rule", sv)); err != nil {
				if !fortiAPIPatch(o["rule"]) {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterAccessListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterAccessListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandRouterAccessListRuleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandRouterAccessListRuleAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandRouterAccessListRulePrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["wildcard"], _ = expandRouterAccessListRuleWildcard(d, i["wildcard"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exact_match"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["exact-match"], _ = expandRouterAccessListRuleExactMatch(d, i["exact_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["flags"], _ = expandRouterAccessListRuleFlags(d, i["flags"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterAccessListRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRuleAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRulePrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRuleWildcard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRuleExactMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterAccessListRuleFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterAccessList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterAccessListName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandRouterAccessListComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok {

		t, err := expandRouterAccessListRule(d, v, "rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
