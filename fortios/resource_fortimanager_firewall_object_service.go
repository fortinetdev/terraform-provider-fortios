package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerFirewallObjectService() *schema.Resource {
	return &schema.Resource{
		Create: createFTMFirewallObjectService,
		Read:   readFTMFirewallObjectService,
		Update: updateFTMFirewallObjectService,
		Delete: deleteFTMFirewallObjectService,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "TCP/UDP/SCTP",
			},
		},
	}
}

func createFTMFirewallObjectService(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMFirewallObjectService")()

	i := &fortimngclient.JSONFirewallObjectService{
		Name:     d.Get("name").(string),
		Comment:  d.Get("comment").(string),
		Category: d.Get("category").(string),
		Protocol: d.Get("protocol").(string),
	}

	err := c.CreateUpdateFirewallObjectService(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object Service: %s", err)
	}

	d.SetId(i.Name)

	return readFTMFirewallObjectService(d, m)
}

func readFTMFirewallObjectService(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMFirewallObjectService")()

	name := d.Id()
	o, err := c.ReadFirewallObjectService(name)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object Service: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("comment", o.Comment)
	d.Set("protocol", o.Protocol)
	d.Set("category", o.Category)

	return nil
}

func updateFTMFirewallObjectService(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMFirewallObjectService")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONFirewallObjectService{
		Name:     d.Get("name").(string),
		Comment:  d.Get("comment").(string),
		Category: d.Get("category").(string),
		Protocol: d.Get("protocol").(string),
	}

	err := c.CreateUpdateFirewallObjectService(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Service: %s", err)
	}

	return readFTMFirewallObjectService(d, m)
}

func deleteFTMFirewallObjectService(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMFirewallObjectService")()

	name := d.Id()

	err := c.DeleteFirewallObjectService(name)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Service: %s", err)
	}

	d.SetId("")

	return nil
}
