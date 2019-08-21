package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerFirewallObjectAddress() *schema.Resource {
	return &schema.Resource{
		Create: createFTMFirewallObjectAddress,
		Read:   readFTMFirewallObjectAddress,
		Update: updateFTMFirewallObjectAddress,
		Delete: deleteFTMFirewallObjectAddress,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createFTMFirewallObjectAddress(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMFirewallObjectAddress")()

	i := &fortimngclient.JSONFirewallObjectAddress{
		Name:           d.Get("name").(string),
		Type:           d.Get("type").(string),
		Comment:        d.Get("comment").(string),
		Fqdn:           d.Get("fqdn").(string),
		AssociatedIntf: d.Get("associated_intf").(string),
	}

	err := c.CreateUpdateFirewallObjectAddress(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object Address: %s", err)
	}

	d.SetId(i.Name)

	return readFTMFirewallObjectAddress(d, m)
}

func readFTMFirewallObjectAddress(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMFirewallObjectAddress")()

	name := d.Id()
	o, err := c.ReadFirewallObjectAddress(name)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object Address: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("type", o.Type)
	d.Set("comment", o.Comment)
	d.Set("fqdn", o.Fqdn)
	d.Set("associated-intf", o.AssociatedIntf)

	return nil
}

func updateFTMFirewallObjectAddress(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMFirewallObjectAddress")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONFirewallObjectAddress{
		Name:           d.Get("name").(string),
		Type:           d.Get("type").(string),
		Comment:        d.Get("comment").(string),
		Fqdn:           d.Get("fqdn").(string),
		AssociatedIntf: d.Get("associated_intf").(string),
	}

	err := c.CreateUpdateFirewallObjectAddress(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Address: %s", err)
	}

	return readFTMFirewallObjectAddress(d, m)
}

func deleteFTMFirewallObjectAddress(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMFirewallObjectAddress")()

	name := d.Id()

	err := c.DeleteFirewallObjectAddress(name)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Address: %s", err)
	}

	d.SetId("")

	return nil
}
