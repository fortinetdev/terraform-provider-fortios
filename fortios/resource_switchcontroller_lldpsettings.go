// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch LLDP settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerLldpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLldpSettingsUpdate,
		Read:   resourceSwitchControllerLldpSettingsRead,
		Update: resourceSwitchControllerLldpSettingsUpdate,
		Delete: resourceSwitchControllerLldpSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tx_hold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16),
				Optional:     true,
				Computed:     true,
			},
			"tx_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 4095),
				Optional:     true,
				Computed:     true,
			},
			"fast_start_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"management_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerLldpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerLldpSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerLldpSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerLldpSettings")
	}

	return resourceSwitchControllerLldpSettingsRead(d, m)
}

func resourceSwitchControllerLldpSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerLldpSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLldpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLldpSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerLldpSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLldpSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLldpSettingsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsTxHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsTxInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsFastStartInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsManagementInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerLldpSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSwitchControllerLldpSettingsStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tx_hold", flattenSwitchControllerLldpSettingsTxHold(o["tx-hold"], d, "tx_hold")); err != nil {
		if !fortiAPIPatch(o["tx-hold"]) {
			return fmt.Errorf("Error reading tx_hold: %v", err)
		}
	}

	if err = d.Set("tx_interval", flattenSwitchControllerLldpSettingsTxInterval(o["tx-interval"], d, "tx_interval")); err != nil {
		if !fortiAPIPatch(o["tx-interval"]) {
			return fmt.Errorf("Error reading tx_interval: %v", err)
		}
	}

	if err = d.Set("fast_start_interval", flattenSwitchControllerLldpSettingsFastStartInterval(o["fast-start-interval"], d, "fast_start_interval")); err != nil {
		if !fortiAPIPatch(o["fast-start-interval"]) {
			return fmt.Errorf("Error reading fast_start_interval: %v", err)
		}
	}

	if err = d.Set("management_interface", flattenSwitchControllerLldpSettingsManagementInterface(o["management-interface"], d, "management_interface")); err != nil {
		if !fortiAPIPatch(o["management-interface"]) {
			return fmt.Errorf("Error reading management_interface: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerLldpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerLldpSettingsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsTxHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsTxInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsFastStartInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsManagementInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLldpSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSwitchControllerLldpSettingsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tx_hold"); ok {
		t, err := expandSwitchControllerLldpSettingsTxHold(d, v, "tx_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tx-hold"] = t
		}
	}

	if v, ok := d.GetOk("tx_interval"); ok {
		t, err := expandSwitchControllerLldpSettingsTxInterval(d, v, "tx_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tx-interval"] = t
		}
	}

	if v, ok := d.GetOk("fast_start_interval"); ok {
		t, err := expandSwitchControllerLldpSettingsFastStartInterval(d, v, "fast_start_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-start-interval"] = t
		}
	}

	if v, ok := d.GetOk("management_interface"); ok {
		t, err := expandSwitchControllerLldpSettingsManagementInterface(d, v, "management_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-interface"] = t
		}
	}

	return &obj, nil
}
