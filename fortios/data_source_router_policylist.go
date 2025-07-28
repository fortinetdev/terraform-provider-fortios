// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRouterPolicyList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterPolicyListRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"seq_numlist": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceRouterPolicyListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	filter := d.Get("filter").(string)
	if filter != "" {
		filter = escapeFilter(filter)
	}

	o, err := c.GenericGroupRead("/api/v2/cmdb/router/policy", filter, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterPolicy: %v", err)
	}

	var tmps []int
	if o != nil {
		for _, r := range o {
			if r == nil {
				continue
			}
			i := r.(map[string]interface{})

			if _, ok := i["seq-num"]; ok {
				tmps = append(tmps, fortiIntValue(i["seq-num"]))
			}
		}
	}
	d.Set("seq_numlist", tmps)

	d.SetId("DataSourceRouterPolicyList" + filter)

	return nil
}
