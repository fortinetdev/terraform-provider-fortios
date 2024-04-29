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
				Computed: true,
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
			"minimum_length": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(8, 128),
				Optional:     true,
				Computed:     true,
			},
			"min_lower_case_letter": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"min_upper_case_letter": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"min_non_alphanumeric": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"min_number": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"min_change_characters": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"expire_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reuse_password": &schema.Schema{
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

func flattenUserPasswordPolicyMinimumLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyMinChangeCharacters(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyExpireStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPasswordPolicyReusePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

	if err = d.Set("minimum_length", flattenUserPasswordPolicyMinimumLength(o["minimum-length"], d, "minimum_length", sv)); err != nil {
		if !fortiAPIPatch(o["minimum-length"]) {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", flattenUserPasswordPolicyMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter", sv)); err != nil {
		if !fortiAPIPatch(o["min-lower-case-letter"]) {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", flattenUserPasswordPolicyMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter", sv)); err != nil {
		if !fortiAPIPatch(o["min-upper-case-letter"]) {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", flattenUserPasswordPolicyMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric", sv)); err != nil {
		if !fortiAPIPatch(o["min-non-alphanumeric"]) {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", flattenUserPasswordPolicyMinNumber(o["min-number"], d, "min_number", sv)); err != nil {
		if !fortiAPIPatch(o["min-number"]) {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("min_change_characters", flattenUserPasswordPolicyMinChangeCharacters(o["min-change-characters"], d, "min_change_characters", sv)); err != nil {
		if !fortiAPIPatch(o["min-change-characters"]) {
			return fmt.Errorf("Error reading min_change_characters: %v", err)
		}
	}

	if err = d.Set("expire_status", flattenUserPasswordPolicyExpireStatus(o["expire-status"], d, "expire_status", sv)); err != nil {
		if !fortiAPIPatch(o["expire-status"]) {
			return fmt.Errorf("Error reading expire_status: %v", err)
		}
	}

	if err = d.Set("reuse_password", flattenUserPasswordPolicyReusePassword(o["reuse-password"], d, "reuse_password", sv)); err != nil {
		if !fortiAPIPatch(o["reuse-password"]) {
			return fmt.Errorf("Error reading reuse_password: %v", err)
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

func expandUserPasswordPolicyMinimumLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinLowerCaseLetter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinUpperCaseLetter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinNonAlphanumeric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyMinChangeCharacters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyExpireStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPasswordPolicyReusePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

	if v, ok := d.GetOk("minimum_length"); ok {
		t, err := expandUserPasswordPolicyMinimumLength(d, v, "minimum_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-length"] = t
		}
	}

	if v, ok := d.GetOkExists("min_lower_case_letter"); ok {
		t, err := expandUserPasswordPolicyMinLowerCaseLetter(d, v, "min_lower_case_letter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-lower-case-letter"] = t
		}
	}

	if v, ok := d.GetOkExists("min_upper_case_letter"); ok {
		t, err := expandUserPasswordPolicyMinUpperCaseLetter(d, v, "min_upper_case_letter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-upper-case-letter"] = t
		}
	}

	if v, ok := d.GetOkExists("min_non_alphanumeric"); ok {
		t, err := expandUserPasswordPolicyMinNonAlphanumeric(d, v, "min_non_alphanumeric", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-non-alphanumeric"] = t
		}
	}

	if v, ok := d.GetOkExists("min_number"); ok {
		t, err := expandUserPasswordPolicyMinNumber(d, v, "min_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-number"] = t
		}
	}

	if v, ok := d.GetOkExists("min_change_characters"); ok {
		t, err := expandUserPasswordPolicyMinChangeCharacters(d, v, "min_change_characters", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-change-characters"] = t
		}
	}

	if v, ok := d.GetOk("expire_status"); ok {
		t, err := expandUserPasswordPolicyExpireStatus(d, v, "expire_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-status"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password"); ok {
		t, err := expandUserPasswordPolicyReusePassword(d, v, "reuse_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password"] = t
		}
	}

	return &obj, nil
}
