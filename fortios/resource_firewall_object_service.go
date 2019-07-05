package fortios

import (
	"fmt"
	"log"

	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirewallObjectService() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectServiceCreate,
		Read:   resourceFirewallObjectServiceRead,
		Update: resourceFirewallObjectServiceUpdate,
		Delete: resourceFirewallObjectServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"iprange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"protocol_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"icmptype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"icmpcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"tcp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"udp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"sctp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceFirewallObjectServiceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	category := d.Get("category").(string)
	protocol := d.Get("protocol").(string)
	fqdn := d.Get("fqdn").(string)
	iprange := d.Get("iprange").(string)
	comment := d.Get("comment").(string)
	protocolNumber := d.Get("protocol_number").(string)
	icmptype := d.Get("icmptype").(string)
	icmpcode := d.Get("icmpcode").(string)
	tcpPortrange := d.Get("tcp_portrange").(string)
	udpPortrange := d.Get("udp_portrange").(string)
	sctpPortrange := d.Get("sctp_portrange").(string)

	if protocol == "ICMP" {
		protocolNumber = "1"
	}

	j1 := &forticlient.JSONFirewallObjectServiceCommon{
		Name:           name,
		Category:       category,
		Protocol:       protocol,
		Comment:        comment,
		ProtocolNumber: protocolNumber,
		Icmptype:       icmptype,
		Icmpcode:       icmpcode,
		TCPPortrange:   tcpPortrange,
		UDPPortrange:   udpPortrange,
		SctpPortrange:  sctpPortrange,
	}
	var j2 *forticlient.JSONFirewallObjectServiceFqdn
	var j3 *forticlient.JSONFirewallObjectServiceIprange

	if fqdn != "" {
		j2 = &forticlient.JSONFirewallObjectServiceFqdn{
			Fqdn: fqdn,
		}
	} else {
		j3 = &forticlient.JSONFirewallObjectServiceIprange{
			Iprange: iprange,
		}
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectService{
		JSONFirewallObjectServiceCommon:  j1,
		JSONFirewallObjectServiceFqdn:    j2,
		JSONFirewallObjectServiceIprange: j3,
	}

	//Call process by sdk
	o, err := c.CreateFirewallObjectService(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object Service: %s", err)
	}

	//Set index for d
	d.SetId(o.Mkey)

	return resourceFirewallObjectServiceRead(d, m)
}

func resourceFirewallObjectServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	category := d.Get("category").(string)
	protocol := d.Get("protocol").(string)
	fqdn := d.Get("fqdn").(string)
	iprange := d.Get("iprange").(string)
	comment := d.Get("comment").(string)
	protocolNumber := d.Get("protocol_number").(string)
	icmptype := d.Get("icmptype").(string)
	icmpcode := d.Get("icmpcode").(string)
	tcpPortrange := d.Get("tcp_portrange").(string)
	udpPortrange := d.Get("udp_portrange").(string)
	sctpPortrange := d.Get("sctp_portrange").(string)

	j1 := &forticlient.JSONFirewallObjectServiceCommon{
		Name:           name,
		Category:       category,
		Protocol:       protocol,
		Comment:        comment,
		ProtocolNumber: protocolNumber,
		Icmptype:       icmptype,
		Icmpcode:       icmpcode,
		TCPPortrange:   tcpPortrange,
		UDPPortrange:   udpPortrange,
		SctpPortrange:  sctpPortrange,
	}
	var j2 *forticlient.JSONFirewallObjectServiceFqdn
	var j3 *forticlient.JSONFirewallObjectServiceIprange

	if fqdn != "" {
		j2 = &forticlient.JSONFirewallObjectServiceFqdn{
			Fqdn: fqdn,
		}
	} else {
		j3 = &forticlient.JSONFirewallObjectServiceIprange{
			Iprange: iprange,
		}
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectService{
		JSONFirewallObjectServiceCommon:  j1,
		JSONFirewallObjectServiceFqdn:    j2,
		JSONFirewallObjectServiceIprange: j3,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallObjectService(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Service: %s", err)
	}

	return resourceFirewallObjectServiceRead(d, m)
}

func resourceFirewallObjectServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallObjectService(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Service: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallObjectServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Get("name").(string)

	if mkey == "" {
		mkey = d.Id()
	}

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallObjectService(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object Service: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.SetId(o.Name)
	d.Set("name", o.Name)
	d.Set("category", o.Category)

	// if o.TCPPortrange == "" && o.UDPPortrange == "" && o.SctpPortrange == "" {
	// 	d.Set("protocol_number", o.ProtocolNumber)
	// }

	d.Set("protocol", o.Protocol)
	d.Set("fqdn", o.Fqdn)
	d.Set("iprange", o.Iprange)
	d.Set("comment", o.Comment)
	d.Set("icmptype", o.Icmptype)
	d.Set("icmpcode", o.Icmpcode)
	d.Set("tcp_portrange", o.TCPPortrange)
	d.Set("udp_portrange", o.UDPPortrange)
	d.Set("sctp_portrange", o.SctpPortrange)

	return nil
}
