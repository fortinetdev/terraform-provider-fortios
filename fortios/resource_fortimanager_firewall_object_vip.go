package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFortimanagerFirewallObjectVip() *schema.Resource {
	return &schema.Resource{
		Create: createFMGFirewallObjectVip,
		Read:   readFMGFirewallObjectVip,
		Update: updateFMGFirewallObjectVip,
		Delete: deleteFMGFirewallObjectVip,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "static-nat",
				ValidateFunc: validation.StringInSlice([]string{
					"static-nat", "dns-translation", "fqdn",
				}, false),
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
			"config_default": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "enable",
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
			},
			"mapped_addr": &schema.Schema{
				Type:     schema.TypeString,
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

func createFMGFirewallObjectVip(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGFirewallObjectVip")()

	cd := ""
	if d.Get("config_default").(string) == "disable" {
		cd = "1"
	}
	i := &fmgclient.JSONFirewallObjectVip{
		Name:          d.Get("name").(string),
		Comment:       d.Get("comment").(string),
		Type:          d.Get("type").(string),
		ArpReply:      d.Get("arp_reply").(string),
		MappedIp:      d.Get("mapped_ip").(string),
		ExtIp:         d.Get("ext_ip").(string),
		ExtIntf:       d.Get("ext_intf").(string),
		ConfigDefault: cd,
		MappedAddr:    d.Get("mapped_addr").(string),
	}

	err := c.CreateUpdateFirewallObjectVip(i, "add", d.Get("adom").(string))
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object Vip: %s", err)
	}

	d.SetId(i.Name)

	return readFMGFirewallObjectVip(d, m)
}

func readFMGFirewallObjectVip(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGFirewallObjectVip")()

	name := d.Id()
	o, err := c.ReadFirewallObjectVip(d.Get("adom").(string), name)
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
	d.Set("config_default", o.ConfigDefault)
	if o.Type == "fqdn" {
		d.Set("mapped_addr", o.MappedAddr)
	}

	return nil
}

func updateFMGFirewallObjectVip(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGFirewallObjectVip")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}
	if d.HasChange("type") {
		return fmt.Errorf("the type argument should not be modified once the virtual ip is created")
	}

	cd := ""
	if d.Get("config_default").(string) == "disable" {
		cd = "1"
	}
	i := &fmgclient.JSONFirewallObjectVip{
		Name:          d.Get("name").(string),
		Comment:       d.Get("comment").(string),
		Type:          d.Get("type").(string),
		ArpReply:      d.Get("arp_reply").(string),
		MappedIp:      d.Get("mapped_ip").(string),
		ExtIp:         d.Get("ext_ip").(string),
		ExtIntf:       d.Get("ext_intf").(string),
		ConfigDefault: cd,
		MappedAddr:    d.Get("mapped_addr").(string),
	}

	err := c.CreateUpdateFirewallObjectVip(i, "update", d.Get("adom").(string))
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Vip: %s", err)
	}

	return readFMGFirewallObjectVip(d, m)
}

func deleteFMGFirewallObjectVip(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGFirewallObjectVip")()

	name := d.Id()

	err := c.DeleteFirewallObjectVip(d.Get("adom").(string), name)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Vip: %s", err)
	}

	d.SetId("")

	return nil
}
