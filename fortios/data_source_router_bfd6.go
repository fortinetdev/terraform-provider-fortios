// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 BFD.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterBfd6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterBfd6Read,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterBfd6Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterBfd6"

	o, err := c.ReadRouterBfd6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterBfd6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterBfd6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterBfd6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterBfd6Neighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["ip6_address"] = dataSourceFlattenRouterBfd6NeighborIp6Address(i["ip6-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterBfd6NeighborInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBfd6NeighborIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6NeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterBfd6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("neighbor", dataSourceFlattenRouterBfd6Neighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("Error reading neighbor: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterBfd6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
