package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/fgtdev/fortimanager-sdk-go/util"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerDVMScript() *schema.Resource {
	return &schema.Resource{
		Create: createFMGDVMScript,
		Read:   readFMGDVMScript,
		Update: updateFMGDVMScript,
		Delete: deleteFMGDVMScript,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "device_database",
				ValidateFunc: util.ValidateStringIn("device_database", "remote_device", "adom_database"),
			},
		},
	}
}

func createFMGDVMScript(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGDVMScript")()

	//Build input data by sdk
	i := &fmgclient.JSONDVMScript{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Content:     d.Get("content").(string),
		Target:      d.Get("target").(string),
		Type:        "cli",
	}

	err := c.CreateUpdateDVMScript(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating devicemanager script: %s", err)
	}

	d.SetId(i.Name)

	return readFMGDVMScript(d, m)
}

func readFMGDVMScript(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGDVMScript")()

	name := d.Id()
	o, err := c.ReadDVMScript(name)
	if err != nil {
		return fmt.Errorf("Error reading devicemanager script: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("content", o.Content)
	d.Set("target", o.Target)

	return nil
}

func updateFMGDVMScript(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGDVMScript")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	//Build input data by sdk
	i := &fmgclient.JSONDVMScript{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Content:     d.Get("content").(string),
		Target:      d.Get("target").(string),
		Type:        "cli",
	}

	err := c.CreateUpdateDVMScript(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating devicemanager script: %s", err)
	}

	return readFMGDVMScript(d, m)
}

func deleteFMGDVMScript(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGDVMScript")()

	name := d.Id()

	err := c.DeleteDVMScript(name)
	if err != nil {
		return fmt.Errorf("Error deleting devicemanager script: %s", err)
	}

	d.SetId("")

	return nil
}
