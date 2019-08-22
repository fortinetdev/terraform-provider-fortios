package fortios

import (
	"fmt"
	"log"
	"strconv"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemAdom() *schema.Resource {
	return &schema.Resource{
		Create: createFTMSystemAdom,
		Read:   readFTMSystemAdom,
		Update: updateFTMSystemAdom,
		Delete: deleteFTMSystemAdom,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func createFTMSystemAdom(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMSystemAdom")()

	i := &fortimngclient.JSONSystemAdom{
		Name:        d.Get("name").(string),
		State:       strconv.Itoa(d.Get("state").(int)),
		Description: d.Get("description").(string),
	}

	err := c.CreateUpdateSystemAdom(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Adom: %s", err)
	}

	d.SetId(i.Name)

	return readFTMSystemAdom(d, m)
}

func readFTMSystemAdom(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMSystemAdom")()

	name := d.Id()
	o, err := c.ReadSystemAdom(name)
	if err != nil {
		return fmt.Errorf("Error reading System Adom: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("state", o.State)

	return nil
}

func updateFTMSystemAdom(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMSystemAdom")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONSystemAdom{
		Name:        d.Get("name").(string),
		State:       strconv.Itoa(d.Get("state").(int)),
		Description: d.Get("description").(string),
	}

	err := c.CreateUpdateSystemAdom(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating System Adom: %s", err)
	}

	return readFTMSystemAdom(d, m)
}

func deleteFTMSystemAdom(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMSystemAdom")()

	name := d.Id()

	err := c.DeleteSystemAdom(name)
	if err != nil {
		return fmt.Errorf("Error deleting System Adom: %s", err)
	}

	d.SetId("")

	return nil
}
