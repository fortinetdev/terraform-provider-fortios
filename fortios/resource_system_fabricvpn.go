// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Setup for self orchestrated fabric auto discovery VPN.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFabricVpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFabricVpnUpdate,
		Read:   resourceSystemFabricVpnRead,
		Update: resourceSystemFabricVpnUpdate,
		Delete: resourceSystemFabricVpnDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"branch_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"policy_rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overlays": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"ipsec_network_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"overlay_tunnel_block": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remote_gw": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"bgp_neighbor": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),
							Optional:     true,
						},
						"overlay_policy": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bgp_network": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"route_policy": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bgp_neighbor_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),
							Optional:     true,
						},
						"bgp_neighbor_range": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipsec_phase1": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"sdwan_member": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"advertised_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bgp_network": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"firewall_address": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"policies": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"loopback_address_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loopback_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"loopback_advertised_subnet": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"psksecret": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"bgp_as_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bgp_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sdwan_zone": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"health_checks": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemFabricVpnUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSystemFabricVpn(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFabricVpn resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFabricVpn(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemFabricVpn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFabricVpn")
	}

	return resourceSystemFabricVpnRead(d, m)
}

func resourceSystemFabricVpnDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFabricVpn(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFabricVpn resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFabricVpn(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFabricVpn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFabricVpnRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSystemFabricVpn(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemFabricVpn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFabricVpn(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFabricVpn resource from API: %v", err)
	}
	return nil
}

func flattenSystemFabricVpnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnSyncMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnBranchName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnPolicyRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnVpnRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlays(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenSystemFabricVpnOverlaysName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ipsec-network-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_network_id"
			}
			tmp["ipsec_network_id"] = flattenSystemFabricVpnOverlaysIpsecNetworkId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["overlay-tunnel-block"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_tunnel_block"
			}
			tmp["overlay_tunnel_block"] = flattenSystemFabricVpnOverlaysOverlayTunnelBlock(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["remote-gw"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_gw"
			}
			tmp["remote_gw"] = flattenSystemFabricVpnOverlaysRemoteGw(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
			}
			tmp["interface"] = flattenSystemFabricVpnOverlaysInterface(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bgp-neighbor"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor"
			}
			tmp["bgp_neighbor"] = flattenSystemFabricVpnOverlaysBgpNeighbor(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["overlay-policy"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_policy"
			}
			tmp["overlay_policy"] = flattenSystemFabricVpnOverlaysOverlayPolicy(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bgp-network"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_network"
			}
			tmp["bgp_network"] = flattenSystemFabricVpnOverlaysBgpNetwork(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["route-policy"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "route_policy"
			}
			tmp["route_policy"] = flattenSystemFabricVpnOverlaysRoutePolicy(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bgp-neighbor-group"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor_group"
			}
			tmp["bgp_neighbor_group"] = flattenSystemFabricVpnOverlaysBgpNeighborGroup(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bgp-neighbor-range"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor_range"
			}
			tmp["bgp_neighbor_range"] = flattenSystemFabricVpnOverlaysBgpNeighborRange(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ipsec-phase1"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_phase1"
			}
			tmp["ipsec_phase1"] = flattenSystemFabricVpnOverlaysIpsecPhase1(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sdwan-member"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sdwan_member"
			}
			tmp["sdwan_member"] = flattenSystemFabricVpnOverlaysSdwanMember(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemFabricVpnOverlaysName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysIpsecNetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnOverlaysOverlayTunnelBlock(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemFabricVpnOverlaysRemoteGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysBgpNeighbor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysOverlayPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnOverlaysBgpNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnOverlaysRoutePolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnOverlaysBgpNeighborGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysBgpNeighborRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnOverlaysIpsecPhase1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnOverlaysSdwanMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnAdvertisedSubnets(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenSystemFabricVpnAdvertisedSubnetsId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["prefix"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
			}
			tmp["prefix"] = flattenSystemFabricVpnAdvertisedSubnetsPrefix(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["access"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "access"
			}
			tmp["access"] = flattenSystemFabricVpnAdvertisedSubnetsAccess(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bgp-network"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_network"
			}
			tmp["bgp_network"] = flattenSystemFabricVpnAdvertisedSubnetsBgpNetwork(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["firewall-address"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "firewall_address"
			}
			tmp["firewall_address"] = flattenSystemFabricVpnAdvertisedSubnetsFirewallAddress(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["policies"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "policies"
			}
			tmp["policies"] = flattenSystemFabricVpnAdvertisedSubnetsPolicies(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemFabricVpnAdvertisedSubnetsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnAdvertisedSubnetsPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemFabricVpnAdvertisedSubnetsAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnAdvertisedSubnetsBgpNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnAdvertisedSubnetsFirewallAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnAdvertisedSubnetsPolicies(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnLoopbackAddressBlock(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemFabricVpnLoopbackInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnLoopbackAdvertisedSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnBgpAsString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnBgpAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFabricVpnSdwanZone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFabricVpnHealthChecks(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFabricVpn(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("status", flattenSystemFabricVpnStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sync_mode", flattenSystemFabricVpnSyncMode(o["sync-mode"], d, "sync_mode", sv)); err != nil {
		if !fortiAPIPatch(o["sync-mode"]) {
			return fmt.Errorf("Error reading sync_mode: %v", err)
		}
	}

	if err = d.Set("branch_name", flattenSystemFabricVpnBranchName(o["branch-name"], d, "branch_name", sv)); err != nil {
		if !fortiAPIPatch(o["branch-name"]) {
			return fmt.Errorf("Error reading branch_name: %v", err)
		}
	}

	if err = d.Set("policy_rule", flattenSystemFabricVpnPolicyRule(o["policy-rule"], d, "policy_rule", sv)); err != nil {
		if !fortiAPIPatch(o["policy-rule"]) {
			return fmt.Errorf("Error reading policy_rule: %v", err)
		}
	}

	if err = d.Set("vpn_role", flattenSystemFabricVpnVpnRole(o["vpn-role"], d, "vpn_role", sv)); err != nil {
		if !fortiAPIPatch(o["vpn-role"]) {
			return fmt.Errorf("Error reading vpn_role: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("overlays", flattenSystemFabricVpnOverlays(o["overlays"], d, "overlays", sv)); err != nil {
			if !fortiAPIPatch(o["overlays"]) {
				return fmt.Errorf("Error reading overlays: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("overlays"); ok {
			if err = d.Set("overlays", flattenSystemFabricVpnOverlays(o["overlays"], d, "overlays", sv)); err != nil {
				if !fortiAPIPatch(o["overlays"]) {
					return fmt.Errorf("Error reading overlays: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("advertised_subnets", flattenSystemFabricVpnAdvertisedSubnets(o["advertised-subnets"], d, "advertised_subnets", sv)); err != nil {
			if !fortiAPIPatch(o["advertised-subnets"]) {
				return fmt.Errorf("Error reading advertised_subnets: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("advertised_subnets"); ok {
			if err = d.Set("advertised_subnets", flattenSystemFabricVpnAdvertisedSubnets(o["advertised-subnets"], d, "advertised_subnets", sv)); err != nil {
				if !fortiAPIPatch(o["advertised-subnets"]) {
					return fmt.Errorf("Error reading advertised_subnets: %v", err)
				}
			}
		}
	}

	if err = d.Set("loopback_address_block", flattenSystemFabricVpnLoopbackAddressBlock(o["loopback-address-block"], d, "loopback_address_block", sv)); err != nil {
		if !fortiAPIPatch(o["loopback-address-block"]) {
			return fmt.Errorf("Error reading loopback_address_block: %v", err)
		}
	}

	if err = d.Set("loopback_interface", flattenSystemFabricVpnLoopbackInterface(o["loopback-interface"], d, "loopback_interface", sv)); err != nil {
		if !fortiAPIPatch(o["loopback-interface"]) {
			return fmt.Errorf("Error reading loopback_interface: %v", err)
		}
	}

	if err = d.Set("loopback_advertised_subnet", flattenSystemFabricVpnLoopbackAdvertisedSubnet(o["loopback-advertised-subnet"], d, "loopback_advertised_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["loopback-advertised-subnet"]) {
			return fmt.Errorf("Error reading loopback_advertised_subnet: %v", err)
		}
	}

	if _, ok := o["bgp-as"].(string); ok {
		if err = d.Set("bgp_as_string", flattenSystemFabricVpnBgpAsString(o["bgp-as"], d, "bgp_as_string", sv)); err != nil {
			if !fortiAPIPatch(o["bgp-as"]) {
				return fmt.Errorf("Error reading bgp_as_string: %v", err)
			}
		}
	}

	if _, ok := o["bgp-as"].(float64); ok {
		if err = d.Set("bgp_as", flattenSystemFabricVpnBgpAs(o["bgp-as"], d, "bgp_as", sv)); err != nil {
			if !fortiAPIPatch(o["bgp-as"]) {
				return fmt.Errorf("Error reading bgp_as: %v", err)
			}
		}
	}

	if err = d.Set("sdwan_zone", flattenSystemFabricVpnSdwanZone(o["sdwan-zone"], d, "sdwan_zone", sv)); err != nil {
		if !fortiAPIPatch(o["sdwan-zone"]) {
			return fmt.Errorf("Error reading sdwan_zone: %v", err)
		}
	}

	if err = d.Set("health_checks", flattenSystemFabricVpnHealthChecks(o["health-checks"], d, "health_checks", sv)); err != nil {
		if !fortiAPIPatch(o["health-checks"]) {
			return fmt.Errorf("Error reading health_checks: %v", err)
		}
	}

	return nil
}

func flattenSystemFabricVpnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemFabricVpnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnSyncMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnBranchName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnPolicyRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnVpnRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemFabricVpnOverlaysName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_network_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-network-id"], _ = expandSystemFabricVpnOverlaysIpsecNetworkId(d, i["ipsec_network_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ipsec-network-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_tunnel_block"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["overlay-tunnel-block"], _ = expandSystemFabricVpnOverlaysOverlayTunnelBlock(d, i["overlay_tunnel_block"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_gw"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-gw"], _ = expandSystemFabricVpnOverlaysRemoteGw(d, i["remote_gw"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSystemFabricVpnOverlaysInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bgp-neighbor"], _ = expandSystemFabricVpnOverlaysBgpNeighbor(d, i["bgp_neighbor"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bgp-neighbor"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["overlay-policy"], _ = expandSystemFabricVpnOverlaysOverlayPolicy(d, i["overlay_policy"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["overlay-policy"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_network"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bgp-network"], _ = expandSystemFabricVpnOverlaysBgpNetwork(d, i["bgp_network"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bgp-network"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-policy"], _ = expandSystemFabricVpnOverlaysRoutePolicy(d, i["route_policy"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-policy"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bgp-neighbor-group"], _ = expandSystemFabricVpnOverlaysBgpNeighborGroup(d, i["bgp_neighbor_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bgp-neighbor-group"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_neighbor_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bgp-neighbor-range"], _ = expandSystemFabricVpnOverlaysBgpNeighborRange(d, i["bgp_neighbor_range"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bgp-neighbor-range"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_phase1"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-phase1"], _ = expandSystemFabricVpnOverlaysIpsecPhase1(d, i["ipsec_phase1"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ipsec-phase1"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sdwan_member"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sdwan-member"], _ = expandSystemFabricVpnOverlaysSdwanMember(d, i["sdwan_member"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sdwan-member"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemFabricVpnOverlaysName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysIpsecNetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysOverlayTunnelBlock(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysRemoteGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysBgpNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysOverlayPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysBgpNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysRoutePolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysBgpNeighborGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysBgpNeighborRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysIpsecPhase1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnOverlaysSdwanMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnets(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemFabricVpnAdvertisedSubnetsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandSystemFabricVpnAdvertisedSubnetsPrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["access"], _ = expandSystemFabricVpnAdvertisedSubnetsAccess(d, i["access"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bgp_network"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bgp-network"], _ = expandSystemFabricVpnAdvertisedSubnetsBgpNetwork(d, i["bgp_network"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bgp-network"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "firewall_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["firewall-address"], _ = expandSystemFabricVpnAdvertisedSubnetsFirewallAddress(d, i["firewall_address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["firewall-address"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policies"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["policies"], _ = expandSystemFabricVpnAdvertisedSubnetsPolicies(d, i["policies"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["policies"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemFabricVpnAdvertisedSubnetsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnetsPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnetsAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnetsBgpNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnetsFirewallAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnAdvertisedSubnetsPolicies(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnLoopbackAddressBlock(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnLoopbackInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnLoopbackAdvertisedSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnPsksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnBgpAsString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnBgpAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnSdwanZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFabricVpnHealthChecks(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFabricVpn(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemFabricVpnStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("sync_mode"); ok {
		if setArgNil {
			obj["sync-mode"] = nil
		} else {
			t, err := expandSystemFabricVpnSyncMode(d, v, "sync_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sync-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("branch_name"); ok {
		if setArgNil {
			obj["branch-name"] = nil
		} else {
			t, err := expandSystemFabricVpnBranchName(d, v, "branch_name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["branch-name"] = t
			}
		}
	} else if d.HasChange("branch_name") {
		obj["branch-name"] = nil
	}

	if v, ok := d.GetOk("policy_rule"); ok {
		if setArgNil {
			obj["policy-rule"] = nil
		} else {
			t, err := expandSystemFabricVpnPolicyRule(d, v, "policy_rule", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["policy-rule"] = t
			}
		}
	}

	if v, ok := d.GetOk("vpn_role"); ok {
		if setArgNil {
			obj["vpn-role"] = nil
		} else {
			t, err := expandSystemFabricVpnVpnRole(d, v, "vpn_role", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vpn-role"] = t
			}
		}
	}

	if v, ok := d.GetOk("overlays"); ok || d.HasChange("overlays") {
		if setArgNil {
			obj["overlays"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemFabricVpnOverlays(d, v, "overlays", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["overlays"] = t
			}
		}
	}

	if v, ok := d.GetOk("advertised_subnets"); ok || d.HasChange("advertised_subnets") {
		if setArgNil {
			obj["advertised-subnets"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemFabricVpnAdvertisedSubnets(d, v, "advertised_subnets", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["advertised-subnets"] = t
			}
		}
	}

	if v, ok := d.GetOk("loopback_address_block"); ok {
		if setArgNil {
			obj["loopback-address-block"] = nil
		} else {
			t, err := expandSystemFabricVpnLoopbackAddressBlock(d, v, "loopback_address_block", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["loopback-address-block"] = t
			}
		}
	}

	if v, ok := d.GetOk("loopback_interface"); ok {
		if setArgNil {
			obj["loopback-interface"] = nil
		} else {
			t, err := expandSystemFabricVpnLoopbackInterface(d, v, "loopback_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["loopback-interface"] = t
			}
		}
	} else if d.HasChange("loopback_interface") {
		obj["loopback-interface"] = nil
	}

	if v, ok := d.GetOkExists("loopback_advertised_subnet"); ok {
		if setArgNil {
			obj["loopback-advertised-subnet"] = nil
		} else {
			t, err := expandSystemFabricVpnLoopbackAdvertisedSubnet(d, v, "loopback_advertised_subnet", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["loopback-advertised-subnet"] = t
			}
		}
	} else if d.HasChange("loopback_advertised_subnet") {
		obj["loopback-advertised-subnet"] = nil
	}

	if v, ok := d.GetOk("psksecret"); ok {
		if setArgNil {
			obj["psksecret"] = nil
		} else {
			t, err := expandSystemFabricVpnPsksecret(d, v, "psksecret", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["psksecret"] = t
			}
		}
	}

	if v, ok := d.GetOk("bgp_as_string"); ok {
		if setArgNil {
			obj["bgp-as"] = nil
		} else {
			new_version_map := map[string][]string{
				">=": []string{"7.6.1"},
			}
			if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
				if _, ok := d.GetOk("bgp_as"); !ok && !d.HasChange("bgp_as") {
					err := fmt.Errorf("Argument 'bgp_as_string' %s.", err)
					return nil, err
				}
			} else {
				t, err := expandSystemFabricVpnBgpAsString(d, v, "bgp_as_string", sv)
				if err != nil {
					return &obj, err
				} else if t != nil {
					obj["bgp-as"] = t
				}
			}
		}
	} else if d.HasChange("bgp_as_string") {
		obj["bgp-as"] = nil
	}

	if v, ok := d.GetOkExists("bgp_as"); ok {
		if setArgNil {
			obj["bgp-as"] = nil
		} else {
			new_version_map := map[string][]string{
				"=": []string{"7.2.4", "7.2.6", "7.2.7", "7.2.8", "7.2.9", "7.2.10", "7.2.11", "7.2.12", "7.4.0", "7.4.1", "7.4.2", "7.4.3", "7.4.4", "7.4.5", "7.4.6", "7.4.7", "7.4.8", "7.6.0"},
			}
			if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
				if _, ok := d.GetOk("bgp_as_string"); !ok && !d.HasChange("bgp_as_string") {
					err := fmt.Errorf("Argument 'bgp_as' %s.", err)
					return nil, err
				}
			} else {
				t, err := expandSystemFabricVpnBgpAs(d, v, "bgp_as", sv)
				if err != nil {
					return &obj, err
				} else if t != nil {
					obj["bgp-as"] = t
				}
			}
		}
	} else if d.HasChange("bgp_as") {
		obj["bgp-as"] = nil
	}

	if v, ok := d.GetOk("sdwan_zone"); ok {
		if setArgNil {
			obj["sdwan-zone"] = nil
		} else {
			t, err := expandSystemFabricVpnSdwanZone(d, v, "sdwan_zone", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sdwan-zone"] = t
			}
		}
	} else if d.HasChange("sdwan_zone") {
		obj["sdwan-zone"] = nil
	}

	if v, ok := d.GetOk("health_checks"); ok {
		if setArgNil {
			obj["health-checks"] = nil
		} else {
			t, err := expandSystemFabricVpnHealthChecks(d, v, "health_checks", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["health-checks"] = t
			}
		}
	} else if d.HasChange("health_checks") {
		obj["health-checks"] = nil
	}

	return &obj, nil
}
