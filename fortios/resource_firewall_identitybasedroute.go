// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure identity based routing.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallIdentityBasedRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIdentityBasedRouteCreate,
		Read:   resourceFirewallIdentityBasedRouteRead,
		Update: resourceFirewallIdentityBasedRouteUpdate,
		Delete: resourceFirewallIdentityBasedRouteDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required: true,
			},
			"comments": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional: true,
				Computed: true,
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
						"gateway": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"groups": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type: schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
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

func resourceFirewallIdentityBasedRouteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallIdentityBasedRoute(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallIdentityBasedRoute resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallIdentityBasedRoute(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallIdentityBasedRoute resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIdentityBasedRoute")
	}

	return resourceFirewallIdentityBasedRouteRead(d, m)
}

func resourceFirewallIdentityBasedRouteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallIdentityBasedRoute(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIdentityBasedRoute resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIdentityBasedRoute(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIdentityBasedRoute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIdentityBasedRoute")
	}

	return resourceFirewallIdentityBasedRouteRead(d, m)
}

func resourceFirewallIdentityBasedRouteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallIdentityBasedRoute(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIdentityBasedRoute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIdentityBasedRouteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallIdentityBasedRoute(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIdentityBasedRoute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIdentityBasedRoute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIdentityBasedRoute resource from API: %v", err)
	}
	return nil
}


func flattenFirewallIdentityBasedRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIdentityBasedRouteComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIdentityBasedRouteRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallIdentityBasedRouteRuleId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			tmp["gateway"] = flattenFirewallIdentityBasedRouteRuleGateway(i["gateway"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			tmp["device"] = flattenFirewallIdentityBasedRouteRuleDevice(i["device"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {
			tmp["groups"] = flattenFirewallIdentityBasedRouteRuleGroups(i["groups"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallIdentityBasedRouteRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIdentityBasedRouteRuleGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIdentityBasedRouteRuleDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIdentityBasedRouteRuleGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenFirewallIdentityBasedRouteRuleGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallIdentityBasedRouteRuleGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectFirewallIdentityBasedRoute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenFirewallIdentityBasedRouteName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallIdentityBasedRouteComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("rule", flattenFirewallIdentityBasedRouteRule(o["rule"], d, "rule")); err != nil {
            if !fortiAPIPatch(o["rule"]) {
                return fmt.Errorf("Error reading rule: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("rule"); ok {
            if err = d.Set("rule", flattenFirewallIdentityBasedRouteRule(o["rule"], d, "rule")); err != nil {
                if !fortiAPIPatch(o["rule"]) {
                    return fmt.Errorf("Error reading rule: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenFirewallIdentityBasedRouteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandFirewallIdentityBasedRouteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIdentityBasedRouteComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIdentityBasedRouteRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallIdentityBasedRouteRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway"], _ = expandFirewallIdentityBasedRouteRuleGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device"], _ = expandFirewallIdentityBasedRouteRuleDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["groups"], _ = expandFirewallIdentityBasedRouteRuleGroups(d, i["groups"], pre_append)
		} else {
			tmp["groups"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallIdentityBasedRouteRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIdentityBasedRouteRuleGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIdentityBasedRouteRuleDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIdentityBasedRouteRuleGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandFirewallIdentityBasedRouteRuleGroupsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallIdentityBasedRouteRuleGroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectFirewallIdentityBasedRoute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallIdentityBasedRouteName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandFirewallIdentityBasedRouteComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok {
		t, err := expandFirewallIdentityBasedRouteRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}


	return &obj, nil
}

