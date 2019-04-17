package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
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

	return resourceSystemLicenseVMRead(d, m)
}

func resourceSystemLicenseVMDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemLicenseVMRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemLicenseVM(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System License VM: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("file_content", o.FileContent)

	return nil
}
