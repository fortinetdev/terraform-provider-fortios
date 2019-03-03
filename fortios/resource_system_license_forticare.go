package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemLicenseFortiCare() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLicenseFortiCareCreate,
		Read:   resourceSystemLicenseFortiCareRead,
		Update: resourceSystemLicenseFortiCareUpdate,
		Delete: resourceSystemLicenseFortiCareDelete,

		Schema: map[string]*schema.Schema{
			"registration_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemLicenseFortiCareCreate(d *schema.ResourceData, m interface{}) error {
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
	// d.SetId(strconv.Itoa(int(o.Mkey)))
	d.SetId("1")

	return nil
}

func resourceSystemLicenseFortiCareUpdate(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Get Params from d

	// registrationCode := d.Get("registration_code").(string)

	// //Build input data by sdk
	// i := &forticlient.JSONSystemLicenseFortiCare{
	// 	RegistrationCode: registrationCode,
	// }

	// //Call process by sdk
	// _, err := c.UpdateSystemLicenseFortiCare(i, mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error updating System License FortiCare: %s", err)
	// }

	// //Set index for d
	// //d.SetId(o.Mkey)
	// d.SetId("1")

	err := resourceSystemLicenseFortiCareCreate(d, m)
	if err != nil {
		return fmt.Errorf("Error creating System License FortiCare: %s", err)
	}

	return nil
}

func resourceSystemLicenseFortiCareDelete(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// err := c.DeleteSystemLicenseFortiCare(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error deleting System License FortiCare: %s", err)
	// }

	// //Set index for d
	// d.SetId("")

	return nil
}

func resourceSystemLicenseFortiCareRead(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// o, err := c.ReadSystemLicenseFortiCare(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error reading System License FortiCare: %s", err)
	// }

	// //Refresh property
	// d.Set("registration_code", o.RegistrationCode)

	return nil
}
