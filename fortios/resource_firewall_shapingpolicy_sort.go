package fortios

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallShapingpolicySort() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShapingpolicySortCreateUpdate,
		Read:   resourceFirewallShapingpolicySortRead,
		Update: resourceFirewallShapingpolicySortCreateUpdate,
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
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					availableOptions := []string{"id", "name"}
					var validValue bool
					for _, ele := range availableOptions {
						if ele == v {
							validValue = true
							break
						}
					}
					if !validValue {
						errs = append(errs, fmt.Errorf("%q must be one of the option of [\"id\", \"name\"], got: \"%v\"", key, v))
					}
					return
				},
			},
			"sortdirection": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					availableOptions := []string{"ascending", "descending", "manual"}
					var validValue bool
					for _, ele := range availableOptions {
						if ele == v {
							validValue = true
							break
						}
					}
					if !validValue {
						errs = append(errs, fmt.Errorf("%q must be one of the option of [\"ascending\", \"descending\", \"manual\"], got: \"%v\"", key, v))
					}
					return
				},
			},
			"manual_order": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state_policy_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceFirewallShapingpolicySortCreateUpdate(d *schema.ResourceData, m interface{}) error {
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
	manual_order_d := d.Get("manual_order").([]interface{})
	manual_order := make([]string, len(manual_order_d))
	for cIndex, cValue := range manual_order_d {
		manual_order[cIndex] = fmt.Sprint(cValue)
	}

	if sortby != "id" && sortby != "name" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" && sortdirection != "manual" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	err := c.CreateUpdateFirewallShapingpolicySort(sortby, sortdirection, vdomparam, manual_order)
	if err != nil {
		return fmt.Errorf("Error sorting FirewallShapingpolicy: %s", err)
	}

	d.SetId(sortby + sortdirection)

	return resourceFirewallShapingpolicySortRead(d, m)
}

func resourceFirewallShapingpolicySortRead(d *schema.ResourceData, m interface{}) error {
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
	manual_order_d := d.Get("manual_order").([]interface{})
	manual_order := make([]string, len(manual_order_d))
	for cIndex, cValue := range manual_order_d {
		manual_order[cIndex] = fmt.Sprint(cValue)
	}

	if sortby != "id" && sortby != "name" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" && sortdirection != "manual" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	sorted, o, err := c.ReadFirewallShapingpolicySort(sortby, sortdirection, vdomparam, manual_order)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShapingpolicy sort status: %s %s", err, mkey)
	}

	if sorted == false {
		d.Set("status", "unsorted")
	} else {
		d.Set("status", "")
	}

	d.Set("force_recreate", "False")

	if o != nil {
		if err := d.Set("state_policy_list", o); err != nil {
			log.Printf("[WARN] Error reading Firewall Security Policy List for (%s): %s", d.Id(), err)
		}
	} else {
		d.Set("state_policy_list", nil)
	}
	d.Set("manual_order", manual_order)

	return nil
}
