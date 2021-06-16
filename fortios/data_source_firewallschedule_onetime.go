// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Onetime schedule configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallScheduleOnetime() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallScheduleOnetimeRead,
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
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"expiration_days": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallScheduleOnetimeRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing FirewallScheduleOnetime: type error")
	}

	o, err := c.ReadFirewallScheduleOnetime(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing FirewallScheduleOnetime: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallScheduleOnetime(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallScheduleOnetime from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallScheduleOnetimeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleOnetimeStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleOnetimeEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleOnetimeColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleOnetimeExpirationDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallScheduleOnetime(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallScheduleOnetimeName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("start", dataSourceFlattenFirewallScheduleOnetimeStart(o["start"], d, "start")); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("end", dataSourceFlattenFirewallScheduleOnetimeEnd(o["end"], d, "end")); err != nil {
		if !fortiAPIPatch(o["end"]) {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("color", dataSourceFlattenFirewallScheduleOnetimeColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("expiration_days", dataSourceFlattenFirewallScheduleOnetimeExpirationDays(o["expiration-days"], d, "expiration_days")); err != nil {
		if !fortiAPIPatch(o["expiration-days"]) {
			return fmt.Errorf("Error reading expiration_days: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallScheduleOnetimeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
