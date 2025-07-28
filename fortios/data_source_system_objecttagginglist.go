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

func dataSourceSystemObjectTaggingList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemObjectTaggingListRead,

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
			"categorylist": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceSystemObjectTaggingListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.GenericGroupRead("/api/v2/cmdb/system/object-tagging", filter, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemObjectTagging: %v", err)
	}

	var tmps []string
	if o != nil {
		for _, r := range o {
			if r == nil {
				continue
			}
			i := r.(map[string]interface{})

			if _, ok := i["category"]; ok {
				tmps = append(tmps, fortiStringValue(i["category"]))
			}
		}
	}
	d.Set("categorylist", tmps)

	d.SetId("DataSourceSystemObjectTaggingList" + filter)

	return nil
}
