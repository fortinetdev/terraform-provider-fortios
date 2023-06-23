// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch LLDP settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"device_detection": &schema.Schema{
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerLldpSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerLldpSettings(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerLldpSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerLldpSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerLldpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLldpSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerLldpSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLldpSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLldpSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsTxHold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsTxInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsFastStartInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsManagementInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpSettingsDeviceDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerLldpSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchControllerLldpSettingsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tx_hold", flattenSwitchControllerLldpSettingsTxHold(o["tx-hold"], d, "tx_hold", sv)); err != nil {
		if !fortiAPIPatch(o["tx-hold"]) {
			return fmt.Errorf("Error reading tx_hold: %v", err)
		}
	}

	if err = d.Set("tx_interval", flattenSwitchControllerLldpSettingsTxInterval(o["tx-interval"], d, "tx_interval", sv)); err != nil {
		if !fortiAPIPatch(o["tx-interval"]) {
			return fmt.Errorf("Error reading tx_interval: %v", err)
		}
	}

	if err = d.Set("fast_start_interval", flattenSwitchControllerLldpSettingsFastStartInterval(o["fast-start-interval"], d, "fast_start_interval", sv)); err != nil {
		if !fortiAPIPatch(o["fast-start-interval"]) {
			return fmt.Errorf("Error reading fast_start_interval: %v", err)
		}
	}

	if err = d.Set("management_interface", flattenSwitchControllerLldpSettingsManagementInterface(o["management-interface"], d, "management_interface", sv)); err != nil {
		if !fortiAPIPatch(o["management-interface"]) {
			return fmt.Errorf("Error reading management_interface: %v", err)
		}
	}

	if err = d.Set("device_detection", flattenSwitchControllerLldpSettingsDeviceDetection(o["device-detection"], d, "device_detection", sv)); err != nil {
		if !fortiAPIPatch(o["device-detection"]) {
			return fmt.Errorf("Error reading device_detection: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerLldpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerLldpSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsTxHold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsTxInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsFastStartInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsManagementInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpSettingsDeviceDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLldpSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSwitchControllerLldpSettingsStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("tx_hold"); ok {
		if setArgNil {
			obj["tx-hold"] = nil
		} else {
			t, err := expandSwitchControllerLldpSettingsTxHold(d, v, "tx_hold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tx-hold"] = t
			}
		}
	}

	if v, ok := d.GetOk("tx_interval"); ok {
		if setArgNil {
			obj["tx-interval"] = nil
		} else {
			t, err := expandSwitchControllerLldpSettingsTxInterval(d, v, "tx_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tx-interval"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("fast_start_interval"); ok {
		if setArgNil {
			obj["fast-start-interval"] = nil
		} else {
			t, err := expandSwitchControllerLldpSettingsFastStartInterval(d, v, "fast_start_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fast-start-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("management_interface"); ok {
		if setArgNil {
			obj["management-interface"] = nil
		} else {
			t, err := expandSwitchControllerLldpSettingsManagementInterface(d, v, "management_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["management-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("device_detection"); ok {
		if setArgNil {
			obj["device-detection"] = nil
		} else {
			t, err := expandSwitchControllerLldpSettingsDeviceDetection(d, v, "device_detection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["device-detection"] = t
			}
		}
	}

	return &obj, nil
}
