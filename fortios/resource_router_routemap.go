// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure route maps.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterRouteMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRouteMapCreate,
		Read:   resourceRouterRouteMapRead,
		Update: resourceRouterRouteMapUpdate,
		Delete: resourceRouterRouteMapDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_as_path": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"match_community": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"match_extcommunity": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"match_community_exact": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_extcommunity_exact": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_origin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"match_ip_address": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"match_ip6_address": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"match_ip_nexthop": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"match_ip6_nexthop": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"match_metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"match_route_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"match_vrf": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 251),
							Optional:     true,
						},
						"set_aggregator_as": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_aggregator_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_aspath_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_aspath": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"as": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"set_atomic_aggregate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_community_delete": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"set_community": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"set_community_additive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_dampening_reachability_half_life": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 45),
							Optional:     true,
						},
						"set_dampening_reuse": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 20000),
							Optional:     true,
						},
						"set_dampening_suppress": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 20000),
							Optional:     true,
						},
						"set_dampening_max_suppress": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"set_dampening_unreachability_half_life": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 45),
							Optional:     true,
						},
						"set_extcommunity_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"set_extcommunity_soo": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"set_ip_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_ip_prefsrc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_vpnv4_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_ip6_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_ip6_nexthop_local": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_vpnv6_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_vpnv6_nexthop_local": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_local_preference": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_originator_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"set_origin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"set_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_flags": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"match_flags": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"set_route_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"set_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
						},
					},
				},
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

func resourceRouterRouteMapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterRouteMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterRouteMap resource while getting object: %v", err)
	}

	o, err := c.CreateRouterRouteMap(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterRouteMap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterRouteMap")
	}

	return resourceRouterRouteMapRead(d, m)
}

func resourceRouterRouteMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterRouteMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterRouteMap resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterRouteMap(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterRouteMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterRouteMap")
	}

	return resourceRouterRouteMapRead(d, m)
}

func resourceRouterRouteMapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterRouteMap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterRouteMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRouteMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterRouteMap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterRouteMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRouteMap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterRouteMap resource from API: %v", err)
	}
	return nil
}

