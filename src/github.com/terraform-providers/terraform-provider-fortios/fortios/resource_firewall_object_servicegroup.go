package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirewallObjectServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectServiceGroupCreate,
		Read:   resourceFirewallObjectServiceGroupRead,
		Update: resourceFirewallObjectServiceGroupUpdate,
		Delete: resourceFirewallObjectServiceGroupDelete,

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

func resourceFirewallObjectServiceGroupCreate(d *schema.ResourceData, m interface{}) error {
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
	i := &forticlient.JSONFirewallObjectServiceGroup{
		Name:    name,
		Member:  members,
		Comment: comment,
	}

	//Call process by sdk
	o, err := c.CreateFirewallObjectServiceGroup(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object ServiceGroup: %s", err)
	}

	// Set index for d
	// d.SetId(strconv.Itoa(int(o.Mkey)))
	d.SetId(o.Mkey)

	return nil
}

func resourceFirewallObjectServiceGroupUpdate(d *schema.ResourceData, m interface{}) error {
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
	i := &forticlient.JSONFirewallObjectServiceGroup{
		Name:    name,
		Member:  members,
		Comment: comment,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallObjectServiceGroup(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object ServiceGroup: %s", err)
	}

	//Set index for d
	//d.SetId(o.Mkey)

	return nil
}

func resourceFirewallObjectServiceGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallObjectServiceGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object ServiceGroup: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallObjectServiceGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallObjectServiceGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object ServiceGroup: %s", err)
	}

	//Refresh property
	d.Set("name", o.Name)
	member := forticlient.ExtractString(o.Member)
	d.Set("member", member)
	d.Set("comment", o.Comment)

	return nil
}
