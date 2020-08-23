// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure push updates.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutoupdatePushUpdate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoupdatePushUpdateUpdate,
		Read:   resourceSystemAutoupdatePushUpdateRead,
		Update: resourceSystemAutoupdatePushUpdateUpdate,
		Delete: resourceSystemAutoupdatePushUpdateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"override": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"address": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Required: true,
			},
		},
	}
}


func resourceSystemAutoupdatePushUpdateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutoupdatePushUpdate(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdatePushUpdate resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutoupdatePushUpdate(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdatePushUpdate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutoupdatePushUpdate")
	}

	return resourceSystemAutoupdatePushUpdateRead(d, m)
}

func resourceSystemAutoupdatePushUpdateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAutoupdatePushUpdate(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoupdatePushUpdate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoupdatePushUpdateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAutoupdatePushUpdate(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdatePushUpdate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoupdatePushUpdate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdatePushUpdate resource from API: %v", err)
	}
	return nil
}


func flattenSystemAutoupdatePushUpdateStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdatePushUpdateOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdatePushUpdateAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdatePushUpdatePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemAutoupdatePushUpdate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("status", flattenSystemAutoupdatePushUpdateStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("override", flattenSystemAutoupdatePushUpdateOverride(o["override"], d, "override")); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("address", flattenSystemAutoupdatePushUpdateAddress(o["address"], d, "address")); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAutoupdatePushUpdatePort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}


	return nil
}

func flattenSystemAutoupdatePushUpdateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemAutoupdatePushUpdateStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdatePushUpdateOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdatePushUpdateAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdatePushUpdatePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemAutoupdatePushUpdate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemAutoupdatePushUpdateStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok {
		t, err := expandSystemAutoupdatePushUpdateOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {
		t, err := expandSystemAutoupdatePushUpdateAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemAutoupdatePushUpdatePort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}


	return &obj, nil
}

