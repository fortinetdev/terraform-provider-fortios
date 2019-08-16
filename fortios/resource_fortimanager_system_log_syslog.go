package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemSyslog() *schema.Resource {
	return &schema.Resource{
		Create: createFTMSystemSyslog,
		Read:   readFTMSystemSyslog,
		Update: updateFTMSystemSyslog,
		Delete: deleteFTMSystemSyslog,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createFTMSystemSyslog(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMSystemSyslog")()

	//Build input data by sdk
	i := &fortimngclient.JSONSystemSyslog{
		Name: d.Get("name").(string),
		Ip:   d.Get("ip").(string),
		Port: d.Get("port").(string),
	}

	err := c.CreateUpdateSystemSyslog(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Syslog: %s", err)
	}

	d.SetId(i.Name)

	return readFTMSystemSyslog(d, m)
}

func readFTMSystemSyslog(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMSystemSyslog")()

	name := d.Id()
	o, err := c.ReadSystemSyslog(name)
	if err != nil {
		return fmt.Errorf("Error reading System Syslog: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("ip", o.Ip)
	d.Set("port", o.Port)

	return nil
}

func updateFTMSystemSyslog(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMSystemSyslog")()

	if d.HasChange("name") {
		return fmt.Errorf("the userid argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONSystemSyslog{
		Name: d.Get("name").(string),
		Ip:   d.Get("ip").(string),
		Port: d.Get("port").(string),
	}

	err := c.CreateUpdateSystemSyslog(i, "update")

	if err != nil {
		return fmt.Errorf("Error updating System Syslog: %s", err)
	}

	return readFTMSystemSyslog(d, m)
}

func deleteFTMSystemSyslog(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMSystemSyslog")()

	name := d.Id()

	err := c.DeleteSystemSyslog(name)
	if err != nil {
		return fmt.Errorf("Error deleting System Syslog: %s", err)
	}

	d.SetId("")

	return nil
}
