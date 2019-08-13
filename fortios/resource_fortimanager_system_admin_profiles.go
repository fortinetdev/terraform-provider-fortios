package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemAdminProfiles() *schema.Resource {
	return &schema.Resource{
		Create: createFMGSystemAdminProfiles,
		Read:   readFMGSystemAdminProfiles,
		Update: updateFMGSystemAdminProfiles,
		Delete: deleteFMGSystemAdminProfiles,

		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Try to give description here...",
			},
			"device_manager": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
			},
		},
	}
}

func createFMGSystemAdminProfiles(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGSystemAdminProfiles")()

	//Get Params from d
	profileId := d.Get("profile_id").(string)
	description := d.Get("description").(string)
	deviceManager := d.Get("device_manager").(string)

	//Build input data by sdk
	i := &fortimngclient.JSONSysAdminProfiles{
		ProfileId:     profileId,
		Description:   description,
		DeviceManager: deviceManager,
	}

	err := c.CreateUpdateSystemAdminProfiles(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Admin Profiles: %s", err)
	}

	d.SetId(profileId)

	return readFMGSystemAdminProfiles(d, m)
}

func readFMGSystemAdminProfiles(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemAdminProfiles")()

	profileId := d.Id()
	o, err := c.ReadSystemAdminProfiles(profileId)
	if err != nil {
		return fmt.Errorf("Error reading System Admin Profiles: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("profile_id", o.ProfileId)
	d.Set("description", o.Description)
	d.Set("device_manager", o.DeviceManager)

	return nil
}

func updateFMGSystemAdminProfiles(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGSystemAdminProfiles")()

	//Get Params from d
	profileId := d.Get("profile_id").(string)
	description := d.Get("description").(string)
	deviceManager := d.Get("device_manager").(string)

	if d.HasChange("profile_id") {
		return fmt.Errorf("the profile_id argument is the key and should not be modified here")
	}

	//Build input data by sdk
	i := &fortimngclient.JSONSysAdminProfiles{
		ProfileId:     profileId,
		Description:   description,
		DeviceManager: deviceManager,
	}

	err := c.CreateUpdateSystemAdminProfiles(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating System Admin Profiles: %s", err)
	}

	return readFMGSystemAdminProfiles(d, m)
}

func deleteFMGSystemAdminProfiles(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGSystemAdminProfiles")()

	profileId := d.Id()

	err := c.DeleteSystemAdminProfiles(profileId)
	if err != nil {
		return fmt.Errorf("Error deleting System Admin Profiles: %s", err)
	}

	d.SetId("")

	return nil
}
