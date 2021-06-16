package fortios

import (
	"fmt"
	"log"
	// "strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallSecurityPolicySort() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSecurityPolicySortCreateUpdate,
		Read:   resourceFirewallSecurityPolicySortRead,
		Update: resourceFirewallSecurityPolicySortCreateUpdate,
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
			"state_policy_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policyid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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

func resourceFirewallSecurityPolicySortCreateUpdate(d *schema.ResourceData, m interface{}) error {
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

	if sortby != "policyid" && sortby != "name" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	err := c.CreateUpdateFirewallSecurityPolicySort(sortby, sortdirection, vdomparam)
	if err != nil {
		return fmt.Errorf("Error Sort Firewall Security Policies: %s", err)
	}

	d.SetId(sortby + sortdirection)

	return resourceFirewallSecurityPolicySortRead(d, m)
}

func resourceFirewallSecurityPolicySortRead(d *schema.ResourceData, m interface{}) error {
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

	if sortby != "policyid" && sortby != "name" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	sorted, err := c.ReadFirewallSecurityPolicySort(sortby, sortdirection, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Security Policy Sort Status: %s %s", err, mkey)
	}

	if sorted == false {
		d.Set("status", "unsorted")
	} else {
		d.Set("status", "")
	}

	o, err := c.GetSecurityPolicyList(vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Security Policy List: %s", err)
	}

	if o != nil {
		i := 1

		var items []interface{}
		for _, z := range o {
			m := make(map[string]interface{})

			m["policyid"] = z.PolicyID
			m["name"] = z.Name
			m["action"] = z.Action

			items = append(items, m)
			i += 1
		}

		if err := d.Set("state_policy_list", items); err != nil {
			log.Printf("[WARN] Error reading Firewall Security Policy List for (%s): %s", d.Id(), err)
		}
	} else {
		d.Set("state_policy_list", nil)
	}

	return nil
}
