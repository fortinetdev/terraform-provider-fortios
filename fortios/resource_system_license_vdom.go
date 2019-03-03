package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemLicenseVDOM() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLicenseVDOMCreate,
		Read:   resourceSystemLicenseVDOMRead,
		Update: resourceSystemLicenseVDOMUpdate,
		Delete: resourceSystemLicenseVDOMDelete,

		Schema: map[string]*schema.Schema{
			"license": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemLicenseVDOMCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	license := d.Get("license").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemLicenseVDOM{
		License: license,
	}

	//Call process by sdk
	o, err := c.CreateSystemLicenseVDOM(i)
	if err != nil {
		return fmt.Errorf("Error creating System License VDOM: %s", err)
	}

	// Set index for d
	// d.SetId(strconv.Itoa(int(o.Mkey)))
	d.SetId(o.Mkey)

	return nil
}

func resourceSystemLicenseVDOMUpdate(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Get Params from d

	// license := d.Get("license").(string)

	// //Build input data by sdk
	// i := &forticlient.JSONSystemLicenseVDOM{
	// 	License:       license,
	// }

	// //Call process by sdk
	// _, err := c.UpdateSystemLicenseVDOM(i, mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error updating System License VDOM: %s", err)
	// }

	// //Set index for d
	// //d.SetId(o.Mkey)
	err := resourceSystemLicenseVDOMCreate(d, m)
	if err != nil {
		return fmt.Errorf("Error creating System License VDOM: %s", err)
	}

	return nil
}

func resourceSystemLicenseVDOMDelete(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// err := c.DeleteSystemLicenseVDOM(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error deleting System License VDOM: %s", err)
	// }

	// //Set index for d
	// d.SetId("")

	return nil
}

func resourceSystemLicenseVDOMRead(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// o, err := c.ReadSystemLicenseVDOM(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error reading System License VDOM: %s", err)
	// }

	// //Refresh property
	// d.Set("license", o.License)

	return nil
}
