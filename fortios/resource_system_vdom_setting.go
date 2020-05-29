package fortios

import (
	"fmt"
	"log"

	"github.com/fortinetdev/forti-sdk-go/fortios/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSystemVdomSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomSettingCreate,
		Read:   resourceSystemVdomSettingRead,
		Update: resourceSystemVdomSettingUpdate,
		Delete: resourceSystemVdomSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"short_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"temporary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemVdomSettingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	shortName := d.Get("short_name").(string)
	temporary := d.Get("temporary").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemVdomSetting{
		Name:      name,
		ShortName: shortName,
		Temporary: temporary,
	}

	//Call process by sdk
	o, err := c.CreateSystemVdomSetting(i)
	if err != nil {
		return fmt.Errorf("Error creating System Vdom Setting: %s", err)
	}

	// Set index for d
	d.SetId(o.Mkey)

	return resourceSystemVdomSettingRead(d, m)
}

func resourceSystemVdomSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	shortName := d.Get("short_name").(string)
	temporary := d.Get("temporary").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemVdomSetting{
		Name:      name,
		ShortName: shortName,
		Temporary: temporary,
	}

	//Call process by sdk
	_, err := c.UpdateSystemVdomSetting(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System Vdom Setting: %s", err)
	}

	return resourceSystemVdomSettingRead(d, m)
}

func resourceSystemVdomSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	err := c.DeleteSystemVdomSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting System Vdom Setting: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceSystemVdomSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemVdomSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System Vdom Setting: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("name", o.Name)
	d.Set("short_name", o.ShortName)
	d.Set("temporary", o.Temporary)

	return nil
}
