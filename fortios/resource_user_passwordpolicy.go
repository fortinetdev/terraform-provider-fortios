// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure user password policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"expired_password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserPasswordPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserPasswordPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserPasswordPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateUserPasswordPolicy(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserPasswordPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserPasswordPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateUserPasswordPolicy(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserPasswordPolicy(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserPasswordPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserPasswordPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserPasswordPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserPasswordPolicy resource from API: %v", err)
	}
	return nil
}

func flattenUserPasswordPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyExpireDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyWarnDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyExpiredPasswordRenewal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserPasswordPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserPasswordPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("expire_days", flattenUserPasswordPolicyExpireDays(o["expire-days"], d, "expire_days", sv)); err != nil {
		if !fortiAPIPatch(o["expire-days"]) {
			return fmt.Errorf("Error reading expire_days: %v", err)
		}
	}

	if err = d.Set("warn_days", flattenUserPasswordPolicyWarnDays(o["warn-days"], d, "warn_days", sv)); err != nil {
		if !fortiAPIPatch(o["warn-days"]) {
			return fmt.Errorf("Error reading warn_days: %v", err)
		}
	}

	if err = d.Set("expired_password_renewal", flattenUserPasswordPolicyExpiredPasswordRenewal(o["expired-password-renewal"], d, "expired_password_renewal", sv)); err != nil {
		if !fortiAPIPatch(o["expired-password-renewal"]) {
			return fmt.Errorf("Error reading expired_password_renewal: %v", err)
		}
	}

	return nil
}

func flattenUserPasswordPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserPasswordPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyExpireDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyWarnDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyExpiredPasswordRenewal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserPasswordPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserPasswordPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("expire_days"); ok {
		t, err := expandUserPasswordPolicyExpireDays(d, v, "expire_days", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-days"] = t
		}
	}

	if v, ok := d.GetOkExists("warn_days"); ok {
		t, err := expandUserPasswordPolicyWarnDays(d, v, "warn_days", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warn-days"] = t
		}
	}

	if v, ok := d.GetOk("expired_password_renewal"); ok {
		t, err := expandUserPasswordPolicyExpiredPasswordRenewal(d, v, "expired_password_renewal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expired-password-renewal"] = t
		}
	}

	return &obj, nil
}
