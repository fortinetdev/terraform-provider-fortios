// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPasswordPolicyUpdate,
		Read:   resourceSystemPasswordPolicyRead,
		Update: resourceSystemPasswordPolicyUpdate,
		Delete: resourceSystemPasswordPolicyDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apply_to": &schema.Schema{
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
			},
			"min_upper_case_letter": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
			},
			"min_non_alphanumeric": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
			},
			"min_number": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
			},
			"min_change_characters": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
			},
			"change_4_characters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expire_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expire_day": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 999),
				Optional:     true,
				Computed:     true,
			},
			"reuse_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reuse_password_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 20),
				Optional:     true,
			},
			"login_lockout_upon_downgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemPasswordPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemPasswordPolicy(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPasswordPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPasswordPolicy")
	}

	return resourceSystemPasswordPolicyRead(d, m)
}

func resourceSystemPasswordPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemPasswordPolicy(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPasswordPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemPasswordPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemPasswordPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPasswordPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSystemPasswordPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPasswordPolicyApplyTo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPasswordPolicyMinimumLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPasswordPolicyMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPasswordPolicyMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPasswordPolicyMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPasswordPolicyMinNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPasswordPolicyMinChangeCharacters(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPasswordPolicyChange4Characters(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPasswordPolicyExpireStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPasswordPolicyExpireDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPasswordPolicyReusePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPasswordPolicyReusePasswordLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPasswordPolicyLoginLockoutUponDowngrade(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemPasswordPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemPasswordPolicyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("apply_to", flattenSystemPasswordPolicyApplyTo(o["apply-to"], d, "apply_to", sv)); err != nil {
		if !fortiAPIPatch(o["apply-to"]) {
			return fmt.Errorf("Error reading apply_to: %v", err)
		}
	}

	if err = d.Set("minimum_length", flattenSystemPasswordPolicyMinimumLength(o["minimum-length"], d, "minimum_length", sv)); err != nil {
		if !fortiAPIPatch(o["minimum-length"]) {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", flattenSystemPasswordPolicyMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter", sv)); err != nil {
		if !fortiAPIPatch(o["min-lower-case-letter"]) {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", flattenSystemPasswordPolicyMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter", sv)); err != nil {
		if !fortiAPIPatch(o["min-upper-case-letter"]) {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", flattenSystemPasswordPolicyMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric", sv)); err != nil {
		if !fortiAPIPatch(o["min-non-alphanumeric"]) {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", flattenSystemPasswordPolicyMinNumber(o["min-number"], d, "min_number", sv)); err != nil {
		if !fortiAPIPatch(o["min-number"]) {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("min_change_characters", flattenSystemPasswordPolicyMinChangeCharacters(o["min-change-characters"], d, "min_change_characters", sv)); err != nil {
		if !fortiAPIPatch(o["min-change-characters"]) {
			return fmt.Errorf("Error reading min_change_characters: %v", err)
		}
	}

	if err = d.Set("change_4_characters", flattenSystemPasswordPolicyChange4Characters(o["change-4-characters"], d, "change_4_characters", sv)); err != nil {
		if !fortiAPIPatch(o["change-4-characters"]) {
			return fmt.Errorf("Error reading change_4_characters: %v", err)
		}
	}

	if err = d.Set("expire_status", flattenSystemPasswordPolicyExpireStatus(o["expire-status"], d, "expire_status", sv)); err != nil {
		if !fortiAPIPatch(o["expire-status"]) {
			return fmt.Errorf("Error reading expire_status: %v", err)
		}
	}

	if err = d.Set("expire_day", flattenSystemPasswordPolicyExpireDay(o["expire-day"], d, "expire_day", sv)); err != nil {
		if !fortiAPIPatch(o["expire-day"]) {
			return fmt.Errorf("Error reading expire_day: %v", err)
		}
	}

	if err = d.Set("reuse_password", flattenSystemPasswordPolicyReusePassword(o["reuse-password"], d, "reuse_password", sv)); err != nil {
		if !fortiAPIPatch(o["reuse-password"]) {
			return fmt.Errorf("Error reading reuse_password: %v", err)
		}
	}

	if err = d.Set("reuse_password_limit", flattenSystemPasswordPolicyReusePasswordLimit(o["reuse-password-limit"], d, "reuse_password_limit", sv)); err != nil {
		if !fortiAPIPatch(o["reuse-password-limit"]) {
			return fmt.Errorf("Error reading reuse_password_limit: %v", err)
		}
	}

	if err = d.Set("login_lockout_upon_downgrade", flattenSystemPasswordPolicyLoginLockoutUponDowngrade(o["login-lockout-upon-downgrade"], d, "login_lockout_upon_downgrade", sv)); err != nil {
		if !fortiAPIPatch(o["login-lockout-upon-downgrade"]) {
			return fmt.Errorf("Error reading login_lockout_upon_downgrade: %v", err)
		}
	}

	return nil
}

func flattenSystemPasswordPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemPasswordPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyApplyTo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinimumLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinLowerCaseLetter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinUpperCaseLetter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinNonAlphanumeric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyMinChangeCharacters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyChange4Characters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyExpireStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyExpireDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyReusePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyReusePasswordLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyLoginLockoutUponDowngrade(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPasswordPolicy(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemPasswordPolicyStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("apply_to"); ok {
		if setArgNil {
			obj["apply-to"] = nil
		} else {
			t, err := expandSystemPasswordPolicyApplyTo(d, v, "apply_to", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["apply-to"] = t
			}
		}
	}

	if v, ok := d.GetOk("minimum_length"); ok {
		if setArgNil {
			obj["minimum-length"] = nil
		} else {
			t, err := expandSystemPasswordPolicyMinimumLength(d, v, "minimum_length", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["minimum-length"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("min_lower_case_letter"); ok {
		if setArgNil {
			obj["min-lower-case-letter"] = nil
		} else {
			t, err := expandSystemPasswordPolicyMinLowerCaseLetter(d, v, "min_lower_case_letter", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["min-lower-case-letter"] = t
			}
		}
	} else if d.HasChange("min_lower_case_letter") {
		obj["min-lower-case-letter"] = nil
	}

	if v, ok := d.GetOkExists("min_upper_case_letter"); ok {
		if setArgNil {
			obj["min-upper-case-letter"] = nil
		} else {
			t, err := expandSystemPasswordPolicyMinUpperCaseLetter(d, v, "min_upper_case_letter", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["min-upper-case-letter"] = t
			}
		}
	} else if d.HasChange("min_upper_case_letter") {
		obj["min-upper-case-letter"] = nil
	}

	if v, ok := d.GetOkExists("min_non_alphanumeric"); ok {
		if setArgNil {
			obj["min-non-alphanumeric"] = nil
		} else {
			t, err := expandSystemPasswordPolicyMinNonAlphanumeric(d, v, "min_non_alphanumeric", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["min-non-alphanumeric"] = t
			}
		}
	} else if d.HasChange("min_non_alphanumeric") {
		obj["min-non-alphanumeric"] = nil
	}

	if v, ok := d.GetOkExists("min_number"); ok {
		if setArgNil {
			obj["min-number"] = nil
		} else {
			t, err := expandSystemPasswordPolicyMinNumber(d, v, "min_number", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["min-number"] = t
			}
		}
	} else if d.HasChange("min_number") {
		obj["min-number"] = nil
	}

	if v, ok := d.GetOkExists("min_change_characters"); ok {
		if setArgNil {
			obj["min-change-characters"] = nil
		} else {
			t, err := expandSystemPasswordPolicyMinChangeCharacters(d, v, "min_change_characters", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["min-change-characters"] = t
			}
		}
	} else if d.HasChange("min_change_characters") {
		obj["min-change-characters"] = nil
	}

	if v, ok := d.GetOk("change_4_characters"); ok {
		if setArgNil {
			obj["change-4-characters"] = nil
		} else {
			t, err := expandSystemPasswordPolicyChange4Characters(d, v, "change_4_characters", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["change-4-characters"] = t
			}
		}
	}

	if v, ok := d.GetOk("expire_status"); ok {
		if setArgNil {
			obj["expire-status"] = nil
		} else {
			t, err := expandSystemPasswordPolicyExpireStatus(d, v, "expire_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["expire-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("expire_day"); ok {
		if setArgNil {
			obj["expire-day"] = nil
		} else {
			t, err := expandSystemPasswordPolicyExpireDay(d, v, "expire_day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["expire-day"] = t
			}
		}
	}

	if v, ok := d.GetOk("reuse_password"); ok {
		if setArgNil {
			obj["reuse-password"] = nil
		} else {
			t, err := expandSystemPasswordPolicyReusePassword(d, v, "reuse_password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["reuse-password"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("reuse_password_limit"); ok {
		if setArgNil {
			obj["reuse-password-limit"] = nil
		} else {
			t, err := expandSystemPasswordPolicyReusePasswordLimit(d, v, "reuse_password_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["reuse-password-limit"] = t
			}
		}
	} else if d.HasChange("reuse_password_limit") {
		obj["reuse-password-limit"] = nil
	}

	if v, ok := d.GetOk("login_lockout_upon_downgrade"); ok {
		if setArgNil {
			obj["login-lockout-upon-downgrade"] = nil
		} else {
			t, err := expandSystemPasswordPolicyLoginLockoutUponDowngrade(d, v, "login_lockout_upon_downgrade", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["login-lockout-upon-downgrade"] = t
			}
		}
	}

	return &obj, nil
}
