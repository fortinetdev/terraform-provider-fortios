// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Trigger for automation stitches.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutomationTrigger() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationTriggerCreate,
		Read:   resourceSystemAutomationTriggerRead,
		Update: resourceSystemAutomationTriggerUpdate,
		Delete: resourceSystemAutomationTriggerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"trigger_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"license_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ioc_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),
				Optional:     true,
				Computed:     true,
			},
			"trigger_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_day": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),
				Optional:     true,
				Computed:     true,
			},
			"trigger_hour": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),
				Optional:     true,
				Computed:     true,
			},
			"trigger_minute": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 60),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemAutomationTriggerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutomationTrigger(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTrigger resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAutomationTrigger(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTrigger resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationTrigger")
	}

	return resourceSystemAutomationTriggerRead(d, m)
}

func resourceSystemAutomationTriggerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutomationTrigger(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTrigger resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutomationTrigger(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTrigger resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationTrigger")
	}

	return resourceSystemAutomationTriggerRead(d, m)
}

func resourceSystemAutomationTriggerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAutomationTrigger(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationTrigger resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationTriggerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAutomationTrigger(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTrigger resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationTrigger(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTrigger resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationTriggerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerEventType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerLicenseType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerIocLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerLogid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerWeekday(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerMinute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutomationTrigger(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemAutomationTriggerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("trigger_type", flattenSystemAutomationTriggerTriggerType(o["trigger-type"], d, "trigger_type")); err != nil {
		if !fortiAPIPatch(o["trigger-type"]) {
			return fmt.Errorf("Error reading trigger_type: %v", err)
		}
	}

	if err = d.Set("event_type", flattenSystemAutomationTriggerEventType(o["event-type"], d, "event_type")); err != nil {
		if !fortiAPIPatch(o["event-type"]) {
			return fmt.Errorf("Error reading event_type: %v", err)
		}
	}

	if err = d.Set("license_type", flattenSystemAutomationTriggerLicenseType(o["license-type"], d, "license_type")); err != nil {
		if !fortiAPIPatch(o["license-type"]) {
			return fmt.Errorf("Error reading license_type: %v", err)
		}
	}

	if err = d.Set("ioc_level", flattenSystemAutomationTriggerIocLevel(o["ioc-level"], d, "ioc_level")); err != nil {
		if !fortiAPIPatch(o["ioc-level"]) {
			return fmt.Errorf("Error reading ioc_level: %v", err)
		}
	}

	if err = d.Set("logid", flattenSystemAutomationTriggerLogid(o["logid"], d, "logid")); err != nil {
		if !fortiAPIPatch(o["logid"]) {
			return fmt.Errorf("Error reading logid: %v", err)
		}
	}

	if err = d.Set("trigger_frequency", flattenSystemAutomationTriggerTriggerFrequency(o["trigger-frequency"], d, "trigger_frequency")); err != nil {
		if !fortiAPIPatch(o["trigger-frequency"]) {
			return fmt.Errorf("Error reading trigger_frequency: %v", err)
		}
	}

	if err = d.Set("trigger_weekday", flattenSystemAutomationTriggerTriggerWeekday(o["trigger-weekday"], d, "trigger_weekday")); err != nil {
		if !fortiAPIPatch(o["trigger-weekday"]) {
			return fmt.Errorf("Error reading trigger_weekday: %v", err)
		}
	}

	if err = d.Set("trigger_day", flattenSystemAutomationTriggerTriggerDay(o["trigger-day"], d, "trigger_day")); err != nil {
		if !fortiAPIPatch(o["trigger-day"]) {
			return fmt.Errorf("Error reading trigger_day: %v", err)
		}
	}

	if err = d.Set("trigger_hour", flattenSystemAutomationTriggerTriggerHour(o["trigger-hour"], d, "trigger_hour")); err != nil {
		if !fortiAPIPatch(o["trigger-hour"]) {
			return fmt.Errorf("Error reading trigger_hour: %v", err)
		}
	}

	if err = d.Set("trigger_minute", flattenSystemAutomationTriggerTriggerMinute(o["trigger-minute"], d, "trigger_minute")); err != nil {
		if !fortiAPIPatch(o["trigger-minute"]) {
			return fmt.Errorf("Error reading trigger_minute: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationTriggerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationTriggerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerEventType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerLicenseType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerIocLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerLogid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerWeekday(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerMinute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationTrigger(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAutomationTriggerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("trigger_type"); ok {
		t, err := expandSystemAutomationTriggerTriggerType(d, v, "trigger_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-type"] = t
		}
	}

	if v, ok := d.GetOk("event_type"); ok {
		t, err := expandSystemAutomationTriggerEventType(d, v, "event_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event-type"] = t
		}
	}

	if v, ok := d.GetOk("license_type"); ok {
		t, err := expandSystemAutomationTriggerLicenseType(d, v, "license_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["license-type"] = t
		}
	}

	if v, ok := d.GetOk("ioc_level"); ok {
		t, err := expandSystemAutomationTriggerIocLevel(d, v, "ioc_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ioc-level"] = t
		}
	}

	if v, ok := d.GetOk("logid"); ok {
		t, err := expandSystemAutomationTriggerLogid(d, v, "logid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logid"] = t
		}
	}

	if v, ok := d.GetOk("trigger_frequency"); ok {
		t, err := expandSystemAutomationTriggerTriggerFrequency(d, v, "trigger_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-frequency"] = t
		}
	}

	if v, ok := d.GetOk("trigger_weekday"); ok {
		t, err := expandSystemAutomationTriggerTriggerWeekday(d, v, "trigger_weekday")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-weekday"] = t
		}
	}

	if v, ok := d.GetOk("trigger_day"); ok {
		t, err := expandSystemAutomationTriggerTriggerDay(d, v, "trigger_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-day"] = t
		}
	}

	if v, ok := d.GetOkExists("trigger_hour"); ok {
		t, err := expandSystemAutomationTriggerTriggerHour(d, v, "trigger_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-hour"] = t
		}
	}

	if v, ok := d.GetOkExists("trigger_minute"); ok {
		t, err := expandSystemAutomationTriggerTriggerMinute(d, v, "trigger_minute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-minute"] = t
		}
	}

	return &obj, nil
}
