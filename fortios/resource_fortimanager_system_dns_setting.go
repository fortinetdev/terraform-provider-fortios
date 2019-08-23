package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemDNSSetting() *schema.Resource {
	return &schema.Resource{
		Create: setFTMSystemDNSSetting,
		Read:   readFTMSystemDNSSetting,
		Update: setFTMSystemDNSSetting,
		Delete: deleteFTMSystemDNSSetting,

		Schema: map[string]*schema.Schema{
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func setFTMSystemDNSSetting(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("setFTMSystemDNSSetting")()

	i := &fortimngclient.JSONSystemDNSSetting{
		Primary:   d.Get("primary").(string),
		Secondary: d.Get("secondary").(string),
	}

	err := c.SetSystemDNSSetting(i)
	if err != nil {
		return fmt.Errorf("Error setting System DNS Setting: %s", err)
	}

	d.SetId("fortimanager-dns-setting")

	return readFTMSystemDNSSetting(d, m)
}

func readFTMSystemDNSSetting(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMSystemDNSSetting")()

	o, err := c.ReadSystemDNSSetting()
	if err != nil {
		return fmt.Errorf("Error reading System DNS Setting: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if d.Get("primary") != "" {
		d.Set("primary", o.Primary)
	}
	if d.Get("secondary") != "" {
		d.Set("secondary", o.Secondary)
	}

	return nil
}

func deleteFTMSystemDNSSetting(d *schema.ResourceData, m interface{}) error {
	return nil
}
