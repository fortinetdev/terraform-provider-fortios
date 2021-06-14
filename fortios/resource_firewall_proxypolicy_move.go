package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallProxypolicyMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProxypolicyMoveCreateUpdate,
		Read:   resourceFirewallProxypolicyMoveRead,
		Update: resourceFirewallProxypolicyMoveCreateUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"policyid_src": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"policyid_dst": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"move": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"state_policy_srcdst_pos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallProxypolicyMoveCreateUpdate(d *schema.ResourceData, m interface{}) error {
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

	srcIdPatch := d.Get("policyid_src").(int)
	dstIdPatch := d.Get("policyid_dst").(int)
	mv := d.Get("move").(string)

	srcId := strconv.Itoa(srcIdPatch)
	dstId := strconv.Itoa(dstIdPatch)

	if mv != "before" && mv != "after" {
		return fmt.Errorf("<move> param should be only 'after' or 'before'")
	}

	err := c.CreateUpdateFirewallProxypolicyMove(srcId, dstId, mv, vdomparam)
	if err != nil {
		return fmt.Errorf("Error Altering FirewallProxypolicy Moveuence: %s", err)
	}

	d.SetId(srcId)

	return resourceFirewallProxypolicyMoveRead(d, m)
}

func resourceFirewallProxypolicyMoveRead(d *schema.ResourceData, m interface{}) error {
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

	sid := d.Get("policyid_src").(int)
	did := d.Get("policyid_dst").(int)
	action := d.Get("move").(string)

	o, err := c.GetFirewallProxypolicyList(vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProxypolicy List: %s", err)
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		for _, z := range o {
			idn := z.Policyid
			if idn == strconv.Itoa(d.Get("policyid_src").(int)) {
				now_sid = i
			}

			if idn == strconv.Itoa(d.Get("policyid_dst").(int)) {
				now_did = i
			}

			i += 1
		}

		state_policy_srcdst_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_policy_srcdst_pos = "FirewallProxypolicy with policyid_src(" + strconv.Itoa(sid) + ")  and policyid_dst(" + strconv.Itoa(did) + ") were deleted"
			} else if now_sid == -1 {
				state_policy_srcdst_pos = "FirewallProxypolicy with policyid_src(" + strconv.Itoa(sid) + ") was deleted"
			} else if now_did == -1 {
				state_policy_srcdst_pos = "FirewallProxypolicy with policyid_dst(" + strconv.Itoa(did) + ") was deleted"
			}
		} else {
			bconsistent := true
			if action == "before" {
				if now_sid != now_did-1 {
					bconsistent = false
				}
			}

			if action == "after" {
				if now_sid != now_did+1 {
					bconsistent = false
				}
			}

			if bconsistent == false {
				relative_pos := now_sid - now_did

				if relative_pos > 0 {
					state_policy_srcdst_pos = "FirewallProxypolicy with policyid_src(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(relative_pos) + " behind FirewallProxypolicy with policyid_dst(" + strconv.Itoa(did) + ")"
				} else {
					state_policy_srcdst_pos = "FirewallProxypolicy with policyid_src(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(-relative_pos) + " ahead of FirewallProxypolicy with policyid_dst(" + strconv.Itoa(did) + ")"
				}
			}
		}

		d.Set("state_policy_srcdst_pos", state_policy_srcdst_pos)

	}

	return nil
}
