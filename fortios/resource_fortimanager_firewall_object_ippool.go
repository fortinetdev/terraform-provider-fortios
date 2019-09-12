package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/fgtdev/fortimanager-sdk-go/util"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerFirewallObjectIppool() *schema.Resource {
	return &schema.Resource{
		Create: createFMGFirewallObjectIppool,
		Read:   readFMGFirewallObjectIppool,
		Update: updateFMGFirewallObjectIppool,
		Delete: deleteFMGFirewallObjectIppool,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"startip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"endip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "overload",
				ValidateFunc: util.ValidateStringIn("overload", "one-to-one"),
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"arp_reply": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "enable",
				ValidateFunc: util.ValidateStringIn("disable", "enable"),
			},
			"arp_intf": &schema.Schema{
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

func createFMGFirewallObjectIppool(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGFirewallObjectIppool")()

	i := &fortimngclient.JSONFirewallObjectIppool{
		Name:           d.Get("name").(string),
		Comment:        d.Get("comment").(string),
		Type:           d.Get("type").(string),
		ArpIntf:        d.Get("arp_intf").(string),
		ArpReply:       d.Get("arp_reply").(string),
		AssociatedIntf: d.Get("associated_intf").(string),
		StartIp:        d.Get("startip").(string),
		EndIp:          d.Get("endip").(string),
	}

	err := c.CreateUpdateFirewallObjectIppool(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object Ippool: %s", err)
	}

	d.SetId(i.Name)

	return readFMGFirewallObjectIppool(d, m)
}

func readFMGFirewallObjectIppool(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGFirewallObjectIppool")()

	name := d.Id()
	o, err := c.ReadFirewallObjectIppool(name)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object Ippool: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("comment", o.Comment)
	d.Set("type", o.Type)
	d.Set("arp-intf", o.ArpIntf)
	d.Set("arp-reply", o.ArpReply)
	d.Set("associated-intf", o.AssociatedIntf)
	d.Set("startip", o.StartIp)
	d.Set("endip", o.EndIp)

	return nil
}

func updateFMGFirewallObjectIppool(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGFirewallObjectIppool")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONFirewallObjectIppool{
		Name:           d.Get("name").(string),
		Comment:        d.Get("comment").(string),
		Type:           d.Get("type").(string),
		ArpIntf:        d.Get("arp_intf").(string),
		ArpReply:       d.Get("arp_reply").(string),
		AssociatedIntf: d.Get("associated_intf").(string),
		StartIp:        d.Get("startip").(string),
		EndIp:          d.Get("endip").(string),
	}

	err := c.CreateUpdateFirewallObjectIppool(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Ippool: %s", err)
	}

	return readFMGFirewallObjectIppool(d, m)
}

func deleteFMGFirewallObjectIppool(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGFirewallObjectIppool")()

	name := d.Id()

	err := c.DeleteFirewallObjectIppool(name)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Ippool: %s", err)
	}

	d.SetId("")

	return nil
}