func flattenRouterRouteMapName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenRouterRouteMapRuleId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenRouterRouteMapRuleAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_as_path"
		if cur_v, ok := i["match-as-path"]; ok {
			tmp["match_as_path"] = flattenRouterRouteMapRuleMatchAsPath(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community"
		if cur_v, ok := i["match-community"]; ok {
			tmp["match_community"] = flattenRouterRouteMapRuleMatchCommunity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity"
		if cur_v, ok := i["match-extcommunity"]; ok {
			tmp["match_extcommunity"] = flattenRouterRouteMapRuleMatchExtcommunity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community_exact"
		if cur_v, ok := i["match-community-exact"]; ok {
			tmp["match_community_exact"] = flattenRouterRouteMapRuleMatchCommunityExact(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity_exact"
		if cur_v, ok := i["match-extcommunity-exact"]; ok {
			tmp["match_extcommunity_exact"] = flattenRouterRouteMapRuleMatchExtcommunityExact(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_origin"
		if cur_v, ok := i["match-origin"]; ok {
			tmp["match_origin"] = flattenRouterRouteMapRuleMatchOrigin(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_interface"
		if cur_v, ok := i["match-interface"]; ok {
			tmp["match_interface"] = flattenRouterRouteMapRuleMatchInterface(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_address"
		if cur_v, ok := i["match-ip-address"]; ok {
			tmp["match_ip_address"] = flattenRouterRouteMapRuleMatchIpAddress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_address"
		if cur_v, ok := i["match-ip6-address"]; ok {
			tmp["match_ip6_address"] = flattenRouterRouteMapRuleMatchIp6Address(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_nexthop"
		if cur_v, ok := i["match-ip-nexthop"]; ok {
			tmp["match_ip_nexthop"] = flattenRouterRouteMapRuleMatchIpNexthop(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_nexthop"
		if cur_v, ok := i["match-ip6-nexthop"]; ok {
			tmp["match_ip6_nexthop"] = flattenRouterRouteMapRuleMatchIp6Nexthop(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_metric"
		if cur_v, ok := i["match-metric"]; ok {
			tmp["match_metric"] = flattenRouterRouteMapRuleMatchMetric(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_route_type"
		if cur_v, ok := i["match-route-type"]; ok {
			tmp["match_route_type"] = flattenRouterRouteMapRuleMatchRouteType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_tag"
		if cur_v, ok := i["match-tag"]; ok {
			tmp["match_tag"] = flattenRouterRouteMapRuleMatchTag(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_vrf"
		if cur_v, ok := i["match-vrf"]; ok {
			tmp["match_vrf"] = flattenRouterRouteMapRuleMatchVrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_as"
		if cur_v, ok := i["set-aggregator-as"]; ok {
			tmp["set_aggregator_as"] = flattenRouterRouteMapRuleSetAggregatorAs(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_ip"
		if cur_v, ok := i["set-aggregator-ip"]; ok {
			tmp["set_aggregator_ip"] = flattenRouterRouteMapRuleSetAggregatorIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath_action"
		if cur_v, ok := i["set-aspath-action"]; ok {
			tmp["set_aspath_action"] = flattenRouterRouteMapRuleSetAspathAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath"
		if cur_v, ok := i["set-aspath"]; ok {
			tmp["set_aspath"] = flattenRouterRouteMapRuleSetAspath(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_atomic_aggregate"
		if cur_v, ok := i["set-atomic-aggregate"]; ok {
			tmp["set_atomic_aggregate"] = flattenRouterRouteMapRuleSetAtomicAggregate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_delete"
		if cur_v, ok := i["set-community-delete"]; ok {
			tmp["set_community_delete"] = flattenRouterRouteMapRuleSetCommunityDelete(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community"
		if cur_v, ok := i["set-community"]; ok {
			tmp["set_community"] = flattenRouterRouteMapRuleSetCommunity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_additive"
		if cur_v, ok := i["set-community-additive"]; ok {
			tmp["set_community_additive"] = flattenRouterRouteMapRuleSetCommunityAdditive(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reachability_half_life"
		if cur_v, ok := i["set-dampening-reachability-half-life"]; ok {
			tmp["set_dampening_reachability_half_life"] = flattenRouterRouteMapRuleSetDampeningReachabilityHalfLife(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reuse"
		if cur_v, ok := i["set-dampening-reuse"]; ok {
			tmp["set_dampening_reuse"] = flattenRouterRouteMapRuleSetDampeningReuse(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_suppress"
		if cur_v, ok := i["set-dampening-suppress"]; ok {
			tmp["set_dampening_suppress"] = flattenRouterRouteMapRuleSetDampeningSuppress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_max_suppress"
		if cur_v, ok := i["set-dampening-max-suppress"]; ok {
			tmp["set_dampening_max_suppress"] = flattenRouterRouteMapRuleSetDampeningMaxSuppress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_unreachability_half_life"
		if cur_v, ok := i["set-dampening-unreachability-half-life"]; ok {
			tmp["set_dampening_unreachability_half_life"] = flattenRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_rt"
		if cur_v, ok := i["set-extcommunity-rt"]; ok {
			tmp["set_extcommunity_rt"] = flattenRouterRouteMapRuleSetExtcommunityRt(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_soo"
		if cur_v, ok := i["set-extcommunity-soo"]; ok {
			tmp["set_extcommunity_soo"] = flattenRouterRouteMapRuleSetExtcommunitySoo(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_nexthop"
		if cur_v, ok := i["set-ip-nexthop"]; ok {
			tmp["set_ip_nexthop"] = flattenRouterRouteMapRuleSetIpNexthop(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_prefsrc"
		if cur_v, ok := i["set-ip-prefsrc"]; ok {
			tmp["set_ip_prefsrc"] = flattenRouterRouteMapRuleSetIpPrefsrc(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv4_nexthop"
		if cur_v, ok := i["set-vpnv4-nexthop"]; ok {
			tmp["set_vpnv4_nexthop"] = flattenRouterRouteMapRuleSetVpnv4Nexthop(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop"
		if cur_v, ok := i["set-ip6-nexthop"]; ok {
			tmp["set_ip6_nexthop"] = flattenRouterRouteMapRuleSetIp6Nexthop(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop_local"
		if cur_v, ok := i["set-ip6-nexthop-local"]; ok {
			tmp["set_ip6_nexthop_local"] = flattenRouterRouteMapRuleSetIp6NexthopLocal(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop"
		if cur_v, ok := i["set-vpnv6-nexthop"]; ok {
			tmp["set_vpnv6_nexthop"] = flattenRouterRouteMapRuleSetVpnv6Nexthop(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop_local"
		if cur_v, ok := i["set-vpnv6-nexthop-local"]; ok {
			tmp["set_vpnv6_nexthop_local"] = flattenRouterRouteMapRuleSetVpnv6NexthopLocal(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_local_preference"
		if cur_v, ok := i["set-local-preference"]; ok {
			tmp["set_local_preference"] = flattenRouterRouteMapRuleSetLocalPreference(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric"
		if cur_v, ok := i["set-metric"]; ok {
			tmp["set_metric"] = flattenRouterRouteMapRuleSetMetric(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric_type"
		if cur_v, ok := i["set-metric-type"]; ok {
			tmp["set_metric_type"] = flattenRouterRouteMapRuleSetMetricType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_originator_id"
		if cur_v, ok := i["set-originator-id"]; ok {
			tmp["set_originator_id"] = flattenRouterRouteMapRuleSetOriginatorId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_origin"
		if cur_v, ok := i["set-origin"]; ok {
			tmp["set_origin"] = flattenRouterRouteMapRuleSetOrigin(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_tag"
		if cur_v, ok := i["set-tag"]; ok {
			tmp["set_tag"] = flattenRouterRouteMapRuleSetTag(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_weight"
		if cur_v, ok := i["set-weight"]; ok {
			tmp["set_weight"] = flattenRouterRouteMapRuleSetWeight(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_flags"
		if cur_v, ok := i["set-flags"]; ok {
			tmp["set_flags"] = flattenRouterRouteMapRuleSetFlags(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_flags"
		if cur_v, ok := i["match-flags"]; ok {
			tmp["match_flags"] = flattenRouterRouteMapRuleMatchFlags(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_route_tag"
		if cur_v, ok := i["set-route-tag"]; ok {
			tmp["set_route_tag"] = flattenRouterRouteMapRuleSetRouteTag(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_priority"
		if cur_v, ok := i["set-priority"]; ok {
			tmp["set_priority"] = flattenRouterRouteMapRuleSetPriority(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterRouteMapRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchAsPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchExtcommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchCommunityExact(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchExtcommunityExact(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchOrigin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchIp6Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchIpNexthop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchIp6Nexthop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleMatchRouteType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleMatchTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleMatchVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetAggregatorAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetAggregatorIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAspathAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAspath(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as"
		if cur_v, ok := i["as"]; ok {
			tmp["as"] = flattenRouterRouteMapRuleSetAspathAs(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "as", d)
	return result
}

func flattenRouterRouteMapRuleSetAspathAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetAtomicAggregate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetCommunityDelete(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "community"
		if cur_v, ok := i["community"]; ok {
			tmp["community"] = flattenRouterRouteMapRuleSetCommunityCommunity(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "community", d)
	return result
}

func flattenRouterRouteMapRuleSetCommunityCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetCommunityAdditive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetDampeningReachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetDampeningReuse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetDampeningSuppress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetDampeningMaxSuppress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetExtcommunityRt(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "community"
		if cur_v, ok := i["community"]; ok {
			tmp["community"] = flattenRouterRouteMapRuleSetExtcommunityRtCommunity(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "community", d)
	return result
}

func flattenRouterRouteMapRuleSetExtcommunityRtCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetExtcommunitySoo(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "community"
		if cur_v, ok := i["community"]; ok {
			tmp["community"] = flattenRouterRouteMapRuleSetExtcommunitySooCommunity(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "community", d)
	return result
}

func flattenRouterRouteMapRuleSetExtcommunitySooCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIpNexthop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIpPrefsrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetVpnv4Nexthop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIp6Nexthop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetIp6NexthopLocal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetVpnv6Nexthop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetVpnv6NexthopLocal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetLocalPreference(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetOriginatorId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetOrigin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRouteMapRuleSetTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleMatchFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetRouteTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterRouteMapRuleSetPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectRouterRouteMap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenRouterRouteMapName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterRouteMapComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("rule", flattenRouterRouteMapRule(o["rule"], d, "rule", sv)); err != nil {
			if !fortiAPIPatch(o["rule"]) {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenRouterRouteMapRule(o["rule"], d, "rule", sv)); err != nil {
				if !fortiAPIPatch(o["rule"]) {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterRouteMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterRouteMapName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterRouteMapRuleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandRouterRouteMapRuleAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_as_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-as-path"], _ = expandRouterRouteMapRuleMatchAsPath(d, i["match_as_path"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-as-path"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-community"], _ = expandRouterRouteMapRuleMatchCommunity(d, i["match_community"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-community"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-extcommunity"], _ = expandRouterRouteMapRuleMatchExtcommunity(d, i["match_extcommunity"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-extcommunity"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community_exact"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-community-exact"], _ = expandRouterRouteMapRuleMatchCommunityExact(d, i["match_community_exact"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_extcommunity_exact"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-extcommunity-exact"], _ = expandRouterRouteMapRuleMatchExtcommunityExact(d, i["match_extcommunity_exact"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_origin"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-origin"], _ = expandRouterRouteMapRuleMatchOrigin(d, i["match_origin"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-interface"], _ = expandRouterRouteMapRuleMatchInterface(d, i["match_interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-ip-address"], _ = expandRouterRouteMapRuleMatchIpAddress(d, i["match_ip_address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-ip-address"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-ip6-address"], _ = expandRouterRouteMapRuleMatchIp6Address(d, i["match_ip6_address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-ip6-address"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_nexthop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-ip-nexthop"], _ = expandRouterRouteMapRuleMatchIpNexthop(d, i["match_ip_nexthop"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-ip-nexthop"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_nexthop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-ip6-nexthop"], _ = expandRouterRouteMapRuleMatchIp6Nexthop(d, i["match_ip6_nexthop"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-ip6-nexthop"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-metric"], _ = expandRouterRouteMapRuleMatchMetric(d, i["match_metric"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-metric"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_route_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-route-type"], _ = expandRouterRouteMapRuleMatchRouteType(d, i["match_route_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-tag"], _ = expandRouterRouteMapRuleMatchTag(d, i["match_tag"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-tag"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-vrf"], _ = expandRouterRouteMapRuleMatchVrf(d, i["match_vrf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-vrf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_as"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-aggregator-as"], _ = expandRouterRouteMapRuleSetAggregatorAs(d, i["set_aggregator_as"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-aggregator-as"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-aggregator-ip"], _ = expandRouterRouteMapRuleSetAggregatorIp(d, i["set_aggregator_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath_action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-aspath-action"], _ = expandRouterRouteMapRuleSetAspathAction(d, i["set_aspath_action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-aspath"], _ = expandRouterRouteMapRuleSetAspath(d, i["set_aspath"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-aspath"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_atomic_aggregate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-atomic-aggregate"], _ = expandRouterRouteMapRuleSetAtomicAggregate(d, i["set_atomic_aggregate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_delete"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-community-delete"], _ = expandRouterRouteMapRuleSetCommunityDelete(d, i["set_community_delete"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-community-delete"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-community"], _ = expandRouterRouteMapRuleSetCommunity(d, i["set_community"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-community"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_additive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-community-additive"], _ = expandRouterRouteMapRuleSetCommunityAdditive(d, i["set_community_additive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reachability_half_life"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-dampening-reachability-half-life"], _ = expandRouterRouteMapRuleSetDampeningReachabilityHalfLife(d, i["set_dampening_reachability_half_life"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-dampening-reachability-half-life"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reuse"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-dampening-reuse"], _ = expandRouterRouteMapRuleSetDampeningReuse(d, i["set_dampening_reuse"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-dampening-reuse"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_suppress"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-dampening-suppress"], _ = expandRouterRouteMapRuleSetDampeningSuppress(d, i["set_dampening_suppress"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-dampening-suppress"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_max_suppress"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-dampening-max-suppress"], _ = expandRouterRouteMapRuleSetDampeningMaxSuppress(d, i["set_dampening_max_suppress"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-dampening-max-suppress"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_unreachability_half_life"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-dampening-unreachability-half-life"], _ = expandRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(d, i["set_dampening_unreachability_half_life"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-dampening-unreachability-half-life"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_rt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-extcommunity-rt"], _ = expandRouterRouteMapRuleSetExtcommunityRt(d, i["set_extcommunity_rt"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-extcommunity-rt"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_soo"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-extcommunity-soo"], _ = expandRouterRouteMapRuleSetExtcommunitySoo(d, i["set_extcommunity_soo"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-extcommunity-soo"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_nexthop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-ip-nexthop"], _ = expandRouterRouteMapRuleSetIpNexthop(d, i["set_ip_nexthop"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-ip-nexthop"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_prefsrc"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-ip-prefsrc"], _ = expandRouterRouteMapRuleSetIpPrefsrc(d, i["set_ip_prefsrc"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-ip-prefsrc"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv4_nexthop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-vpnv4-nexthop"], _ = expandRouterRouteMapRuleSetVpnv4Nexthop(d, i["set_vpnv4_nexthop"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-vpnv4-nexthop"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-ip6-nexthop"], _ = expandRouterRouteMapRuleSetIp6Nexthop(d, i["set_ip6_nexthop"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-ip6-nexthop"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop_local"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-ip6-nexthop-local"], _ = expandRouterRouteMapRuleSetIp6NexthopLocal(d, i["set_ip6_nexthop_local"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-ip6-nexthop-local"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-vpnv6-nexthop"], _ = expandRouterRouteMapRuleSetVpnv6Nexthop(d, i["set_vpnv6_nexthop"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-vpnv6-nexthop"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_vpnv6_nexthop_local"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-vpnv6-nexthop-local"], _ = expandRouterRouteMapRuleSetVpnv6NexthopLocal(d, i["set_vpnv6_nexthop_local"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-vpnv6-nexthop-local"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_local_preference"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-local-preference"], _ = expandRouterRouteMapRuleSetLocalPreference(d, i["set_local_preference"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-local-preference"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-metric"], _ = expandRouterRouteMapRuleSetMetric(d, i["set_metric"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-metric"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-metric-type"], _ = expandRouterRouteMapRuleSetMetricType(d, i["set_metric_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_originator_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-originator-id"], _ = expandRouterRouteMapRuleSetOriginatorId(d, i["set_originator_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-originator-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_origin"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-origin"], _ = expandRouterRouteMapRuleSetOrigin(d, i["set_origin"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-tag"], _ = expandRouterRouteMapRuleSetTag(d, i["set_tag"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-tag"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-weight"], _ = expandRouterRouteMapRuleSetWeight(d, i["set_weight"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-weight"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_flags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-flags"], _ = expandRouterRouteMapRuleSetFlags(d, i["set_flags"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_flags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-flags"], _ = expandRouterRouteMapRuleMatchFlags(d, i["match_flags"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_route_tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-route-tag"], _ = expandRouterRouteMapRuleSetRouteTag(d, i["set_route_tag"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-route-tag"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["set-priority"], _ = expandRouterRouteMapRuleSetPriority(d, i["set_priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["set-priority"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRouteMapRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchAsPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchExtcommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchCommunityExact(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchExtcommunityExact(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchOrigin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchIp6Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchIpNexthop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchIp6Nexthop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchRouteType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAggregatorAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAggregatorIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAspathAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAspath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["as"], _ = expandRouterRouteMapRuleSetAspathAs(d, i["as"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRouteMapRuleSetAspathAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetAtomicAggregate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetCommunityDelete(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["community"], _ = expandRouterRouteMapRuleSetCommunityCommunity(d, i["community"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRouteMapRuleSetCommunityCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetCommunityAdditive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningReachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningReuse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningSuppress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningMaxSuppress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetExtcommunityRt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["community"], _ = expandRouterRouteMapRuleSetExtcommunityRtCommunity(d, i["community"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRouteMapRuleSetExtcommunityRtCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetExtcommunitySoo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["community"], _ = expandRouterRouteMapRuleSetExtcommunitySooCommunity(d, i["community"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRouteMapRuleSetExtcommunitySooCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIpNexthop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIpPrefsrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetVpnv4Nexthop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIp6Nexthop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetIp6NexthopLocal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetVpnv6Nexthop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetVpnv6NexthopLocal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetLocalPreference(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetOriginatorId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetOrigin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleMatchFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetRouteTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRouteMapRuleSetPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterRouteMap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandRouterRouteMapName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandRouterRouteMapComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	} else if d.HasChange("comments") {
		obj["comments"] = nil
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandRouterRouteMapRule(d, v, "rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
