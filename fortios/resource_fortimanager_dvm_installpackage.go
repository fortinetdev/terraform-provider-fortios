package fortios

import (
	"fmt"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerDVMInstallPolicyPackage() *schema.Resource {
	return &schema.Resource{
		Create: createFTMDVMInstallPolicyPackage,
		Read:   readFTMDVMInstallPolicyPackage,
		Update: updateFTMDVMInstallPolicyPackage,
		Delete: deleteFTMDVMInstallPolicyPackage,

		Schema: map[string]*schema.Schema{
			"package_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createFTMDVMInstallPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMDVMInstallPolicyPackage")()

	i := &fortimngclient.JSONDVMInstallPolicyPackage{
		Name: d.Get("package_name").(string),
	}

	err := c.CreateDVMInstallPolicyPackage(i)
	if err != nil {
		return fmt.Errorf("Error creating devicemanager install policy package action: %s", err)
	}

	d.SetId(i.Name)

	return nil
}

func readFTMDVMInstallPolicyPackage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateFTMDVMInstallPolicyPackage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFTMDVMInstallPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMDVMInstallPolicyPackage")()

	d.SetId("")

	return nil
}
