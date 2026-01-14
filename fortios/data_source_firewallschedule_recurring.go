// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Recurring schedule configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceFirewallScheduleRecurring() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallScheduleRecurringRead,
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
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"label_day": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallScheduleRecurringRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing FirewallScheduleRecurring: type error")
	}

	o, err := c.ReadFirewallScheduleRecurring(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing FirewallScheduleRecurring: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallScheduleRecurring(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallScheduleRecurring from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallScheduleRecurringName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleRecurringUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleRecurringStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleRecurringEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleRecurringDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleRecurringLabelDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleRecurringColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallScheduleRecurringFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallScheduleRecurring(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallScheduleRecurringName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", dataSourceFlattenFirewallScheduleRecurringUuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("start", dataSourceFlattenFirewallScheduleRecurringStart(o["start"], d, "start")); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("end", dataSourceFlattenFirewallScheduleRecurringEnd(o["end"], d, "end")); err != nil {
		if !fortiAPIPatch(o["end"]) {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("day", dataSourceFlattenFirewallScheduleRecurringDay(o["day"], d, "day")); err != nil {
		if !fortiAPIPatch(o["day"]) {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	if err = d.Set("label_day", dataSourceFlattenFirewallScheduleRecurringLabelDay(o["label-day"], d, "label_day")); err != nil {
		if !fortiAPIPatch(o["label-day"]) {
			return fmt.Errorf("Error reading label_day: %v", err)
		}
	}

	if err = d.Set("color", dataSourceFlattenFirewallScheduleRecurringColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("fabric_object", dataSourceFlattenFirewallScheduleRecurringFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if !fortiAPIPatch(o["fabric-object"]) {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallScheduleRecurringFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
