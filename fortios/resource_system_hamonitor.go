// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure HA monitor.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemHaMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaMonitorUpdate,
		Read:   resourceSystemHaMonitorRead,
		Update: resourceSystemHaMonitorUpdate,
		Delete: resourceSystemHaMonitorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"monitor_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_hb_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"vlan_hb_lost_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemHaMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemHaMonitor(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitor resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemHaMonitor(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemHaMonitor")
	}

	return resourceSystemHaMonitorRead(d, m)
}

func resourceSystemHaMonitorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemHaMonitor(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaMonitorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemHaMonitor(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaMonitor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitor resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaMonitorMonitorVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitorVlanHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitorVlanHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHaMonitor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("monitor_vlan", flattenSystemHaMonitorMonitorVlan(o["monitor-vlan"], d, "monitor_vlan")); err != nil {
		if !fortiAPIPatch(o["monitor-vlan"]) {
			return fmt.Errorf("Error reading monitor_vlan: %v", err)
		}
	}

	if err = d.Set("vlan_hb_interval", flattenSystemHaMonitorVlanHbInterval(o["vlan-hb-interval"], d, "vlan_hb_interval")); err != nil {
		if !fortiAPIPatch(o["vlan-hb-interval"]) {
			return fmt.Errorf("Error reading vlan_hb_interval: %v", err)
		}
	}

	if err = d.Set("vlan_hb_lost_threshold", flattenSystemHaMonitorVlanHbLostThreshold(o["vlan-hb-lost-threshold"], d, "vlan_hb_lost_threshold")); err != nil {
		if !fortiAPIPatch(o["vlan-hb-lost-threshold"]) {
			return fmt.Errorf("Error reading vlan_hb_lost_threshold: %v", err)
		}
	}

	return nil
}

func flattenSystemHaMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaMonitorMonitorVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitorVlanHbInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitorVlanHbLostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaMonitor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("monitor_vlan"); ok {
		t, err := expandSystemHaMonitorMonitorVlan(d, v, "monitor_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-vlan"] = t
		}
	}

	if v, ok := d.GetOk("vlan_hb_interval"); ok {
		t, err := expandSystemHaMonitorVlanHbInterval(d, v, "vlan_hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("vlan_hb_lost_threshold"); ok {
		t, err := expandSystemHaMonitorVlanHbLostThreshold(d, v, "vlan_hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-hb-lost-threshold"] = t
		}
	}

	return &obj, nil
}
