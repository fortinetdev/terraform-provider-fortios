// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 address templates.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallAddress6Template() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallAddress6TemplateRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_segment_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"subnet_segment": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bits": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"exclusive": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"values": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
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

func dataSourceFirewallAddress6TemplateRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing FirewallAddress6Template: type error")
	}

	o, err := c.ReadFirewallAddress6Template(mkey)
	if err != nil {
		return fmt.Errorf("Error describing FirewallAddress6Template: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallAddress6Template(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallAddress6Template from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallAddress6TemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TemplateIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TemplateSubnetSegmentCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TemplateSubnetSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallAddress6TemplateSubnetSegmentId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenFirewallAddress6TemplateSubnetSegmentName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bits"
		if _, ok := i["bits"]; ok {
			tmp["bits"] = dataSourceFlattenFirewallAddress6TemplateSubnetSegmentBits(i["bits"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusive"
		if _, ok := i["exclusive"]; ok {
			tmp["exclusive"] = dataSourceFlattenFirewallAddress6TemplateSubnetSegmentExclusive(i["exclusive"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := i["values"]; ok {
			tmp["values"] = dataSourceFlattenFirewallAddress6TemplateSubnetSegmentValues(i["values"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallAddress6TemplateSubnetSegmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TemplateSubnetSegmentName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TemplateSubnetSegmentBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TemplateSubnetSegmentExclusive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TemplateSubnetSegmentValues(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallAddress6TemplateSubnetSegmentValuesName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = dataSourceFlattenFirewallAddress6TemplateSubnetSegmentValuesValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallAddress6TemplateSubnetSegmentValuesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TemplateSubnetSegmentValuesValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallAddress6Template(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallAddress6TemplateName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ip6", dataSourceFlattenFirewallAddress6TemplateIp6(o["ip6"], d, "ip6")); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("subnet_segment_count", dataSourceFlattenFirewallAddress6TemplateSubnetSegmentCount(o["subnet-segment-count"], d, "subnet_segment_count")); err != nil {
		if !fortiAPIPatch(o["subnet-segment-count"]) {
			return fmt.Errorf("Error reading subnet_segment_count: %v", err)
		}
	}

	if err = d.Set("subnet_segment", dataSourceFlattenFirewallAddress6TemplateSubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
		if !fortiAPIPatch(o["subnet-segment"]) {
			return fmt.Errorf("Error reading subnet_segment: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallAddress6TemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
