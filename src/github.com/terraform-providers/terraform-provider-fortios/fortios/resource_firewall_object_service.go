package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirewallObjectService() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectServiceCreate,
		Read:   resourceFirewallObjectServiceRead,
		Update: resourceFirewallObjectServiceUpdate,
		Delete: resourceFirewallObjectServiceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
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
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iprange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

	j1 := &forticlient.JSONFirewallObjectServiceCommon{
		Name:     name,
		Category: category,
		Protocol: protocol,
		Comment:  comment,
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

	// Set index for d
	// d.SetId(strconv.Itoa(int(o.Mkey)))
	d.SetId(o.Mkey)

	return nil
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

	j1 := &forticlient.JSONFirewallObjectServiceCommon{
		Name:     name,
		Category: category,
		Protocol: protocol,
		Comment:  comment,
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

	//Set index for d
	//d.SetId(o.Mkey)

	return nil
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
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallObjectService(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object Service: %s", err)
	}

	//Refresh property
	d.Set("name", o.Name)
	d.Set("category", o.Category)
	d.Set("protocol", o.Protocol)
	d.Set("fqdn", o.Fqdn)
	d.Set("iprange", o.Iprange)
	d.Set("comment", o.Comment)

	return nil
}
