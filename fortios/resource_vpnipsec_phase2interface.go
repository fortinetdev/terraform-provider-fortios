// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure VPN autokey tunnel.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnIpsecPhase2Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecPhase2InterfaceCreate,
		Read:   resourceVpnIpsecPhase2InterfaceRead,
		Update: resourceVpnIpsecPhase2InterfaceUpdate,
		Delete: resourceVpnIpsecPhase2InterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"phase1name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"dhcp_ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proposal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_df": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_sender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_forwarder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keylifeseconds": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(120, 172800),
				Optional:     true,
				Computed:     true,
			},
			"keylifekbs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"keylife_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"single_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encapsulation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2tp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"initiator_ts_narrow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"src_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"src_name6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"src_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_start_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_end_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_subnet6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"dst_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"dst_name6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"dst_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_start_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_end_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_subnet6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceVpnIpsecPhase2InterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnIpsecPhase2Interface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase2Interface resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecPhase2Interface(obj)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase2Interface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase2Interface")
	}

	return resourceVpnIpsecPhase2InterfaceRead(d, m)
}

func resourceVpnIpsecPhase2InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnIpsecPhase2Interface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase2Interface resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecPhase2Interface(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase2Interface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase2Interface")
	}

	return resourceVpnIpsecPhase2InterfaceRead(d, m)
}

