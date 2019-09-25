package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
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
		},
	}
}

func resourceFirewallSecurityPolicySeqCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	srcId := d.Get("policy_src_id").(int)
	dstId := d.Get("policy_dst_id").(int)
	alterPos := d.Get("alter_position").(string)

	if alterPos != "before" && alterPos != "after" {
		return fmt.Errorf("<alter_position> param should be only 'after' or 'before'")
	}

	err := c.CreateUpdateFirewallSecurityPolicySeq(srcId, dstId, alterPos)
	if err != nil {
		return fmt.Errorf("Error Altering Firewall Security Policy Sequence: %s", err)
	}

	d.SetId(strconv.Itoa(srcId))

	return nil
}

// Not suitable operation
func resourceFirewallSecurityPolicySeqRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	return c.ReadFirewallSecurityPolicySeq()
}

// Not suitable operation
func resourceFirewallSecurityPolicySeqDel(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	return c.DelFirewallSecurityPolicySeq()
}
