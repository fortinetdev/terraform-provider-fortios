package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemAdmin() *schema.Resource {
	return &schema.Resource{
		Create: setFMGSystemAdmin,
		Read:   readFMGSystemAdmin,
		Update: setFMGSystemAdmin,
		Delete: deleteFMGSystemAdmin,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func setFMGSystemAdmin(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("setFMGSystemAdmin")()

	i := &fmgclient.JSONSystemAdmin{
		HttpPort:    d.Get("http_port").(int),
		HttpsPort:   d.Get("https_port").(int),
		IdleTimeout: d.Get("idle_timeout").(int),
	}

	err := c.SetSystemAdmin(i)
	if err != nil {
		return fmt.Errorf("Error setting System Admin : %s", err)
	}

	// Set an id for this resource
	d.SetId("fortimanager-system-admin")

	return readFMGSystemAdmin(d, m)
}

func readFMGSystemAdmin(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemAdmin")()

	o, err := c.ReadSystemAdmin()
	if err != nil {
		return fmt.Errorf("Error reading System Admin : %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if d.Get("http_port").(int) != 0 {
		d.Set("http_port", o.HttpPort)
	}
	if d.Get("https_port").(int) != 0 {
		d.Set("https_port", o.HttpsPort)
	}
	if d.Get("idle_timeout").(int) != 0 {
		d.Set("idle_timeout", o.IdleTimeout)
	}

	return nil
}

func deleteFMGSystemAdmin(d *schema.ResourceData, m interface{}) error {
	return nil
}
