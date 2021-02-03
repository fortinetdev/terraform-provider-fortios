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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemHaMonitor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemHaMonitorRead,
		Schema: map[string]*schema.Schema{
			"monitor_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vlan_hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemHaMonitorRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemHaMonitor"

	o, err := c.ReadSystemHaMonitor(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemHaMonitor: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemHaMonitor(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemHaMonitor from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemHaMonitorMonitorVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMonitorVlanHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMonitorVlanHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemHaMonitor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("monitor_vlan", dataSourceFlattenSystemHaMonitorMonitorVlan(o["monitor-vlan"], d, "monitor_vlan")); err != nil {
		if !fortiAPIPatch(o["monitor-vlan"]) {
			return fmt.Errorf("Error reading monitor_vlan: %v", err)
		}
	}

	if err = d.Set("vlan_hb_interval", dataSourceFlattenSystemHaMonitorVlanHbInterval(o["vlan-hb-interval"], d, "vlan_hb_interval")); err != nil {
		if !fortiAPIPatch(o["vlan-hb-interval"]) {
			return fmt.Errorf("Error reading vlan_hb_interval: %v", err)
		}
	}

	if err = d.Set("vlan_hb_lost_threshold", dataSourceFlattenSystemHaMonitorVlanHbLostThreshold(o["vlan-hb-lost-threshold"], d, "vlan_hb_lost_threshold")); err != nil {
		if !fortiAPIPatch(o["vlan-hb-lost-threshold"]) {
			return fmt.Errorf("Error reading vlan_hb_lost_threshold: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemHaMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
