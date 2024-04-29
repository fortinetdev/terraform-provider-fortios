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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutoupdateSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoupdateScheduleUpdate,
		Read:   resourceSystemAutoupdateScheduleRead,
		Update: resourceSystemAutoupdateScheduleUpdate,
		Delete: resourceSystemAutoupdateScheduleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"frequency": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutoupdateScheduleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSystemAutoupdateSchedule(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateSchedule resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutoupdateSchedule(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateSchedule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutoupdateSchedule")
	}

	return resourceSystemAutoupdateScheduleRead(d, m)
}

func resourceSystemAutoupdateScheduleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemAutoupdateSchedule(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateSchedule resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoupdateSchedule(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemAutoupdateSchedule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoupdateScheduleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSystemAutoupdateSchedule(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdateSchedule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoupdateSchedule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdateSchedule resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoupdateScheduleStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoupdateScheduleFrequency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoupdateScheduleTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoupdateScheduleDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAutoupdateSchedule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemAutoupdateScheduleStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("frequency", flattenSystemAutoupdateScheduleFrequency(o["frequency"], d, "frequency", sv)); err != nil {
		if !fortiAPIPatch(o["frequency"]) {
			return fmt.Errorf("Error reading frequency: %v", err)
		}
	}

	if err = d.Set("time", flattenSystemAutoupdateScheduleTime(o["time"], d, "time", sv)); err != nil {
		if !fortiAPIPatch(o["time"]) {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	if err = d.Set("day", flattenSystemAutoupdateScheduleDay(o["day"], d, "day", sv)); err != nil {
		if !fortiAPIPatch(o["day"]) {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoupdateScheduleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAutoupdateScheduleStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateScheduleFrequency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateScheduleTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateScheduleDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoupdateSchedule(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemAutoupdateScheduleStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("frequency"); ok {
		if setArgNil {
			obj["frequency"] = nil
		} else {
			t, err := expandSystemAutoupdateScheduleFrequency(d, v, "frequency", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["frequency"] = t
			}
		}
	}

	if v, ok := d.GetOk("time"); ok {
		if setArgNil {
			obj["time"] = nil
		} else {
			t, err := expandSystemAutoupdateScheduleTime(d, v, "time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["time"] = t
			}
		}
	}

	if v, ok := d.GetOk("day"); ok {
		if setArgNil {
			obj["day"] = nil
		} else {
			t, err := expandSystemAutoupdateScheduleDay(d, v, "day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["day"] = t
			}
		}
	}

	return &obj, nil
}
