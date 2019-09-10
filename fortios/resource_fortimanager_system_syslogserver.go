package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemSyslogServer() *schema.Resource {
	return &schema.Resource{
		Create: createFMGSystemSyslogServer,
		Read:   readFMGSystemSyslogServer,
		Update: updateFMGSystemSyslogServer,
		Delete: deleteFMGSystemSyslogServer,

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
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func createFMGSystemSyslogServer(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGSystemSyslogServer")()

	i := &fortimngclient.JSONSystemSyslogServer{
		Name: d.Get("name").(string),
		Ip:   d.Get("ip").(string),
		Port: d.Get("port").(int),
	}

	err := c.CreateUpdateSystemSyslogServer(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Syslog Server: %s", err)
	}

	d.SetId(i.Name)

	return readFMGSystemSyslogServer(d, m)
}

func readFMGSystemSyslogServer(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemSyslogServer")()

	name := d.Id()
	o, err := c.ReadSystemSyslogServer(name)
	if err != nil {
		return fmt.Errorf("Error reading System Syslog Server: %s", err)
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

func updateFMGSystemSyslogServer(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGSystemSyslogServer")()

	if d.HasChange("name") {
		return fmt.Errorf("the userid argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONSystemSyslogServer{
		Name: d.Get("name").(string),
		Ip:   d.Get("ip").(string),
		Port: d.Get("port").(int),
	}

	err := c.CreateUpdateSystemSyslogServer(i, "update")

	if err != nil {
		return fmt.Errorf("Error updating System Syslog Server: %s", err)
	}

	return readFMGSystemSyslogServer(d, m)
}

func deleteFMGSystemSyslogServer(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGSystemSyslogServer")()

	name := d.Id()

	err := c.DeleteSystemSyslogServer(name)
	if err != nil {
		return fmt.Errorf("Error deleting System Syslog Server: %s", err)
	}

	d.SetId("")

	return nil
}
