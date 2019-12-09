package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallObjectVipGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectVipGroupCreate,
		Read:   resourceFirewallObjectVipGroupRead,
		Update: resourceFirewallObjectVipGroupUpdate,
		Delete: resourceFirewallObjectVipGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "any",
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceFirewallObjectVipGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	comments := d.Get("comments").(string)
	interfacef := d.Get("interface").(string)
	member := d.Get("member").([]interface{})

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
	i := &forticlient.JSONFirewallObjectVipGroup{
		Name:      name,
		Comments:  comments,
		Interface: interfacef,
		Member:    members,
	}

	//Call process by sdk
	o, err := c.CreateFirewallObjectVipGroup(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object VipGroup: %s", err)
	}

	//Set index for d
	d.SetId(o.Mkey)

	return resourceFirewallObjectVipGroupRead(d, m)
}

func resourceFirewallObjectVipGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	comments := d.Get("comments").(string)
	interfacef := d.Get("interface").(string)
	member := d.Get("member").([]interface{})

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
	i := &forticlient.JSONFirewallObjectVipGroup{
		Name:      name,
		Comments:  comments,
		Interface: interfacef,
		Member:    members,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallObjectVipGroup(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object VipGroup: %s", err)
	}

	return resourceFirewallObjectVipGroupRead(d, m)
}

func resourceFirewallObjectVipGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallObjectVipGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object VipGroup: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallObjectVipGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Get("name").(string)

	if mkey == "" {
		mkey = d.Id()
	}

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallObjectVipGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object VipGroup: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.SetId(o.Name)
	d.Set("name", o.Name)
	d.Set("comments", o.Comments)
	d.Set("interface", o.Interface)
	member := forticlient.ExtractString(o.Member)
	if err := d.Set("member", member); err != nil {
		log.Printf("[WARN] Error setting Firewall Object VipGroup for (%s): %s", d.Id(), err)
	}

	return nil
}
