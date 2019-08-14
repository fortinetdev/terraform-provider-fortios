package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerDVMScript() *schema.Resource {
	return &schema.Resource{
		Create: createFTMDVMScript,
		Read:   readFTMDVMScript,
		Update: updateFTMDVMScript,
		Delete: deleteFTMDVMScript,

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
			"execute": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"exec_target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createFTMDVMScript(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMDVMScript")()

	//Build input data by sdk
	i := &fortimngclient.JSONDVMScript{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Content:     d.Get("content").(string),
		Type:        "cli",
	}

	err := c.CreateUpdateDVMScript(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating devicemanager script: %s", err)
	}

	d.SetId(i.Name)

	if d.Get("execute").(bool) {
		target := d.Get("exec_target").(string)
		if target == "" {
			return fmt.Errorf("exec_target should be set when execute is true: %s", err)
		}

		params := &fortimngclient.JSONDVMScriptExecute{
			Target:     target,
			ScriptName: d.Get("name").(string),
		}

		err := c.ExecuteDVMScript(params)

		if err != nil {
			return fmt.Errorf("Error executing devicemanager script: %s", err)
		}
	}

	return readFTMDVMScript(d, m)
}

func readFTMDVMScript(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMDVMScript")()

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

	return nil
}

func updateFTMDVMScript(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMDVMScript")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	//Build input data by sdk
	i := &fortimngclient.JSONDVMScript{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Content:     d.Get("content").(string),
		Type:        "cli",
	}

	err := c.CreateUpdateDVMScript(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating devicemanager script: %s", err)
	}

	if d.Get("execute").(bool) {
		target := d.Get("exec_target").(string)
		if target == "" {
			return fmt.Errorf("exec_target should be set when execute is true: %s", err)
		}

		params := &fortimngclient.JSONDVMScriptExecute{
			Target:     target,
			ScriptName: d.Get("name").(string),
		}

		err := c.ExecuteDVMScript(params)

		if err != nil {
			return fmt.Errorf("Error executing devicemanager script: %s", err)
		}
	}

	return readFTMDVMScript(d, m)
}

func deleteFTMDVMScript(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMDVMScript")()

	name := d.Id()

	err := c.DeleteDVMScript(name)
	if err != nil {
		return fmt.Errorf("Error deleting devicemanager script: %s", err)
	}

	d.SetId("")

	return nil
}
