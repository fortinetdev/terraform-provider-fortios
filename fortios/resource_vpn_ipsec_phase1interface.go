package fortios

import (
	"fmt"
	"log"

	"github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVPNIPsecPhase1Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPNIPsecPhase1InterfaceCreate,
		Read:   resourceVPNIPsecPhase1InterfaceRead,
		Update: resourceVPNIPsecPhase1InterfaceUpdate,
		Delete: resourceVPNIPsecPhase1InterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"psksecret": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"peertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proposal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"wizard_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"peerid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peergrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_split_include": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_include_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_split_exclude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode_cfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authmethod_remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVPNIPsecPhase1InterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	typef := d.Get("type").(string)
	interfacef := d.Get("interface").(string)
	peertype := d.Get("peertype").(string)
	proposal := d.Get("proposal").(string)
	comments := d.Get("comments").(string)
	wizardType := d.Get("wizard_type").(string)
	remoteGw := d.Get("remote_gw").(string)
	psksecret := d.Get("psksecret").(string)
	certificate := d.Get("certificate").([]interface{})
	peerid := d.Get("peerid").(string)
	peer := d.Get("peer").(string)
	peergrp := d.Get("peergrp").(string)
	ipv4SplitInclude := d.Get("ipv4_split_include").(string)
	splitIncludeService := d.Get("split_include_service").(string)
	ipv4SplitExclude := d.Get("ipv4_split_exclude").(string)
	modeCfg := d.Get("mode_cfg").(string)
	authmethod := d.Get("authmethod").(string)
	authmethodRemote := d.Get("authmethod_remote").(string)

	var certificates []forticlient.MultValue

	for _, v := range certificate {
		if v == nil {
			return fmt.Errorf("null value")
		}
		certificates = append(certificates,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONVPNIPsecPhase1Interface{
		Name:                name,
		Type:                typef,
		Interface:           interfacef,
		Peertype:            peertype,
		Proposal:            proposal,
		Comments:            comments,
		WizardType:          wizardType,
		RemoteGw:            remoteGw,
		Psksecret:           psksecret,
		Certificate:         certificates,
		Peerid:              peerid,
		Peer:                peer,
		Peergrp:             peergrp,
		IPv4SplitInclude:    ipv4SplitInclude,
		SplitIncludeService: splitIncludeService,
		IPv4SplitExclude:    ipv4SplitExclude,
		ModeCfg:             modeCfg,
		Authmethod:          authmethod,
		AuthmethodRemote:    authmethodRemote,
	}

	//Call process by sdk
	o, err := c.CreateVPNIPsecPhase1Interface(i)
	if err != nil {
		return fmt.Errorf("Error creating VPN IPsec Phase1Interface: %s", err)
	}

	// Set index for d
	d.SetId(o.Mkey)

	return resourceVPNIPsecPhase1InterfaceRead(d, m)
}

func resourceVPNIPsecPhase1InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	typef := d.Get("type").(string)
	interfacef := d.Get("interface").(string)
	peertype := d.Get("peertype").(string)
	proposal := d.Get("proposal").(string)
	comments := d.Get("comments").(string)
	wizardType := d.Get("wizard_type").(string)
	remoteGw := d.Get("remote_gw").(string)
	psksecret := d.Get("psksecret").(string)
	certificate := d.Get("certificate").([]interface{})
	peerid := d.Get("peerid").(string)
	peer := d.Get("peer").(string)
	peergrp := d.Get("peergrp").(string)
	ipv4SplitInclude := d.Get("ipv4_split_include").(string)
	splitIncludeService := d.Get("split_include_service").(string)
	ipv4SplitExclude := d.Get("ipv4_split_exclude").(string)
	modeCfg := d.Get("mode_cfg").(string)
	authmethod := d.Get("authmethod").(string)
	authmethodRemote := d.Get("authmethod_remote").(string)

	var certificates []forticlient.MultValue

	for _, v := range certificate {
		if v == nil {
			return fmt.Errorf("null value")
		}
		certificates = append(certificates,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONVPNIPsecPhase1Interface{
		Name:                name,
		Type:                typef,
		Interface:           interfacef,
		Peertype:            peertype,
		Proposal:            proposal,
		Comments:            comments,
		WizardType:          wizardType,
		RemoteGw:            remoteGw,
		Psksecret:           psksecret,
		Certificate:         certificates,
		Peerid:              peerid,
		Peer:                peer,
		Peergrp:             peergrp,
		IPv4SplitInclude:    ipv4SplitInclude,
		SplitIncludeService: splitIncludeService,
		IPv4SplitExclude:    ipv4SplitExclude,
		ModeCfg:             modeCfg,
		Authmethod:          authmethod,
		AuthmethodRemote:    authmethodRemote,
	}

	//Call process by sdk
	_, err := c.UpdateVPNIPsecPhase1Interface(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VPN IPsec Phase1Interface: %s", err)
	}

	return resourceVPNIPsecPhase1InterfaceRead(d, m)
}

func resourceVPNIPsecPhase1InterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	err := c.DeleteVPNIPsecPhase1Interface(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VPN IPsec Phase1Interface: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceVPNIPsecPhase1InterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadVPNIPsecPhase1Interface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VPN IPsec Phase1Interface: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("name", o.Name)
	d.Set("type", o.Type)
	d.Set("interface", o.Interface)
	d.Set("peertype", o.Peertype)
	d.Set("proposal", o.Proposal)
	d.Set("comments", o.Comments)
	d.Set("wizard_type", o.WizardType)
	d.Set("remote_gw", o.RemoteGw)
	//d.Set("psksecret", o.Psksecret)
	certificate := forticlient.ExtractString(o.Certificate)
	if err := d.Set("certificate", certificate); err != nil {
		log.Printf("[WARN] Error setting VPN IPsec Phase1Interface for (%s): %s", d.Id(), err)
	}
	d.Set("peerid", o.Peerid)
	d.Set("peer", o.Peer)
	d.Set("peergrp", o.Peergrp)
	d.Set("ipv4_split_include", o.IPv4SplitInclude)
	d.Set("split_include_service", o.SplitIncludeService)
	d.Set("ipv4_split_exclude", o.IPv4SplitExclude)
	d.Set("mode_cfg", o.ModeCfg)
	d.Set("authmethod", o.Authmethod)
	d.Set("authmethod_remote", o.AuthmethodRemote)

	return nil
}
