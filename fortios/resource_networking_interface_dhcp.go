package fortios

import (
	"fmt"
	"log"
	"strings"

	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNetworkingInterfaceDHCP() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkingInterfaceDHCPCreate,
		Read:   resourceNetworkingInterfaceDHCPRead,
		Update: resourceNetworkingInterfaceDHCPUpdate,
		Delete: resourceNetworkingInterfaceDHCPDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
			"default_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "enable",
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"lease_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  604800,
			},
			// default == Same as System DNS, local == Same as interface, specify == assign dns servers
			"dns_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"dns_server_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"dns_server_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"dns_server_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
		},
	}
}

func resourceNetworkingInterfaceDHCPCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	var ipRange []forticlient.JSONNetworkingInterfaceDHCPIPRange
	ipr := forticlient.JSONNetworkingInterfaceDHCPIPRange{
		StartIP: strings.Split(d.Get("ip_range").(string), "-")[0],
		EndIP:   strings.Split(d.Get("ip_range").(string), "-")[1],
	}
	ipRange = append(ipRange, ipr)

	i := &forticlient.JSONNetworkingInterfaceDHCP{
		Interface:      d.Get("interface").(string),
		Vdom:           d.Get("vdom").(string),
		DefaultGateway: d.Get("default_gw").(string),
		Netmask:        d.Get("netmask").(string),
		IPRange:        ipRange,
		DNSService:     d.Get("dns_service").(string),
		DNSServer1:     d.Get("dns_server_1").(string),
		DNSServer2:     d.Get("dns_server_2").(string),
		DNSServer3:     d.Get("dns_server_3").(string),
		LeaseTime:      d.Get("lease_time").(int),
	}

	//Call process by sdk
	o, err := c.CreateNetworkingInterfaceDHCP(i)
	if err != nil {
		return fmt.Errorf("error creating Networking Interface DHCP Server: %s", err)
	}

	//Set index for d
	log.Printf("FOS-id is %v\n", o.Mkey)
	d.SetId(o.Mkey)

	return resourceNetworkingInterfaceDHCPRead(d, m)
}

func resourceNetworkingInterfaceDHCPUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	var ipRange []forticlient.JSONNetworkingInterfaceDHCPIPRange
	ipr := forticlient.JSONNetworkingInterfaceDHCPIPRange{
		StartIP: strings.Split(d.Get("ip_range").(string), "-")[0],
		EndIP:   strings.Split(d.Get("ip_range").(string), "-")[1],
	}
	ipRange = append(ipRange, ipr)

	i := &forticlient.JSONNetworkingInterfaceDHCP{
		Interface:      d.Get("interface").(string),
		Vdom:           d.Get("vdom").(string),
		DefaultGateway: d.Get("default_gw").(string),
		Netmask:        d.Get("netmask").(string),
		IPRange:        ipRange,
		DNSService:     d.Get("dns_service").(string),
		DNSServer1:     d.Get("dns_server_1").(string),
		DNSServer2:     d.Get("dns_server_2").(string),
		DNSServer3:     d.Get("dns_server_3").(string),
		LeaseTime:      d.Get("lease_time").(int),
	}

	//Call process by sdk
	_, err := c.UpdateNetworkingInterfaceDHCP(i, mkey)
	if err != nil {
		return fmt.Errorf("error updating Networking Interface DHCP Server: %s", err)
	}

	return resourceNetworkingInterfaceDHCPRead(d, m)
}

func resourceNetworkingInterfaceDHCPDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteNetworkingInterfaceDHCP(mkey)
	if err != nil {
		return fmt.Errorf("error deleting Networking Interface DHCP Server: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceNetworkingInterfaceDHCPRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadNetworkingInterfaceDHCP(mkey)
	if err != nil {
		return fmt.Errorf("error reading Networking Interface DHCP Server: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property

	d.Set("interface", o.Interface)
	d.Set("vdom", o.Vdom)
	d.Set("default_gw", o.DefaultGateway)
	d.Set("netmask", o.Netmask)
	d.Set("ip_range", o.IPRange[0].StartIP+"-"+o.IPRange[0].EndIP)
	d.Set("dns_service", o.DNSService)
	d.Set("dns_server_1", o.DNSServer1)
	d.Set("dns_server_2", o.DNSServer2)
	d.Set("dns_server_3", o.DNSServer3)
	d.Set("status", o.Status)
	d.Set("lease_time", o.LeaseTime)

	return nil
}
