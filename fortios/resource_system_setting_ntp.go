package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemSettingNTP() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSettingNTPCreateUpdate,
		Read:   resourceSystemSettingNTPRead,
		Update: resourceSystemSettingNTPCreateUpdate,
		Delete: resourceSystemSettingNTPDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ntpserver": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ntpsync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "enable",
			},
		},
	}
}

func resourceSystemSettingNTPCreateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	typef := d.Get("type").(string)
	ntpserver := d.Get("ntpserver").([]interface{})
	ntpsync := d.Get("ntpsync").(string)

	var ntpservers []forticlient.NTPMultValue

	for _, v := range ntpserver {
		if v == nil {
			return fmt.Errorf("null ntpserver")
		}
		ntpservers = append(ntpservers,
			forticlient.NTPMultValue{
				Server: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONSystemSettingNTP{
		Type:      typef,
		Ntpserver: ntpservers,
		Ntpsync:   ntpsync,
	}

	//Call process by sdk
	_, err := c.UpdateSystemSettingNTP(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System Setting NTP: %s", err)
	}

	d.SetId(typef)

	return resourceSystemSettingNTPRead(d, m)
}

func resourceSystemSettingNTPDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemSettingNTPRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemSettingNTP(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System Setting NTP: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("type", o.Type)
	d.Set("ntpsync", o.Ntpsync)

	return nil
}
