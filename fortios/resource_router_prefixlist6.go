// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 prefix lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterPrefixList6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterPrefixList6Create,
		Read:   resourceRouterPrefixList6Read,
		Update: resourceRouterPrefixList6Update,
		Delete: resourceRouterPrefixList6Delete,

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
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ge": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 128),
							Optional:     true,
							Computed:     true,
						},
						"le": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 128),
							Optional:     true,
							Computed:     true,
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

func resourceRouterPrefixList6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterPrefixList6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixList6 resource while getting object: %v", err)
	}

	o, err := c.CreateRouterPrefixList6(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterPrefixList6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterPrefixList6")
	}

	return resourceRouterPrefixList6Read(d, m)
}

func resourceRouterPrefixList6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterPrefixList6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterPrefixList6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterPrefixList6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterPrefixList6")
	}

	return resourceRouterPrefixList6Read(d, m)
}

func resourceRouterPrefixList6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterPrefixList6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPrefixList6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPrefixList6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadRouterPrefixList6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixList6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPrefixList6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterPrefixList6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterPrefixList6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6Comments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6Rule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterPrefixList6RuleId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenRouterPrefixList6RuleAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {

			tmp["prefix6"] = flattenRouterPrefixList6RulePrefix6(i["prefix6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ge"
		if _, ok := i["ge"]; ok {

			tmp["ge"] = flattenRouterPrefixList6RuleGe(i["ge"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "le"
		if _, ok := i["le"]; ok {

			tmp["le"] = flattenRouterPrefixList6RuleLe(i["le"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {

			tmp["flags"] = flattenRouterPrefixList6RuleFlags(i["flags"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterPrefixList6RuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6RulePrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleGe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleLe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPrefixList6RuleFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterPrefixList6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenRouterPrefixList6Name(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterPrefixList6Comments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenRouterPrefixList6Rule(o["rule"], d, "rule", sv)); err != nil {
			if !fortiAPIPatch(o["rule"]) {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenRouterPrefixList6Rule(o["rule"], d, "rule", sv)); err != nil {
				if !fortiAPIPatch(o["rule"]) {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterPrefixList6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterPrefixList6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6Comments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6Rule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandRouterPrefixList6RuleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandRouterPrefixList6RuleAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix6"], _ = expandRouterPrefixList6RulePrefix6(d, i["prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ge"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ge"], _ = expandRouterPrefixList6RuleGe(d, i["ge"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "le"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["le"], _ = expandRouterPrefixList6RuleLe(d, i["le"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["flags"], _ = expandRouterPrefixList6RuleFlags(d, i["flags"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPrefixList6RuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RulePrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleGe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleLe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPrefixList6RuleFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPrefixList6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterPrefixList6Name(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandRouterPrefixList6Comments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok {

		t, err := expandRouterPrefixList6Rule(d, v, "rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
