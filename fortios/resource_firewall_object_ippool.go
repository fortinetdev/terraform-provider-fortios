package fortios

import (
	"fmt"
	"log"

	"github.com/fortinetdev/forti-sdk-go/fortios/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallObjectIPPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectIPPoolCreate,
		Read:   resourceFirewallObjectIPPoolRead,
		Update: resourceFirewallObjectIPPoolUpdate,
		Delete: resourceFirewallObjectIPPoolDelete,

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
			"startip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"endip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
		},
	}
}

func resourceFirewallObjectIPPoolCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	typef := d.Get("type").(string)
	startip := d.Get("startip").(string)
	endip := d.Get("endip").(string)
	arpReply := d.Get("arp_reply").(string)
	comments := d.Get("comments").(string)

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectIPPool{
		Name:     name,
		Type:     typef,
		Startip:  startip,
		Endip:    endip,
		ArpReply: arpReply,
		Comments: comments,
	}

	//Call process by sdk
	o, err := c.CreateFirewallObjectIPPool(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object IPPool: %s", err)
	}

	//Set index for d
	d.SetId(o.Mkey)

	return resourceFirewallObjectIPPoolRead(d, m)
}

func resourceFirewallObjectIPPoolUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	typef := d.Get("type").(string)
	startip := d.Get("startip").(string)
	endip := d.Get("endip").(string)
	arpReply := d.Get("arp_reply").(string)
	comments := d.Get("comments").(string)

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectIPPool{
		Name:     name,
		Type:     typef,
		Startip:  startip,
		Endip:    endip,
		ArpReply: arpReply,
		Comments: comments,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallObjectIPPool(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object IPPool: %s", err)
	}

	return resourceFirewallObjectIPPoolRead(d, m)
}

func resourceFirewallObjectIPPoolDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallObjectIPPool(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object IPPool: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallObjectIPPoolRead(d *schema.ResourceData, m interface{}) error {
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
	o, err := c.ReadFirewallObjectIPPool(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object IPPool: %s", err)
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
	d.Set("startip", o.Startip)
	d.Set("endip", o.Endip)
	d.Set("arp_reply", o.ArpReply)
	d.Set("comments", o.Comments)

	return nil
}
