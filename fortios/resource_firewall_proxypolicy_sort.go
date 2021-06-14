package fortios

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallProxypolicySort() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProxypolicySortCreateUpdate,
		Read:   resourceFirewallProxypolicySortRead,
		Update: resourceFirewallProxypolicySortCreateUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"sortby": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sortdirection": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallProxypolicySortCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	sortby := d.Get("sortby").(string)
	sortdirection := d.Get("sortdirection").(string)

	if sortby != "policyid" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	err := c.CreateUpdateFirewallProxypolicySort(sortby, sortdirection, vdomparam)
	if err != nil {
		return fmt.Errorf("Error sorting FirewallProxypolicy: %s", err)
	}

	d.SetId(sortby + sortdirection)

	return resourceFirewallProxypolicySortRead(d, m)
}

func resourceFirewallProxypolicySortRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	sortby := d.Get("sortby").(string)
	sortdirection := d.Get("sortdirection").(string)

	if sortby != "policyid" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	sorted, err := c.ReadFirewallProxypolicySort(sortby, sortdirection, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProxypolicy sort status: %s %s", err, mkey)
	}

	if sorted == false {
		d.Set("status", "unsorted")
	} else {
		d.Set("status", "")
	}

	return nil
}
