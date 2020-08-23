// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Recurring schedule configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallScheduleRecurring() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallScheduleRecurringCreate,
		Read:   resourceFirewallScheduleRecurringRead,
		Update: resourceFirewallScheduleRecurringUpdate,
		Delete: resourceFirewallScheduleRecurringDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Required: true,
			},
			"start": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"end": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"day": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallScheduleRecurringCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallScheduleRecurring(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallScheduleRecurring resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallScheduleRecurring(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallScheduleRecurring resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallScheduleRecurring")
	}

	return resourceFirewallScheduleRecurringRead(d, m)
}

func resourceFirewallScheduleRecurringUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallScheduleRecurring(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallScheduleRecurring resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallScheduleRecurring(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallScheduleRecurring resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallScheduleRecurring")
	}

	return resourceFirewallScheduleRecurringRead(d, m)
}

func resourceFirewallScheduleRecurringDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallScheduleRecurring(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallScheduleRecurring resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallScheduleRecurringRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallScheduleRecurring(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallScheduleRecurring resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallScheduleRecurring(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallScheduleRecurring resource from API: %v", err)
	}
	return nil
}


func flattenFirewallScheduleRecurringName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectFirewallScheduleRecurring(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenFirewallScheduleRecurringName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("start", flattenFirewallScheduleRecurringStart(o["start"], d, "start")); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("end", flattenFirewallScheduleRecurringEnd(o["end"], d, "end")); err != nil {
		if !fortiAPIPatch(o["end"]) {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("day", flattenFirewallScheduleRecurringDay(o["day"], d, "day")); err != nil {
		if !fortiAPIPatch(o["day"]) {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallScheduleRecurringColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}


	return nil
}

func flattenFirewallScheduleRecurringFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandFirewallScheduleRecurringName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectFirewallScheduleRecurring(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallScheduleRecurringName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("start"); ok {
		t, err := expandFirewallScheduleRecurringStart(d, v, "start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	if v, ok := d.GetOk("end"); ok {
		t, err := expandFirewallScheduleRecurringEnd(d, v, "end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end"] = t
		}
	}

	if v, ok := d.GetOk("day"); ok {
		t, err := expandFirewallScheduleRecurringDay(d, v, "day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok {
		t, err := expandFirewallScheduleRecurringColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}


	return &obj, nil
}

