package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerFirewallSecurityPolicyPackage() *schema.Resource {
	return &schema.Resource{
		Create: createFTMFirewallSecurityPolicyPackage,
		Read:   readFTMFirewallSecurityPolicyPackage,
		Update: updateFTMFirewallSecurityPolicyPackage,
		Delete: deleteFTMFirewallSecurityPolicyPackage,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createFTMFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMFirewallSecurityPolicyPackage")()

	i := &fortimngclient.JSONFirewallSecurityPolicyPackage{
		Name:   d.Get("name").(string),
		Target: d.Get("target").(string),
	}

	err := c.CreateUpdateFirewallSecurityPolicyPackage(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating security policy package: %s", err)
	}

	d.SetId(i.Name)

	return readFTMFirewallSecurityPolicyPackage(d, m)
}

func readFTMFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMFirewallSecurityPolicyPackage")()

	name := d.Id()

	o, err := c.ReadFirewallSecurityPolicyPackage(name)
	if err != nil {
		return fmt.Errorf("Error reading security policy package: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("target", o.Target)

	return nil
}

func updateFTMFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMFirewallSecurtyPolicyPackage")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONFirewallSecurityPolicyPackage{
		Name:   d.Get("name").(string),
		Target: d.Get("target").(string),
	}

	err := c.CreateUpdateFirewallSecurityPolicyPackage(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating firewall security policy package: %s", err)
	}

	return readFTMFirewallSecurityPolicyPackage(d, m)
}

func deleteFTMFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMFirewallSecurityPolicyPackage")()

	name := d.Id()

	err := c.DeleteFirewallSecurityPolicyPackage(name)
	if err != nil {
		return fmt.Errorf("Error deleting firewall security policy package: %s", err)
	}

	d.SetId("")

	return nil
}
