package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSystemLicenseFortiCare() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLicenseFortiCareCreateUpdate,
		Read:   resourceSystemLicenseFortiCareRead,
		Update: resourceSystemLicenseFortiCareCreateUpdate,
		Delete: resourceSystemLicenseFortiCareDelete,

		Schema: map[string]*schema.Schema{
			"registration_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemLicenseFortiCareCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	registrationCode := d.Get("registration_code").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemLicenseFortiCare{
		RegistrationCode: registrationCode,
	}

	//Call process by sdk
	_, err := c.CreateSystemLicenseFortiCare(i)
	if err != nil {
		return fmt.Errorf("Error creating System License FortiCare: %s", err)
	}

	// Set index for d
	d.SetId("FortiCareLicense")
	return resourceSystemLicenseFortiCareRead(d, m)
}

func resourceSystemLicenseFortiCareDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemLicenseFortiCareRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemLicenseFortiCare(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System License FortiCare: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("registration_code", o.RegistrationCode)

	return nil
}
