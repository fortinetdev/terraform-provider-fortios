// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SNMP system info configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSnmpSysinfoRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"contact_info": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_high_cpu_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trap_low_memory_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trap_log_full_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemSnmpSysinfoRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemSnmpSysinfo"

	o, err := c.ReadSystemSnmpSysinfo(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpSysinfo: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemSnmpSysinfo(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpSysinfo from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemSnmpSysinfoStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpSysinfoEngineId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpSysinfoDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpSysinfoContactInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpSysinfoLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpSysinfoTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpSysinfoTrapLowMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpSysinfoTrapLogFullThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemSnmpSysinfo(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemSnmpSysinfoStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("engine_id", dataSourceFlattenSystemSnmpSysinfoEngineId(o["engine-id"], d, "engine_id")); err != nil {
		if !fortiAPIPatch(o["engine-id"]) {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}

	if err = d.Set("description", dataSourceFlattenSystemSnmpSysinfoDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("contact_info", dataSourceFlattenSystemSnmpSysinfoContactInfo(o["contact-info"], d, "contact_info")); err != nil {
		if !fortiAPIPatch(o["contact-info"]) {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}

	if err = d.Set("location", dataSourceFlattenSystemSnmpSysinfoLocation(o["location"], d, "location")); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("trap_high_cpu_threshold", dataSourceFlattenSystemSnmpSysinfoTrapHighCpuThreshold(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold")); err != nil {
		if !fortiAPIPatch(o["trap-high-cpu-threshold"]) {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_low_memory_threshold", dataSourceFlattenSystemSnmpSysinfoTrapLowMemoryThreshold(o["trap-low-memory-threshold"], d, "trap_low_memory_threshold")); err != nil {
		if !fortiAPIPatch(o["trap-low-memory-threshold"]) {
			return fmt.Errorf("Error reading trap_low_memory_threshold: %v", err)
		}
	}

	if err = d.Set("trap_log_full_threshold", dataSourceFlattenSystemSnmpSysinfoTrapLogFullThreshold(o["trap-log-full-threshold"], d, "trap_log_full_threshold")); err != nil {
		if !fortiAPIPatch(o["trap-log-full-threshold"]) {
			return fmt.Errorf("Error reading trap_log_full_threshold: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemSnmpSysinfoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
