// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure global MAC synchronization settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerMacSyncSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerMacSyncSettingsUpdate,
		Read:   resourceSwitchControllerMacSyncSettingsRead,
		Update: resourceSwitchControllerMacSyncSettingsUpdate,
		Delete: resourceSwitchControllerMacSyncSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"mac_sync_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 1800),
				Optional:     true,
			},
		},
	}
}

func resourceSwitchControllerMacSyncSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSwitchControllerMacSyncSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerMacSyncSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerMacSyncSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerMacSyncSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerMacSyncSettings")
	}

	return resourceSwitchControllerMacSyncSettingsRead(d, m)
}

func resourceSwitchControllerMacSyncSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerMacSyncSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerMacSyncSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerMacSyncSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerMacSyncSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerMacSyncSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSwitchControllerMacSyncSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerMacSyncSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerMacSyncSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerMacSyncSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerMacSyncSettingsMacSyncInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSwitchControllerMacSyncSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mac_sync_interval", flattenSwitchControllerMacSyncSettingsMacSyncInterval(o["mac-sync-interval"], d, "mac_sync_interval", sv)); err != nil {
		if !fortiAPIPatch(o["mac-sync-interval"]) {
			return fmt.Errorf("Error reading mac_sync_interval: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerMacSyncSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerMacSyncSettingsMacSyncInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerMacSyncSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mac_sync_interval"); ok {
		if setArgNil {
			obj["mac-sync-interval"] = nil
		} else {
			t, err := expandSwitchControllerMacSyncSettingsMacSyncInterval(d, v, "mac_sync_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mac-sync-interval"] = t
			}
		}
	} else if d.HasChange("mac_sync_interval") {
		obj["mac-sync-interval"] = nil
	}

	return &obj, nil
}
