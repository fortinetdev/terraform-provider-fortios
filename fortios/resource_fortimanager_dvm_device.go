package fortios

import (
	"fmt"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerDVMDevice() *schema.Resource {
	return &schema.Resource{
		Create: AddFMGDVMDevice,
		Read:   ReadFMGDVMDevice,
		Update: UpdateFMGDVMDevice,
		Delete: DeleteFMGDVMDevice,

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
			"dev_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func AddFMGDVMDevice(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("AddDVMDevice")()

	//Get Params from d
	userId := d.Get("userid").(string)
	passwd := d.Get("password").(string)
	ip := d.Get("ipaddr").(string)
	name := d.Get("dev_name").(string)

	//Build input data by sdk
	i := &fortimngclient.JSONDVMDeviceAdd{
		UserId:   userId,
		Passwd:   passwd,
		Ipaddr:   ip,
		Name:     name,
		MgmtMode: "fmg",
	}

	err := c.AddDVMDevice(i)
	if err != nil {
		return fmt.Errorf("Error adding dvm device: %s", err)
	}

	d.SetId(name)

	return nil
}

func DeleteFMGDVMDevice(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("DeleteDVMDevice")()

	name := d.Id()

	i := &fortimngclient.JSONDVMDeviceDel{
		Adom: "root",
		Name: name,
	}

	err := c.DeleteDVMDevice(i)
	if err != nil {
		return fmt.Errorf("Error deleting dvm device: %s", err)
	}

	return nil
}
func ReadFMGDVMDevice(d *schema.ResourceData, m interface{}) error {
	return nil
}

func UpdateFMGDVMDevice(d *schema.ResourceData, m interface{}) error {
	return nil
}
