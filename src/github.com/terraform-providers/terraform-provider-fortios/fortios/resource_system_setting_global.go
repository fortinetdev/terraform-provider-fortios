package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemSettingGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSettingGlobalCreate,
		Read:   resourceSystemSettingGlobalRead,
		Update: resourceSystemSettingGlobalUpdate,
		Delete: resourceSystemSettingGlobalDelete,

		Schema: map[string]*schema.Schema{
			"admintimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"admin_sport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_ssh_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSettingGlobalCreate(d *schema.ResourceData, m interface{}) error {
	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Get Params from d
	// admintimeout := d.Get("admintimeout").(string)
	// timezone := d.Get("timezone").(string)
	// hostname := d.Get("hostname").(string)
	// adminSport := d.Get("admin_sport").(string)
	// adminSSHPort := d.Get("admin_ssh_port").(string)

	// //Build input data by sdk
	// i := &forticlient.JSONSystemSettingGlobal{
	// 	Admintimeout: admintimeout,
	// 	Timezone:     timezone,
	// 	Hostname:     hostname,
	// 	AdminSport:   adminSport,
	// 	AdminSSHPort: adminSSHPort,
	// }

	// //Call process by sdk
	// o, err := c.CreateSystemSettingGlobal(i)
	// if err != nil {
	// 	return fmt.Errorf("Error creating System Setting Global: %s", err)
	// }

	// // Set index for d
	// // d.SetId(strconv.Itoa(int(o.Mkey)))
	// d.SetId(o.Mkey)

	err := resourceSystemSettingGlobalUpdate(d, m)
	if err != nil {
		return fmt.Errorf("Error updating System Setting Global: %s", err)
	}

	return nil
}

func resourceSystemSettingGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d

	admintimeout := d.Get("admintimeout").(string)
	timezone := d.Get("timezone").(string)
	hostname := d.Get("hostname").(string)
	adminSport := d.Get("admin_sport").(string)
	adminSSHPort := d.Get("admin_ssh_port").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemSettingGlobal{
		Admintimeout: admintimeout,
		Timezone:     timezone,
		Hostname:     hostname,
		AdminSport:   adminSport,
		AdminSSHPort: adminSSHPort,
	}

	//Call process by sdk
	_, err := c.UpdateSystemSettingGlobal(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System Setting Global: %s", err)
	}

	//Set index for d
	//d.SetId(o.Mkey)
	d.SetId("1")

	return nil
}

func resourceSystemSettingGlobalDelete(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// err := c.DeleteSystemSettingGlobal(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error deleting System Setting Global: %s", err)
	// }

	// //Set index for d
	d.SetId("")

	return nil
}

func resourceSystemSettingGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemSettingGlobal(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System Setting Global: %s", err)
	}

	//Refresh property
	d.Set("admintimeout", o.Admintimeout)
	d.Set("timezone", o.Timezone)
	d.Set("hostname", o.Hostname)
	d.Set("admin_sport", o.AdminSport)
	d.Set("admin_ssh_port", o.AdminSSHPort)

	return nil
}
