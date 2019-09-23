package fortios

import (
	"fmt"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerDVMInstallDev() *schema.Resource {
	return &schema.Resource{
		Create: createFMGDVMInstallDev,
		Read:   readFMGDVMInstallDev,
		Update: updateFMGDVMInstallDev,
		Delete: deleteFMGDVMInstallDev,

		Schema: map[string]*schema.Schema{
			"target_devname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     3,
				Description: "Timeout for installing the script to the target, default: 3 minutes",
			},
		},
	}
}

func createFMGDVMInstallDev(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGDVMInstallDev")()

	//Build input data by sdk
	i := &fmgclient.JSONDVMInstallDev{
		Name:    d.Get("target_devname").(string),
		Timeout: d.Get("timeout").(int),
	}

	err := c.CreateDVMInstallDev(i)
	if err != nil {
		return fmt.Errorf("Error creating devicemanager install device action: %s", err)
	}

	d.SetId(i.Name)

	return nil
}

func readFMGDVMInstallDev(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateFMGDVMInstallDev(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFMGDVMInstallDev(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGDVMInstallDev")()

	d.SetId("")

	return nil
}
