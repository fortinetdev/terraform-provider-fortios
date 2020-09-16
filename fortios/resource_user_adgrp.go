// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FSSO groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserAdgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserAdgrpCreate,
		Read:   resourceUserAdgrpRead,
		Update: resourceUserAdgrpUpdate,
		Delete: resourceUserAdgrpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required:     true,
			},
			"server_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceUserAdgrpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserAdgrp(d)
	if err != nil {
		return fmt.Errorf("Error creating UserAdgrp resource while getting object: %v", err)
	}

	o, err := c.CreateUserAdgrp(obj)

	if err != nil {
		return fmt.Errorf("Error creating UserAdgrp resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserAdgrp")
	}

	return resourceUserAdgrpRead(d, m)
}

func resourceUserAdgrpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserAdgrp(d)
	if err != nil {
		return fmt.Errorf("Error updating UserAdgrp resource while getting object: %v", err)
	}

	o, err := c.UpdateUserAdgrp(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserAdgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserAdgrp")
	}

	return resourceUserAdgrpRead(d, m)
}

func resourceUserAdgrpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteUserAdgrp(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserAdgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserAdgrpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserAdgrp(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserAdgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserAdgrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserAdgrp resource from API: %v", err)
	}
	return nil
}

func flattenUserAdgrpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserAdgrpServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserAdgrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenUserAdgrpName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_name", flattenUserAdgrpServerName(o["server-name"], d, "server_name")); err != nil {
		if !fortiAPIPatch(o["server-name"]) {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	return nil
}

func flattenUserAdgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserAdgrpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserAdgrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserAdgrpName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok {
		t, err := expandUserAdgrpServerName(d, v, "server_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	return &obj, nil
}
