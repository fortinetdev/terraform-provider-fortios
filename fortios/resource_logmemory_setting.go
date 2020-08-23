// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Settings for memory buffer.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogMemorySetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogMemorySettingUpdate,
		Read:   resourceLogMemorySettingRead,
		Update: resourceLogMemorySettingUpdate,
		Delete: resourceLogMemorySettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diskfull": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceLogMemorySettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogMemorySetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogMemorySetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogMemorySetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogMemorySetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogMemorySetting")
	}

	return resourceLogMemorySettingRead(d, m)
}

func resourceLogMemorySettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteLogMemorySetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting LogMemorySetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogMemorySettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogMemorySetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogMemorySetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogMemorySetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogMemorySetting resource from API: %v", err)
	}
	return nil
}


func flattenLogMemorySettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogMemorySettingDiskfull(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectLogMemorySetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("status", flattenLogMemorySettingStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("diskfull", flattenLogMemorySettingDiskfull(o["diskfull"], d, "diskfull")); err != nil {
		if !fortiAPIPatch(o["diskfull"]) {
			return fmt.Errorf("Error reading diskfull: %v", err)
		}
	}


	return nil
}

func flattenLogMemorySettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandLogMemorySettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogMemorySettingDiskfull(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectLogMemorySetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("status"); ok {
		t, err := expandLogMemorySettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("diskfull"); ok {
		t, err := expandLogMemorySettingDiskfull(d, v, "diskfull")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskfull"] = t
		}
	}


	return &obj, nil
}

