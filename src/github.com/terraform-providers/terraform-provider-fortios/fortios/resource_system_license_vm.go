package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemLicenseVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLicenseVMCreate,
		Read:   resourceSystemLicenseVMRead,
		Update: resourceSystemLicenseVMUpdate,
		Delete: resourceSystemLicenseVMDelete,

		Schema: map[string]*schema.Schema{
			"file_content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemLicenseVMCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	fileContent := d.Get("file_content").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemLicenseVM{
		FileContent: fileContent,
	}

	//Call process by sdk
	o, err := c.CreateSystemLicenseVM(i)
	if err != nil {
		return fmt.Errorf("Error creating System License VM: %s", err)
	}

	// Set index for d
	// d.SetId(strconv.Itoa(int(o.Mkey)))
	d.SetId(o.Mkey)

	return nil
}

func resourceSystemLicenseVMUpdate(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Get Params from d

	// fileContent := d.Get("file_content").(string)

	// //Build input data by sdk
	// i := &forticlient.JSONSystemLicenseVM{
	// 	FileContent: fileContent,
	// }

	// //Call process by sdk
	// _, err := c.UpdateSystemLicenseVM(i, mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error updating System License VM: %s", err)
	// }

	// //Set index for d
	// //d.SetId(o.Mkey)

	err := resourceSystemLicenseVMCreate(d, m)
	if err != nil {
		return fmt.Errorf("Error creating System License VM: %s", err)
	}

	return nil
}

func resourceSystemLicenseVMDelete(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// err := c.DeleteSystemLicenseVM(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error deleting System License VM: %s", err)
	// }

	// //Set index for d
	// d.SetId("")

	return nil
}

func resourceSystemLicenseVMRead(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// o, err := c.ReadSystemLicenseVM(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error reading System License VM: %s", err)
	// }

	// //Refresh property
	// d.Set("file_content", o.FileContent)

	return nil
}
