// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Settings for Security Rating.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSecurityRatingSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSecurityRatingSettingsUpdate,
		Read:   resourceSystemSecurityRatingSettingsRead,
		Update: resourceSystemSecurityRatingSettingsUpdate,
		Delete: resourceSystemSecurityRatingSettingsDelete,

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
			"override_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSecurityRatingSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSecurityRatingSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSecurityRatingSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSecurityRatingSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSecurityRatingSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSecurityRatingSettings")
	}

	return resourceSystemSecurityRatingSettingsRead(d, m)
}

func resourceSystemSecurityRatingSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSecurityRatingSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemSecurityRatingSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSecurityRatingSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSecurityRatingSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSecurityRatingSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSecurityRatingSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSecurityRatingSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSecurityRatingSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSecurityRatingSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemSecurityRatingSettingsOverrideSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSecurityRatingSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("override_sync", flattenSystemSecurityRatingSettingsOverrideSync(o["override-sync"], d, "override_sync", sv)); err != nil {
		if !fortiAPIPatch(o["override-sync"]) {
			return fmt.Errorf("Error reading override_sync: %v", err)
		}
	}

	return nil
}

func flattenSystemSecurityRatingSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSecurityRatingSettingsOverrideSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSecurityRatingSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("override_sync"); ok {
		if setArgNil {
			obj["override-sync"] = nil
		} else {
			t, err := expandSystemSecurityRatingSettingsOverrideSync(d, v, "override_sync", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override-sync"] = t
			}
		}
	}

	return &obj, nil
}
