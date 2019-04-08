package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemAdminAdministrator() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminAdministratorCreate,
		Read:   resourceSystemAdminAdministratorRead,
		Update: resourceSystemAdminAdministratorUpdate,
		Delete: resourceSystemAdminAdministratorDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"accprofile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceSystemAdminAdministratorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	password := d.Get("password").(string)
	trusthost1 := d.Get("trusthost1").(string)
	trusthost2 := d.Get("trusthost2").(string)
	trusthost3 := d.Get("trusthost3").(string)
	trusthost4 := d.Get("trusthost4").(string)
	trusthost5 := d.Get("trusthost5").(string)
	trusthost6 := d.Get("trusthost6").(string)
	trusthost7 := d.Get("trusthost7").(string)
	trusthost8 := d.Get("trusthost8").(string)
	trusthost9 := d.Get("trusthost9").(string)
	trusthost10 := d.Get("trusthost10").(string)
	accprofile := d.Get("accprofile").(string)
	comments := d.Get("comments").(string)
	vdom := d.Get("vdom").([]interface{})

	var vdoms []forticlient.MultValue

	for _, v := range vdom {
		if v == nil {
			return fmt.Errorf("null value")
		}
		vdoms = append(vdoms,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONSystemAdminAdministrator{
		Name:        name,
		Password:    password,
		Trusthost1:  trusthost1,
		Trusthost2:  trusthost2,
		Trusthost3:  trusthost3,
		Trusthost4:  trusthost4,
		Trusthost5:  trusthost5,
		Trusthost6:  trusthost6,
		Trusthost7:  trusthost7,
		Trusthost8:  trusthost8,
		Trusthost9:  trusthost9,
		Trusthost10: trusthost10,
		Accprofile:  accprofile,
		Comments:    comments,
		Vdom:        vdoms,
	}

	//Call process by sdk
	o, err := c.CreateSystemAdminAdministrator(i)
	if err != nil {
		return fmt.Errorf("Error creating System Admin Administrator: %s", err)
	}

	//Set index for d
	d.SetId(o.Mkey)

	return resourceSystemAdminAdministratorRead(d, m)
}

func resourceSystemAdminAdministratorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d

	name := d.Get("name").(string)
	//password := d.Get("password").(string)
	trusthost1 := d.Get("trusthost1").(string)
	trusthost2 := d.Get("trusthost2").(string)
	trusthost3 := d.Get("trusthost3").(string)
	trusthost4 := d.Get("trusthost4").(string)
	trusthost5 := d.Get("trusthost5").(string)
	trusthost6 := d.Get("trusthost6").(string)
	trusthost7 := d.Get("trusthost7").(string)
	trusthost8 := d.Get("trusthost8").(string)
	trusthost9 := d.Get("trusthost9").(string)
	trusthost10 := d.Get("trusthost10").(string)
	accprofile := d.Get("accprofile").(string)
	comments := d.Get("comments").(string)
	vdom := d.Get("vdom").([]interface{})

	var vdoms []forticlient.MultValue

	for _, v := range vdom {
		if v == nil {
			return fmt.Errorf("null value")
		}
		vdoms = append(vdoms,
			forticlient.MultValue{
				Name: v.(string),
			})
	}
	//Build input data by sdk
	i := &forticlient.JSONSystemAdminAdministrator2{
		Name: name,
		// Password:    password,
		Trusthost1:  trusthost1,
		Trusthost2:  trusthost2,
		Trusthost3:  trusthost3,
		Trusthost4:  trusthost4,
		Trusthost5:  trusthost5,
		Trusthost6:  trusthost6,
		Trusthost7:  trusthost7,
		Trusthost8:  trusthost8,
		Trusthost9:  trusthost9,
		Trusthost10: trusthost10,
		Accprofile:  accprofile,
		Comments:    comments,
		Vdom:        vdoms,
	}

	//Call process by sdk
	_, err := c.UpdateSystemAdminAdministrator(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System Admin Administrator: %s", err)
	}

	return resourceSystemAdminAdministratorRead(d, m)
}

func resourceSystemAdminAdministratorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteSystemAdminAdministrator(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting System Admin Administrator: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceSystemAdminAdministratorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemAdminAdministrator(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System Admin Administrator: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("name", o.Name)
	//d.Set("password", o.Password)
	d.Set("trusthost1", o.Trusthost1)
	d.Set("trusthost2", o.Trusthost2)
	d.Set("trusthost3", o.Trusthost3)
	d.Set("trusthost4", o.Trusthost4)
	d.Set("trusthost5", o.Trusthost5)
	d.Set("trusthost6", o.Trusthost6)
	d.Set("trusthost7", o.Trusthost7)
	d.Set("trusthost8", o.Trusthost8)
	d.Set("trusthost9", o.Trusthost9)
	d.Set("trusthost10", o.Trusthost10)
	d.Set("accprofile", o.Accprofile)
	d.Set("comments", o.Comments)
	vdom := forticlient.ExtractString(o.Vdom)
	if err := d.Set("vdom", vdom); err != nil {
		log.Printf("[WARN] Error setting System Admin Administrator for (%s): %s", d.Id(), err)
	}

	return nil
}
