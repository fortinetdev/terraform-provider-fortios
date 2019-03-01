package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirewallObjectVip() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectVipCreate,
		Read:   resourceFirewallObjectVipRead,
		Update: resourceFirewallObjectVipUpdate,
		Delete: resourceFirewallObjectVipDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"extip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mappedip": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"extintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "any",
			},
			"portforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "tcp",
			},
			"extport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
		},
	}
}

func resourceFirewallObjectVipCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	extip := d.Get("extip").(string)
	mappedip := d.Get("mappedip").([]interface{})
	extintf := d.Get("extintf").(string)
	portforward := d.Get("portforward").(string)
	protocol := d.Get("protocol").(string)
	extport := d.Get("extport").(string)
	mappedport := d.Get("mappedport").(string)

	var mappedips []forticlient.VIPMultValue

	for _, v := range mappedip {
		mappedips = append(mappedips,
			forticlient.VIPMultValue{
				Range: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectVip{
		Name:        name,
		Comment:     comment,
		Extip:       extip,
		Mappedip:    mappedips,
		Extintf:     extintf,
		Portforward: portforward,
		Protocol:    protocol,
		Extport:     extport,
		Mappedport:  mappedport,
	}

	//Call process by sdk
	o, err := c.CreateFirewallObjectVip(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object Vip: %s", err)
	}

	// Set index for d
	// d.SetId(strconv.Itoa(int(o.Mkey)))
	d.SetId(o.Mkey)

	return nil
}

func resourceFirewallObjectVipUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d

	name := d.Get("name").(string)
	comment := d.Get("comment").(string)
	extip := d.Get("extip").(string)
	mappedip := d.Get("mappedip").([]interface{})
	extintf := d.Get("extintf").(string)
	portforward := d.Get("portforward").(string)
	protocol := d.Get("protocol").(string)
	extport := d.Get("extport").(string)
	mappedport := d.Get("mappedport").(string)

	var mappedips []forticlient.VIPMultValue

	for _, v := range mappedip {
		mappedips = append(mappedips,
			forticlient.VIPMultValue{
				Range: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectVip{
		Name:        name,
		Comment:     comment,
		Extip:       extip,
		Mappedip:    mappedips,
		Extintf:     extintf,
		Portforward: portforward,
		Protocol:    protocol,
		Extport:     extport,
		Mappedport:  mappedport,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallObjectVip(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Vip: %s", err)
	}

	//Set index for d
	//d.SetId(o.Mkey)

	return nil
}

func resourceFirewallObjectVipDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallObjectVip(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Vip: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallObjectVipRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallObjectVip(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object Vip: %s", err)
	}

	//Refresh property
	d.Set("name", o.Name)
	d.Set("comment", o.Comment)
	d.Set("extip", o.Extip)

	vs := make([]string, 0, len(o.Mappedip))
	for _, v := range o.Mappedip {
		c := v.Range
		vs = append(vs, c)
	}

	d.Set("mappedip", vs)
	d.Set("extintf", o.Extintf)
	d.Set("portforward", o.Portforward)
	d.Set("protocol", o.Protocol)
	d.Set("extport", o.Extport)
	d.Set("mappedport", o.Mappedport)

	return nil
}
