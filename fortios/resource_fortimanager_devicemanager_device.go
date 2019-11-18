package fortios

import (
	"fmt"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerDVMDevice() *schema.Resource {
	return &schema.Resource{
		Create: createFMGDVMDevice,
		Read:   readFMGDVMDevice,
		Update: updateFMGDVMDevice,
		Delete: deleteFMGDVMDevice,

		Schema: map[string]*schema.Schema{
			"userid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipaddr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
		},
	}
}

func createFMGDVMDevice(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGDVMDevice")()

	//Build input data by sdk
	i := &fmgclient.JSONDVMDeviceCreate{
		UserId:   d.Get("userid").(string),
		Passwd:   d.Get("password").(string),
		Ipaddr:   d.Get("ipaddr").(string),
		Name:     d.Get("device_name").(string),
		MgmtMode: "fmg",
	}

	err := c.CreateDVMDevice(d.Get("adom").(string), i)
	if err != nil {
		return fmt.Errorf("Error adding dvm device: %s", err)
	}

	d.SetId(i.Name)

	return nil
}

func deleteFMGDVMDevice(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGDVMDevice")()

	name := d.Id()

	i := &fmgclient.JSONDVMDeviceDel{
		Adom: d.Get("adom").(string),
		Name: name,
	}

	err := c.DeleteDVMDevice(i)
	if err != nil {
		return fmt.Errorf("Error deleting dvm device: %s", err)
	}

	return nil
}
func readFMGDVMDevice(d *schema.ResourceData, m interface{}) error {
	return nil
}

func updateFMGDVMDevice(d *schema.ResourceData, m interface{}) error {
	return nil
}
