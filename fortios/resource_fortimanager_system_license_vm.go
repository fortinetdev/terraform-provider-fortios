package fortios

import (
	"fmt"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemLicenseVM() *schema.Resource {
	return &schema.Resource{
		Create: addFTMSystemLicenseVM,
		Read:   readFTMSystemLicenseVM,
		Update: updateFTMSystemLicenseVM,
		Delete: deleteFTMSystemLicenseVM,

		Schema: map[string]*schema.Schema{
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file_content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func addFTMSystemLicenseVM(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("addFTMSystemLicenseVM")()

	i := &fortimngclient.JSONSystemLicenseVM{
		Target:      d.Get("target").(string),
		FileContent: d.Get("file_content").(string),
	}

	err := c.AddSystemLicenseVM(i)
	if err != nil {
		return fmt.Errorf("Error adding License for %s: %s", i.Target, err)
	}

	d.SetId("fortimanager-license-vm")

	return nil
}

func updateFTMSystemLicenseVM(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readFTMSystemLicenseVM(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFTMSystemLicenseVM(d *schema.ResourceData, m interface{}) error {
	return nil
}
