// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure user password policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPasswordPolicyCreate,
		Read:   resourceUserPasswordPolicyRead,
		Update: resourceUserPasswordPolicyUpdate,
		Delete: resourceUserPasswordPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"expire_days": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 999),
				Optional:     true,
				Computed:     true,
			},
			"warn_days": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 30),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceUserPasswordPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating UserPasswordPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateUserPasswordPolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating UserPasswordPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserPasswordPolicy")
	}

	return resourceUserPasswordPolicyRead(d, m)
}

func resourceUserPasswordPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserPasswordPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating UserPasswordPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateUserPasswordPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserPasswordPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserPasswordPolicy")
	}

	return resourceUserPasswordPolicyRead(d, m)
}

func resourceUserPasswordPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteUserPasswordPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserPasswordPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserPasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserPasswordPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserPasswordPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserPasswordPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserPasswordPolicy resource from API: %v", err)
	}
	return nil
}

func flattenUserPasswordPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyExpireDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPasswordPolicyWarnDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserPasswordPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenUserPasswordPolicyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("expire_days", flattenUserPasswordPolicyExpireDays(o["expire-days"], d, "expire_days")); err != nil {
		if !fortiAPIPatch(o["expire-days"]) {
			return fmt.Errorf("Error reading expire_days: %v", err)
		}
	}

	if err = d.Set("warn_days", flattenUserPasswordPolicyWarnDays(o["warn-days"], d, "warn_days")); err != nil {
		if !fortiAPIPatch(o["warn-days"]) {
			return fmt.Errorf("Error reading warn_days: %v", err)
		}
	}

	return nil
}

func flattenUserPasswordPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserPasswordPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyExpireDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyWarnDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserPasswordPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserPasswordPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("expire_days"); ok {
		t, err := expandUserPasswordPolicyExpireDays(d, v, "expire_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-days"] = t
		}
	}

	if v, ok := d.GetOkExists("warn_days"); ok {
		t, err := expandUserPasswordPolicyWarnDays(d, v, "warn_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warn-days"] = t
		}
	}

	return &obj, nil
}
