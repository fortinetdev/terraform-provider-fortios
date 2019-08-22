package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerFirewallSecurityAdomRevision() *schema.Resource {
	return &schema.Resource{
		Create: createFTMFirewallSecurityAdomRevision,
		Read:   readFTMFirewallSecurityAdomRevision,
		Update: updateFTMFirewallSecurityAdomRevision,
		Delete: deleteFTMFirewallSecurityAdomRevision,

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
		},
	}
}

func createFTMFirewallSecurityAdomRevision(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMFirewallSecurityAdomRevision")()

	i := &fortimngclient.JSONFirewallSecurityAdomRevision{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		CreatedBy:   d.Get("created_by").(string),
		Locked:      d.Get("locked").(int),
		Version:     "0",
	}

	version, err := c.CreateUpdateFirewallSecurityAdomRevision(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating Firewall Security Adom Revision: %s", err)
	}

	d.SetId(version)

	return readFTMFirewallSecurityAdomRevision(d, m)
}

func readFTMFirewallSecurityAdomRevision(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMFirewallSecurityAdomRevision")()

	version := d.Id()
	o, err := c.ReadFirewallSecurityAdomRevision(version)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Security Adom Revision: %s", err)
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

func updateFTMFirewallSecurityAdomRevision(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMFirewallSecurityAdomRevision")()

	i := &fortimngclient.JSONFirewallSecurityAdomRevision{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		CreatedBy:   d.Get("created_by").(string),
		Locked:      d.Get("locked").(int),
		Version:     d.Id(),
	}

	_, err := c.CreateUpdateFirewallSecurityAdomRevision(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating Firewall Security Adom Revision: %s", err)
	}

	return readFTMFirewallSecurityAdomRevision(d, m)
}

func deleteFTMFirewallSecurityAdomRevision(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMFirewallSecurityAdomRevision")()

	version := d.Id()

	err := c.DeleteFirewallSecurityAdomRevision(version)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Security Adom Revision: %s", err)
	}

	d.SetId("")

	return nil
}
