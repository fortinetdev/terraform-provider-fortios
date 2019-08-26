package fortios

import (
	"fmt"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemLicenseFortiCare() *schema.Resource {
	return &schema.Resource{
		Create: addFTMSystemLicenseFortiCare,
		Read:   readFTMSystemLicenseFortiCare,
		Update: updateFTMSystemLicenseFortiCare,
		Delete: deleteFTMSystemLicenseFortiCare,

		Schema: map[string]*schema.Schema{
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"registration_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func addFTMSystemLicenseFortiCare(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("addFTMSystemLicenseFortiCare")()

	i := &fortimngclient.SystemLicenseFortiCare{
		Target:           d.Get("target").(string),
		RegistrationCode: d.Get("registration_code").(string),
	}

	err := c.AddSystemLicenseFortiCare(i)
	if err != nil {
		return fmt.Errorf("Error adding License for %s: %s", i.Target, err)
	}

	d.SetId("fortimanager-license-forticare")

	return nil
}

func updateFTMSystemLicenseFortiCare(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readFTMSystemLicenseFortiCare(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFTMSystemLicenseFortiCare(d *schema.ResourceData, m interface{}) error {
	return nil
}
