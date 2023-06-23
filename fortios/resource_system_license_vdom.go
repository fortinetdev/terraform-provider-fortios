package fortios

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-fortios/sdk/sdkcore"
)

func resourceSystemLicenseVDOM() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLicenseVDOMCreateUpdate,
		Read:   resourceSystemLicenseVDOMRead,
		Update: resourceSystemLicenseVDOMCreateUpdate,
		Delete: resourceSystemLicenseVDOMDelete,

		Schema: map[string]*schema.Schema{
			"license": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemLicenseVDOMCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	license := d.Get("license").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemLicenseVDOM{
		License: license,
	}

	//Call process by sdk
	_, err := c.CreateSystemLicenseVDOM(i)
	if err != nil {
		return fmt.Errorf("Error creating System License VDOM: %s", err)
	}

	//Set index for d
	d.SetId("LicenseVDOM")

	return resourceSystemLicenseVDOMRead(d, m)
}

func resourceSystemLicenseVDOMDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemLicenseVDOMRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemLicenseVDOM(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System License VDOM: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("license", o.License)

	return nil
}
