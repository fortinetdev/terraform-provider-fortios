// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 BFD.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterBfd6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBfd6Update,
		Read:   resourceRouterBfd6Read,
		Update: resourceRouterBfd6Update,
		Delete: resourceRouterBfd6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterBfd6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterBfd6(d, false)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterBfd6(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterBfd6")
	}

	return resourceRouterBfd6Read(d, m)
}

func resourceRouterBfd6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterBfd6(d, true)

	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBfd6(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing RouterBfd6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBfd6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterBfd6(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBfd6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterBfd6Neighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			tmp["ip6_address"] = flattenRouterBfd6NeighborIp6Address(i["ip6-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterBfd6NeighborInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterBfd6NeighborIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterBfd6NeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterBfd6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterBfd6Neighbor(o["neighbor"], d, "neighbor")); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterBfd6Neighbor(o["neighbor"], d, "neighbor")); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterBfd6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterBfd6Neighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6-address"], _ = expandRouterBfd6NeighborIp6Address(d, i["ip6_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterBfd6NeighborInterface(d, i["interface"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBfd6NeighborIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6NeighborInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBfd6(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if bemptysontable {
		obj["neighbor"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("neighbor"); ok {
			t, err := expandRouterBfd6Neighbor(d, v, "neighbor")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor"] = t
			}
		}
	}

	return &obj, nil
}
