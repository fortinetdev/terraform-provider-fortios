package fortios

import (
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemLicenseFortiFlex() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLicenseFortiFlexCreateUpdate,
		Read:   resourceSystemLicenseFortiFlexRead,
		Update: resourceSystemLicenseFortiFlexCreateUpdate,
		Delete: resourceSystemLicenseFortiFlexDelete,

		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"token_writeonly": &schema.Schema{
				Type:      schema.TypeString,
				WriteOnly: true,
				Optional:  true,
			},
			"proxy_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemLicenseFortiFlexCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Build input data
	para := make(map[string]interface{})
	woToken, diags := d.GetRawConfigAt(cty.GetAttrPath("token_writeonly"))
	if diags.HasError() {
		return fmt.Errorf(diags[0].Summary + "\n" + diags[0].Detail)
	}
	if !woToken.Type().Equals(cty.String) {
		return fmt.Errorf("error retrieving write-only attribute: token_writeonly - retrieved config value is not a string")
	}

	if !woToken.IsNull() {
		para["token"] = woToken.AsString()
	} else if v, ok := d.GetOk("token"); ok {
		para["token"] = v.(string)
	}
	if v, ok := d.GetOk("proxy_url"); ok {
		para["proxy_url"] = v
	}

	//Call process by sdk
	_, err := c.CreateSystemLicenseFortiFlex(&para)
	if err != nil {
		return fmt.Errorf("Error creating System License VM: %s", err)
	}

	//Set index for d
	d.SetId("LicenseFortiFlex")

	return nil
}

func resourceSystemLicenseFortiFlexDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemLicenseFortiFlexRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
