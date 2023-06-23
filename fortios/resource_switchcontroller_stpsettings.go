// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch spanning tree protocol (STP).

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerStpSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStpSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerStpSettings(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerStpSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerStpSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerStpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStpSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerStpSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerStpSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerStpSettingsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsRevision(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsHelloTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsForwardTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsMaxAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsMaxHops(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStpSettingsPendingTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerStpSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerStpSettingsName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSwitchControllerStpSettingsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("revision", flattenSwitchControllerStpSettingsRevision(o["revision"], d, "revision", sv)); err != nil {
		if !fortiAPIPatch(o["revision"]) {
			return fmt.Errorf("Error reading revision: %v", err)
		}
	}

	if err = d.Set("hello_time", flattenSwitchControllerStpSettingsHelloTime(o["hello-time"], d, "hello_time", sv)); err != nil {
		if !fortiAPIPatch(o["hello-time"]) {
			return fmt.Errorf("Error reading hello_time: %v", err)
		}
	}

	if err = d.Set("forward_time", flattenSwitchControllerStpSettingsForwardTime(o["forward-time"], d, "forward_time", sv)); err != nil {
		if !fortiAPIPatch(o["forward-time"]) {
			return fmt.Errorf("Error reading forward_time: %v", err)
		}
	}

	if err = d.Set("max_age", flattenSwitchControllerStpSettingsMaxAge(o["max-age"], d, "max_age", sv)); err != nil {
		if !fortiAPIPatch(o["max-age"]) {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("max_hops", flattenSwitchControllerStpSettingsMaxHops(o["max-hops"], d, "max_hops", sv)); err != nil {
		if !fortiAPIPatch(o["max-hops"]) {
			return fmt.Errorf("Error reading max_hops: %v", err)
		}
	}

	if err = d.Set("pending_timer", flattenSwitchControllerStpSettingsPendingTimer(o["pending-timer"], d, "pending_timer", sv)); err != nil {
		if !fortiAPIPatch(o["pending-timer"]) {
			return fmt.Errorf("Error reading pending_timer: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerStpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerStpSettingsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsRevision(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsHelloTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsForwardTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsMaxAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsMaxHops(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStpSettingsPendingTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerStpSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		if setArgNil {
			obj["name"] = nil
		} else {
			t, err := expandSwitchControllerStpSettingsName(d, v, "name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["name"] = t
			}
		}
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSwitchControllerStpSettingsStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("revision"); ok {
		if setArgNil {
			obj["revision"] = nil
		} else {
			t, err := expandSwitchControllerStpSettingsRevision(d, v, "revision", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["revision"] = t
			}
		}
	}

	if v, ok := d.GetOk("hello_time"); ok {
		if setArgNil {
			obj["hello-time"] = nil
		} else {
			t, err := expandSwitchControllerStpSettingsHelloTime(d, v, "hello_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hello-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("forward_time"); ok {
		if setArgNil {
			obj["forward-time"] = nil
		} else {
			t, err := expandSwitchControllerStpSettingsForwardTime(d, v, "forward_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forward-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_age"); ok {
		if setArgNil {
			obj["max-age"] = nil
		} else {
			t, err := expandSwitchControllerStpSettingsMaxAge(d, v, "max_age", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-age"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_hops"); ok {
		if setArgNil {
			obj["max-hops"] = nil
		} else {
			t, err := expandSwitchControllerStpSettingsMaxHops(d, v, "max_hops", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-hops"] = t
			}
		}
	}

	if v, ok := d.GetOk("pending_timer"); ok {
		if setArgNil {
			obj["pending-timer"] = nil
		} else {
			t, err := expandSwitchControllerStpSettingsPendingTimer(d, v, "pending_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pending-timer"] = t
			}
		}
	}

	return &obj, nil
}
