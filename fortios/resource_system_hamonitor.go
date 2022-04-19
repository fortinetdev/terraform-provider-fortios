// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure HA monitor.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemHaMonitor(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitor resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemHaMonitor(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemHaMonitor(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitor resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaMonitor(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemHaMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaMonitorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemHaMonitor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaMonitor(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitor resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaMonitorMonitorVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMonitorVlanHbInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMonitorVlanHbLostThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemHaMonitor(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("monitor_vlan", flattenSystemHaMonitorMonitorVlan(o["monitor-vlan"], d, "monitor_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-vlan"]) {
			return fmt.Errorf("Error reading monitor_vlan: %v", err)
		}
	}

	if err = d.Set("vlan_hb_interval", flattenSystemHaMonitorVlanHbInterval(o["vlan-hb-interval"], d, "vlan_hb_interval", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-hb-interval"]) {
			return fmt.Errorf("Error reading vlan_hb_interval: %v", err)
		}
	}

	if err = d.Set("vlan_hb_lost_threshold", flattenSystemHaMonitorVlanHbLostThreshold(o["vlan-hb-lost-threshold"], d, "vlan_hb_lost_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-hb-lost-threshold"]) {
			return fmt.Errorf("Error reading vlan_hb_lost_threshold: %v", err)
		}
	}

	return nil
}

func flattenSystemHaMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemHaMonitorMonitorVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitorVlanHbInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitorVlanHbLostThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaMonitor(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("monitor_vlan"); ok {
		if setArgNil {
			obj["monitor-vlan"] = nil
		} else {

			t, err := expandSystemHaMonitorMonitorVlan(d, v, "monitor_vlan", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["monitor-vlan"] = t
			}
		}
	}

	if v, ok := d.GetOk("vlan_hb_interval"); ok {
		if setArgNil {
			obj["vlan-hb-interval"] = nil
		} else {

			t, err := expandSystemHaMonitorVlanHbInterval(d, v, "vlan_hb_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vlan-hb-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("vlan_hb_lost_threshold"); ok {
		if setArgNil {
			obj["vlan-hb-lost-threshold"] = nil
		} else {

			t, err := expandSystemHaMonitorVlanHbLostThreshold(d, v, "vlan_hb_lost_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vlan-hb-lost-threshold"] = t
			}
		}
	}

	return &obj, nil
}
