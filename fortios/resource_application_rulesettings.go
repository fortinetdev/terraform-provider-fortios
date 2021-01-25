// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure application rule settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceApplicationRuleSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationRuleSettingsCreate,
		Read:   resourceApplicationRuleSettingsRead,
		Update: resourceApplicationRuleSettingsUpdate,
		Delete: resourceApplicationRuleSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceApplicationRuleSettingsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectApplicationRuleSettings(d)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationRuleSettings resource while getting object: %v", err)
	}

	o, err := c.CreateApplicationRuleSettings(obj)

	if err != nil {
		return fmt.Errorf("Error creating ApplicationRuleSettings resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("ApplicationRuleSettings")
	}

	return resourceApplicationRuleSettingsRead(d, m)
}

func resourceApplicationRuleSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectApplicationRuleSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationRuleSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateApplicationRuleSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationRuleSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("ApplicationRuleSettings")
	}

	return resourceApplicationRuleSettingsRead(d, m)
}

func resourceApplicationRuleSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteApplicationRuleSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationRuleSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationRuleSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadApplicationRuleSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationRuleSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationRuleSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationRuleSettings resource from API: %v", err)
	}
	return nil
}

func flattenApplicationRuleSettingsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectApplicationRuleSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenApplicationRuleSettingsId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenApplicationRuleSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandApplicationRuleSettingsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationRuleSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandApplicationRuleSettingsId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
