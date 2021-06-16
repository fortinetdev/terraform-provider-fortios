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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Required:     true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceFirewallScheduleRecurringCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallScheduleRecurring(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallScheduleRecurring resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallScheduleRecurring(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallScheduleRecurring(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallScheduleRecurring resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallScheduleRecurring(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallScheduleRecurring(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallScheduleRecurring(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallScheduleRecurring resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallScheduleRecurring(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallScheduleRecurring resource from API: %v", err)
	}
	return nil
}

func flattenFirewallScheduleRecurringName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringEnd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallScheduleRecurringColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallScheduleRecurring(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallScheduleRecurringName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("start", flattenFirewallScheduleRecurringStart(o["start"], d, "start", sv)); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("end", flattenFirewallScheduleRecurringEnd(o["end"], d, "end", sv)); err != nil {
		if !fortiAPIPatch(o["end"]) {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("day", flattenFirewallScheduleRecurringDay(o["day"], d, "day", sv)); err != nil {
		if !fortiAPIPatch(o["day"]) {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallScheduleRecurringColor(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	return nil
}

func flattenFirewallScheduleRecurringFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallScheduleRecurringName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringEnd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleRecurringColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallScheduleRecurring(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallScheduleRecurringName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("start"); ok {

		t, err := expandFirewallScheduleRecurringStart(d, v, "start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	if v, ok := d.GetOk("end"); ok {

		t, err := expandFirewallScheduleRecurringEnd(d, v, "end", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end"] = t
		}
	}

	if v, ok := d.GetOk("day"); ok {

		t, err := expandFirewallScheduleRecurringDay(d, v, "day", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {

		t, err := expandFirewallScheduleRecurringColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	return &obj, nil
}
