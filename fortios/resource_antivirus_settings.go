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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"default_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed:     true,
			},
		},
	}
}

func resourceAntivirusSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectAntivirusSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateAntivirusSettings(obj, mkey)
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

	err := c.DeleteAntivirusSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadAntivirusSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusSettings resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusSettingsDefaultDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusSettingsGrayware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusSettingsOverrideTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_db", flattenAntivirusSettingsDefaultDb(o["default-db"], d, "default_db")); err != nil {
		if !fortiAPIPatch(o["default-db"]) {
			return fmt.Errorf("Error reading default_db: %v", err)
		}
	}

	if err = d.Set("grayware", flattenAntivirusSettingsGrayware(o["grayware"], d, "grayware")); err != nil {
		if !fortiAPIPatch(o["grayware"]) {
			return fmt.Errorf("Error reading grayware: %v", err)
		}
	}

	if err = d.Set("override_timeout", flattenAntivirusSettingsOverrideTimeout(o["override-timeout"], d, "override_timeout")); err != nil {
		if !fortiAPIPatch(o["override-timeout"]) {
			return fmt.Errorf("Error reading override_timeout: %v", err)
		}
	}

	return nil
}

func flattenAntivirusSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusSettingsDefaultDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsGrayware(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusSettingsOverrideTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_db"); ok {
		t, err := expandAntivirusSettingsDefaultDb(d, v, "default_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-db"] = t
		}
	}

	if v, ok := d.GetOk("grayware"); ok {
		t, err := expandAntivirusSettingsGrayware(d, v, "grayware")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["grayware"] = t
		}
	}

	if v, ok := d.GetOk("override_timeout"); ok {
		t, err := expandAntivirusSettingsOverrideTimeout(d, v, "override_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-timeout"] = t
		}
	}

	return &obj, nil
}
