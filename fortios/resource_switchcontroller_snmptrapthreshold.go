// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch SNMP trap threshold values globally.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerSnmpTrapThreshold() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSnmpTrapThresholdUpdate,
		Read:   resourceSwitchControllerSnmpTrapThresholdRead,
		Update: resourceSwitchControllerSnmpTrapThresholdUpdate,
		Delete: resourceSwitchControllerSnmpTrapThresholdDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"trap_high_cpu_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_low_memory_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trap_log_full_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSnmpTrapThresholdUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerSnmpTrapThreshold(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpTrapThreshold resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSnmpTrapThreshold(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpTrapThreshold resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSnmpTrapThreshold")
	}

	return resourceSwitchControllerSnmpTrapThresholdRead(d, m)
}

func resourceSwitchControllerSnmpTrapThresholdDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerSnmpTrapThreshold(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpTrapThreshold resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSnmpTrapThreshold(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerSnmpTrapThreshold resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSnmpTrapThresholdRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerSnmpTrapThreshold(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpTrapThreshold resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSnmpTrapThreshold(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpTrapThreshold resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSnmpTrapThresholdTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSnmpTrapThresholdTrapLowMemoryThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSnmpTrapThresholdTrapLogFullThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerSnmpTrapThreshold(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("trap_high_cpu_threshold", flattenSwitchControllerSnmpTrapThresholdTrapHighCpuThreshold(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-high-cpu-threshold"]) {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_low_memory_threshold", flattenSwitchControllerSnmpTrapThresholdTrapLowMemoryThreshold(o["trap-low-memory-threshold"], d, "trap_low_memory_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-low-memory-threshold"]) {
			return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
		}
	}

	if err = d.Set("trap_log_full_threshold", flattenSwitchControllerSnmpTrapThresholdTrapLogFullThreshold(o["trap-log-full-threshold"], d, "trap_log_full_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-log-full-threshold"]) {
			return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSnmpTrapThresholdFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerSnmpTrapThresholdTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpTrapThresholdTrapLowMemoryThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpTrapThresholdTrapLogFullThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSnmpTrapThreshold(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("trap_high_cpu_threshold"); ok {
		if setArgNil {
			obj["trap-high-cpu-threshold"] = nil
		} else {

			t, err := expandSwitchControllerSnmpTrapThresholdTrapHighCpuThreshold(d, v, "trap_high_cpu_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trap-high-cpu-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("trap_low_memory_threshold"); ok {
		if setArgNil {
			obj["trap-low-memory-threshold"] = nil
		} else {

			t, err := expandSwitchControllerSnmpTrapThresholdTrapLowMemoryThreshold(d, v, "trap_low_memory_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trap-low-memory-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("trap_log_full_threshold"); ok {
		if setArgNil {
			obj["trap-log-full-threshold"] = nil
		} else {

			t, err := expandSwitchControllerSnmpTrapThresholdTrapLogFullThreshold(d, v, "trap_log_full_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trap-log-full-threshold"] = t
			}
		}
	}

	return &obj, nil
}
