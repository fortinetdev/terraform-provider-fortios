// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure network monitor settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerNetworkMonitorSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerNetworkMonitorSettingsUpdate,
		Read:   resourceSwitchControllerNetworkMonitorSettingsRead,
		Update: resourceSwitchControllerNetworkMonitorSettingsUpdate,
		Delete: resourceSwitchControllerNetworkMonitorSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"network_monitoring": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceSwitchControllerNetworkMonitorSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerNetworkMonitorSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNetworkMonitorSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerNetworkMonitorSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNetworkMonitorSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerNetworkMonitorSettings")
	}

	return resourceSwitchControllerNetworkMonitorSettingsRead(d, m)
}

func resourceSwitchControllerNetworkMonitorSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerNetworkMonitorSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerNetworkMonitorSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerNetworkMonitorSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerNetworkMonitorSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNetworkMonitorSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerNetworkMonitorSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNetworkMonitorSettings resource from API: %v", err)
	}
	return nil
}


func flattenSwitchControllerNetworkMonitorSettingsNetworkMonitoring(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSwitchControllerNetworkMonitorSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("network_monitoring", flattenSwitchControllerNetworkMonitorSettingsNetworkMonitoring(o["network-monitoring"], d, "network_monitoring")); err != nil {
		if !fortiAPIPatch(o["network-monitoring"]) {
			return fmt.Errorf("Error reading network_monitoring: %v", err)
		}
	}


	return nil
}

func flattenSwitchControllerNetworkMonitorSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSwitchControllerNetworkMonitorSettingsNetworkMonitoring(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSwitchControllerNetworkMonitorSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("network_monitoring"); ok {
		t, err := expandSwitchControllerNetworkMonitorSettingsNetworkMonitoring(d, v, "network_monitoring")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-monitoring"] = t
		}
	}


	return &obj, nil
}

