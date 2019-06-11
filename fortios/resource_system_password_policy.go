package fortios

import (
	"fmt"
	"log"

	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPasswordPolicyCreateUpdate,
		Read:   resourceSystemPasswordPolicyRead,
		Update: resourceSystemPasswordPolicyCreateUpdate,
		Delete: resourceSystemPasswordPolicyDelete,

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apply_to": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"minimum_length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"min_lower_case_letter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"min_upper_case_letter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"min_non_alphanumeric": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"min_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"change_4_characters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"expire_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
			},
			"expire_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "90",
			},
			"reuse_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
			},
		},
	}
}

func resourceSystemPasswordPolicyCreateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	status := d.Get("status").(string)
	applyTo := d.Get("apply_to").(string)
	minimumLength := d.Get("minimum_length").(string)
	minLowerCaseLetter := d.Get("min_lower_case_letter").(string)
	minUpperCaseLetter := d.Get("min_upper_case_letter").(string)
	minNonAlphanumeric := d.Get("min_non_alphanumeric").(string)
	minNumber := d.Get("min_number").(string)
	change4Characters := d.Get("change_4_characters").(string)
	expireStatus := d.Get("expire_status").(string)
	expireDay := d.Get("expire_day").(string)
	reusePassword := d.Get("reuse_password").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemPasswordPolicy{
		Status:             status,
		ApplyTo:            applyTo,
		MinimunLength:      minimumLength,
		MinLowerCaseLetter: minLowerCaseLetter,
		MinUpperCaseLetter: minUpperCaseLetter,
		MinNonAlphanumeric: minNonAlphanumeric,
		MinNumber:          minNumber,
		Change4Characters:  change4Characters,
		ExpireStatus:       expireStatus,
		ExpireDay:          expireDay,
		ReusePassword:      reusePassword,
	}

	//Call process by sdk
	_, err := c.UpdateSystemPasswordPolicy(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System Setting Global: %s", err)
	}

	d.SetId(applyTo)

	return resourceSystemSettingGlobalRead(d, m)
}
func resourceSystemPasswordPolicyDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemPasswordPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemPasswordPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System Password Policy: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("status", o.Status)
	d.Set("apply_to", o.ApplyTo)
	d.Set("minimum_length", o.MinimunLength)
	d.Set("min_lower_case_letter", o.MinLowerCaseLetter)
	d.Set("min_upper_case_letter", o.MinUpperCaseLetter)
	d.Set("min_non_alphanumeric", o.MinNonAlphanumeric)
	d.Set("min_number", o.MinNumber)
	d.Set("change_4_characters", o.Change4Characters)
	d.Set("expire_status", o.ExpireStatus)
	d.Set("expire_day", o.ExpireDay)
	d.Set("reuse_password", o.ReusePassword)

	return nil
}
