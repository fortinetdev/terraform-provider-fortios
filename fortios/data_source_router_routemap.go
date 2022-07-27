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

func dataSourceRouterRouteMap() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterRouteMapRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_as_path": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_community_exact": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_origin": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_ip_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_ip6_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"match_route_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"match_vrf": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_aggregator_as": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_aggregator_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_aspath_action": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_aspath": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"as": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"set_atomic_aggregate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_community_delete": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_community": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"set_community_additive": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_dampening_reachability_half_life": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_dampening_reuse": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_dampening_suppress": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_dampening_max_suppress": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_dampening_unreachability_half_life": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_extcommunity_rt": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"set_extcommunity_soo": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"set_ip_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_ip6_nexthop": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_ip6_nexthop_local": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_local_preference": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_originator_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_origin": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"set_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_weight": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_flags": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"match_flags": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_route_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"set_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterRouteMapRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterRouteMap: type error")
	}

	o, err := c.ReadRouterRouteMap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterRouteMap: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterRouteMap(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterRouteMap from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterRouteMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenRouterRouteMapRuleId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = dataSourceFlattenRouterRouteMapRuleAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_as_path"
		if _, ok := i["match-as-path"]; ok {
			tmp["match_as_path"] = dataSourceFlattenRouterRouteMapRuleMatchAsPath(i["match-as-path"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community"
		if _, ok := i["match-community"]; ok {
			tmp["match_community"] = dataSourceFlattenRouterRouteMapRuleMatchCommunity(i["match-community"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_community_exact"
		if _, ok := i["match-community-exact"]; ok {
			tmp["match_community_exact"] = dataSourceFlattenRouterRouteMapRuleMatchCommunityExact(i["match-community-exact"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_origin"
		if _, ok := i["match-origin"]; ok {
			tmp["match_origin"] = dataSourceFlattenRouterRouteMapRuleMatchOrigin(i["match-origin"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_interface"
		if _, ok := i["match-interface"]; ok {
			tmp["match_interface"] = dataSourceFlattenRouterRouteMapRuleMatchInterface(i["match-interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_address"
		if _, ok := i["match-ip-address"]; ok {
			tmp["match_ip_address"] = dataSourceFlattenRouterRouteMapRuleMatchIpAddress(i["match-ip-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_address"
		if _, ok := i["match-ip6-address"]; ok {
			tmp["match_ip6_address"] = dataSourceFlattenRouterRouteMapRuleMatchIp6Address(i["match-ip6-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip_nexthop"
		if _, ok := i["match-ip-nexthop"]; ok {
			tmp["match_ip_nexthop"] = dataSourceFlattenRouterRouteMapRuleMatchIpNexthop(i["match-ip-nexthop"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_ip6_nexthop"
		if _, ok := i["match-ip6-nexthop"]; ok {
			tmp["match_ip6_nexthop"] = dataSourceFlattenRouterRouteMapRuleMatchIp6Nexthop(i["match-ip6-nexthop"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_metric"
		if _, ok := i["match-metric"]; ok {
			tmp["match_metric"] = dataSourceFlattenRouterRouteMapRuleMatchMetric(i["match-metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_route_type"
		if _, ok := i["match-route-type"]; ok {
			tmp["match_route_type"] = dataSourceFlattenRouterRouteMapRuleMatchRouteType(i["match-route-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_tag"
		if _, ok := i["match-tag"]; ok {
			tmp["match_tag"] = dataSourceFlattenRouterRouteMapRuleMatchTag(i["match-tag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_vrf"
		if _, ok := i["match-vrf"]; ok {
			tmp["match_vrf"] = dataSourceFlattenRouterRouteMapRuleMatchVrf(i["match-vrf"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_as"
		if _, ok := i["set-aggregator-as"]; ok {
			tmp["set_aggregator_as"] = dataSourceFlattenRouterRouteMapRuleSetAggregatorAs(i["set-aggregator-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aggregator_ip"
		if _, ok := i["set-aggregator-ip"]; ok {
			tmp["set_aggregator_ip"] = dataSourceFlattenRouterRouteMapRuleSetAggregatorIp(i["set-aggregator-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath_action"
		if _, ok := i["set-aspath-action"]; ok {
			tmp["set_aspath_action"] = dataSourceFlattenRouterRouteMapRuleSetAspathAction(i["set-aspath-action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_aspath"
		if _, ok := i["set-aspath"]; ok {
			tmp["set_aspath"] = dataSourceFlattenRouterRouteMapRuleSetAspath(i["set-aspath"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_atomic_aggregate"
		if _, ok := i["set-atomic-aggregate"]; ok {
			tmp["set_atomic_aggregate"] = dataSourceFlattenRouterRouteMapRuleSetAtomicAggregate(i["set-atomic-aggregate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_delete"
		if _, ok := i["set-community-delete"]; ok {
			tmp["set_community_delete"] = dataSourceFlattenRouterRouteMapRuleSetCommunityDelete(i["set-community-delete"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community"
		if _, ok := i["set-community"]; ok {
			tmp["set_community"] = dataSourceFlattenRouterRouteMapRuleSetCommunity(i["set-community"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_community_additive"
		if _, ok := i["set-community-additive"]; ok {
			tmp["set_community_additive"] = dataSourceFlattenRouterRouteMapRuleSetCommunityAdditive(i["set-community-additive"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reachability_half_life"
		if _, ok := i["set-dampening-reachability-half-life"]; ok {
			tmp["set_dampening_reachability_half_life"] = dataSourceFlattenRouterRouteMapRuleSetDampeningReachabilityHalfLife(i["set-dampening-reachability-half-life"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_reuse"
		if _, ok := i["set-dampening-reuse"]; ok {
			tmp["set_dampening_reuse"] = dataSourceFlattenRouterRouteMapRuleSetDampeningReuse(i["set-dampening-reuse"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_suppress"
		if _, ok := i["set-dampening-suppress"]; ok {
			tmp["set_dampening_suppress"] = dataSourceFlattenRouterRouteMapRuleSetDampeningSuppress(i["set-dampening-suppress"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_max_suppress"
		if _, ok := i["set-dampening-max-suppress"]; ok {
			tmp["set_dampening_max_suppress"] = dataSourceFlattenRouterRouteMapRuleSetDampeningMaxSuppress(i["set-dampening-max-suppress"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_dampening_unreachability_half_life"
		if _, ok := i["set-dampening-unreachability-half-life"]; ok {
			tmp["set_dampening_unreachability_half_life"] = dataSourceFlattenRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(i["set-dampening-unreachability-half-life"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_rt"
		if _, ok := i["set-extcommunity-rt"]; ok {
			tmp["set_extcommunity_rt"] = dataSourceFlattenRouterRouteMapRuleSetExtcommunityRt(i["set-extcommunity-rt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_extcommunity_soo"
		if _, ok := i["set-extcommunity-soo"]; ok {
			tmp["set_extcommunity_soo"] = dataSourceFlattenRouterRouteMapRuleSetExtcommunitySoo(i["set-extcommunity-soo"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip_nexthop"
		if _, ok := i["set-ip-nexthop"]; ok {
			tmp["set_ip_nexthop"] = dataSourceFlattenRouterRouteMapRuleSetIpNexthop(i["set-ip-nexthop"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop"
		if _, ok := i["set-ip6-nexthop"]; ok {
			tmp["set_ip6_nexthop"] = dataSourceFlattenRouterRouteMapRuleSetIp6Nexthop(i["set-ip6-nexthop"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_ip6_nexthop_local"
		if _, ok := i["set-ip6-nexthop-local"]; ok {
			tmp["set_ip6_nexthop_local"] = dataSourceFlattenRouterRouteMapRuleSetIp6NexthopLocal(i["set-ip6-nexthop-local"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_local_preference"
		if _, ok := i["set-local-preference"]; ok {
			tmp["set_local_preference"] = dataSourceFlattenRouterRouteMapRuleSetLocalPreference(i["set-local-preference"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric"
		if _, ok := i["set-metric"]; ok {
			tmp["set_metric"] = dataSourceFlattenRouterRouteMapRuleSetMetric(i["set-metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_metric_type"
		if _, ok := i["set-metric-type"]; ok {
			tmp["set_metric_type"] = dataSourceFlattenRouterRouteMapRuleSetMetricType(i["set-metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_originator_id"
		if _, ok := i["set-originator-id"]; ok {
			tmp["set_originator_id"] = dataSourceFlattenRouterRouteMapRuleSetOriginatorId(i["set-originator-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_origin"
		if _, ok := i["set-origin"]; ok {
			tmp["set_origin"] = dataSourceFlattenRouterRouteMapRuleSetOrigin(i["set-origin"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_tag"
		if _, ok := i["set-tag"]; ok {
			tmp["set_tag"] = dataSourceFlattenRouterRouteMapRuleSetTag(i["set-tag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_weight"
		if _, ok := i["set-weight"]; ok {
			tmp["set_weight"] = dataSourceFlattenRouterRouteMapRuleSetWeight(i["set-weight"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_flags"
		if _, ok := i["set-flags"]; ok {
			tmp["set_flags"] = dataSourceFlattenRouterRouteMapRuleSetFlags(i["set-flags"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_flags"
		if _, ok := i["match-flags"]; ok {
			tmp["match_flags"] = dataSourceFlattenRouterRouteMapRuleMatchFlags(i["match-flags"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_route_tag"
		if _, ok := i["set-route-tag"]; ok {
			tmp["set_route_tag"] = dataSourceFlattenRouterRouteMapRuleSetRouteTag(i["set-route-tag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "set_priority"
		if _, ok := i["set-priority"]; ok {
			tmp["set_priority"] = dataSourceFlattenRouterRouteMapRuleSetPriority(i["set-priority"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRouteMapRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchAsPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchCommunityExact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchOrigin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchIpNexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchIp6Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchRouteType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetAggregatorAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetAggregatorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetAspathAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetAspath(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
		if _, ok := i["as"]; ok {
			tmp["as"] = dataSourceFlattenRouterRouteMapRuleSetAspathAs(i["as"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRouteMapRuleSetAspathAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetAtomicAggregate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetCommunityDelete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetCommunity(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
		if _, ok := i["community"]; ok {
			tmp["community"] = dataSourceFlattenRouterRouteMapRuleSetCommunityCommunity(i["community"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRouteMapRuleSetCommunityCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetCommunityAdditive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetDampeningReachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetDampeningReuse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetDampeningSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetDampeningMaxSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetDampeningUnreachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetExtcommunityRt(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
		if _, ok := i["community"]; ok {
			tmp["community"] = dataSourceFlattenRouterRouteMapRuleSetExtcommunityRtCommunity(i["community"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRouteMapRuleSetExtcommunityRtCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetExtcommunitySoo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
		if _, ok := i["community"]; ok {
			tmp["community"] = dataSourceFlattenRouterRouteMapRuleSetExtcommunitySooCommunity(i["community"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRouteMapRuleSetExtcommunitySooCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetIpNexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetIp6Nexthop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetIp6NexthopLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetLocalPreference(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetOriginatorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetOrigin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleMatchFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRouteMapRuleSetPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterRouteMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenRouterRouteMapName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenRouterRouteMapComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("rule", dataSourceFlattenRouterRouteMapRule(o["rule"], d, "rule")); err != nil {
		if !fortiAPIPatch(o["rule"]) {
			return fmt.Errorf("Error reading rule: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterRouteMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
