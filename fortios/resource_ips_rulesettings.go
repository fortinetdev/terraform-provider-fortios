// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPS rule setting.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIpsRuleSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsRuleSettingsCreate,
		Read:   resourceIpsRuleSettingsRead,
		Update: resourceIpsRuleSettingsUpdate,
		Delete: resourceIpsRuleSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIpsRuleSettingsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIpsRuleSettings(d)
	if err != nil {
		return fmt.Errorf("Error creating IpsRuleSettings resource while getting object: %v", err)
	}

	o, err := c.CreateIpsRuleSettings(obj)

	if err != nil {
		return fmt.Errorf("Error creating IpsRuleSettings resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("IpsRuleSettings")
	}

	return resourceIpsRuleSettingsRead(d, m)
}

func resourceIpsRuleSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIpsRuleSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsRuleSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsRuleSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating IpsRuleSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("IpsRuleSettings")
	}

	return resourceIpsRuleSettingsRead(d, m)
}

func resourceIpsRuleSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteIpsRuleSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting IpsRuleSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsRuleSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadIpsRuleSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading IpsRuleSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsRuleSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsRuleSettings resource from API: %v", err)
	}
	return nil
}

func flattenIpsRuleSettingsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsRuleSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenIpsRuleSettingsId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenIpsRuleSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsRuleSettingsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsRuleSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandIpsRuleSettingsId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
