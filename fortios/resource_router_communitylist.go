// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure community lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterCommunityList() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterCommunityListCreate,
		Read:   resourceRouterCommunityListRead,
		Update: resourceRouterCommunityListUpdate,
		Delete: resourceRouterCommunityListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required: true,
			},
			"type": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"rule": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"regexp": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional: true,
							Computed: true,
						},
						"match": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterCommunityListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterCommunityList(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterCommunityList resource while getting object: %v", err)
	}

	o, err := c.CreateRouterCommunityList(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterCommunityList resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterCommunityList")
	}

	return resourceRouterCommunityListRead(d, m)
}

func resourceRouterCommunityListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterCommunityList(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterCommunityList resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterCommunityList(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterCommunityList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterCommunityList")
	}

	return resourceRouterCommunityListRead(d, m)
}

func resourceRouterCommunityListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterCommunityList(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterCommunityList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterCommunityListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterCommunityList(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterCommunityList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterCommunityList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterCommunityList resource from API: %v", err)
	}
	return nil
}


func flattenRouterCommunityListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterCommunityListRuleId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenRouterCommunityListRuleAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := i["regexp"]; ok {
			tmp["regexp"] = flattenRouterCommunityListRuleRegexp(i["regexp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match"
		if _, ok := i["match"]; ok {
			tmp["match"] = flattenRouterCommunityListRuleMatch(i["match"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterCommunityListRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRuleRegexp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterCommunityListRuleMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectRouterCommunityList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenRouterCommunityListName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenRouterCommunityListType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("rule", flattenRouterCommunityListRule(o["rule"], d, "rule")); err != nil {
            if !fortiAPIPatch(o["rule"]) {
                return fmt.Errorf("Error reading rule: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("rule"); ok {
            if err = d.Set("rule", flattenRouterCommunityListRule(o["rule"], d, "rule")); err != nil {
                if !fortiAPIPatch(o["rule"]) {
                    return fmt.Errorf("Error reading rule: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenRouterCommunityListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandRouterCommunityListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterCommunityListRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandRouterCommunityListRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["regexp"], _ = expandRouterCommunityListRuleRegexp(d, i["regexp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match"], _ = expandRouterCommunityListRuleMatch(d, i["match"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterCommunityListRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRuleRegexp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterCommunityListRuleMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectRouterCommunityList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandRouterCommunityListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandRouterCommunityListType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok {
		t, err := expandRouterCommunityListRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}


	return &obj, nil
}

