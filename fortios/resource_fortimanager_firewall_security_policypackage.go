package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerFirewallSecurityPolicyPackage() *schema.Resource {
	return &schema.Resource{
		Create: createFMGFirewallSecurityPolicyPackage,
		Read:   readFMGFirewallSecurityPolicyPackage,
		Update: updateFMGFirewallSecurityPolicyPackage,
		Delete: deleteFMGFirewallSecurityPolicyPackage,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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

func createFMGFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGFirewallSecurityPolicyPackage")()

	i := &fmgclient.JSONFirewallSecurityPolicyPackage{
		Name:   d.Get("name").(string),
		Target: d.Get("target").(string),
	}

	err := c.CreateUpdateFirewallSecurityPolicyPackage(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating security policy package: %s", err)
	}

	d.SetId(i.Name)

	return readFMGFirewallSecurityPolicyPackage(d, m)
}

func readFMGFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGFirewallSecurityPolicyPackage")()

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

func updateFMGFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGFirewallSecurtyPolicyPackage")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fmgclient.JSONFirewallSecurityPolicyPackage{
		Name:   d.Get("name").(string),
		Target: d.Get("target").(string),
	}

	err := c.CreateUpdateFirewallSecurityPolicyPackage(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating firewall security policy package: %s", err)
	}

	return readFMGFirewallSecurityPolicyPackage(d, m)
}

func deleteFMGFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGFirewallSecurityPolicyPackage")()

	name := d.Id()

	err := c.DeleteFirewallSecurityPolicyPackage(name)
	if err != nil {
		return fmt.Errorf("Error deleting firewall security policy package: %s", err)
	}

	d.SetId("")

	return nil
}
