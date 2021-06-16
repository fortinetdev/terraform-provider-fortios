// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure shared traffic shaper.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallShaperTrafficShaper() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallShaperTrafficShaperRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"guaranteed_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"maximum_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bandwidth_unit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"per_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dscp_marking_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"exceed_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"exceed_dscp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"maximum_dscp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"overhead": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"exceed_class_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallShaperTrafficShaperRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing FirewallShaperTrafficShaper: type error")
	}

	o, err := c.ReadFirewallShaperTrafficShaper(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing FirewallShaperTrafficShaper: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallShaperTrafficShaper(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallShaperTrafficShaper from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallShaperTrafficShaperName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperGuaranteedBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperMaximumBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperBandwidthUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperPerPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperDiffserv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperDscpMarkingMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperExceedBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperExceedDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperMaximumDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperOverhead(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallShaperTrafficShaperExceedClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallShaperTrafficShaper(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallShaperTrafficShaperName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth", dataSourceFlattenFirewallShaperTrafficShaperGuaranteedBandwidth(o["guaranteed-bandwidth"], d, "guaranteed_bandwidth")); err != nil {
		if !fortiAPIPatch(o["guaranteed-bandwidth"]) {
			return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
		}
	}

	if err = d.Set("maximum_bandwidth", dataSourceFlattenFirewallShaperTrafficShaperMaximumBandwidth(o["maximum-bandwidth"], d, "maximum_bandwidth")); err != nil {
		if !fortiAPIPatch(o["maximum-bandwidth"]) {
			return fmt.Errorf("Error reading maximum_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_unit", dataSourceFlattenFirewallShaperTrafficShaperBandwidthUnit(o["bandwidth-unit"], d, "bandwidth_unit")); err != nil {
		if !fortiAPIPatch(o["bandwidth-unit"]) {
			return fmt.Errorf("Error reading bandwidth_unit: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenFirewallShaperTrafficShaperPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("per_policy", dataSourceFlattenFirewallShaperTrafficShaperPerPolicy(o["per-policy"], d, "per_policy")); err != nil {
		if !fortiAPIPatch(o["per-policy"]) {
			return fmt.Errorf("Error reading per_policy: %v", err)
		}
	}

	if err = d.Set("diffserv", dataSourceFlattenFirewallShaperTrafficShaperDiffserv(o["diffserv"], d, "diffserv")); err != nil {
		if !fortiAPIPatch(o["diffserv"]) {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("diffservcode", dataSourceFlattenFirewallShaperTrafficShaperDiffservcode(o["diffservcode"], d, "diffservcode")); err != nil {
		if !fortiAPIPatch(o["diffservcode"]) {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dscp_marking_method", dataSourceFlattenFirewallShaperTrafficShaperDscpMarkingMethod(o["dscp-marking-method"], d, "dscp_marking_method")); err != nil {
		if !fortiAPIPatch(o["dscp-marking-method"]) {
			return fmt.Errorf("Error reading dscp_marking_method: %v", err)
		}
	}

	if err = d.Set("exceed_bandwidth", dataSourceFlattenFirewallShaperTrafficShaperExceedBandwidth(o["exceed-bandwidth"], d, "exceed_bandwidth")); err != nil {
		if !fortiAPIPatch(o["exceed-bandwidth"]) {
			return fmt.Errorf("Error reading exceed_bandwidth: %v", err)
		}
	}

	if err = d.Set("exceed_dscp", dataSourceFlattenFirewallShaperTrafficShaperExceedDscp(o["exceed-dscp"], d, "exceed_dscp")); err != nil {
		if !fortiAPIPatch(o["exceed-dscp"]) {
			return fmt.Errorf("Error reading exceed_dscp: %v", err)
		}
	}

	if err = d.Set("maximum_dscp", dataSourceFlattenFirewallShaperTrafficShaperMaximumDscp(o["maximum-dscp"], d, "maximum_dscp")); err != nil {
		if !fortiAPIPatch(o["maximum-dscp"]) {
			return fmt.Errorf("Error reading maximum_dscp: %v", err)
		}
	}

	if err = d.Set("overhead", dataSourceFlattenFirewallShaperTrafficShaperOverhead(o["overhead"], d, "overhead")); err != nil {
		if !fortiAPIPatch(o["overhead"]) {
			return fmt.Errorf("Error reading overhead: %v", err)
		}
	}

	if err = d.Set("exceed_class_id", dataSourceFlattenFirewallShaperTrafficShaperExceedClassId(o["exceed-class-id"], d, "exceed_class_id")); err != nil {
		if !fortiAPIPatch(o["exceed-class-id"]) {
			return fmt.Errorf("Error reading exceed_class_id: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallShaperTrafficShaperFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
