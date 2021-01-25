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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterAspathList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterAspathListRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"regexp": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterAspathListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterAspathList: type error")
	}

	o, err := c.ReadRouterAspathList(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterAspathList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterAspathList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterAspathList from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterAspathListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterAspathListRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterAspathListRuleId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = dataSourceFlattenRouterAspathListRuleAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := i["regexp"]; ok {
			tmp["regexp"] = dataSourceFlattenRouterAspathListRuleRegexp(i["regexp"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterAspathListRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterAspathListRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterAspathListRuleRegexp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterAspathList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenRouterAspathListName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("rule", dataSourceFlattenRouterAspathListRule(o["rule"], d, "rule")); err != nil {
		if !fortiAPIPatch(o["rule"]) {
			return fmt.Errorf("Error reading rule: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterAspathListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
