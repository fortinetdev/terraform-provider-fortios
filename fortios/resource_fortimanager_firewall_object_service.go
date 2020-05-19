package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/fortinetdev/forti-sdk-go/fortimanager/util"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFortimanagerFirewallObjectService() *schema.Resource {
	return &schema.Resource{
		Create: createFTMFirewallObjectService,
		Read:   readFTMFirewallObjectService,
		Update: updateFTMFirewallObjectService,
		Delete: deleteFTMFirewallObjectService,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
				Default:  "",
				ValidateFunc: validation.StringInSlice([]string{
					"", "File Access", "Authentication", "Email", "General", "Network Services", "Remote Access", "Tunneling", "VoIP, Messaging & Other Applications", "Web Access", "Web Proxy",
				}, false),
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "TCP/UDP/SCTP",
				ValidateFunc: validation.StringInSlice([]string{
					"TCP/UDP/SCTP", "ICMP", "ICMP6", "IP",
				}, false),
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iprange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_portrange": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"udp_portrange": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"sctp_portrange": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"icmp_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"icmp_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol_number": &schema.Schema{
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

func createFTMFirewallObjectService(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMFirewallObjectService")()

	i := &fmgclient.JSONFirewallObjectService{
		Name:          d.Get("name").(string),
		Comment:       d.Get("comment").(string),
		Category:      d.Get("category").(string),
		Protocol:      d.Get("protocol").(string),
		Proxy:         d.Get("proxy").(string),
		Fqdn:          d.Get("fqdn").(string),
		Iprange:       d.Get("iprange").(string),
		TcpPortRange:  util.InterfaceArray2StrArray(d.Get("tcp_portrange").([]interface{})),
		UdpPortRange:  util.InterfaceArray2StrArray(d.Get("udp_portrange").([]interface{})),
		SctpPortRange: util.InterfaceArray2StrArray(d.Get("sctp_portrange").([]interface{})),
		IcmpCode:      d.Get("icmp_code").(int),
		IcmpType:      d.Get("icmp_type").(int),
		ProtocolNum:   d.Get("protocol_number").(int),
	}

	err := c.CreateUpdateFirewallObjectService(i, "add", d.Get("adom").(string))
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
	o, err := c.ReadFirewallObjectService(d.Get("adom").(string), name)
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
	d.Set("proxy", o.Proxy)
	if o.Protocol == "TCP/UDP/SCTP" {
		d.Set("fqdn", o.Fqdn)
		d.Set("iprange", o.Iprange)
		d.Set("tcp_portrange", o.TcpPortRange)
		d.Set("udp_portrange", o.UdpPortRange)
		d.Set("sctp_portrange", o.SctpPortRange)
	} else if o.Protocol == "ICMP" || o.Protocol == "ICMP6" {
		d.Set("icmp_type", o.IcmpType)
		d.Set("icmp_code", o.IcmpCode)
	} else if o.Protocol == "IP" {
		d.Set("protocol_number", o.ProtocolNum)
	}

	return nil
}

func updateFTMFirewallObjectService(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMFirewallObjectService")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}
	if d.HasChange("proxy") {
		return fmt.Errorf("proxy can't be modified once the service is created")
	}

	i := &fmgclient.JSONFirewallObjectService{
		Name:          d.Get("name").(string),
		Comment:       d.Get("comment").(string),
		Category:      d.Get("category").(string),
		Protocol:      d.Get("protocol").(string),
		Proxy:         d.Get("proxy").(string),
		Fqdn:          d.Get("fqdn").(string),
		Iprange:       d.Get("iprange").(string),
		TcpPortRange:  util.InterfaceArray2StrArray(d.Get("tcp_portrange").([]interface{})),
		UdpPortRange:  util.InterfaceArray2StrArray(d.Get("udp_portrange").([]interface{})),
		SctpPortRange: util.InterfaceArray2StrArray(d.Get("sctp_portrange").([]interface{})),
		IcmpCode:      d.Get("icmp_code").(int),
		IcmpType:      d.Get("icmp_type").(int),
		ProtocolNum:   d.Get("protocol_number").(int),
	}

	err := c.CreateUpdateFirewallObjectService(i, "update", d.Get("adom").(string))
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Service: %s", err)
	}

	return readFTMFirewallObjectService(d, m)
}

func deleteFTMFirewallObjectService(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMFirewallObjectService")()

	name := d.Id()

	err := c.DeleteFirewallObjectService(d.Get("adom").(string), name)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Service: %s", err)
	}

	d.SetId("")

	return nil
}
