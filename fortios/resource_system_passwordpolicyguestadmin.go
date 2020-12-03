// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure the password policy for guest administrators.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemPasswordPolicyGuestAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPasswordPolicyGuestAdminUpdate,
		Read:   resourceSystemPasswordPolicyGuestAdminRead,
		Update: resourceSystemPasswordPolicyGuestAdminUpdate,
		Delete: resourceSystemPasswordPolicyGuestAdminDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceSystemPasswordPolicyGuestAdminUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemPasswordPolicyGuestAdmin(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicyGuestAdmin resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPasswordPolicyGuestAdmin(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemPasswordPolicyGuestAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPasswordPolicyGuestAdmin")
	}

	return resourceSystemPasswordPolicyGuestAdminRead(d, m)
}

func resourceSystemPasswordPolicyGuestAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemPasswordPolicyGuestAdmin(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPasswordPolicyGuestAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPasswordPolicyGuestAdminRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemPasswordPolicyGuestAdmin(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicyGuestAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPasswordPolicyGuestAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPasswordPolicyGuestAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemPasswordPolicyGuestAdminStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminApplyTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinimumLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminMinNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminChange4Characters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminExpireStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminExpireDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPasswordPolicyGuestAdminReusePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemPasswordPolicyGuestAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemPasswordPolicyGuestAdminStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("apply_to", flattenSystemPasswordPolicyGuestAdminApplyTo(o["apply-to"], d, "apply_to")); err != nil {
		if !fortiAPIPatch(o["apply-to"]) {
			return fmt.Errorf("Error reading apply_to: %v", err)
		}
	}

	if err = d.Set("minimum_length", flattenSystemPasswordPolicyGuestAdminMinimumLength(o["minimum-length"], d, "minimum_length")); err != nil {
		if !fortiAPIPatch(o["minimum-length"]) {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", flattenSystemPasswordPolicyGuestAdminMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter")); err != nil {
		if !fortiAPIPatch(o["min-lower-case-letter"]) {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", flattenSystemPasswordPolicyGuestAdminMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter")); err != nil {
		if !fortiAPIPatch(o["min-upper-case-letter"]) {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", flattenSystemPasswordPolicyGuestAdminMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric")); err != nil {
		if !fortiAPIPatch(o["min-non-alphanumeric"]) {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", flattenSystemPasswordPolicyGuestAdminMinNumber(o["min-number"], d, "min_number")); err != nil {
		if !fortiAPIPatch(o["min-number"]) {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("change_4_characters", flattenSystemPasswordPolicyGuestAdminChange4Characters(o["change-4-characters"], d, "change_4_characters")); err != nil {
		if !fortiAPIPatch(o["change-4-characters"]) {
			return fmt.Errorf("Error reading change_4_characters: %v", err)
		}
	}

	if err = d.Set("expire_status", flattenSystemPasswordPolicyGuestAdminExpireStatus(o["expire-status"], d, "expire_status")); err != nil {
		if !fortiAPIPatch(o["expire-status"]) {
			return fmt.Errorf("Error reading expire_status: %v", err)
		}
	}

	if err = d.Set("expire_day", flattenSystemPasswordPolicyGuestAdminExpireDay(o["expire-day"], d, "expire_day")); err != nil {
		if !fortiAPIPatch(o["expire-day"]) {
			return fmt.Errorf("Error reading expire_day: %v", err)
		}
	}

	if err = d.Set("reuse_password", flattenSystemPasswordPolicyGuestAdminReusePassword(o["reuse-password"], d, "reuse_password")); err != nil {
		if !fortiAPIPatch(o["reuse-password"]) {
			return fmt.Errorf("Error reading reuse_password: %v", err)
		}
	}

	return nil
}

func flattenSystemPasswordPolicyGuestAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPasswordPolicyGuestAdminStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminApplyTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinimumLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinLowerCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinUpperCaseLetter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinNonAlphanumeric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminMinNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminChange4Characters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminExpireStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminExpireDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPasswordPolicyGuestAdminReusePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPasswordPolicyGuestAdmin(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("apply_to"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminApplyTo(d, v, "apply_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apply-to"] = t
		}
	}

	if v, ok := d.GetOk("minimum_length"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminMinimumLength(d, v, "minimum_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-length"] = t
		}
	}

	if v, ok := d.GetOkExists("min_lower_case_letter"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminMinLowerCaseLetter(d, v, "min_lower_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-lower-case-letter"] = t
		}
	}

	if v, ok := d.GetOkExists("min_upper_case_letter"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminMinUpperCaseLetter(d, v, "min_upper_case_letter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-upper-case-letter"] = t
		}
	}

	if v, ok := d.GetOkExists("min_non_alphanumeric"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminMinNonAlphanumeric(d, v, "min_non_alphanumeric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-non-alphanumeric"] = t
		}
	}

	if v, ok := d.GetOkExists("min_number"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminMinNumber(d, v, "min_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-number"] = t
		}
	}

	if v, ok := d.GetOk("change_4_characters"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminChange4Characters(d, v, "change_4_characters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-4-characters"] = t
		}
	}

	if v, ok := d.GetOk("expire_status"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminExpireStatus(d, v, "expire_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-status"] = t
		}
	}

	if v, ok := d.GetOk("expire_day"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminExpireDay(d, v, "expire_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire-day"] = t
		}
	}

	if v, ok := d.GetOk("reuse_password"); ok {
		t, err := expandSystemPasswordPolicyGuestAdminReusePassword(d, v, "reuse_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-password"] = t
		}
	}

	return &obj, nil
}
