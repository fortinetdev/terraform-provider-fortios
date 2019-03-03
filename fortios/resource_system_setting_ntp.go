package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemSettingNTP() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSettingNTPCreate,
		Read:   resourceSystemSettingNTPRead,
		Update: resourceSystemSettingNTPUpdate,
		Delete: resourceSystemSettingNTPDelete,

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

func resourceSystemSettingNTPCreate(d *schema.ResourceData, m interface{}) error {
	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Get Params from d
	// typef := d.Get("type").(string)
	// ntpserver := d.Get("ntpserver").([]interface{})

	// var ntpservers []forticlient.MultValue

	// for _, v := range ntpserver {
	// 	ntpservers = append(ntpservers,
	// 		forticlient.MultValue{
	// 			Name: v.(string),
	// 		})
	// }

	// //Build input data by sdk
	// i := &forticlient.JSONSystemSettingNTP{
	// 	Type:      typef,
	// 	Ntpserver: ntpservers,
	// 	Ntpsync:       ntpsync,
	// }

	// //Call process by sdk
	// o, err := c.CreateSystemSettingNTP(i)
	// if err != nil {
	// 	return fmt.Errorf("Error creating System Setting NTP: %s", err)
	// }

	// // Set index for d
	// // d.SetId(strconv.Itoa(int(o.Mkey)))
	// d.SetId(o.Mkey)

	err := resourceSystemSettingNTPUpdate(d, m)
	if err != nil {
		return fmt.Errorf("Error updating System Setting NTP: %s", err)
	}

	return nil
}

func resourceSystemSettingNTPUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d

	typef := d.Get("type").(string)
	ntpserver := d.Get("ntpserver").([]interface{})
	ntpsync := d.Get("ntpsync").(string)

	var ntpservers []forticlient.NTPMultValue

	for _, v := range ntpserver {
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

	//Set index for d
	//d.SetId(o.Mkey)
	d.SetId("1")

	return nil
}

func resourceSystemSettingNTPDelete(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// err := c.DeleteSystemSettingNTP(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error deleting System Setting NTP: %s", err)
	// }

	// //Set index for d
	d.SetId("")

	return nil
}

func resourceSystemSettingNTPRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemSettingNTP(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System Setting NTP: %s", err)
	}

	//Refresh property
	d.Set("type", o.Type)

	// vs := make([]string, 0, len(o.Ntpserver))
	// for _, v := range o.Ntpserver {
	// 	c := v.Server
	// 	vs = append(vs, c)
	// }

	// d.Set("ntpserver", vs)
	d.Set("ntpsync", o.Ntpsync)

	return nil
}
