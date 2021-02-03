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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemPasswordPolicyRead,
		Schema: map[string]*schema.Schema{
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

func dataSourceSystemPasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemPasswordPolicy"

	o, err := c.ReadSystemPasswordPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemPasswordPolicy: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemPasswordPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemPasswordPolicy from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemPasswordPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyApplyTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyMinimumLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyMinLowerCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyMinUpperCaseLetter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyMinNonAlphanumeric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyMinNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyChange4Characters(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyExpireStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyExpireDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPasswordPolicyReusePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemPasswordPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemPasswordPolicyStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("apply_to", dataSourceFlattenSystemPasswordPolicyApplyTo(o["apply-to"], d, "apply_to")); err != nil {
		if !fortiAPIPatch(o["apply-to"]) {
			return fmt.Errorf("Error reading apply_to: %v", err)
		}
	}

	if err = d.Set("minimum_length", dataSourceFlattenSystemPasswordPolicyMinimumLength(o["minimum-length"], d, "minimum_length")); err != nil {
		if !fortiAPIPatch(o["minimum-length"]) {
			return fmt.Errorf("Error reading minimum_length: %v", err)
		}
	}

	if err = d.Set("min_lower_case_letter", dataSourceFlattenSystemPasswordPolicyMinLowerCaseLetter(o["min-lower-case-letter"], d, "min_lower_case_letter")); err != nil {
		if !fortiAPIPatch(o["min-lower-case-letter"]) {
			return fmt.Errorf("Error reading min_lower_case_letter: %v", err)
		}
	}

	if err = d.Set("min_upper_case_letter", dataSourceFlattenSystemPasswordPolicyMinUpperCaseLetter(o["min-upper-case-letter"], d, "min_upper_case_letter")); err != nil {
		if !fortiAPIPatch(o["min-upper-case-letter"]) {
			return fmt.Errorf("Error reading min_upper_case_letter: %v", err)
		}
	}

	if err = d.Set("min_non_alphanumeric", dataSourceFlattenSystemPasswordPolicyMinNonAlphanumeric(o["min-non-alphanumeric"], d, "min_non_alphanumeric")); err != nil {
		if !fortiAPIPatch(o["min-non-alphanumeric"]) {
			return fmt.Errorf("Error reading min_non_alphanumeric: %v", err)
		}
	}

	if err = d.Set("min_number", dataSourceFlattenSystemPasswordPolicyMinNumber(o["min-number"], d, "min_number")); err != nil {
		if !fortiAPIPatch(o["min-number"]) {
			return fmt.Errorf("Error reading min_number: %v", err)
		}
	}

	if err = d.Set("change_4_characters", dataSourceFlattenSystemPasswordPolicyChange4Characters(o["change-4-characters"], d, "change_4_characters")); err != nil {
		if !fortiAPIPatch(o["change-4-characters"]) {
			return fmt.Errorf("Error reading change_4_characters: %v", err)
		}
	}

	if err = d.Set("expire_status", dataSourceFlattenSystemPasswordPolicyExpireStatus(o["expire-status"], d, "expire_status")); err != nil {
		if !fortiAPIPatch(o["expire-status"]) {
			return fmt.Errorf("Error reading expire_status: %v", err)
		}
	}

	if err = d.Set("expire_day", dataSourceFlattenSystemPasswordPolicyExpireDay(o["expire-day"], d, "expire_day")); err != nil {
		if !fortiAPIPatch(o["expire-day"]) {
			return fmt.Errorf("Error reading expire_day: %v", err)
		}
	}

	if err = d.Set("reuse_password", dataSourceFlattenSystemPasswordPolicyReusePassword(o["reuse-password"], d, "reuse_password")); err != nil {
		if !fortiAPIPatch(o["reuse-password"]) {
			return fmt.Errorf("Error reading reuse_password: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemPasswordPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
