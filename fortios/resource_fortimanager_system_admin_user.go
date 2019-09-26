package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/fgtdev/fortimanager-sdk-go/util"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemAdminUser() *schema.Resource {
	return &schema.Resource{
		Create: createFMGSystemAdminUser,
		Read:   readFMGSystemAdminUser,
		Update: updateFMGSystemAdminUser,
		Delete: deleteFMGSystemAdminUser,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "local",
				ValidateFunc: util.ValidateStringIn("local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"),
			},
			"profileid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Restricted_User",
			},
			"rpc_permit": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "none",
				ValidateFunc: util.ValidateStringIn("none", "read", "read-write"),
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0 0.0.0.0",
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "255.255.255.255 255.255.255.255",
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "255.255.255.255 255.255.255.255",
			},
		},
	}
}

func createFMGSystemAdminUser(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGSystemAdminUser")()

	//Build input data by sdk
	i := &fmgclient.JSONSysAdminUser{
		UserId:      d.Get("userid").(string),
		Passwd:      d.Get("password").(string),
		Description: d.Get("description").(string),
		UserType:    d.Get("user_type").(string),
		ProfileId:   d.Get("profileid").(string),
		RpcPermit:   d.Get("rpc_permit").(string),
		Trusthost1:  d.Get("trusthost1").(string),
		Trusthost2:  d.Get("trusthost2").(string),
		Trusthost3:  d.Get("trusthost3").(string),
	}

	err := c.CreateUpdateSystemAdminUser(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Admin User: %s", err)
	}

	d.SetId(i.UserId)

	return readFMGSystemAdminUser(d, m)
}

func readFMGSystemAdminUser(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemAdminUser")()

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
	d.Set("trusthost3", o.Trusthost3)

	return nil
}

func updateFMGSystemAdminUser(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGSystemAdminUser")()

	if d.HasChange("userid") {
		return fmt.Errorf("the userid argument is the key and should not be modified here")
	}

	//Build input data by sdk
	i := &fmgclient.JSONSysAdminUser{
		UserId:      d.Get("userid").(string),
		Passwd:      d.Get("password").(string),
		Description: d.Get("description").(string),
		UserType:    d.Get("user_type").(string),
		ProfileId:   d.Get("profileid").(string),
		RpcPermit:   d.Get("rpc_permit").(string),
		Trusthost1:  d.Get("trusthost1").(string),
		Trusthost2:  d.Get("trusthost2").(string),
		Trusthost3:  d.Get("trusthost3").(string),
	}

	err := c.CreateUpdateSystemAdminUser(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating System Admin User: %s", err)
	}

	return readFMGSystemAdminUser(d, m)
}

func deleteFMGSystemAdminUser(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGSystemAdminUser")()

	userid := d.Id()

	err := c.DeleteSystemAdminUser(userid)
	if err != nil {
		return fmt.Errorf("Error deleting System Admin User: %s", err)
	}

	d.SetId("")

	return nil
}
