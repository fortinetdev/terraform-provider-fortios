package fortios

import (
	"fmt"
	"log"

	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemSettingGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSettingGlobalCreateUpdate,
		Read:   resourceSystemSettingGlobalRead,
		Update: resourceSystemSettingGlobalCreateUpdate,
		Delete: resourceSystemSettingGlobalDelete,

		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"admintimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "5",
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "04",
			},
			"admin_sport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "443",
			},
			"admin_ssh_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "22",
			},
			"admin_scp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
			},
		},
	}
}

func resourceSystemSettingGlobalCreateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	admintimeout := d.Get("admintimeout").(string)
	timezone := d.Get("timezone").(string)
	hostname := d.Get("hostname").(string)
	adminSport := d.Get("admin_sport").(string)
	adminSSHPort := d.Get("admin_ssh_port").(string)
	adminScp := d.Get("admin_scp").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemSettingGlobal{
		Admintimeout: admintimeout,
		Timezone:     timezone,
		Hostname:     hostname,
		AdminSport:   adminSport,
		AdminSSHPort: adminSSHPort,
		AdminScp:     adminScp,
	}

	//Call process by sdk
	_, err := c.UpdateSystemSettingGlobal(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System Setting Global: %s", err)
	}

	d.SetId(hostname)

	return resourceSystemSettingGlobalRead(d, m)
}

func resourceSystemSettingGlobalDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
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

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("admintimeout", o.Admintimeout)
	d.Set("timezone", o.Timezone)
	d.Set("hostname", o.Hostname)
	d.Set("admin_sport", o.AdminSport)
	d.Set("admin_ssh_port", o.AdminSSHPort)
	d.Set("admin_scp", o.AdminScp)

	return nil
}
