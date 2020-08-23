// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: SNMP system info configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpSysinfoUpdate,
		Read:   resourceSystemSnmpSysinfoRead,
		Update: resourceSystemSnmpSysinfoUpdate,
		Delete: resourceSystemSnmpSysinfoDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"engine_id": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 24),
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
			},
			"contact_info": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
			},
			"location": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
			},
			"trap_high_cpu_threshold": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional: true,
				Computed: true,
			},
			"trap_low_memory_threshold": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional: true,
				Computed: true,
			},
			"trap_log_full_threshold": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceSystemSnmpSysinfoUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSnmpSysinfo(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSnmpSysinfo(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSnmpSysinfo")
	}

	return resourceSystemSnmpSysinfoRead(d, m)
}

func resourceSystemSnmpSysinfoDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemSnmpSysinfo(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpSysinfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpSysinfoRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemSnmpSysinfo(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpSysinfo(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource from API: %v", err)
	}
	return nil
}


func flattenSystemSnmpSysinfoStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoEngineId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoContactInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapLowMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpSysinfoTrapLogFullThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemSnmpSysinfo(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("status", flattenSystemSnmpSysinfoStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("engine_id", flattenSystemSnmpSysinfoEngineId(o["engine-id"], d, "engine_id")); err != nil {
		if !fortiAPIPatch(o["engine-id"]) {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemSnmpSysinfoDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("contact_info", flattenSystemSnmpSysinfoContactInfo(o["contact-info"], d, "contact_info")); err != nil {
		if !fortiAPIPatch(o["contact-info"]) {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}

	if err = d.Set("location", flattenSystemSnmpSysinfoLocation(o["location"], d, "location")); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("trap_high_cpu_threshold", flattenSystemSnmpSysinfoTrapHighCpuThreshold(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold")); err != nil {
		if !fortiAPIPatch(o["trap-high-cpu-threshold"]) {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_low_memory_threshold", flattenSystemSnmpSysinfoTrapLowMemoryThreshold(o["trap-low-memory-threshold"], d, "trap_low_memory_threshold")); err != nil {
		if !fortiAPIPatch(o["trap-low-memory-threshold"]) {
			return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
		}
	}

	if err = d.Set("trap_log_full_threshold", flattenSystemSnmpSysinfoTrapLogFullThreshold(o["trap-log-full-threshold"], d, "trap_log_full_threshold")); err != nil {
		if !fortiAPIPatch(o["trap-log-full-threshold"]) {
			return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
		}
	}


	return nil
}

func flattenSystemSnmpSysinfoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemSnmpSysinfoStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoEngineId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoContactInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapLowMemoryThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpSysinfoTrapLogFullThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemSnmpSysinfo(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSnmpSysinfoStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("engine_id"); ok {
		t, err := expandSystemSnmpSysinfoEngineId(d, v, "engine_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-id"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemSnmpSysinfoDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("contact_info"); ok {
		t, err := expandSystemSnmpSysinfoContactInfo(d, v, "contact_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact-info"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok {
		t, err := expandSystemSnmpSysinfoLocation(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("trap_high_cpu_threshold"); ok {
		t, err := expandSystemSnmpSysinfoTrapHighCpuThreshold(d, v, "trap_high_cpu_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-high-cpu-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_low_memory_threshold"); ok {
		t, err := expandSystemSnmpSysinfoTrapLowMemoryThreshold(d, v, "trap_low_memory_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-low-memory-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_log_full_threshold"); ok {
		t, err := expandSystemSnmpSysinfoTrapLogFullThreshold(d, v, "trap_log_full_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-log-full-threshold"] = t
		}
	}


	return &obj, nil
}

