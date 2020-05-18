package fortios

import (
	"fmt"
	"log"
	"strconv"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/fortinetdev/forti-sdk-go/fortimanager/util"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemAdom() *schema.Resource {
	return &schema.Resource{
		Create: createFMGSystemAdom,
		Read:   readFMGSystemAdom,
		Update: updateFMGSystemAdom,
		Delete: deleteFMGSystemAdom,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "FortiGate",
				ValidateFunc: util.ValidateStringIn("FortiGate", "FortiCarrier"),
			},
			"central_management_vpn": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"central_management_fortiap": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"central_management_sdwan": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"mode": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Normal",
				ValidateFunc: util.ValidateStringIn("Normal", "Backup"),
			},
			"perform_policy_check_before_every_install": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"action_when_conflicts_occur_during_policy_check": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Continue",
				ValidateFunc: util.ValidateStringIn("Continue", "Stop"),
			},
			"auto_push_policy_packages_when_device_back_online": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Disable",
				ValidateFunc: util.ValidateStringIn("Disable", "Enable"),
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func createFMGSystemAdom(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGSystemAdom")()

	flags := []string{}
	if !d.Get("central_management_vpn").(bool) {
		flags = append(flags, "no_vpn_console")
	}
	if !d.Get("central_management_fortiap").(bool) {
		flags = append(flags, "per_device_wtp")
	}
	if d.Get("central_management_sdwan").(bool) {
		flags = append(flags, "central_sdwan")
	}
	if d.Get("mode").(string) == "Backup" {
		flags = append(flags, "backup")
	}
	if d.Get("perform_policy_check_before_every_install").(bool) {
		flags = append(flags, "policy_check_on_install")
	}
	if d.Get("action_when_conflicts_occur_during_policy_check").(string) == "Stop" {
		flags = append(flags, "install_on_policy_check_fail")
	}
	if d.Get("auto_push_policy_packages_when_device_back_online").(string) == "Enable" {
		flags = append(flags, "auto_push_cfg")
	}

	type_s := ""
	if d.Get("type").(string) == "FortiCarrier" {
		type_s = "foc"
	}

	i := &fmgclient.JSONSystemAdom{
		Name:           d.Get("name").(string),
		RestrictedPrds: type_s,
		Status:         strconv.Itoa(d.Get("status").(int)),
		Flags:          flags,
	}

	err := c.CreateUpdateSystemAdom(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Adom: %s", err)
	}

	d.SetId(i.Name)

	return readFMGSystemAdom(d, m)
}

func readFMGSystemAdom(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemAdom")()

	name := d.Id()
	o, err := c.ReadSystemAdom(name)
	if err != nil {
		return fmt.Errorf("Error reading System Adom: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from status", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("status", o.Status)
	d.Set("type", o.RestrictedPrds)

	return nil
}

func updateFMGSystemAdom(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGSystemAdom")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	if d.HasChange("type") {
		return fmt.Errorf("type can't be modified here")
	}

	flags := []string{}
	if !d.Get("central_management_vpn").(bool) {
		flags = append(flags, "no_vpn_console")
	}
	if !d.Get("central_management_fortiap").(bool) {
		flags = append(flags, "per_device_wtp")
	}
	if d.Get("central_management_sdwan").(bool) {
		flags = append(flags, "central_sdwan")
	}
	if d.Get("mode").(string) == "Backup" {
		flags = append(flags, "backup")
	}
	if d.Get("perform_policy_check_before_every_install").(bool) {
		flags = append(flags, "policy_check_on_install")
	}
	if d.Get("action_when_conflicts_occur_during_policy_check").(string) == "Stop" {
		flags = append(flags, "install_on_policy_check_fail")
	}
	if d.Get("auto_push_policy_packages_when_device_back_online").(string) == "Enable" {
		flags = append(flags, "auto_push_cfg")
	}

	type_s := ""
	if d.Get("type").(string) == "FortiCarrier" {
		type_s = "foc"
	}

	i := &fmgclient.JSONSystemAdom{
		Name:           d.Get("name").(string),
		RestrictedPrds: type_s,
		Status:         strconv.Itoa(d.Get("status").(int)),
		Flags:          flags,
	}

	err := c.CreateUpdateSystemAdom(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating System Adom: %s", err)
	}

	return readFMGSystemAdom(d, m)
}

func deleteFMGSystemAdom(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGSystemAdom")()

	name := d.Id()

	err := c.DeleteSystemAdom(name)
	if err != nil {
		return fmt.Errorf("Error deleting System Adom: %s", err)
	}

	d.SetId("")

	return nil
}
