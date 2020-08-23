// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch spanning tree protocol (STP).

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerStpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerStpSettingsUpdate,
		Read:   resourceSwitchControllerStpSettingsRead,
		Update: resourceSwitchControllerStpSettingsUpdate,
		Delete: resourceSwitchControllerStpSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revision": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"hello_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"forward_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(4, 30),
				Optional:     true,
				Computed:     true,
			},
			"max_age": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(6, 40),
				Optional:     true,
				Computed:     true,
			},
			"max_hops": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 40),
				Optional:     true,
				Computed:     true,
			},
			"pending_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchControllerStpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerStpSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStpSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerStpSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerStpSettings")
	}

	return resourceSwitchControllerStpSettingsRead(d, m)
}

func resourceSwitchControllerStpSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerStpSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerStpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStpSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerStpSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerStpSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerStpSettingsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsRevision(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsHelloTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsForwardTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsMaxAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsMaxHops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsPendingTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerStpSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerStpSettingsName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerStpSettingsStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("revision", flattenSwitchControllerStpSettingsRevision(o["revision"], d, "revision")); err != nil {
		if !fortiAPIPatch(o["revision"]) {
			return fmt.Errorf("Error reading revision: %v", err)
		}
	}

	if err = d.Set("hello_time", flattenSwitchControllerStpSettingsHelloTime(o["hello-time"], d, "hello_time")); err != nil {
		if !fortiAPIPatch(o["hello-time"]) {
			return fmt.Errorf("Error reading hello_time: %v", err)
		}
	}

	if err = d.Set("forward_time", flattenSwitchControllerStpSettingsForwardTime(o["forward-time"], d, "forward_time")); err != nil {
		if !fortiAPIPatch(o["forward-time"]) {
			return fmt.Errorf("Error reading forward_time: %v", err)
		}
	}

	if err = d.Set("max_age", flattenSwitchControllerStpSettingsMaxAge(o["max-age"], d, "max_age")); err != nil {
		if !fortiAPIPatch(o["max-age"]) {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("max_hops", flattenSwitchControllerStpSettingsMaxHops(o["max-hops"], d, "max_hops")); err != nil {
		if !fortiAPIPatch(o["max-hops"]) {
			return fmt.Errorf("Error reading max_hops: %v", err)
		}
	}

	if err = d.Set("pending_timer", flattenSwitchControllerStpSettingsPendingTimer(o["pending-timer"], d, "pending_timer")); err != nil {
		if !fortiAPIPatch(o["pending-timer"]) {
			return fmt.Errorf("Error reading pending_timer: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerStpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerStpSettingsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsRevision(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsHelloTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsForwardTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsMaxAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsMaxHops(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsPendingTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerStpSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerStpSettingsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSwitchControllerStpSettingsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("revision"); ok {
		t, err := expandSwitchControllerStpSettingsRevision(d, v, "revision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision"] = t
		}
	}

	if v, ok := d.GetOk("hello_time"); ok {
		t, err := expandSwitchControllerStpSettingsHelloTime(d, v, "hello_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-time"] = t
		}
	}

	if v, ok := d.GetOk("forward_time"); ok {
		t, err := expandSwitchControllerStpSettingsForwardTime(d, v, "forward_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-time"] = t
		}
	}

	if v, ok := d.GetOk("max_age"); ok {
		t, err := expandSwitchControllerStpSettingsMaxAge(d, v, "max_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-age"] = t
		}
	}

	if v, ok := d.GetOk("max_hops"); ok {
		t, err := expandSwitchControllerStpSettingsMaxHops(d, v, "max_hops")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-hops"] = t
		}
	}

	if v, ok := d.GetOk("pending_timer"); ok {
		t, err := expandSwitchControllerStpSettingsPendingTimer(d, v, "pending_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pending-timer"] = t
		}
	}

	return &obj, nil
}
