package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallSecurityPolicySeq() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSecurityPolicySeqCreateUpdate,
		Read:   resourceFirewallSecurityPolicySeqRead,
		Update: resourceFirewallSecurityPolicySeqCreateUpdate,
		Delete: resourceFirewallSecurityPolicySeqDel,

		Schema: map[string]*schema.Schema{
			"policy_src_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"policy_dst_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"alter_position": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enable_state_checking": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"state_policy_srcdst_pos": &schema.Schema{
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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallSecurityPolicySeqCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	srcIdPatch := d.Get("policy_src_id").(int)
	dstIdPatch := d.Get("policy_dst_id").(int)
	alterPos := d.Get("alter_position").(string)

	srcId := strconv.Itoa(srcIdPatch)
	dstId := strconv.Itoa(dstIdPatch)

	if alterPos != "before" && alterPos != "after" {
		return fmt.Errorf("<alter_position> param should be only 'after' or 'before'")
	}

	err := c.CreateUpdateFirewallSecurityPolicySeq(srcId, dstId, alterPos)
	if err != nil {
		return fmt.Errorf("Error Altering Firewall Security Policy Sequence: %s", err)
	}

	d.SetId(srcId)

	return resourceFirewallSecurityPolicySeqRead(d, m)
}

// Not suitable operation
func resourceFirewallSecurityPolicySeqDel(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	return c.DelFirewallSecurityPolicySeq()
}

func resourceFirewallSecurityPolicySeqRead(d *schema.ResourceData, m interface{}) error {
	enable_state_checking := d.Get("enable_state_checking").(bool)

	if enable_state_checking == false {
		d.Set("state_policy_list", nil)
		return nil
	}

	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	err := c.ReadFirewallSecurityPolicySeq()
	if err != nil {
		return fmt.Errorf("Error reading Firewall Security Policy List: %s %s", err, mkey)
	}

	sid := d.Get("policy_src_id").(int)
	did := d.Get("policy_dst_id").(int)
	action := d.Get("alter_position").(string)

	o, err := c.GetSecurityPolicyList()
	if err != nil {
		return fmt.Errorf("Error reading Firewall Security Policy List: %s", err)
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		var items []interface{}
		for _, z := range o {
			m := make(map[string]interface{})

			idn := z.PolicyID
			if idn == strconv.Itoa(d.Get("policy_src_id").(int)) {
				now_sid = i
				idn = "*" + idn
			}

			if idn == strconv.Itoa(d.Get("policy_dst_id").(int)) {
				now_did = i
				idn = "*" + idn
			}

			m["policyid"] = idn
			m["name"] = z.Name
			m["action"] = z.Action

			items = append(items, m)
			i += 1
		}

		if err := d.Set("state_policy_list", items); err != nil {
			log.Printf("[WARN] Error reading Firewall Security Policy List for (%s): %s", d.Id(), err)
		}

		state_policy_srcdst_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_policy_srcdst_pos = "policies with policy_src_id(" + strconv.Itoa(sid) + ")  and policy_dst_id(" + strconv.Itoa(did) + ") were deleted"
			} else if now_sid == -1 {
				state_policy_srcdst_pos = "policy with policy_src_id(" + strconv.Itoa(sid) + ") was deleted"
			} else if now_did == -1 {
				state_policy_srcdst_pos = "policy with policy_dst_id(" + strconv.Itoa(did) + ") was deleted"
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
					state_policy_srcdst_pos = "policy with policy_src_id(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(relative_pos) + " behind policy with policy_dst_id(" + strconv.Itoa(did) + ")"
				} else {
					state_policy_srcdst_pos = "policy with policy_src_id(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(-relative_pos) + " ahead of policy with policy_dst_id(" + strconv.Itoa(did) + ")"
				}
			}
		}

		d.Set("state_policy_srcdst_pos", state_policy_srcdst_pos)

	}

	return nil
}
