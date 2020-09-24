// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure web proxy tunnelling for the FDN.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutoupdateTunneling() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoupdateTunnelingUpdate,
		Read:   resourceSystemAutoupdateTunnelingRead,
		Update: resourceSystemAutoupdateTunnelingUpdate,
		Delete: resourceSystemAutoupdateTunnelingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 49),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceSystemAutoupdateTunnelingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutoupdateTunneling(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateTunneling resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutoupdateTunneling(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdateTunneling resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutoupdateTunneling")
	}

	return resourceSystemAutoupdateTunnelingRead(d, m)
}

func resourceSystemAutoupdateTunnelingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAutoupdateTunneling(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoupdateTunneling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoupdateTunnelingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAutoupdateTunneling(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdateTunneling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoupdateTunneling(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdateTunneling resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoupdateTunnelingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdateTunnelingAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdateTunnelingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdateTunnelingUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoupdateTunnelingPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoupdateTunneling(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemAutoupdateTunnelingStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("address", flattenSystemAutoupdateTunnelingAddress(o["address"], d, "address")); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAutoupdateTunnelingPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemAutoupdateTunnelingUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoupdateTunnelingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutoupdateTunnelingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateTunnelingAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateTunnelingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateTunnelingUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdateTunnelingPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoupdateTunneling(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemAutoupdateTunnelingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {
		t, err := expandSystemAutoupdateTunnelingAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemAutoupdateTunnelingPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandSystemAutoupdateTunnelingUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemAutoupdateTunnelingPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	return &obj, nil
}
