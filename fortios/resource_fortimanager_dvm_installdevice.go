package fortios

import (
	"fmt"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerDVMInstallDev() *schema.Resource {
	return &schema.Resource{
		Create: createFTMDVMInstallDev,
		Read:   readFTMDVMInstallDev,
		Update: updateFTMDVMInstallDev,
		Delete: deleteFTMDVMInstallDev,

		Schema: map[string]*schema.Schema{
			"target_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createFTMDVMInstallDev(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMDVMInstallDev")()

	//Build input data by sdk
	i := &fortimngclient.JSONDVMInstallDev{
		Name: d.Get("target_name").(string),
	}

	err := c.CreateDVMInstallDev(i)
	if err != nil {
		return fmt.Errorf("Error creating devicemanager install device action: %s", err)
	}

	d.SetId(i.Name)

	return nil
}

func readFTMDVMInstallDev(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateFTMDVMInstallDev(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFTMDVMInstallDev(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMDVMInstallDev")()

	d.SetId("")

	return nil
}
