// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure integrated NAC settings for FortiSwitch.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerNacSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerNacSettingsCreate,
		Read:   resourceSwitchControllerNacSettingsRead,
		Update: resourceSwitchControllerNacSettingsUpdate,
		Delete: resourceSwitchControllerNacSettingsDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inactive_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1440),
				Optional:     true,
				Computed:     true,
			},
			"onboarding_vlan": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"auto_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bounce_nac_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_down_flush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerNacSettingsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerNacSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerNacSettings resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerNacSettings(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerNacSettings resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerNacSettings")
	}

	return resourceSwitchControllerNacSettingsRead(d, m)
}

func resourceSwitchControllerNacSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerNacSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNacSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerNacSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerNacSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerNacSettings")
	}

	return resourceSwitchControllerNacSettingsRead(d, m)
}

func resourceSwitchControllerNacSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerNacSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerNacSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerNacSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerNacSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNacSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerNacSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerNacSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerNacSettingsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacSettingsMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacSettingsInactiveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacSettingsOnboardingVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacSettingsAutoAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacSettingsBounceNacPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerNacSettingsLinkDownFlush(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerNacSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerNacSettingsName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("mode", flattenSwitchControllerNacSettingsMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("inactive_timer", flattenSwitchControllerNacSettingsInactiveTimer(o["inactive-timer"], d, "inactive_timer", sv)); err != nil {
		if !fortiAPIPatch(o["inactive-timer"]) {
			return fmt.Errorf("Error reading inactive_timer: %v", err)
		}
	}

	if err = d.Set("onboarding_vlan", flattenSwitchControllerNacSettingsOnboardingVlan(o["onboarding-vlan"], d, "onboarding_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["onboarding-vlan"]) {
			return fmt.Errorf("Error reading onboarding_vlan: %v", err)
		}
	}

	if err = d.Set("auto_auth", flattenSwitchControllerNacSettingsAutoAuth(o["auto-auth"], d, "auto_auth", sv)); err != nil {
		if !fortiAPIPatch(o["auto-auth"]) {
			return fmt.Errorf("Error reading auto_auth: %v", err)
		}
	}

	if err = d.Set("bounce_nac_port", flattenSwitchControllerNacSettingsBounceNacPort(o["bounce-nac-port"], d, "bounce_nac_port", sv)); err != nil {
		if !fortiAPIPatch(o["bounce-nac-port"]) {
			return fmt.Errorf("Error reading bounce_nac_port: %v", err)
		}
	}

	if err = d.Set("link_down_flush", flattenSwitchControllerNacSettingsLinkDownFlush(o["link-down-flush"], d, "link_down_flush", sv)); err != nil {
		if !fortiAPIPatch(o["link-down-flush"]) {
			return fmt.Errorf("Error reading link_down_flush: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerNacSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerNacSettingsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacSettingsMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacSettingsInactiveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacSettingsOnboardingVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacSettingsAutoAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacSettingsBounceNacPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerNacSettingsLinkDownFlush(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerNacSettings(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerNacSettingsName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSwitchControllerNacSettingsMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOkExists("inactive_timer"); ok {
		t, err := expandSwitchControllerNacSettingsInactiveTimer(d, v, "inactive_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inactive-timer"] = t
		}
	}

	if v, ok := d.GetOk("onboarding_vlan"); ok {
		t, err := expandSwitchControllerNacSettingsOnboardingVlan(d, v, "onboarding_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["onboarding-vlan"] = t
		}
	}

	if v, ok := d.GetOk("auto_auth"); ok {
		t, err := expandSwitchControllerNacSettingsAutoAuth(d, v, "auto_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-auth"] = t
		}
	}

	if v, ok := d.GetOk("bounce_nac_port"); ok {
		t, err := expandSwitchControllerNacSettingsBounceNacPort(d, v, "bounce_nac_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-nac-port"] = t
		}
	}

	if v, ok := d.GetOk("link_down_flush"); ok {
		t, err := expandSwitchControllerNacSettingsLinkDownFlush(d, v, "link_down_flush", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-flush"] = t
		}
	}

	return &obj, nil
}
