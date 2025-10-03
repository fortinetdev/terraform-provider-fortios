// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch ip source guard log.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerIpSourceGuardLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerIpSourceGuardLogUpdate,
		Read:   resourceSwitchControllerIpSourceGuardLogRead,
		Update: resourceSwitchControllerIpSourceGuardLogUpdate,
		Delete: resourceSwitchControllerIpSourceGuardLogDelete,

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
			"log_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"violation_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1500),
				Optional:     true,
			},
		},
	}
}

func resourceSwitchControllerIpSourceGuardLogUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerIpSourceGuardLog(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIpSourceGuardLog resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerIpSourceGuardLog(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIpSourceGuardLog resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerIpSourceGuardLog")
	}

	return resourceSwitchControllerIpSourceGuardLogRead(d, m)
}

func resourceSwitchControllerIpSourceGuardLogDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerIpSourceGuardLog(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerIpSourceGuardLog resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerIpSourceGuardLog(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerIpSourceGuardLog resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerIpSourceGuardLogRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerIpSourceGuardLog(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIpSourceGuardLog resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerIpSourceGuardLog(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerIpSourceGuardLog resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerIpSourceGuardLogLogViolations(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerIpSourceGuardLogViolationTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSwitchControllerIpSourceGuardLog(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("log_violations", flattenSwitchControllerIpSourceGuardLogLogViolations(o["log-violations"], d, "log_violations", sv)); err != nil {
		if !fortiAPIPatch(o["log-violations"]) {
			return fmt.Errorf("Error reading log_violations: %v", err)
		}
	}

	if err = d.Set("violation_timer", flattenSwitchControllerIpSourceGuardLogViolationTimer(o["violation-timer"], d, "violation_timer", sv)); err != nil {
		if !fortiAPIPatch(o["violation-timer"]) {
			return fmt.Errorf("Error reading violation_timer: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerIpSourceGuardLogFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerIpSourceGuardLogLogViolations(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerIpSourceGuardLogViolationTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerIpSourceGuardLog(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log_violations"); ok {
		if setArgNil {
			obj["log-violations"] = nil
		} else {
			t, err := expandSwitchControllerIpSourceGuardLogLogViolations(d, v, "log_violations", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-violations"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("violation_timer"); ok {
		if setArgNil {
			obj["violation-timer"] = nil
		} else {
			t, err := expandSwitchControllerIpSourceGuardLogViolationTimer(d, v, "violation_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["violation-timer"] = t
			}
		}
	} else if d.HasChange("violation_timer") {
		obj["violation-timer"] = nil
	}

	return &obj, nil
}
