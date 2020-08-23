// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Settings for null device logging.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogNullDeviceSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogNullDeviceSettingUpdate,
		Read:   resourceLogNullDeviceSettingRead,
		Update: resourceLogNullDeviceSettingUpdate,
		Delete: resourceLogNullDeviceSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}


func resourceLogNullDeviceSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogNullDeviceSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogNullDeviceSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogNullDeviceSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogNullDeviceSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogNullDeviceSetting")
	}

	return resourceLogNullDeviceSettingRead(d, m)
}

func resourceLogNullDeviceSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteLogNullDeviceSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting LogNullDeviceSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogNullDeviceSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogNullDeviceSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogNullDeviceSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogNullDeviceSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogNullDeviceSetting resource from API: %v", err)
	}
	return nil
}


func flattenLogNullDeviceSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectLogNullDeviceSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("status", flattenLogNullDeviceSettingStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}


	return nil
}

func flattenLogNullDeviceSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandLogNullDeviceSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectLogNullDeviceSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("status"); ok {
		t, err := expandLogNullDeviceSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}


	return &obj, nil
}

