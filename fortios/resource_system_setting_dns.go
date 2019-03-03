package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemSettingDNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSettingDNSCreate,
		Read:   resourceSystemSettingDNSRead,
		Update: resourceSystemSettingDNSUpdate,
		Delete: resourceSystemSettingDNSDelete,

		Schema: map[string]*schema.Schema{
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "208.91.112.53",
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "208.91.112.52",
			},
		},
	}
}

func resourceSystemSettingDNSCreate(d *schema.ResourceData, m interface{}) error {
	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Get Params from d
	// primary := d.Get("primary").(string)
	// secondary := d.Get("secondary").(string)

	// //Build input data by sdk
	// i := &forticlient.JSONSystemSettingDNS{
	// 	Primary:   primary,
	// 	Secondary: secondary,
	// }

	// //Call process by sdk
	// o, err := c.CreateSystemSettingDNS(i)
	// if err != nil {
	// 	return fmt.Errorf("Error creating System Setting DNS: %s", err)
	// }

	// // Set index for d
	// // d.SetId(strconv.Itoa(int(o.Mkey)))
	// d.SetId(o.Mkey)

	err := resourceSystemSettingDNSUpdate(d, m)
	if err != nil {
		return fmt.Errorf("Error updating System Setting DNS: %s", err)
	}

	return nil
}

func resourceSystemSettingDNSUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d

	primary := d.Get("primary").(string)
	secondary := d.Get("secondary").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemSettingDNS{
		Primary:   primary,
		Secondary: secondary,
	}

	//Call process by sdk
	_, err := c.UpdateSystemSettingDNS(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System Setting DNS: %s", err)
	}

	//Set index for d
	//d.SetId(o.Mkey)
	d.SetId("1")

	return nil
}

func resourceSystemSettingDNSDelete(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// err := c.DeleteSystemSettingDNS(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error deleting System Setting DNS: %s", err)
	// }

	// //Set index for d
	d.SetId("")

	return nil
}

func resourceSystemSettingDNSRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemSettingDNS(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System Setting DNS: %s", err)
	}

	//Refresh property
	d.Set("primary", o.Primary)
	d.Set("secondary", o.Secondary)

	return nil
}
