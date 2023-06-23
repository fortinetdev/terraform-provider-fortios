// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure network monitor settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"network_monitoring": &schema.Schema{
				Type:     schema.TypeString,
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerNetworkMonitorSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNetworkMonitorSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerNetworkMonitorSettings(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerNetworkMonitorSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNetworkMonitorSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerNetworkMonitorSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerNetworkMonitorSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerNetworkMonitorSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerNetworkMonitorSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNetworkMonitorSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerNetworkMonitorSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNetworkMonitorSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerNetworkMonitorSettingsNetworkMonitoring(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerNetworkMonitorSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("network_monitoring", flattenSwitchControllerNetworkMonitorSettingsNetworkMonitoring(o["network-monitoring"], d, "network_monitoring", sv)); err != nil {
		if !fortiAPIPatch(o["network-monitoring"]) {
			return fmt.Errorf("Error reading network_monitoring: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerNetworkMonitorSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerNetworkMonitorSettingsNetworkMonitoring(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerNetworkMonitorSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("network_monitoring"); ok {
		if setArgNil {
			obj["network-monitoring"] = nil
		} else {
			t, err := expandSwitchControllerNetworkMonitorSettingsNetworkMonitoring(d, v, "network_monitoring", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["network-monitoring"] = t
			}
		}
	}

	return &obj, nil
}
