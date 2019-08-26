package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemNTPSetting() *schema.Resource {
	return &schema.Resource{
		Create: setFTMSystemNTPSetting,
		Read:   readFTMSystemNTPSetting,
		Update: setFTMSystemNTPSetting,
		Delete: deleteFTMSystemNTPSetting,

		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sync_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func setFTMSystemNTPSetting(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("setFTMSystemNTPSetting")()

	i := &fortimngclient.JSONSystemNTPSetting{
		Id:           1,
		Server:       d.Get("server").(string),
		Status:       d.Get("status").(string),
		SyncInterval: d.Get("sync_interval").(int),
	}

	err := c.SetSystemNTPSetting(i)
	if err != nil {
		return fmt.Errorf("Error setting System NTP Setting: %s", err)
	}

	d.SetId("fortimanager-ntp-setting")

	return readFTMSystemNTPSetting(d, m)
}

func readFTMSystemNTPSetting(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMSystemNTPSetting")()

	o, err := c.ReadSystemNTPSetting()
	if err != nil {
		return fmt.Errorf("Error reading System NTP Setting: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("server", o.Server)
	if d.Get("status") != "" {
		d.Set("status", o.Status)
	}
	if d.Get("sync_interval") != "" {
		d.Set("sync_interval", o.SyncInterval)
	}

	return nil
}

func deleteFTMSystemNTPSetting(d *schema.ResourceData, m interface{}) error {
	return nil
}
