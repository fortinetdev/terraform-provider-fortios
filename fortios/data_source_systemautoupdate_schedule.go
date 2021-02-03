// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure update schedule.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemAutoupdateSchedule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAutoupdateScheduleRead,
		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"frequency": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemAutoupdateScheduleRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemAutoupdateSchedule"

	o, err := c.ReadSystemAutoupdateSchedule(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoupdateSchedule: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAutoupdateSchedule(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutoupdateSchedule from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAutoupdateScheduleStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoupdateScheduleFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoupdateScheduleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutoupdateScheduleDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAutoupdateSchedule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemAutoupdateScheduleStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("frequency", dataSourceFlattenSystemAutoupdateScheduleFrequency(o["frequency"], d, "frequency")); err != nil {
		if !fortiAPIPatch(o["frequency"]) {
			return fmt.Errorf("Error reading frequency: %v", err)
		}
	}

	if err = d.Set("time", dataSourceFlattenSystemAutoupdateScheduleTime(o["time"], d, "time")); err != nil {
		if !fortiAPIPatch(o["time"]) {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	if err = d.Set("day", dataSourceFlattenSystemAutoupdateScheduleDay(o["day"], d, "day")); err != nil {
		if !fortiAPIPatch(o["day"]) {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAutoupdateScheduleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
