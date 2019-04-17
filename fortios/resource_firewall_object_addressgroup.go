package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirewallObjectAddressGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectAddressGroupCreate,
		Read:   resourceFirewallObjectAddressGroupRead,
		Update: resourceFirewallObjectAddressGroupUpdate,
		Delete: resourceFirewallObjectAddressGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
		},
	}
}

func resourceFirewallObjectAddressGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	member := d.Get("member").([]interface{})
	comment := d.Get("comment").(string)

	var members []forticlient.MultValue

	for _, v := range member {
		if v == nil {
			return fmt.Errorf("null value")
		}
		members = append(members,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectAddressGroup{
		Name:    name,
		Member:  members,
		Comment: comment,
	}

	//Call process by sdk
	o, err := c.CreateFirewallObjectAddressGroup(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object AddressGroup: %s", err)
	}

	//Set index for d
	d.SetId(o.Mkey)

	return resourceFirewallObjectAddressGroupRead(d, m)
}

func resourceFirewallObjectAddressGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	member := d.Get("member").([]interface{})
	comment := d.Get("comment").(string)

	var members []forticlient.MultValue

	for _, v := range member {
		if v == nil {
			return fmt.Errorf("null value")
		}
		members = append(members,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectAddressGroup{
		Name:    name,
		Member:  members,
		Comment: comment,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallObjectAddressGroup(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object AddressGroup: %s", err)
	}

	return resourceFirewallObjectAddressGroupRead(d, m)
}

func resourceFirewallObjectAddressGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallObjectAddressGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object AddressGroup: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallObjectAddressGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallObjectAddressGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object AddressGroup: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("name", o.Name)
	member := forticlient.ExtractString(o.Member)
	if err := d.Set("member", member); err != nil {
		log.Printf("[WARN] Error setting Firewall Object AddressGroup for (%s): %s", d.Id(), err)
	}
	d.Set("comment", o.Comment)

	return nil
}
