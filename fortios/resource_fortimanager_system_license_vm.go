package fortios

import (
	"fmt"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemLicenseVM() *schema.Resource {
	return &schema.Resource{
		Create: addFMGSystemLicenseVM,
		Read:   readFMGSystemLicenseVM,
		Update: updateFMGSystemLicenseVM,
		Delete: deleteFMGSystemLicenseVM,

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

func addFMGSystemLicenseVM(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("addFMGSystemLicenseVM")()

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

func updateFMGSystemLicenseVM(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readFMGSystemLicenseVM(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFMGSystemLicenseVM(d *schema.ResourceData, m interface{}) error {
	return nil
}
