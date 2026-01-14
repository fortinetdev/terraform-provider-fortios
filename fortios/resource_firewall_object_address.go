package fortios

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	forticlient "github.com/terraform-providers/terraform-provider-fortios/sdk/sdkcore"
)

func resourceFirewallObjectAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectAddressCreate,
		Read:   resourceFirewallObjectAddressRead,
		Update: resourceFirewallObjectAddressUpdate,
		Delete: resourceFirewallObjectAddressDelete,

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
				Required: true,
			},
			"subnet": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ToDiagFunc(isValidSubnet),
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"associated_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"show_in_address_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"static_route_configure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallObjectAddressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	typef := d.Get("type").(string)
	subnet := d.Get("subnet").(string)
	startIP := d.Get("start_ip").(string)
	endIP := d.Get("end_ip").(string)
	fqdn := d.Get("fqdn").(string)
	country := d.Get("country").(string)
	comment := d.Get("comment").(string)
	associatedInterface := d.Get("associated_interface").(string)
	showInAddressList := d.Get("show_in_address_list").(string)
	staticRouteConfigure := d.Get("static_route_configure").(string)

	j1 := &forticlient.JSONFirewallObjectAddressCommon{
		Name:              name,
		Type:              typef,
		Comment:           comment,
		AssociatedIntf:    associatedInterface,
		ShowInAddressList: showInAddressList,
		AllowRouting:      staticRouteConfigure,
	}
	var j2 *forticlient.JSONFirewallObjectAddressIPRange
	var j3 *forticlient.JSONFirewallObjectAddressCountry
	var j4 *forticlient.JSONFirewallObjectAddressFqdn
	var j5 *forticlient.JSONFirewallObjectAddressIPMask

	switch typef {
	case "iprange":
		j2 = &forticlient.JSONFirewallObjectAddressIPRange{
			StartIP: startIP,
			EndIP:   endIP,
		}

	case "geography":
		j3 = &forticlient.JSONFirewallObjectAddressCountry{
			Country: country,
		}

	case "fqdn":
		j4 = &forticlient.JSONFirewallObjectAddressFqdn{
			Fqdn: fqdn,
		}

	case "ipmask":
		j5 = &forticlient.JSONFirewallObjectAddressIPMask{
			Subnet: subnet,
		}

	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectAddress{
		JSONFirewallObjectAddressCommon:  j1,
		JSONFirewallObjectAddressIPRange: j2,
		JSONFirewallObjectAddressCountry: j3,
		JSONFirewallObjectAddressFqdn:    j4,
		JSONFirewallObjectAddressIPMask:  j5,
	}

	//Call process by sdk
	o, err := c.CreateFirewallObjectAddress(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object Address: %s", err)
	}

	//Set index for d
	d.SetId(o.Mkey)

	return resourceFirewallObjectAddressRead(d, m)
}

func resourceFirewallObjectAddressUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	typef := d.Get("type").(string)
	subnet := d.Get("subnet").(string)
	startIP := d.Get("start_ip").(string)
	endIP := d.Get("end_ip").(string)
	fqdn := d.Get("fqdn").(string)
	country := d.Get("country").(string)
	comment := d.Get("comment").(string)
	associatedInterface := d.Get("associated_interface").(string)
	showInAddressList := d.Get("show_in_address_list").(string)
	staticRouteConfigure := d.Get("static_route_configure").(string)

	j1 := &forticlient.JSONFirewallObjectAddressCommon{
		Name:              name,
		Type:              typef,
		Comment:           comment,
		AssociatedIntf:    associatedInterface,
		ShowInAddressList: showInAddressList,
		AllowRouting:      staticRouteConfigure,
	}
	var j2 *forticlient.JSONFirewallObjectAddressIPRange
	var j3 *forticlient.JSONFirewallObjectAddressCountry
	var j4 *forticlient.JSONFirewallObjectAddressFqdn
	var j5 *forticlient.JSONFirewallObjectAddressIPMask

	switch typef {
	case "iprange":
		j2 = &forticlient.JSONFirewallObjectAddressIPRange{
			StartIP: startIP,
			EndIP:   endIP,
		}

	case "geography":
		j3 = &forticlient.JSONFirewallObjectAddressCountry{
			Country: country,
		}

	case "fqdn":
		j4 = &forticlient.JSONFirewallObjectAddressFqdn{
			Fqdn: fqdn,
		}

	case "ipmask":
		j5 = &forticlient.JSONFirewallObjectAddressIPMask{
			Subnet: subnet,
		}

	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectAddress{
		JSONFirewallObjectAddressCommon:  j1,
		JSONFirewallObjectAddressIPRange: j2,
		JSONFirewallObjectAddressCountry: j3,
		JSONFirewallObjectAddressFqdn:    j4,
		JSONFirewallObjectAddressIPMask:  j5,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallObjectAddress(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Address: %s", err)
	}

	return resourceFirewallObjectAddressRead(d, m)
}

func resourceFirewallObjectAddressDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallObjectAddress(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Address: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallObjectAddressRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Get("name").(string)

	if mkey == "" {
		mkey = d.Id() //for import
	}

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallObjectAddress(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object Address: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.SetId(o.Name)
	d.Set("name", o.Name)
	d.Set("type", o.Type)
	if o.Type == "ipmask" {
		d.Set("subnet", validateConvIPMask2CIDR(d.Get("subnet").(string), o.Subnet))
	}

	if o.Type == "iprange" {
		d.Set("start_ip", o.StartIP)
		d.Set("end_ip", o.EndIP)
	}

	d.Set("fqdn", o.Fqdn)
	d.Set("country", o.Country)
	d.Set("comment", o.Comment)
	d.Set("associated_interface", o.AssociatedIntf)
	d.Set("show_in_address_list", o.ShowInAddressList)
	if o.Type != "iprange" && o.Type != "geography" {
		d.Set("static_route_configure", o.AllowRouting)
	}

	return nil
}
