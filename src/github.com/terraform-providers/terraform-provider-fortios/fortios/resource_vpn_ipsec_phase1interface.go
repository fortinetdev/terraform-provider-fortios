package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVPNIPsecPhase1Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPNIPsecPhase1InterfaceCreate,
		Read:   resourceVPNIPsecPhase1InterfaceRead,
		Update: resourceVPNIPsecPhase1InterfaceUpdate,
		Delete: resourceVPNIPsecPhase1InterfaceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"peertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proposal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wizard_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"psksecret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"peerid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"peer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"peergrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_split_include": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"split_include_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_split_exclude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVPNIPsecPhase1InterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
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

	var certificates []forticlient.MultValue

	for _, v := range certificate {
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
	}

	//Call process by sdk
	o, err := c.CreateVPNIPsecPhase1Interface(i)
	if err != nil {
		return fmt.Errorf("Error creating VPN IPsec Phase1Interface: %s", err)
	}

	// Set index for d
	// d.SetId(strconv.Itoa(int(o.Mkey)))
	d.SetId(o.Mkey)

	return nil
}

func resourceVPNIPsecPhase1InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
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

	var certificates []forticlient.MultValue

	for _, v := range certificate {
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
	}

	//Call process by sdk
	_, err := c.UpdateVPNIPsecPhase1Interface(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VPN IPsec Phase1Interface: %s", err)
	}

	//Set index for d
	//d.SetId(o.Mkey)

	return nil
}

func resourceVPNIPsecPhase1InterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
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
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadVPNIPsecPhase1Interface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VPN IPsec Phase1Interface: %s", err)
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
	d.Set("psksecret", o.Psksecret)
	certificate := forticlient.ExpandStringList(o.Certificate)
	d.Set("certificate", certificate)
	d.Set("peerid", o.Peerid)
	d.Set("peer", o.Peer)
	d.Set("peergrp", o.Peergrp)
	d.Set("ipv4_split_include", o.IPv4SplitInclude)
	d.Set("split_include_service", o.SplitIncludeService)
	d.Set("ipv4_split_exclude", o.IPv4SplitExclude)

	return nil
}
