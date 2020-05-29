package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "flow",
				ValidateFunc: validation.StringInSlice([]string{
					"flow", "proxy",
				}, false),
			},
		},
	}
}

func createFMGFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGFirewallSecurityPolicyPackage")()

	i := &fmgclient.JSONFirewallSecurityPolicyPackage{
		Name:           d.Get("name").(string),
		Vdom:           d.Get("vdom").(string),
		Target:         d.Get("target").(string),
		InspectionMode: d.Get("inspection_mode").(string),
	}

	err := c.CreateUpdateFirewallSecurityPolicyPackage(i, "add", d.Get("adom").(string))
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

	o, err := c.ReadFirewallSecurityPolicyPackage(d.Get("adom").(string), name)
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
	if o.InspectionMode != "" {
		d.Set("inspection_mode", o.InspectionMode)
	}

	return nil
}

func updateFMGFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGFirewallSecurtyPolicyPackage")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fmgclient.JSONFirewallSecurityPolicyPackage{
		Name:           d.Get("name").(string),
		Vdom:           d.Get("vdom").(string),
		Target:         d.Get("target").(string),
		InspectionMode: d.Get("inspection_mode").(string),
	}

	err := c.CreateUpdateFirewallSecurityPolicyPackage(i, "update", d.Get("adom").(string))
	if err != nil {
		return fmt.Errorf("Error updating firewall security policy package: %s", err)
	}

	return readFMGFirewallSecurityPolicyPackage(d, m)
}

func deleteFMGFirewallSecurityPolicyPackage(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGFirewallSecurityPolicyPackage")()

	name := d.Id()

	err := c.DeleteFirewallSecurityPolicyPackage(d.Get("adom").(string), name)
	if err != nil {
		return fmt.Errorf("Error deleting firewall security policy package: %s", err)
	}

	d.SetId("")

	return nil
}
