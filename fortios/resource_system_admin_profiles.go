package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSystemAdminProfiles() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminProfilesCreate,
		Read:   resourceSystemAdminProfilesRead,
		Update: resourceSystemAdminProfilesUpdate,
		Delete: resourceSystemAdminProfilesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"secfabgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftviewgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sysgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loggrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpngrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utmgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wanoptgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admintimeout_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAdminProfilesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	scope := d.Get("scope").(string)
	comments := d.Get("comments").(string)
	secfabgrp := d.Get("secfabgrp").(string)
	ftviewgrp := d.Get("ftviewgrp").(string)
	authgrp := d.Get("authgrp").(string)
	sysgrp := d.Get("sysgrp").(string)
	netgrp := d.Get("netgrp").(string)
	loggrp := d.Get("loggrp").(string)
	fwgrp := d.Get("fwgrp").(string)
	vpngrp := d.Get("vpngrp").(string)
	utmgrp := d.Get("utmgrp").(string)
	wanoptgrp := d.Get("wanoptgrp").(string)
	wifi := d.Get("wifi").(string)
	admintimeoutOverride := d.Get("admintimeout_override").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemAdminProfiles{
		Name:                 name,
		Scope:                scope,
		Comments:             comments,
		Secfabgrp:            secfabgrp,
		Ftviewgrp:            ftviewgrp,
		Authgrp:              authgrp,
		Sysgrp:               sysgrp,
		Netgrp:               netgrp,
		Loggrp:               loggrp,
		Fwgrp:                fwgrp,
		Vpngrp:               vpngrp,
		Utmgrp:               utmgrp,
		Wanoptgrp:            wanoptgrp,
		Wifi:                 wifi,
		AdmintimeoutOverride: admintimeoutOverride,
	}

	//Call process by sdk
	o, err := c.CreateSystemAdminProfiles(i)
	if err != nil {
		return fmt.Errorf("Error creating System Admin Profiles: %s", err)
	}

	//Set index for d
	d.SetId(o.Mkey)

	return resourceSystemAdminProfilesRead(d, m)
}

func resourceSystemAdminProfilesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	scope := d.Get("scope").(string)
	comments := d.Get("comments").(string)
	secfabgrp := d.Get("secfabgrp").(string)
	ftviewgrp := d.Get("ftviewgrp").(string)
	authgrp := d.Get("authgrp").(string)
	sysgrp := d.Get("sysgrp").(string)
	netgrp := d.Get("netgrp").(string)
	loggrp := d.Get("loggrp").(string)
	fwgrp := d.Get("fwgrp").(string)
	vpngrp := d.Get("vpngrp").(string)
	utmgrp := d.Get("utmgrp").(string)
	wanoptgrp := d.Get("wanoptgrp").(string)
	wifi := d.Get("wifi").(string)
	admintimeoutOverride := d.Get("admintimeout_override").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemAdminProfiles{
		Name:                 name,
		Scope:                scope,
		Comments:             comments,
		Secfabgrp:            secfabgrp,
		Ftviewgrp:            ftviewgrp,
		Authgrp:              authgrp,
		Sysgrp:               sysgrp,
		Netgrp:               netgrp,
		Loggrp:               loggrp,
		Fwgrp:                fwgrp,
		Vpngrp:               vpngrp,
		Utmgrp:               utmgrp,
		Wanoptgrp:            wanoptgrp,
		Wifi:                 wifi,
		AdmintimeoutOverride: admintimeoutOverride,
	}

	//Call process by sdk
	_, err := c.UpdateSystemAdminProfiles(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System Admin Profiles: %s", err)
	}

	return resourceSystemAdminProfilesRead(d, m)
}

func resourceSystemAdminProfilesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	err := c.DeleteSystemAdminProfiles(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting System Admin Profiles: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceSystemAdminProfilesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemAdminProfiles(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System Admin Profiles: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("name", o.Name)
	d.Set("scope", o.Scope)
	d.Set("comments", o.Comments)
	d.Set("secfabgrp", o.Secfabgrp)
	d.Set("ftviewgrp", o.Ftviewgrp)
	d.Set("authgrp", o.Authgrp)
	d.Set("sysgrp", o.Sysgrp)
	d.Set("netgrp", o.Netgrp)
	d.Set("loggrp", o.Loggrp)
	d.Set("fwgrp", o.Fwgrp)
	d.Set("vpngrp", o.Vpngrp)
	d.Set("utmgrp", o.Utmgrp)
	d.Set("wanoptgrp", o.Wanoptgrp)
	d.Set("wifi", o.Wifi)
	d.Set("admintimeout_override", o.AdmintimeoutOverride)

	return nil
}
