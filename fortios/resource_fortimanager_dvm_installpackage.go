package fortios

import (
	"fmt"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerDVMInstallPolicyPackage() *schema.Resource {
	return &schema.Resource{
		Create: createFMGDVMInstallPolicyPackage,
		Read:   readFMGDVMInstallPolicyPackage,
		Update: updateFMGDVMInstallPolicyPackage,
		Delete: deleteFMGDVMInstallPolicyPackage,

		Schema: map[string]*schema.Schema{
			"package_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     3,
				Description: "Timeout for installing the package to the target, default: 3 minutes",
			},
		},
	}
}

func createFMGDVMInstallPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGDVMInstallPolicyPackage")()

	i := &fortimngclient.JSONDVMInstallPolicyPackage{
		Name:    d.Get("package_name").(string),
		Timeout: d.Get("timeout").(int),
	}

	err := c.CreateDVMInstallPolicyPackage(i)
	if err != nil {
		return fmt.Errorf("Error creating devicemanager install policy package action: %s", err)
	}

	d.SetId(i.Name)

	return nil
}

func readFMGDVMInstallPolicyPackage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateFMGDVMInstallPolicyPackage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFMGDVMInstallPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGDVMInstallPolicyPackage")()

	d.SetId("")

	return nil
}
