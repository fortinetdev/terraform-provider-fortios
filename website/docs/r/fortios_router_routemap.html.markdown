---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_routemap"
description: |-
  Configure route maps.
---

# fortios_router_routemap
Configure route maps.

## Example Usage

```hcl
resource "fortios_router_routemap" "trname" {
  name = "1"

  rule {
    action                                 = "deny"
    match_community_exact                  = "disable"
    match_flags                            = 0
    match_metric                           = 0
    match_origin                           = "none"
    match_route_type                       = "No type specified"
    match_tag                              = 0
    set_aggregator_as                      = 0
    set_aggregator_ip                      = "0.0.0.0"
    set_aspath_action                      = "prepend"
    set_atomic_aggregate                   = "disable"
    set_community_additive                 = "disable"
    set_dampening_max_suppress             = 0
    set_dampening_reachability_half_life   = 0
    set_dampening_reuse                    = 0
    set_dampening_suppress                 = 0
    set_dampening_unreachability_half_life = 0
    set_flags                              = 128
    set_ip6_nexthop                        = "::"
    set_ip6_nexthop_local                  = "::"
    set_ip_nexthop                         = "0.0.0.0"
    set_local_preference                   = 0
    set_metric                             = 0
    set_metric_type                        = "No type specified"
    set_origin                             = "none"
    set_originator_id                      = "0.0.0.0"
    set_route_tag                          = 0
    set_tag                                = 0
    set_weight                             = 21
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Optional comments.
* `rule` - Rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `rule` block supports:

* `id` - Rule ID.
* `action` - Action. Valid values: `permit`, `deny`.
* `match_as_path` - Match BGP AS path list.
* `match_community` - Match BGP community list.
* `match_community_exact` - Enable/disable exact matching of communities. Valid values: `enable`, `disable`.
* `match_origin` - Match BGP origin code. Valid values: `none`, `egp`, `igp`, `incomplete`.
* `match_interface` - Match interface configuration.
* `match_ip_address` - Match IP address permitted by access-list or prefix-list.
* `match_ip6_address` - Match IPv6 address permitted by access-list6 or prefix-list6.
* `match_ip_nexthop` - Match next hop IP address passed by access-list or prefix-list.
* `match_ip6_nexthop` - Match next hop IPv6 address passed by access-list6 or prefix-list6.
* `match_metric` - Match metric for redistribute routes.
* `match_route_type` - Match route type. Valid values: `1`, `2`, `none`.
* `match_tag` - Match tag.
* `match_vrf` - Match VRF ID.
* `set_aggregator_as` - BGP aggregator AS.
* `set_aggregator_ip` - BGP aggregator IP.
* `set_aspath_action` - Specify preferred action of set-aspath. Valid values: `prepend`, `replace`.
* `set_aspath` - Prepend BGP AS path attribute. The structure of `set_aspath` block is documented below.
* `set_atomic_aggregate` - Enable/disable BGP atomic aggregate attribute. Valid values: `enable`, `disable`.
* `set_community_delete` - Delete communities matching community list.
* `set_community` - BGP community attribute. The structure of `set_community` block is documented below.
* `set_community_additive` - Enable/disable adding set-community to existing community. Valid values: `enable`, `disable`.
* `set_dampening_reachability_half_life` - Reachability half-life time for the penalty (1 - 45 min, 0 = unset).
* `set_dampening_reuse` - Value to start reusing a route (1 - 20000, 0 = unset).
* `set_dampening_suppress` - Value to start suppressing a route (1 - 20000, 0 = unset).
* `set_dampening_max_suppress` - Maximum duration to suppress a route (1 - 255 min, 0 = unset).
* `set_dampening_unreachability_half_life` - Unreachability Half-life time for the penalty (1 - 45 min, 0 = unset)
* `set_extcommunity_rt` - Route Target extended community. The structure of `set_extcommunity_rt` block is documented below.
* `set_extcommunity_soo` - Site-of-Origin extended community. The structure of `set_extcommunity_soo` block is documented below.
* `set_ip_nexthop` - IP address of next hop.
* `set_ip6_nexthop` - IPv6 global address of next hop.
* `set_ip6_nexthop_local` - IPv6 local address of next hop.
* `set_local_preference` - BGP local preference path attribute.
* `set_metric` - Metric value.
* `set_metric_type` - Metric type. Valid values: `1`, `2`, `none`.
* `set_originator_id` - BGP originator ID attribute.
* `set_origin` - BGP origin code. Valid values: `none`, `egp`, `igp`, `incomplete`.
* `set_tag` - Tag value.
* `set_weight` - BGP weight for routing table.
* `set_flags` - BGP flags value (0 - 65535)
* `match_flags` - BGP flag value to match (0 - 65535)
* `set_route_tag` - Route tag for routing table.

The `set_aspath` block supports:

* `as` - AS number (0 - 42949672). NOTE: Use quotes for repeating numbers, e.g.: "1 1 2"


The `set_community` block supports:

* `community` - Attribute: AA|AA:NN|internet|local-AS|no-advertise|no-export.

The `set_extcommunity_rt` block supports:

* `community` - AA:NN.

The `set_extcommunity_soo` block supports:

* `community` - AA:NN


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router RouteMap can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_routemap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
