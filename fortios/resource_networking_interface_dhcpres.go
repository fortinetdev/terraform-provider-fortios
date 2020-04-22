package fortios

import (
	"fmt"
	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceNetworkingInterfaceDHCPIpRes() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkingInterfaceDHCPIpResCreate,
		Read:   resourceNetworkingInterfaceDHCPIpResRead,
		Update: resourceNetworkingInterfaceDHCPIpResUpdate,
		Delete: resourceNetworkingInterfaceDHCPiPResDelete,

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
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "enable",
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Object Managed by Terraform",
			},
		},
	}
}

func resourceNetworkingInterfaceDHCPIpResCreate(d *schema.ResourceData, m interface{}) error {
	logPrefix := "resourceNetworkingInterfaceDHCPIpResCreate"
	c := m.(*FortiClient).Client
	c.Retries = 1

	skey, err := c.NetworkingInterfaceDHCPSrvByName(d.Get("interface").(string))
	if err != nil {
		return fmt.Errorf("error retrieving DHCP Server from interface: %s", err)
	}

	i := &forticlient.JSONNetworkingInterfaceDHCPIpRes{
		IP:          d.Get("ip").(string),
		Mac:         d.Get("mac").(string),
		Description: d.Get("description").(string),
		Vdom:        d.Get("vdom").(string),
	}

	//Call process by sdk
	o, err := c.CreateNetworkInterfaceDHCPIpPRes(i, skey)
	if err != nil {
		return fmt.Errorf("error creating DHCP IP Reservation: %s", err)
	}

	//Set index for d
	log.Printf(logPrefix+"FOS-id is %v\n", o.Mkey)
	d.SetId(o.Mkey)

	return resourceNetworkingInterfaceDHCPIpResRead(d, m)
}

func resourceNetworkingInterfaceDHCPIpResUpdate(d *schema.ResourceData, m interface{}) error {
	logPrefix := "resourceNetworkingInterfaceDHCPIpResUpdate"
	c := m.(*FortiClient).Client
	c.Retries = 1

	skey, err := c.NetworkingInterfaceDHCPSrvByName(d.Get("interface").(string))
	if err != nil {
		return fmt.Errorf("error retrieving DHCP Server from interface: %s", err)
	}
	mkey := d.Id()

	i := &forticlient.JSONNetworkingInterfaceDHCPIpRes{
		IP:          d.Get("ip").(string),
		Mac:         d.Get("mac").(string),
		Description: d.Get("description").(string),
		Vdom:        d.Get("vdom").(string),
	}

	//Call process by sdk
	o, err := c.UpdateNetworkInterfaceDHCPIpPRes(i, skey, mkey)
	if err != nil {
		return fmt.Errorf("error creating DHCP IP Reservation: %s", err)
	}

	//Set index for d
	log.Printf(logPrefix+"FOS-id is %v\n", o.Mkey)
	d.SetId(o.Mkey)

	return resourceNetworkingInterfaceDHCPIpResRead(d, m)
}

func resourceNetworkingInterfaceDHCPiPResDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	skey, err := c.NetworkingInterfaceDHCPSrvByName(d.Get("interface").(string))
	if err != nil {
		return fmt.Errorf("error retrieving DHCP Server from interface: %s", err)
	}
	mkey := d.Id()

	//Call process by sdk
	_, err = c.DeleteNetworkingInterfaceDHCPIpRes(skey, mkey)
	if err != nil {
		return fmt.Errorf("error deleting DHCPm IP Reservation: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceNetworkingInterfaceDHCPIpResRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	skey, err := c.NetworkingInterfaceDHCPSrvByName(d.Get("interface").(string))
	if err != nil {
		return fmt.Errorf("error retrieving DHCP Server from interface: %s", err)
	}
	mkey := d.Id()

	//Call process by sdk
	o, err := c.ReadNetworkingInterfaceDHCPIpRes(skey, mkey)
	if err != nil {
		return fmt.Errorf("error reading DHCP IP Reservation: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("ip", o.IP)
	d.Set("mac", o.Mac)
	d.Set("description", o.Description)
	d.Set("vdom", o.Vdom)

	return nil
}
