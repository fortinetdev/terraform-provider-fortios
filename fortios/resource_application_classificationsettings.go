// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure app classification setting.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationClassificationSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationClassificationSettingsUpdate,
		Read:   resourceApplicationClassificationSettingsRead,
		Update: resourceApplicationClassificationSettingsUpdate,
		Delete: resourceApplicationClassificationSettingsDelete,

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
			"default_app_classification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceApplicationClassificationSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationClassificationSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationClassificationSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateApplicationClassificationSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationClassificationSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ApplicationClassificationSettings")
	}

	return resourceApplicationClassificationSettingsRead(d, m)
}

func resourceApplicationClassificationSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectApplicationClassificationSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating ApplicationClassificationSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateApplicationClassificationSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing ApplicationClassificationSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationClassificationSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadApplicationClassificationSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationClassificationSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationClassificationSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationClassificationSettings resource from API: %v", err)
	}
	return nil
}

func flattenApplicationClassificationSettingsDefaultAppClassification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectApplicationClassificationSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("default_app_classification", flattenApplicationClassificationSettingsDefaultAppClassification(o["default-app-classification"], d, "default_app_classification", sv)); err != nil {
		if !fortiAPIPatch(o["default-app-classification"]) {
			return fmt.Errorf("Error reading default_app_classification: %v", err)
		}
	}

	return nil
}

func flattenApplicationClassificationSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandApplicationClassificationSettingsDefaultAppClassification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationClassificationSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_app_classification"); ok {
		if setArgNil {
			obj["default-app-classification"] = nil
		} else {
			t, err := expandApplicationClassificationSettingsDefaultAppClassification(d, v, "default_app_classification", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-app-classification"] = t
			}
		}
	}

	return &obj, nil
}
