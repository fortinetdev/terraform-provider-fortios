package fortios

import (
	"fmt"
	"log"

	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallObjectServiceCategory() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallObjectServiceCategoryCreate,
		Read:   resourceFirewallObjectServiceCategoryRead,
		Update: resourceFirewallObjectServiceCategoryUpdate,
		Delete: resourceFirewallObjectServiceCategoryDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
		},
	}
}

func resourceFirewallObjectServiceCategoryCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	name := d.Get("name").(string)
	comment := d.Get("comment").(string)

	j1 := &forticlient.JSONFirewallObjectServiceCategoryItem{
		Name:    name,
		Comment: comment,
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectServiceCategory{
		JSONFirewallObjectServiceCategoryItem: j1,
	}

	//Call process by sdk
	o, err := c.CreateFirewallObjectServiceCategory(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Object Service Category: %s", err)
	}

	//Set index for d
	d.SetId(o.Mkey)

	return resourceFirewallObjectServiceCategoryRead(d, m)
}

func resourceFirewallObjectServiceCategoryUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	name := d.Get("name").(string)
	comment := d.Get("comment").(string)

	if d.HasChange("name") {
		return fmt.Errorf("the name key should not be modified")
	}

	j1 := &forticlient.JSONFirewallObjectServiceCategoryItem{
		Name:    name,
		Comment: comment,
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallObjectServiceCategory{
		JSONFirewallObjectServiceCategoryItem: j1,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallObjectServiceCategory(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Object Service Category: %s", err)
	}

	return resourceFirewallObjectServiceRead(d, m)
}

func resourceFirewallObjectServiceCategoryDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallObjectServiceCategory(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Object Service Category: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallObjectServiceCategoryRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallObjectServiceCategory(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Object Service Category: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("name", o.Name)
	d.Set("comment", o.Comment)

	return nil
}
