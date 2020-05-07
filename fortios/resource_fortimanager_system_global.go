package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/fgtdev/fortimanager-sdk-go/util"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemGlobal() *schema.Resource {
	return &schema.Resource{
		Create: setFMGSystemGlobal,
		Read:   readFMGSystemGlobal,
		Update: setFMGSystemGlobal,
		Delete: deleteFMGSystemGlobal,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fortianalyzer_status": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.ValidateStringIn("disable", "enable"),
			},
			"adom_status": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.ValidateStringIn("disable", "enable"),
			},
			"adom_mode": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.ValidateStringIn("normal", "advanced"),
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timezone": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.ValidateStringIn("00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89"),
			},
		},
	}
}

func setFMGSystemGlobal(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("setFMGSystemGlobal")()

	log.Printf("shengh: setFMGSystemGlobal Init")

	i := &fmgclient.JSONSystemGlobal{
		FazStatus:  d.Get("fortianalyzer_status").(string),
		AdomStatus: d.Get("adom_status").(string),
		AdomMode:   d.Get("adom_mode").(string),
		Hostname:   d.Get("hostname").(string),
		TimeZone:   d.Get("timezone").(string),
	}

	err := c.SetSystemGlobal(i)
	if err != nil {
		return fmt.Errorf("Error setting System Global : %s", err)
	}

	c.DebugNum++
	log.Printf("shengh: ClientFortimanager setFMGSystemGlobal: %d", c.DebugNum)

	d.SetId("fortimanager-global-setting")

	return readFMGSystemGlobal(d, m)
}

func readFMGSystemGlobal(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemGlobal")()

	o, err := c.ReadSystemGlobal()
	if err != nil {
		return fmt.Errorf("Error reading System Global : %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if d.Get("fortianalyzer_status").(string) != "" {
		d.Set("fortianalyzer_status", o.FazStatus)
	}
	if d.Get("hostname").(string) != "" {
		d.Set("hostname", o.Hostname)
	}
	if d.Get("adom_status").(string) != "" {
		d.Set("adom_status", o.AdomStatus)
	}
	if d.Get("adom_mode").(string) != "" {
		d.Set("adom_mode", o.AdomMode)
	}
	if d.Get("timezone").(string) != "" {
		d.Set("timezone", o.TimeZone)
	}

	return nil
}

func deleteFMGSystemGlobal(d *schema.ResourceData, m interface{}) error {
	return nil
}
