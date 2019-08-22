package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerFirewallObjectVip() *schema.Resource {
	return &schema.Resource{
		Create: createFTMFirewallObjectVip,
		Read:   readFTMFirewallObjectVip,
		Update: updateFTMFirewallObjectVip,
		Delete: deleteFTMFirewallObjectVip,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "static-nat",
			},
			"ext_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ext_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
			},
			"mapped_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createFTMFirewallObjectVip(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMFirewallObjectVip")()

	i := &fortimngclient.JSONFirewallObjectVip{
		Name:     d.Get("name").(string),
		Comment:  d.Get("comment").(string),
		Type:     d.Get("type").(string),
		ArpReply: d.Get("arp_reply").(string),
		MappedIp: d.Get("mapped_ip").(string),
		ExtIp:    d.Get("ext_ip").(string),
		ExtIntf:  d.Get("ext_intf").(string),
	}

	err := c.CreateUpdateFirewallObjectVip(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object Vip: %s", err)
	}

	d.SetId(i.Name)

	return readFTMFirewallObjectVip(d, m)
}

func readFTMFirewallObjectVip(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMFirewallObjectVip")()

	name := d.Id()
	o, err := c.ReadFirewallObjectVip(name)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object Vip: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("comment", o.Comment)
	d.Set("type", o.Type)
	d.Set("arp-reply", o.ArpReply)
	d.Set("ext_ip", o.ExtIp)
	d.Set("ext_intf", o.ExtIntf)
	d.Set("mapped_ip", o.MappedIp)

	return nil
}

func updateFTMFirewallObjectVip(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMFirewallObjectVip")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONFirewallObjectVip{
		Name:     d.Get("name").(string),
		Comment:  d.Get("comment").(string),
		Type:     d.Get("type").(string),
		ArpReply: d.Get("arp_reply").(string),
		MappedIp: d.Get("mapped_ip").(string),
		ExtIp:    d.Get("ext_ip").(string),
		ExtIntf:  d.Get("ext_intf").(string),
	}

	err := c.CreateUpdateFirewallObjectVip(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Vip: %s", err)
	}

	return readFTMFirewallObjectVip(d, m)
}

func deleteFTMFirewallObjectVip(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMFirewallObjectVip")()

	name := d.Id()

	err := c.DeleteFirewallObjectVip(name)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Vip: %s", err)
	}

	d.SetId("")

	return nil
}
