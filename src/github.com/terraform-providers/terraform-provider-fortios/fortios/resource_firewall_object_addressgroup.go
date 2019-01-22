package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
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
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

	// Set index for d
	// d.SetId(strconv.Itoa(int(o.Mkey)))
	d.SetId(o.Mkey)

	return nil
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

	//Set index for d
	//d.SetId(o.Mkey)

	return nil
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

	//Refresh property
	d.Set("name", o.Name)
	member := forticlient.ExtractString(o.Member)
	d.Set("member", member)
	d.Set("comment", o.Comment)

	return nil
}
