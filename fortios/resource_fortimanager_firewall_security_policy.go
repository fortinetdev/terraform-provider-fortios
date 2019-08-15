package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerFirewallSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: createFTMFirewallSecurityPolicy,
		Read:   readFTMFirewallSecurityPolicy,
		Update: updateFTMFirewallSecurityPolicy,
		Delete: deleteFTMFirewallSecurityPolicy,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"srcaddr": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"srcintf": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"dstaddr": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"dstintf": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"service": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"schedule_time": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"package_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createFTMFirewallSecurityPolicy(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMFirewallSecurityPolicy")()

	p := &fortimngclient.JSONFirewallSecurityPolicy{
		Name:     d.Get("name").(string),
		Action:   d.Get("action").(string),
		SrcAddr:  c.InterfaceArray2StrArray(d.Get("srcaddr").([]interface{})),
		SrcIntf:  c.InterfaceArray2StrArray(d.Get("srcintf").([]interface{})),
		DstAddr:  c.InterfaceArray2StrArray(d.Get("dstaddr").([]interface{})),
		DstIntf:  c.InterfaceArray2StrArray(d.Get("dstintf").([]interface{})),
		Service:  c.InterfaceArray2StrArray(d.Get("service").([]interface{})),
		Schedule: c.InterfaceArray2StrArray(d.Get("schedule_time").([]interface{})),
		PolicyId: "0",
	}

	i := &fortimngclient.FirewallSecurityPolicyInput{
		Policy:      p,
		PackageName: d.Get("package_name").(string),
	}

	policyid, err := c.CreateUpdateFirewallSecurityPolicy(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating devicemanager security policy: %s", err)
	}

	d.SetId(policyid)

	return readFTMFirewallSecurityPolicy(d, m)
}

func readFTMFirewallSecurityPolicy(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMFirewallSecurityPolicy")()

	policyid := d.Id()
	pkg_name := d.Get("package_name").(string)

	o, err := c.ReadFirewallSecurityPolicy(policyid, pkg_name)
	if err != nil {
		return fmt.Errorf("Error reading devicemanager security policy: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("srcaddr", o.SrcAddr)
	d.Set("srcintf", o.SrcIntf)
	d.Set("dstaddr", o.DstAddr)
	d.Set("dstintf", o.DstIntf)
	d.Set("service", o.Service)
	d.Set("action", o.Action)
	d.Set("schedule_time", o.Schedule)

	return nil
}

func updateFTMFirewallSecurityPolicy(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMFirewallSecurtyPolicy")()

	/*
	   // policy name can be modified on Fortimanager, so skip the check here
	   	if d.HasChange("name") {
	   		return fmt.Errorf("the name argument is the key and should not be modified here")
	   	}
	*/

	p := &fortimngclient.JSONFirewallSecurityPolicy{
		Name:     d.Get("name").(string),
		Action:   d.Get("action").(string),
		SrcAddr:  c.InterfaceArray2StrArray(d.Get("srcaddr").([]interface{})),
		SrcIntf:  c.InterfaceArray2StrArray(d.Get("srcintf").([]interface{})),
		DstAddr:  c.InterfaceArray2StrArray(d.Get("dstaddr").([]interface{})),
		DstIntf:  c.InterfaceArray2StrArray(d.Get("dstintf").([]interface{})),
		Service:  c.InterfaceArray2StrArray(d.Get("service").([]interface{})),
		Schedule: c.InterfaceArray2StrArray(d.Get("schedule_time").([]interface{})),
		PolicyId: d.Id(),
	}

	i := &fortimngclient.FirewallSecurityPolicyInput{
		Policy:      p,
		PackageName: d.Get("package_name").(string),
	}

	_, err := c.CreateUpdateFirewallSecurityPolicy(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating devicemanager security policy: %s", err)
	}

	return readFTMFirewallSecurityPolicy(d, m)
}

func deleteFTMFirewallSecurityPolicy(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMFirewallSecurityPolicy")()

	policyid := d.Id()
	pkg_name := d.Get("package_name").(string)

	err := c.DeleteFirewallSecurityPolicy(policyid, pkg_name)
	if err != nil {
		return fmt.Errorf("Error deleting devicemanager policy: %s", err)
	}

	d.SetId("")

	return nil
}
