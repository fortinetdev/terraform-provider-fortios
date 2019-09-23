package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/fgtdev/fortimanager-sdk-go/util"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemNTP() *schema.Resource {
	return &schema.Resource{
		Create: setFMGSystemNTP,
		Read:   readFMGSystemNTP,
		Update: setFMGSystemNTP,
		Delete: deleteFMGSystemNTP,

		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.ValidateStringIn("disable", "enable"),
			},
			"sync_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func setFMGSystemNTP(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("setFMGSystemNTP")()

	i := &fmgclient.JSONSystemNTP{
		Id:           1,
		Server:       d.Get("server").(string),
		Status:       d.Get("status").(string),
		SyncInterval: d.Get("sync_interval").(int),
	}

	err := c.SetSystemNTP(i)
	if err != nil {
		return fmt.Errorf("Error setting System NTP : %s", err)
	}

	d.SetId("fortimanager-ntp-setting")

	return readFMGSystemNTP(d, m)
}

func readFMGSystemNTP(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemNTP")()

	o, err := c.ReadSystemNTP()
	if err != nil {
		return fmt.Errorf("Error reading System NTP : %s", err)
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

func deleteFMGSystemNTP(d *schema.ResourceData, m interface{}) error {
	return nil
}
