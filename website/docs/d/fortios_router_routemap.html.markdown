---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_routemap"
description: |-
  Get information on an fortios router routemap.
---

# Data Source: fortios_router_routemap
Use this data source to get information on an fortios router routemap

## Argument Reference

* `name` - (Required) Specify the name of the desired router routemap.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comments` - Optional comments.
* `rule` - Rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - Rule ID.
* `action` - Action.
* `match_as_path` - Match BGP AS path list.
* `match_community` - Match BGP community list.
* `match_extcommunity` - Match BGP extended community list.
* `match_community_exact` - Enable/disable exact matching of communities.
* `match_extcommunity_exact` - Enable/disable exact matching of extended communities.
* `match_origin` - Match BGP origin code.
* `match_interface` - Match interface configuration.
* `match_ip_address` - Match IP address permitted by access-list or prefix-list.
* `match_ip6_address` - Match IPv6 address permitted by access-list6 or prefix-list6.
* `match_ip_nexthop` - Match next hop IP address passed by access-list or prefix-list.
* `match_ip6_nexthop` - Match next hop IPv6 address passed by access-list6 or prefix-list6.
* `match_metric` - Match metric for redistribute routes.
* `match_route_type` - Match route type.
* `match_tag` - Match tag.
* `match_vrf` - Match VRF ID.
* `match_suppress` - Enable/disable matching of suppressed original neighbor.
* `set_aggregator_as` - BGP aggregator AS.
* `set_aggregator_ip` - BGP aggregator IP.
* `set_aspath_action` - Specify preferred action of set-aspath.
* `set_aspath` - Prepend BGP AS path attribute. The structure of `set_aspath` block is documented below.
* `set_atomic_aggregate` - Enable/disable BGP atomic aggregate attribute.
* `set_community_delete` - Delete communities matching community list.
* `set_community` - BGP community attribute. The structure of `set_community` block is documented below.
* `set_community_additive` - Enable/disable adding set-community to existing community.
* `set_dampening_reachability_half_life` - Reachability half-life time for the penalty (1 - 45 min, 0 = unset).
* `set_dampening_reuse` - Value to start reusing a route (1 - 20000, 0 = unset).
* `set_dampening_suppress` - Value to start suppressing a route (1 - 20000, 0 = unset).
* `set_dampening_max_suppress` - Maximum duration to suppress a route (1 - 255 min, 0 = unset).
* `set_dampening_unreachability_half_life` - Unreachability Half-life time for the penalty (1 - 45 min, 0 = unset)
* `set_extcommunity_rt` - Route Target extended community. The structure of `set_extcommunity_rt` block is documented below.
* `set_extcommunity_soo` - Site-of-Origin extended community. The structure of `set_extcommunity_soo` block is documented below.
* `set_ip_nexthop` - IP address of next hop.
* `set_ip_prefsrc` - IP address of preferred source.
* `set_vpnv4_nexthop` - IP address of VPNv4 next-hop.
* `set_ip6_nexthop` - IPv6 global address of next hop.
* `set_ip6_nexthop_local` - IPv6 local address of next hop.
* `set_vpnv6_nexthop` - IPv6 global address of VPNv6 next-hop.
* `set_vpnv6_nexthop_local` - IPv6 link-local address of VPNv6 next-hop.
* `set_local_preference` - BGP local preference path attribute.
* `set_metric` - Metric value.
* `set_metric_type` - Metric type.
* `set_originator_id` - BGP originator ID attribute.
* `set_origin` - BGP origin code.
* `set_tag` - Tag value.
* `set_weight` - BGP weight for routing table.
* `set_flags` - BGP flags value (0 - 65535)
* `match_flags` - BGP flag value to match (0 - 65535)
* `set_route_tag` - Route tag for routing table.
* `set_priority` - Priority for routing table.

The `set_aspath` block contains:

* `as` - AS number (0 - 42949672). NOTE: Use quotes for repeating numbers, e.g.: "1 1 2"


The `set_community` block contains:

* `community` - Attribute: AA|AA:NN|internet|local-AS|no-advertise|no-export.

The `set_extcommunity_rt` block contains:

* `community` - AA:NN.

The `set_extcommunity_soo` block contains:

* `community` - AA:NN

