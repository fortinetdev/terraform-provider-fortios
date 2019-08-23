package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemGlobalSetting() *schema.Resource {
	return &schema.Resource{
		Create: setFTMSystemGlobalSetting,
		Read:   readFTMSystemGlobalSetting,
		Update: setFTMSystemGlobalSetting,
		Delete: deleteFTMSystemGlobalSetting,

		Schema: map[string]*schema.Schema{
			"faz_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"adom_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func setFTMSystemGlobalSetting(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("setFTMSystemGlobalSetting")()

	i := &fortimngclient.JSONSystemGlobalSetting{
		Hostname:   d.Get("hostname").(string),
		FazStatus:  d.Get("faz_status").(string),
		AdomStatus: d.Get("adom_status").(string),
	}

	err := c.SetSystemGlobalSetting(i)
	if err != nil {
		return fmt.Errorf("Error setting System Global Setting: %s", err)
	}

	// Set an id for this resource
	d.SetId("99")

	return readFTMSystemGlobalSetting(d, m)
}

func readFTMSystemGlobalSetting(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMSystemGlobalSetting")()

	o, err := c.ReadSystemGlobalSetting()
	if err != nil {
		return fmt.Errorf("Error reading System Global Setting: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("hostname", o.Hostname)
	d.Set("faz_status", o.FazStatus)
	d.Set("adom_status", o.AdomStatus)

	return nil
}

func deleteFTMSystemGlobalSetting(d *schema.ResourceData, m interface{}) error {
	return nil
}
