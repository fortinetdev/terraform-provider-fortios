package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirewallObjectServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectServiceGroupCreate,
		Read:   resourceFirewallObjectServiceGroupRead,
		Update: resourceFirewallObjectServiceGroupUpdate,
		Delete: resourceFirewallObjectServiceGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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

func resourceFirewallObjectServiceGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

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

	//Set index for d
	d.SetId(o.Mkey)

	return resourceFirewallObjectServiceGroupRead(d, m)
}

func resourceFirewallObjectServiceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

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

	return resourceFirewallObjectServiceGroupRead(d, m)
}

func resourceFirewallObjectServiceGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

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
	mkey := d.Get("name").(string)

	if mkey == "" {
		mkey = d.Id()
	}

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallObjectServiceGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object ServiceGroup: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.SetId(o.Name)
	d.Set("name", o.Name)
	member := forticlient.ExtractString(o.Member)
	if err := d.Set("member", member); err != nil {
		log.Printf("[WARN] Error setting Firewall Object ServiceGroup for (%s): %s", d.Id(), err)
	}
	d.Set("comment", o.Comment)

	return nil
}
