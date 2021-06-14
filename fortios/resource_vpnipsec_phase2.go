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

func resourceVpnIpsecPhase2() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecPhase2Create,
		Read:   resourceVpnIpsecPhase2Read,
		Update: resourceVpnIpsecPhase2Update,
		Delete: resourceVpnIpsecPhase2Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"phase1name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"dhcp_ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_natip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"selector_match": &schema.Schema{
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

func resourceVpnIpsecPhase2Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnIpsecPhase2(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase2 resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecPhase2(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase2 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase2")
	}

	return resourceVpnIpsecPhase2Read(d, m)
}

func resourceVpnIpsecPhase2Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnIpsecPhase2(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase2 resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecPhase2(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase2 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase2")
	}

	return resourceVpnIpsecPhase2Read(d, m)
}

func resourceVpnIpsecPhase2Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnIpsecPhase2(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecPhase2 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase2Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnIpsecPhase2(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase2 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecPhase2(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase2 resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecPhase2Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Phase1Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DhcpIpsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2UseNatip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SelectorMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Proposal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Pfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Ipv4Df(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Dhgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Replay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Keepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2AutoNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2AddRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Keylifeseconds(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Keylifekbs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2KeylifeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SingleSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2RouteOverlap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Encapsulation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2L2Tp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Comments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2InitiatorTsNarrow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Diffserv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Diffservcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2Protocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcName6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcStartIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcEndIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenVpnIpsecPhase2SrcSubnet6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2SrcPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstName6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstStartIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstEndIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenVpnIpsecPhase2DstSubnet6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase2DstPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnIpsecPhase2(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnIpsecPhase2Name(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("phase1name", flattenVpnIpsecPhase2Phase1Name(o["phase1name"], d, "phase1name", sv)); err != nil {
		if !fortiAPIPatch(o["phase1name"]) {
			return fmt.Errorf("Error reading phase1name: %v", err)
		}
	}

	if err = d.Set("dhcp_ipsec", flattenVpnIpsecPhase2DhcpIpsec(o["dhcp-ipsec"], d, "dhcp_ipsec", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-ipsec"]) {
			return fmt.Errorf("Error reading dhcp_ipsec: %v", err)
		}
	}

	if err = d.Set("use_natip", flattenVpnIpsecPhase2UseNatip(o["use-natip"], d, "use_natip", sv)); err != nil {
		if !fortiAPIPatch(o["use-natip"]) {
			return fmt.Errorf("Error reading use_natip: %v", err)
		}
	}

	if err = d.Set("selector_match", flattenVpnIpsecPhase2SelectorMatch(o["selector-match"], d, "selector_match", sv)); err != nil {
		if !fortiAPIPatch(o["selector-match"]) {
			return fmt.Errorf("Error reading selector_match: %v", err)
		}
	}

	if err = d.Set("proposal", flattenVpnIpsecPhase2Proposal(o["proposal"], d, "proposal", sv)); err != nil {
		if !fortiAPIPatch(o["proposal"]) {
			return fmt.Errorf("Error reading proposal: %v", err)
		}
	}

	if err = d.Set("pfs", flattenVpnIpsecPhase2Pfs(o["pfs"], d, "pfs", sv)); err != nil {
		if !fortiAPIPatch(o["pfs"]) {
			return fmt.Errorf("Error reading pfs: %v", err)
		}
	}

	if err = d.Set("ipv4_df", flattenVpnIpsecPhase2Ipv4Df(o["ipv4-df"], d, "ipv4_df", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-df"]) {
			return fmt.Errorf("Error reading ipv4_df: %v", err)
		}
	}

	if err = d.Set("dhgrp", flattenVpnIpsecPhase2Dhgrp(o["dhgrp"], d, "dhgrp", sv)); err != nil {
		if !fortiAPIPatch(o["dhgrp"]) {
			return fmt.Errorf("Error reading dhgrp: %v", err)
		}
	}

	if err = d.Set("replay", flattenVpnIpsecPhase2Replay(o["replay"], d, "replay", sv)); err != nil {
		if !fortiAPIPatch(o["replay"]) {
			return fmt.Errorf("Error reading replay: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenVpnIpsecPhase2Keepalive(o["keepalive"], d, "keepalive", sv)); err != nil {
		if !fortiAPIPatch(o["keepalive"]) {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	if err = d.Set("auto_negotiate", flattenVpnIpsecPhase2AutoNegotiate(o["auto-negotiate"], d, "auto_negotiate", sv)); err != nil {
		if !fortiAPIPatch(o["auto-negotiate"]) {
			return fmt.Errorf("Error reading auto_negotiate: %v", err)
		}
	}

	if err = d.Set("add_route", flattenVpnIpsecPhase2AddRoute(o["add-route"], d, "add_route", sv)); err != nil {
		if !fortiAPIPatch(o["add-route"]) {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("keylifeseconds", flattenVpnIpsecPhase2Keylifeseconds(o["keylifeseconds"], d, "keylifeseconds", sv)); err != nil {
		if !fortiAPIPatch(o["keylifeseconds"]) {
			return fmt.Errorf("Error reading keylifeseconds: %v", err)
		}
	}

	if err = d.Set("keylifekbs", flattenVpnIpsecPhase2Keylifekbs(o["keylifekbs"], d, "keylifekbs", sv)); err != nil {
		if !fortiAPIPatch(o["keylifekbs"]) {
			return fmt.Errorf("Error reading keylifekbs: %v", err)
		}
	}

	if err = d.Set("keylife_type", flattenVpnIpsecPhase2KeylifeType(o["keylife-type"], d, "keylife_type", sv)); err != nil {
		if !fortiAPIPatch(o["keylife-type"]) {
			return fmt.Errorf("Error reading keylife_type: %v", err)
		}
	}

	if err = d.Set("single_source", flattenVpnIpsecPhase2SingleSource(o["single-source"], d, "single_source", sv)); err != nil {
		if !fortiAPIPatch(o["single-source"]) {
			return fmt.Errorf("Error reading single_source: %v", err)
		}
	}

	if err = d.Set("route_overlap", flattenVpnIpsecPhase2RouteOverlap(o["route-overlap"], d, "route_overlap", sv)); err != nil {
		if !fortiAPIPatch(o["route-overlap"]) {
			return fmt.Errorf("Error reading route_overlap: %v", err)
		}
	}

	if err = d.Set("encapsulation", flattenVpnIpsecPhase2Encapsulation(o["encapsulation"], d, "encapsulation", sv)); err != nil {
		if !fortiAPIPatch(o["encapsulation"]) {
			return fmt.Errorf("Error reading encapsulation: %v", err)
		}
	}

	if err = d.Set("l2tp", flattenVpnIpsecPhase2L2Tp(o["l2tp"], d, "l2tp", sv)); err != nil {
		if !fortiAPIPatch(o["l2tp"]) {
			return fmt.Errorf("Error reading l2tp: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnIpsecPhase2Comments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("initiator_ts_narrow", flattenVpnIpsecPhase2InitiatorTsNarrow(o["initiator-ts-narrow"], d, "initiator_ts_narrow", sv)); err != nil {
		if !fortiAPIPatch(o["initiator-ts-narrow"]) {
			return fmt.Errorf("Error reading initiator_ts_narrow: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenVpnIpsecPhase2Diffserv(o["diffserv"], d, "diffserv", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv"]) {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenVpnIpsecPhase2Diffservcode(o["diffservcode"], d, "diffservcode", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode"]) {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("protocol", flattenVpnIpsecPhase2Protocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("src_name", flattenVpnIpsecPhase2SrcName(o["src-name"], d, "src_name", sv)); err != nil {
		if !fortiAPIPatch(o["src-name"]) {
			return fmt.Errorf("Error reading src_name: %v", err)
		}
	}

	if err = d.Set("src_name6", flattenVpnIpsecPhase2SrcName6(o["src-name6"], d, "src_name6", sv)); err != nil {
		if !fortiAPIPatch(o["src-name6"]) {
			return fmt.Errorf("Error reading src_name6: %v", err)
		}
	}

	if err = d.Set("src_addr_type", flattenVpnIpsecPhase2SrcAddrType(o["src-addr-type"], d, "src_addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["src-addr-type"]) {
			return fmt.Errorf("Error reading src_addr_type: %v", err)
		}
	}

	if err = d.Set("src_start_ip", flattenVpnIpsecPhase2SrcStartIp(o["src-start-ip"], d, "src_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["src-start-ip"]) {
			return fmt.Errorf("Error reading src_start_ip: %v", err)
		}
	}

	if err = d.Set("src_start_ip6", flattenVpnIpsecPhase2SrcStartIp6(o["src-start-ip6"], d, "src_start_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["src-start-ip6"]) {
			return fmt.Errorf("Error reading src_start_ip6: %v", err)
		}
	}

	if err = d.Set("src_end_ip", flattenVpnIpsecPhase2SrcEndIp(o["src-end-ip"], d, "src_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["src-end-ip"]) {
			return fmt.Errorf("Error reading src_end_ip: %v", err)
		}
	}

	if err = d.Set("src_end_ip6", flattenVpnIpsecPhase2SrcEndIp6(o["src-end-ip6"], d, "src_end_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["src-end-ip6"]) {
			return fmt.Errorf("Error reading src_end_ip6: %v", err)
		}
	}

	if err = d.Set("src_subnet", flattenVpnIpsecPhase2SrcSubnet(o["src-subnet"], d, "src_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["src-subnet"]) {
			return fmt.Errorf("Error reading src_subnet: %v", err)
		}
	}

	if err = d.Set("src_subnet6", flattenVpnIpsecPhase2SrcSubnet6(o["src-subnet6"], d, "src_subnet6", sv)); err != nil {
		if !fortiAPIPatch(o["src-subnet6"]) {
			return fmt.Errorf("Error reading src_subnet6: %v", err)
		}
	}

	if err = d.Set("src_port", flattenVpnIpsecPhase2SrcPort(o["src-port"], d, "src_port", sv)); err != nil {
		if !fortiAPIPatch(o["src-port"]) {
			return fmt.Errorf("Error reading src_port: %v", err)
		}
	}

	if err = d.Set("dst_name", flattenVpnIpsecPhase2DstName(o["dst-name"], d, "dst_name", sv)); err != nil {
		if !fortiAPIPatch(o["dst-name"]) {
			return fmt.Errorf("Error reading dst_name: %v", err)
		}
	}

	if err = d.Set("dst_name6", flattenVpnIpsecPhase2DstName6(o["dst-name6"], d, "dst_name6", sv)); err != nil {
		if !fortiAPIPatch(o["dst-name6"]) {
			return fmt.Errorf("Error reading dst_name6: %v", err)
		}
	}

	if err = d.Set("dst_addr_type", flattenVpnIpsecPhase2DstAddrType(o["dst-addr-type"], d, "dst_addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["dst-addr-type"]) {
			return fmt.Errorf("Error reading dst_addr_type: %v", err)
		}
	}

	if err = d.Set("dst_start_ip", flattenVpnIpsecPhase2DstStartIp(o["dst-start-ip"], d, "dst_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dst-start-ip"]) {
			return fmt.Errorf("Error reading dst_start_ip: %v", err)
		}
	}

	if err = d.Set("dst_start_ip6", flattenVpnIpsecPhase2DstStartIp6(o["dst-start-ip6"], d, "dst_start_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["dst-start-ip6"]) {
			return fmt.Errorf("Error reading dst_start_ip6: %v", err)
		}
	}

	if err = d.Set("dst_end_ip", flattenVpnIpsecPhase2DstEndIp(o["dst-end-ip"], d, "dst_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dst-end-ip"]) {
			return fmt.Errorf("Error reading dst_end_ip: %v", err)
		}
	}

	if err = d.Set("dst_end_ip6", flattenVpnIpsecPhase2DstEndIp6(o["dst-end-ip6"], d, "dst_end_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["dst-end-ip6"]) {
			return fmt.Errorf("Error reading dst_end_ip6: %v", err)
		}
	}

	if err = d.Set("dst_subnet", flattenVpnIpsecPhase2DstSubnet(o["dst-subnet"], d, "dst_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["dst-subnet"]) {
			return fmt.Errorf("Error reading dst_subnet: %v", err)
		}
	}

	if err = d.Set("dst_subnet6", flattenVpnIpsecPhase2DstSubnet6(o["dst-subnet6"], d, "dst_subnet6", sv)); err != nil {
		if !fortiAPIPatch(o["dst-subnet6"]) {
			return fmt.Errorf("Error reading dst_subnet6: %v", err)
		}
	}

	if err = d.Set("dst_port", flattenVpnIpsecPhase2DstPort(o["dst-port"], d, "dst_port", sv)); err != nil {
		if !fortiAPIPatch(o["dst-port"]) {
			return fmt.Errorf("Error reading dst_port: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecPhase2FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnIpsecPhase2Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Phase1Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DhcpIpsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2UseNatip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SelectorMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Proposal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Pfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Ipv4Df(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Dhgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Replay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Keepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2AutoNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2AddRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Keylifeseconds(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Keylifekbs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2KeylifeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SingleSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2RouteOverlap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Encapsulation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2L2Tp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Comments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2InitiatorTsNarrow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Diffserv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Diffservcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2Protocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcName6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcStartIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcEndIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcSubnet6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2SrcPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstName6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstStartIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstEndIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstSubnet6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase2DstPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecPhase2(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandVpnIpsecPhase2Name(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("phase1name"); ok {

		t, err := expandVpnIpsecPhase2Phase1Name(d, v, "phase1name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phase1name"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ipsec"); ok {

		t, err := expandVpnIpsecPhase2DhcpIpsec(d, v, "dhcp_ipsec", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ipsec"] = t
		}
	}

	if v, ok := d.GetOk("use_natip"); ok {

		t, err := expandVpnIpsecPhase2UseNatip(d, v, "use_natip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-natip"] = t
		}
	}

	if v, ok := d.GetOk("selector_match"); ok {

		t, err := expandVpnIpsecPhase2SelectorMatch(d, v, "selector_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selector-match"] = t
		}
	}

	if v, ok := d.GetOk("proposal"); ok {

		t, err := expandVpnIpsecPhase2Proposal(d, v, "proposal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proposal"] = t
		}
	}

	if v, ok := d.GetOk("pfs"); ok {

		t, err := expandVpnIpsecPhase2Pfs(d, v, "pfs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pfs"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_df"); ok {

		t, err := expandVpnIpsecPhase2Ipv4Df(d, v, "ipv4_df", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-df"] = t
		}
	}

	if v, ok := d.GetOk("dhgrp"); ok {

		t, err := expandVpnIpsecPhase2Dhgrp(d, v, "dhgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhgrp"] = t
		}
	}

	if v, ok := d.GetOk("replay"); ok {

		t, err := expandVpnIpsecPhase2Replay(d, v, "replay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replay"] = t
		}
	}

	if v, ok := d.GetOk("keepalive"); ok {

		t, err := expandVpnIpsecPhase2Keepalive(d, v, "keepalive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	if v, ok := d.GetOk("auto_negotiate"); ok {

		t, err := expandVpnIpsecPhase2AutoNegotiate(d, v, "auto_negotiate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("add_route"); ok {

		t, err := expandVpnIpsecPhase2AddRoute(d, v, "add_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("keylifeseconds"); ok {

		t, err := expandVpnIpsecPhase2Keylifeseconds(d, v, "keylifeseconds", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylifeseconds"] = t
		}
	}

	if v, ok := d.GetOk("keylifekbs"); ok {

		t, err := expandVpnIpsecPhase2Keylifekbs(d, v, "keylifekbs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylifekbs"] = t
		}
	}

	if v, ok := d.GetOk("keylife_type"); ok {

		t, err := expandVpnIpsecPhase2KeylifeType(d, v, "keylife_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylife-type"] = t
		}
	}

	if v, ok := d.GetOk("single_source"); ok {

		t, err := expandVpnIpsecPhase2SingleSource(d, v, "single_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-source"] = t
		}
	}

	if v, ok := d.GetOk("route_overlap"); ok {

		t, err := expandVpnIpsecPhase2RouteOverlap(d, v, "route_overlap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-overlap"] = t
		}
	}

	if v, ok := d.GetOk("encapsulation"); ok {

		t, err := expandVpnIpsecPhase2Encapsulation(d, v, "encapsulation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encapsulation"] = t
		}
	}

	if v, ok := d.GetOk("l2tp"); ok {

		t, err := expandVpnIpsecPhase2L2Tp(d, v, "l2tp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2tp"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandVpnIpsecPhase2Comments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("initiator_ts_narrow"); ok {

		t, err := expandVpnIpsecPhase2InitiatorTsNarrow(d, v, "initiator_ts_narrow", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiator-ts-narrow"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok {

		t, err := expandVpnIpsecPhase2Diffserv(d, v, "diffserv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok {

		t, err := expandVpnIpsecPhase2Diffservcode(d, v, "diffservcode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOkExists("protocol"); ok {

		t, err := expandVpnIpsecPhase2Protocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("src_name"); ok {

		t, err := expandVpnIpsecPhase2SrcName(d, v, "src_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-name"] = t
		}
	}

	if v, ok := d.GetOk("src_name6"); ok {

		t, err := expandVpnIpsecPhase2SrcName6(d, v, "src_name6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-name6"] = t
		}
	}

	if v, ok := d.GetOk("src_addr_type"); ok {

		t, err := expandVpnIpsecPhase2SrcAddrType(d, v, "src_addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("src_start_ip"); ok {

		t, err := expandVpnIpsecPhase2SrcStartIp(d, v, "src_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_start_ip6"); ok {

		t, err := expandVpnIpsecPhase2SrcStartIp6(d, v, "src_start_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-start-ip6"] = t
		}
	}

	if v, ok := d.GetOk("src_end_ip"); ok {

		t, err := expandVpnIpsecPhase2SrcEndIp(d, v, "src_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_end_ip6"); ok {

		t, err := expandVpnIpsecPhase2SrcEndIp6(d, v, "src_end_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-end-ip6"] = t
		}
	}

	if v, ok := d.GetOk("src_subnet"); ok {

		t, err := expandVpnIpsecPhase2SrcSubnet(d, v, "src_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-subnet"] = t
		}
	}

	if v, ok := d.GetOk("src_subnet6"); ok {

		t, err := expandVpnIpsecPhase2SrcSubnet6(d, v, "src_subnet6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-subnet6"] = t
		}
	}

	if v, ok := d.GetOkExists("src_port"); ok {

		t, err := expandVpnIpsecPhase2SrcPort(d, v, "src_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-port"] = t
		}
	}

	if v, ok := d.GetOk("dst_name"); ok {

		t, err := expandVpnIpsecPhase2DstName(d, v, "dst_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-name"] = t
		}
	}

	if v, ok := d.GetOk("dst_name6"); ok {

		t, err := expandVpnIpsecPhase2DstName6(d, v, "dst_name6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-name6"] = t
		}
	}

	if v, ok := d.GetOk("dst_addr_type"); ok {

		t, err := expandVpnIpsecPhase2DstAddrType(d, v, "dst_addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("dst_start_ip"); ok {

		t, err := expandVpnIpsecPhase2DstStartIp(d, v, "dst_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("dst_start_ip6"); ok {

		t, err := expandVpnIpsecPhase2DstStartIp6(d, v, "dst_start_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-start-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dst_end_ip"); ok {

		t, err := expandVpnIpsecPhase2DstEndIp(d, v, "dst_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("dst_end_ip6"); ok {

		t, err := expandVpnIpsecPhase2DstEndIp6(d, v, "dst_end_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-end-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dst_subnet"); ok {

		t, err := expandVpnIpsecPhase2DstSubnet(d, v, "dst_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-subnet"] = t
		}
	}

	if v, ok := d.GetOk("dst_subnet6"); ok {

		t, err := expandVpnIpsecPhase2DstSubnet6(d, v, "dst_subnet6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-subnet6"] = t
		}
	}

	if v, ok := d.GetOkExists("dst_port"); ok {

		t, err := expandVpnIpsecPhase2DstPort(d, v, "dst_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-port"] = t
		}
	}

	return &obj, nil
}
