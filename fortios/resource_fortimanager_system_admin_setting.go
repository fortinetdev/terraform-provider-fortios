package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemAdminSetting() *schema.Resource {
	return &schema.Resource{
		Create: setFTMSystemAdminSetting,
		Read:   readFTMSystemAdminSetting,
		Update: setFTMSystemAdminSetting,
		Delete: deleteFTMSystemAdminSetting,

		Schema: map[string]*schema.Schema{
			"http_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
		},
	}
}

func setFTMSystemAdminSetting(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("setFTMSystemAdminSetting")()

	i := &fortimngclient.JSONSystemAdminSetting{
		HttpPort:    d.Get("http_port").(int),
		HttpsPort:   d.Get("https_port").(int),
		IdleTimeout: d.Get("idle_timeout").(int),
	}

	err := c.SetSystemAdminSetting(i)
	if err != nil {
		return fmt.Errorf("Error setting System Admin Setting: %s", err)
	}

	// Set an id for this resource
	d.SetId("98")

	return readFTMSystemAdminSetting(d, m)
}

func readFTMSystemAdminSetting(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMSystemAdminSetting")()

	o, err := c.ReadSystemAdminSetting()
	if err != nil {
		return fmt.Errorf("Error reading System Admin Setting: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("http_port", o.HttpPort)
	d.Set("https_port", o.HttpsPort)
	d.Set("idle_timeout", o.IdleTimeout)

	return nil
}

func deleteFTMSystemAdminSetting(d *schema.ResourceData, m interface{}) error {
	return nil
}
