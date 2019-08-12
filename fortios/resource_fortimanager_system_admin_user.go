package fortios

import (
	"fmt"
	"log"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemAdminUser() *schema.Resource {
	return &schema.Resource{
		Create: createFTMSystemAdminUser,
		Read:   readFTMSystemAdminUser,
		Update: updateFTMSystemAdminUser,
		Delete: deleteFTMSystemAdminUser,

		Schema: map[string]*schema.Schema{
			"userid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "local",
			},
			"profileid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Restricted_User",
			},
			"rpc_permit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
		},
	}
}

func createFTMSystemAdminUser(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMSystemAdminUser")()

	//Get Params from d
	userId := d.Get("userid").(string)
	passwd := d.Get("password").(string)
	description := d.Get("description").(string)
	userType := d.Get("user_type").(string)
	profileId := d.Get("profileid").(string)
	rpcPermit := d.Get("rpc_permit").(string)
	trusthost1 := d.Get("trusthost1").(string)
	trusthost2 := d.Get("trusthost2").(string)

	//Build input data by sdk
	i := &fortimngclient.JSONSysAdminUser{
		UserId:      userId,
		Passwd:      passwd,
		Description: description,
		UserType:    userType,
		ProfileId:   profileId,
		RpcPermit:   rpcPermit,
		Trusthost1:  trusthost1,
		Trusthost2:  trusthost2,
	}

	err := c.CreateUpdateSystemAdminUser(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Admin User: %s", err)
	}

	d.SetId(userId)

	return readFTMSystemAdminUser(d, m)
}

func readFTMSystemAdminUser(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMSystemAdminUser")()

	userid := d.Id()
	o, err := c.ReadSystemAdminUser(userid)
	if err != nil {
		return fmt.Errorf("Error reading System Admin User: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("userid", o.UserId)
	d.Set("description", o.Description)
	d.Set("user_type", o.UserType)
	d.Set("profileid", o.ProfileId)
	d.Set("rpc_permit", o.RpcPermit)
	d.Set("trusthost1", o.Trusthost1)
	d.Set("trusthost2", o.Trusthost2)

	return nil
}

func updateFTMSystemAdminUser(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMSystemAdminUser")()

	//Get Params from d
	userId := d.Get("userid").(string)
	passwd := d.Get("password").(string)
	description := d.Get("description").(string)
	userType := d.Get("user_type").(string)
	profileId := d.Get("profileid").(string)
	rpcPermit := d.Get("rpc_permit").(string)
	trusthost1 := d.Get("trusthost1").(string)
	trusthost2 := d.Get("trusthost2").(string)

	if d.HasChange("userid") {
		return fmt.Errorf("the userid argument is the key and should not be modified here")
	}

	//Build input data by sdk
	i := &fortimngclient.JSONSysAdminUser{
		UserId:      userId,
		Passwd:      passwd,
		Description: description,
		UserType:    userType,
		ProfileId:   profileId,
		RpcPermit:   rpcPermit,
		Trusthost1:  trusthost1,
		Trusthost2:  trusthost2,
	}
	err := c.CreateUpdateSystemAdminUser(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating System Admin User: %s", err)
	}

	return readFTMSystemAdminUser(d, m)
}

func deleteFTMSystemAdminUser(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMSystemAdminUser")()

	userid := d.Id()

	err := c.DeleteSystemAdminUser(userid)
	if err != nil {
		return fmt.Errorf("Error deleting System Admin User: %s", err)
	}

	d.SetId("")

	return nil
}