func resourceVpnIpsecPhase2InterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteVpnIpsecPhase2Interface(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecPhase2Interface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase2InterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadVpnIpsecPhase2Interface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase2Interface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecPhase2Interface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase2Interface resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecPhase2InterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfacePhase1Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDhcpIpsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceProposal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfacePfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceIpv4Df(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDhgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceReplay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceKeepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceAutoNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceAddRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceAutoDiscoverySender(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceAutoDiscoveryForwarder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceKeylifeseconds(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceKeylifekbs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceKeylifeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSingleSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceRouteOverlap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceEncapsulation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceL2Tp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceInitiatorTsNarrow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDiffserv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDiffservcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcName6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcStartIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcEndIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenVpnIpsecPhase2InterfaceSrcSubnet6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceSrcPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstName6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstStartIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstEndIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenVpnIpsecPhase2InterfaceDstSubnet6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InterfaceDstPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnIpsecPhase2Interface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnIpsecPhase2InterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("phase1name", flattenVpnIpsecPhase2InterfacePhase1Name(o["phase1name"], d, "phase1name", sv)); err != nil {
		if !fortiAPIPatch(o["phase1name"]) {
			return fmt.Errorf("Error reading phase1name: %v", err)
		}
	}

	if err = d.Set("dhcp_ipsec", flattenVpnIpsecPhase2InterfaceDhcpIpsec(o["dhcp-ipsec"], d, "dhcp_ipsec", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-ipsec"]) {
			return fmt.Errorf("Error reading dhcp_ipsec: %v", err)
		}
	}

	if err = d.Set("proposal", flattenVpnIpsecPhase2InterfaceProposal(o["proposal"], d, "proposal", sv)); err != nil {
		if !fortiAPIPatch(o["proposal"]) {
			return fmt.Errorf("Error reading proposal: %v", err)
		}
	}

	if err = d.Set("pfs", flattenVpnIpsecPhase2InterfacePfs(o["pfs"], d, "pfs", sv)); err != nil {
		if !fortiAPIPatch(o["pfs"]) {
			return fmt.Errorf("Error reading pfs: %v", err)
		}
	}

	if err = d.Set("ipv4_df", flattenVpnIpsecPhase2InterfaceIpv4Df(o["ipv4-df"], d, "ipv4_df", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-df"]) {
			return fmt.Errorf("Error reading ipv4_df: %v", err)
		}
	}

	if err = d.Set("dhgrp", flattenVpnIpsecPhase2InterfaceDhgrp(o["dhgrp"], d, "dhgrp", sv)); err != nil {
		if !fortiAPIPatch(o["dhgrp"]) {
			return fmt.Errorf("Error reading dhgrp: %v", err)
		}
	}

	if err = d.Set("replay", flattenVpnIpsecPhase2InterfaceReplay(o["replay"], d, "replay", sv)); err != nil {
		if !fortiAPIPatch(o["replay"]) {
			return fmt.Errorf("Error reading replay: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenVpnIpsecPhase2InterfaceKeepalive(o["keepalive"], d, "keepalive", sv)); err != nil {
		if !fortiAPIPatch(o["keepalive"]) {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	if err = d.Set("auto_negotiate", flattenVpnIpsecPhase2InterfaceAutoNegotiate(o["auto-negotiate"], d, "auto_negotiate", sv)); err != nil {
		if !fortiAPIPatch(o["auto-negotiate"]) {
			return fmt.Errorf("Error reading auto_negotiate: %v", err)
		}
	}

	if err = d.Set("add_route", flattenVpnIpsecPhase2InterfaceAddRoute(o["add-route"], d, "add_route", sv)); err != nil {
		if !fortiAPIPatch(o["add-route"]) {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("auto_discovery_sender", flattenVpnIpsecPhase2InterfaceAutoDiscoverySender(o["auto-discovery-sender"], d, "auto_discovery_sender", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-sender"]) {
			return fmt.Errorf("Error reading auto_discovery_sender: %v", err)
		}
	}

	if err = d.Set("auto_discovery_forwarder", flattenVpnIpsecPhase2InterfaceAutoDiscoveryForwarder(o["auto-discovery-forwarder"], d, "auto_discovery_forwarder", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-forwarder"]) {
			return fmt.Errorf("Error reading auto_discovery_forwarder: %v", err)
		}
	}

	if err = d.Set("keylifeseconds", flattenVpnIpsecPhase2InterfaceKeylifeseconds(o["keylifeseconds"], d, "keylifeseconds", sv)); err != nil {
		if !fortiAPIPatch(o["keylifeseconds"]) {
			return fmt.Errorf("Error reading keylifeseconds: %v", err)
		}
	}

	if err = d.Set("keylifekbs", flattenVpnIpsecPhase2InterfaceKeylifekbs(o["keylifekbs"], d, "keylifekbs", sv)); err != nil {
		if !fortiAPIPatch(o["keylifekbs"]) {
			return fmt.Errorf("Error reading keylifekbs: %v", err)
		}
	}

	if err = d.Set("keylife_type", flattenVpnIpsecPhase2InterfaceKeylifeType(o["keylife-type"], d, "keylife_type", sv)); err != nil {
		if !fortiAPIPatch(o["keylife-type"]) {
			return fmt.Errorf("Error reading keylife_type: %v", err)
		}
	}

	if err = d.Set("single_source", flattenVpnIpsecPhase2InterfaceSingleSource(o["single-source"], d, "single_source", sv)); err != nil {
		if !fortiAPIPatch(o["single-source"]) {
			return fmt.Errorf("Error reading single_source: %v", err)
		}
	}

	if err = d.Set("route_overlap", flattenVpnIpsecPhase2InterfaceRouteOverlap(o["route-overlap"], d, "route_overlap", sv)); err != nil {
		if !fortiAPIPatch(o["route-overlap"]) {
			return fmt.Errorf("Error reading route_overlap: %v", err)
		}
	}

	if err = d.Set("encapsulation", flattenVpnIpsecPhase2InterfaceEncapsulation(o["encapsulation"], d, "encapsulation", sv)); err != nil {
		if !fortiAPIPatch(o["encapsulation"]) {
			return fmt.Errorf("Error reading encapsulation: %v", err)
		}
	}

	if err = d.Set("l2tp", flattenVpnIpsecPhase2InterfaceL2Tp(o["l2tp"], d, "l2tp", sv)); err != nil {
		if !fortiAPIPatch(o["l2tp"]) {
			return fmt.Errorf("Error reading l2tp: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnIpsecPhase2InterfaceComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("initiator_ts_narrow", flattenVpnIpsecPhase2InterfaceInitiatorTsNarrow(o["initiator-ts-narrow"], d, "initiator_ts_narrow", sv)); err != nil {
		if !fortiAPIPatch(o["initiator-ts-narrow"]) {
			return fmt.Errorf("Error reading initiator_ts_narrow: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenVpnIpsecPhase2InterfaceDiffserv(o["diffserv"], d, "diffserv", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv"]) {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenVpnIpsecPhase2InterfaceDiffservcode(o["diffservcode"], d, "diffservcode", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode"]) {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("protocol", flattenVpnIpsecPhase2InterfaceProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("src_name", flattenVpnIpsecPhase2InterfaceSrcName(o["src-name"], d, "src_name", sv)); err != nil {
		if !fortiAPIPatch(o["src-name"]) {
			return fmt.Errorf("Error reading src_name: %v", err)
		}
	}

	if err = d.Set("src_name6", flattenVpnIpsecPhase2InterfaceSrcName6(o["src-name6"], d, "src_name6", sv)); err != nil {
		if !fortiAPIPatch(o["src-name6"]) {
			return fmt.Errorf("Error reading src_name6: %v", err)
		}
	}

	if err = d.Set("src_addr_type", flattenVpnIpsecPhase2InterfaceSrcAddrType(o["src-addr-type"], d, "src_addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["src-addr-type"]) {
			return fmt.Errorf("Error reading src_addr_type: %v", err)
		}
	}

	if err = d.Set("src_start_ip", flattenVpnIpsecPhase2InterfaceSrcStartIp(o["src-start-ip"], d, "src_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["src-start-ip"]) {
			return fmt.Errorf("Error reading src_start_ip: %v", err)
		}
	}

	if err = d.Set("src_start_ip6", flattenVpnIpsecPhase2InterfaceSrcStartIp6(o["src-start-ip6"], d, "src_start_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["src-start-ip6"]) {
			return fmt.Errorf("Error reading src_start_ip6: %v", err)
		}
	}

	if err = d.Set("src_end_ip", flattenVpnIpsecPhase2InterfaceSrcEndIp(o["src-end-ip"], d, "src_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["src-end-ip"]) {
			return fmt.Errorf("Error reading src_end_ip: %v", err)
		}
	}

	if err = d.Set("src_end_ip6", flattenVpnIpsecPhase2InterfaceSrcEndIp6(o["src-end-ip6"], d, "src_end_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["src-end-ip6"]) {
			return fmt.Errorf("Error reading src_end_ip6: %v", err)
		}
	}

	if err = d.Set("src_subnet", flattenVpnIpsecPhase2InterfaceSrcSubnet(o["src-subnet"], d, "src_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["src-subnet"]) {
			return fmt.Errorf("Error reading src_subnet: %v", err)
		}
	}

	if err = d.Set("src_subnet6", flattenVpnIpsecPhase2InterfaceSrcSubnet6(o["src-subnet6"], d, "src_subnet6", sv)); err != nil {
		if !fortiAPIPatch(o["src-subnet6"]) {
			return fmt.Errorf("Error reading src_subnet6: %v", err)
		}
	}

	if err = d.Set("src_port", flattenVpnIpsecPhase2InterfaceSrcPort(o["src-port"], d, "src_port", sv)); err != nil {
		if !fortiAPIPatch(o["src-port"]) {
			return fmt.Errorf("Error reading src_port: %v", err)
		}
	}

	if err = d.Set("dst_name", flattenVpnIpsecPhase2InterfaceDstName(o["dst-name"], d, "dst_name", sv)); err != nil {
		if !fortiAPIPatch(o["dst-name"]) {
			return fmt.Errorf("Error reading dst_name: %v", err)
		}
	}

	if err = d.Set("dst_name6", flattenVpnIpsecPhase2InterfaceDstName6(o["dst-name6"], d, "dst_name6", sv)); err != nil {
		if !fortiAPIPatch(o["dst-name6"]) {
			return fmt.Errorf("Error reading dst_name6: %v", err)
		}
	}

	if err = d.Set("dst_addr_type", flattenVpnIpsecPhase2InterfaceDstAddrType(o["dst-addr-type"], d, "dst_addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["dst-addr-type"]) {
			return fmt.Errorf("Error reading dst_addr_type: %v", err)
		}
	}

	if err = d.Set("dst_start_ip", flattenVpnIpsecPhase2InterfaceDstStartIp(o["dst-start-ip"], d, "dst_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dst-start-ip"]) {
			return fmt.Errorf("Error reading dst_start_ip: %v", err)
		}
	}

	if err = d.Set("dst_start_ip6", flattenVpnIpsecPhase2InterfaceDstStartIp6(o["dst-start-ip6"], d, "dst_start_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["dst-start-ip6"]) {
			return fmt.Errorf("Error reading dst_start_ip6: %v", err)
		}
	}

	if err = d.Set("dst_end_ip", flattenVpnIpsecPhase2InterfaceDstEndIp(o["dst-end-ip"], d, "dst_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dst-end-ip"]) {
			return fmt.Errorf("Error reading dst_end_ip: %v", err)
		}
	}

	if err = d.Set("dst_end_ip6", flattenVpnIpsecPhase2InterfaceDstEndIp6(o["dst-end-ip6"], d, "dst_end_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["dst-end-ip6"]) {
			return fmt.Errorf("Error reading dst_end_ip6: %v", err)
		}
	}

	if err = d.Set("dst_subnet", flattenVpnIpsecPhase2InterfaceDstSubnet(o["dst-subnet"], d, "dst_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["dst-subnet"]) {
			return fmt.Errorf("Error reading dst_subnet: %v", err)
		}
	}

	if err = d.Set("dst_subnet6", flattenVpnIpsecPhase2InterfaceDstSubnet6(o["dst-subnet6"], d, "dst_subnet6", sv)); err != nil {
		if !fortiAPIPatch(o["dst-subnet6"]) {
			return fmt.Errorf("Error reading dst_subnet6: %v", err)
		}
	}

	if err = d.Set("dst_port", flattenVpnIpsecPhase2InterfaceDstPort(o["dst-port"], d, "dst_port", sv)); err != nil {
		if !fortiAPIPatch(o["dst-port"]) {
			return fmt.Errorf("Error reading dst_port: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecPhase2InterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnIpsecPhase2InterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfacePhase1Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDhcpIpsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceProposal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfacePfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceIpv4Df(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDhgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceReplay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceKeepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceAutoNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceAddRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceAutoDiscoverySender(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceAutoDiscoveryForwarder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceKeylifeseconds(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceKeylifekbs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceKeylifeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSingleSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceRouteOverlap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceEncapsulation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceL2Tp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceInitiatorTsNarrow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDiffserv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDiffservcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcName6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcStartIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcEndIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcSubnet6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceSrcPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstName6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstStartIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstEndIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstSubnet6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InterfaceDstPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecPhase2Interface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandVpnIpsecPhase2InterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("phase1name"); ok {

		t, err := expandVpnIpsecPhase2InterfacePhase1Name(d, v, "phase1name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phase1name"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ipsec"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDhcpIpsec(d, v, "dhcp_ipsec", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ipsec"] = t
		}
	}

	if v, ok := d.GetOk("proposal"); ok {

		t, err := expandVpnIpsecPhase2InterfaceProposal(d, v, "proposal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proposal"] = t
		}
	}

	if v, ok := d.GetOk("pfs"); ok {

		t, err := expandVpnIpsecPhase2InterfacePfs(d, v, "pfs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pfs"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_df"); ok {

		t, err := expandVpnIpsecPhase2InterfaceIpv4Df(d, v, "ipv4_df", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-df"] = t
		}
	}

	if v, ok := d.GetOk("dhgrp"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDhgrp(d, v, "dhgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhgrp"] = t
		}
	}

	if v, ok := d.GetOk("replay"); ok {

		t, err := expandVpnIpsecPhase2InterfaceReplay(d, v, "replay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replay"] = t
		}
	}

	if v, ok := d.GetOk("keepalive"); ok {

		t, err := expandVpnIpsecPhase2InterfaceKeepalive(d, v, "keepalive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	if v, ok := d.GetOk("auto_negotiate"); ok {

		t, err := expandVpnIpsecPhase2InterfaceAutoNegotiate(d, v, "auto_negotiate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("add_route"); ok {

		t, err := expandVpnIpsecPhase2InterfaceAddRoute(d, v, "add_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_sender"); ok {

		t, err := expandVpnIpsecPhase2InterfaceAutoDiscoverySender(d, v, "auto_discovery_sender", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-sender"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_forwarder"); ok {

		t, err := expandVpnIpsecPhase2InterfaceAutoDiscoveryForwarder(d, v, "auto_discovery_forwarder", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-forwarder"] = t
		}
	}

	if v, ok := d.GetOk("keylifeseconds"); ok {

		t, err := expandVpnIpsecPhase2InterfaceKeylifeseconds(d, v, "keylifeseconds", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylifeseconds"] = t
		}
	}

	if v, ok := d.GetOk("keylifekbs"); ok {

		t, err := expandVpnIpsecPhase2InterfaceKeylifekbs(d, v, "keylifekbs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylifekbs"] = t
		}
	}

	if v, ok := d.GetOk("keylife_type"); ok {

		t, err := expandVpnIpsecPhase2InterfaceKeylifeType(d, v, "keylife_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylife-type"] = t
		}
	}

	if v, ok := d.GetOk("single_source"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSingleSource(d, v, "single_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-source"] = t
		}
	}

	if v, ok := d.GetOk("route_overlap"); ok {

		t, err := expandVpnIpsecPhase2InterfaceRouteOverlap(d, v, "route_overlap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-overlap"] = t
		}
	}

	if v, ok := d.GetOk("encapsulation"); ok {

		t, err := expandVpnIpsecPhase2InterfaceEncapsulation(d, v, "encapsulation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encapsulation"] = t
		}
	}

	if v, ok := d.GetOk("l2tp"); ok {

		t, err := expandVpnIpsecPhase2InterfaceL2Tp(d, v, "l2tp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2tp"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandVpnIpsecPhase2InterfaceComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("initiator_ts_narrow"); ok {

		t, err := expandVpnIpsecPhase2InterfaceInitiatorTsNarrow(d, v, "initiator_ts_narrow", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiator-ts-narrow"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDiffserv(d, v, "diffserv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDiffservcode(d, v, "diffservcode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOkExists("protocol"); ok {

		t, err := expandVpnIpsecPhase2InterfaceProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("src_name"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcName(d, v, "src_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-name"] = t
		}
	}

	if v, ok := d.GetOk("src_name6"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcName6(d, v, "src_name6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-name6"] = t
		}
	}

	if v, ok := d.GetOk("src_addr_type"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcAddrType(d, v, "src_addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("src_start_ip"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcStartIp(d, v, "src_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_start_ip6"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcStartIp6(d, v, "src_start_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-start-ip6"] = t
		}
	}

	if v, ok := d.GetOk("src_end_ip"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcEndIp(d, v, "src_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_end_ip6"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcEndIp6(d, v, "src_end_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-end-ip6"] = t
		}
	}

	if v, ok := d.GetOk("src_subnet"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcSubnet(d, v, "src_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-subnet"] = t
		}
	}

	if v, ok := d.GetOk("src_subnet6"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcSubnet6(d, v, "src_subnet6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-subnet6"] = t
		}
	}

	if v, ok := d.GetOkExists("src_port"); ok {

		t, err := expandVpnIpsecPhase2InterfaceSrcPort(d, v, "src_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-port"] = t
		}
	}

	if v, ok := d.GetOk("dst_name"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstName(d, v, "dst_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-name"] = t
		}
	}

	if v, ok := d.GetOk("dst_name6"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstName6(d, v, "dst_name6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-name6"] = t
		}
	}

	if v, ok := d.GetOk("dst_addr_type"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstAddrType(d, v, "dst_addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("dst_start_ip"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstStartIp(d, v, "dst_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("dst_start_ip6"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstStartIp6(d, v, "dst_start_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-start-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dst_end_ip"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstEndIp(d, v, "dst_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("dst_end_ip6"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstEndIp6(d, v, "dst_end_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-end-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dst_subnet"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstSubnet(d, v, "dst_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-subnet"] = t
		}
	}

	if v, ok := d.GetOk("dst_subnet6"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstSubnet6(d, v, "dst_subnet6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-subnet6"] = t
		}
	}

	if v, ok := d.GetOkExists("dst_port"); ok {

		t, err := expandVpnIpsecPhase2InterfaceDstPort(d, v, "dst_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-port"] = t
		}
	}

	return &obj, nil
}
