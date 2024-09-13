// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure AntiVirus settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusSettingsUpdate,
		Read:   resourceAntivirusSettingsRead,
		Update: resourceAntivirusSettingsUpdate,
		Delete: resourceAntivirusSettingsDelete,

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
			"machine_learning_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_extreme_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"grayware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 3600),
				Optional:     true,
			},
			"cache_infected_result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_clean_result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAntivirusSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAntivirusSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateAntivirusSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AntivirusSettings")
	}

	return resourceAntivirusSettingsRead(d, m)
}

func resourceAntivirusSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectAntivirusSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating AntivirusSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateAntivirusSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing AntivirusSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAntivirusSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusSettings resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusSettingsMachineLearningDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusSettingsUseExtremeDb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusSettingsDefaultDb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusSettingsGrayware(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusSettingsOverrideTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenAntivirusSettingsCacheInfectedResult(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusSettingsCacheCleanResult(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectAntivirusSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("machine_learning_detection", flattenAntivirusSettingsMachineLearningDetection(o["machine-learning-detection"], d, "machine_learning_detection", sv)); err != nil {
		if !fortiAPIPatch(o["machine-learning-detection"]) {
			return fmt.Errorf("Error reading machine_learning_detection: %v", err)
		}
	}

	if err = d.Set("use_extreme_db", flattenAntivirusSettingsUseExtremeDb(o["use-extreme-db"], d, "use_extreme_db", sv)); err != nil {
		if !fortiAPIPatch(o["use-extreme-db"]) {
			return fmt.Errorf("Error reading use_extreme_db: %v", err)
		}
	}

	if err = d.Set("default_db", flattenAntivirusSettingsDefaultDb(o["default-db"], d, "default_db", sv)); err != nil {
		if !fortiAPIPatch(o["default-db"]) {
			return fmt.Errorf("Error reading default_db: %v", err)
		}
	}

	if err = d.Set("grayware", flattenAntivirusSettingsGrayware(o["grayware"], d, "grayware", sv)); err != nil {
		if !fortiAPIPatch(o["grayware"]) {
			return fmt.Errorf("Error reading grayware: %v", err)
		}
	}

	if err = d.Set("override_timeout", flattenAntivirusSettingsOverrideTimeout(o["override-timeout"], d, "override_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["override-timeout"]) {
			return fmt.Errorf("Error reading override_timeout: %v", err)
		}
	}

	if err = d.Set("cache_infected_result", flattenAntivirusSettingsCacheInfectedResult(o["cache-infected-result"], d, "cache_infected_result", sv)); err != nil {
		if !fortiAPIPatch(o["cache-infected-result"]) {
			return fmt.Errorf("Error reading cache_infected_result: %v", err)
		}
	}

	if err = d.Set("cache_clean_result", flattenAntivirusSettingsCacheCleanResult(o["cache-clean-result"], d, "cache_clean_result", sv)); err != nil {
		if !fortiAPIPatch(o["cache-clean-result"]) {
			return fmt.Errorf("Error reading cache_clean_result: %v", err)
		}
	}

	return nil
}

func flattenAntivirusSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandAntivirusSettingsMachineLearningDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsUseExtremeDb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsDefaultDb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsGrayware(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsOverrideTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsCacheInfectedResult(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsCacheCleanResult(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("machine_learning_detection"); ok {
		if setArgNil {
			obj["machine-learning-detection"] = nil
		} else {
			t, err := expandAntivirusSettingsMachineLearningDetection(d, v, "machine_learning_detection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["machine-learning-detection"] = t
			}
		}
	}

	if v, ok := d.GetOk("use_extreme_db"); ok {
		if setArgNil {
			obj["use-extreme-db"] = nil
		} else {
			t, err := expandAntivirusSettingsUseExtremeDb(d, v, "use_extreme_db", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["use-extreme-db"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_db"); ok {
		if setArgNil {
			obj["default-db"] = nil
		} else {
			t, err := expandAntivirusSettingsDefaultDb(d, v, "default_db", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-db"] = t
			}
		}
	} else if d.HasChange("default_db") {
		obj["default-db"] = nil
	}

	if v, ok := d.GetOk("grayware"); ok {
		if setArgNil {
			obj["grayware"] = nil
		} else {
			t, err := expandAntivirusSettingsGrayware(d, v, "grayware", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["grayware"] = t
			}
		}
	}

	if v, ok := d.GetOk("override_timeout"); ok {
		if setArgNil {
			obj["override-timeout"] = nil
		} else {
			t, err := expandAntivirusSettingsOverrideTimeout(d, v, "override_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override-timeout"] = t
			}
		}
	} else if d.HasChange("override_timeout") {
		obj["override-timeout"] = nil
	}

	if v, ok := d.GetOk("cache_infected_result"); ok {
		if setArgNil {
			obj["cache-infected-result"] = nil
		} else {
			t, err := expandAntivirusSettingsCacheInfectedResult(d, v, "cache_infected_result", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cache-infected-result"] = t
			}
		}
	}

	if v, ok := d.GetOk("cache_clean_result"); ok {
		if setArgNil {
			obj["cache-clean-result"] = nil
		} else {
			t, err := expandAntivirusSettingsCacheCleanResult(d, v, "cache_clean_result", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cache-clean-result"] = t
			}
		}
	}

	return &obj, nil
}
