// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure the password policy for guest administrators.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemPasswordPolicyGuestAdmin() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemPasswordPolicyGuestAdminRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"apply_to": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"minimum_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_lower_case_letter": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_upper_case_letter": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_non_alphanumeric": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_change_characters": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"change_4_characters": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"expire_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"expire_day": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"reuse_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemPasswordPolicyGuestAdminRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemPasswordPolicyGuestAdmin"

	o, err := c.ReadSystemPasswordPolicyGuestAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemPasswordPolicyGuestAdmin: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemPasswordPolicyGuestAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemPasswordPolicyGuestAdmin from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminApplyTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminMinimumLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminMinNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminMinChangeCharacters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminChange4Characters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminExpireStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminExpireDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminReusePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemPasswordPolicyGuestAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemPasswordPolicyGuestAdminStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("apply_to", dataSourceFlattenSystemPasswordPolicyGuestAdminApplyTo(o["apply-to"], d, "apply_to")); err != nil {
		if !fortiAPIPatch(o["apply-to"]) {
			return fmt.Errorf("Error reading apply_to: %v", err)
		}
	}

	if err = d.Set("minimum_length", dataSourceFlattenSystemPasswordPolicyGuestAdminMinimumLength(o["minimum-length"], d, "minimum_length")); err != nil {
		if !fortiAPIPatch(o["minimum-length"]) {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", dataSourceFlattenSystemPasswordPolicyGuestAdminMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter")); err != nil {
		if !fortiAPIPatch(o["min-lower-case-letter"]) {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", dataSourceFlattenSystemPasswordPolicyGuestAdminMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter")); err != nil {
		if !fortiAPIPatch(o["min-upper-case-letter"]) {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", dataSourceFlattenSystemPasswordPolicyGuestAdminMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric")); err != nil {
		if !fortiAPIPatch(o["min-non-alphanumeric"]) {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", dataSourceFlattenSystemPasswordPolicyGuestAdminMinNumber(o["min-number"], d, "min_number")); err != nil {
		if !fortiAPIPatch(o["min-number"]) {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("min_change_characters", dataSourceFlattenSystemPasswordPolicyGuestAdminMinChangeCharacters(o["min-change-characters"], d, "min_change_characters")); err != nil {
		if !fortiAPIPatch(o["min-change-characters"]) {
			return fmt.Errorf("Error reading min_change_characters: %v", err)
		}
	}

	if err = d.Set("change_4_characters", dataSourceFlattenSystemPasswordPolicyGuestAdminChange4Characters(o["change-4-characters"], d, "change_4_characters")); err != nil {
		if !fortiAPIPatch(o["change-4-characters"]) {
			return fmt.Errorf("Error reading change_4_characters: %v", err)
		}
	}

	if err = d.Set("expire_status", dataSourceFlattenSystemPasswordPolicyGuestAdminExpireStatus(o["expire-status"], d, "expire_status")); err != nil {
		if !fortiAPIPatch(o["expire-status"]) {
			return fmt.Errorf("Error reading expire_status: %v", err)
		}
	}

	if err = d.Set("expire_day", dataSourceFlattenSystemPasswordPolicyGuestAdminExpireDay(o["expire-day"], d, "expire_day")); err != nil {
		if !fortiAPIPatch(o["expire-day"]) {
			return fmt.Errorf("Error reading expire_day: %v", err)
		}
	}

	if err = d.Set("reuse_password", dataSourceFlattenSystemPasswordPolicyGuestAdminReusePassword(o["reuse-password"], d, "reuse_password")); err != nil {
		if !fortiAPIPatch(o["reuse-password"]) {
			return fmt.Errorf("Error reading reuse_password: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemPasswordPolicyGuestAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
