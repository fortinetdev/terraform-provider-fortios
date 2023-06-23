package fortios

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	forticlient "github.com/terraform-providers/terraform-provider-fortios/sdk/sdkcore"
)

func resourceSystemLicenseVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLicenseVMCreateUpdate,
		Read:   resourceSystemLicenseVMRead,
		Update: resourceSystemLicenseVMCreateUpdate,
		Delete: resourceSystemLicenseVMDelete,

		Schema: map[string]*schema.Schema{
			"file_content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemLicenseVMCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	fileContent := d.Get("file_content").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemLicenseVM{
		FileContent: fileContent,
	}

	//Call process by sdk
	_, err := c.CreateSystemLicenseVM(i)
	if err != nil {
		return fmt.Errorf("Error creating System License VM: %s", err)
	}

	//Set index for d
	d.SetId("LicenseVM")

	return nil
}

func resourceSystemLicenseVMDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemLicenseVMRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
