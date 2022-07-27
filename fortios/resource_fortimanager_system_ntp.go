package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFortimanagerSystemNTP() *schema.Resource {
	return &schema.Resource{
		Create: setFMGSystemNTP,
		Read:   readFMGSystemNTP,
		Update: setFMGSystemNTP,
		Delete: deleteFMGSystemNTP,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
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

	c.DebugNum++
	log.Printf("ClientFortimanager setFMGSystemNTP: %d", c.DebugNum)

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
