package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerObjectAdomRevision() *schema.Resource {
	return &schema.Resource{
		Create: createFMGObjectAdomRevision,
		Read:   readFMGObjectAdomRevision,
		Update: updateFMGObjectAdomRevision,
		Delete: deleteFMGObjectAdomRevision,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"locked": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
		},
	}
}

func createFMGObjectAdomRevision(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGObjectAdomRevision")()

	i := &fmgclient.JSONObjectAdomRevision{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		CreatedBy:   d.Get("created_by").(string),
		Locked:      d.Get("locked").(int),
		Version:     "0",
	}

	version, err := c.CreateUpdateObjectAdomRevision(i, "add", d.Get("adom").(string))
	if err != nil {
		return fmt.Errorf("Error creating Object Adom Revision: %s", err)
	}

	d.SetId(version)

	return readFMGObjectAdomRevision(d, m)
}

func readFMGObjectAdomRevision(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGObjectAdomRevision")()

	version := d.Id()
	o, err := c.ReadObjectAdomRevision(d.Get("adom").(string), version)
	if err != nil {
		return fmt.Errorf("Error reading Object Adom Revision: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("locked", o.Locked)
	d.Set("created_by", o.CreatedBy)

	return nil
}

func updateFMGObjectAdomRevision(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGObjectAdomRevision")()

	i := &fmgclient.JSONObjectAdomRevision{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		CreatedBy:   d.Get("created_by").(string),
		Locked:      d.Get("locked").(int),
		Version:     d.Id(),
	}

	_, err := c.CreateUpdateObjectAdomRevision(i, "update", d.Get("adom").(string))
	if err != nil {
		return fmt.Errorf("Error updating Object Adom Revision: %s", err)
	}

	return readFMGObjectAdomRevision(d, m)
}

func deleteFMGObjectAdomRevision(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGObjectAdomRevision")()

	version := d.Id()

	err := c.DeleteObjectAdomRevision(d.Get("adom").(string), version)
	if err != nil {
		return fmt.Errorf("Error deleting Object Adom Revision: %s", err)
	}

	d.SetId("")

	return nil
}
