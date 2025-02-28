---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp"
description: |-
  Configure BGP.
---

# fortios_router_bgp
Configure BGP.

~> The provider supports the definition of Neighbor in Router Bgp `fortios_router_bgp`, and also allows the definition of separate Neighbor resources `fortios_routerbgp_neighbor`, but do not use a `fortios_router_bgp` with in-line Neighbor in conjunction with any `fortios_routerbgp_neighbor` resources, otherwise conflicts and overwrite will occur.

~> The provider supports the definition of Network in Router Bgp `fortios_router_bgp`, and also allows the definition of separate Network resources `fortios_routerbgp_network`, but do not use a `fortios_router_bgp` with in-line Network in conjunction with any `fortios_routerbgp_network` resources, otherwise conflicts and overwrite will occur.

~> The provider supports the definition of Network6 in Router Bgp `fortios_router_bgp`, and also allows the definition of separate Network6 resources `fortios_routerbgp_network6`, but do not use a `fortios_router_bgp` with in-line Network6 in conjunction with any `fortios_routerbgp_network6` resources, otherwise conflicts and overwrite will occur.



## Example Usage

```hcl
resource "fortios_router_bgp" "trname" {
  additional_path_select             = 2
  additional_path_select6            = 2
  always_compare_med                 = "disable"
  as                                 = 0
  client_to_client_reflection        = "enable"
  cluster_id                         = "0.0.0.0"
  dampening                          = "disable"
  dampening_max_suppress_time        = 60
  dampening_reachability_half_life   = 15
  dampening_reuse                    = 750
  dampening_suppress                 = 2000
  dampening_unreachability_half_life = 15
  default_local_preference           = 100
  deterministic_med                  = "disable"
  distance_external                  = 20
  distance_internal                  = 200
  distance_local                     = 200
  graceful_restart_time              = 120
  graceful_stalepath_time            = 360
  graceful_update_delay              = 120
  holdtime_timer                     = 180
  ibgp_multipath                     = "disable"
  ignore_optional_capability         = "enable"
  keepalive_timer                    = 60
  log_neighbour_changes              = "enable"
  network_import_check               = "enable"
  scan_time                          = 60
  synchronization                    = "disable"

  redistribute {
    name   = "connected"
    status = "disable"
  }
  redistribute {
    name   = "rip"
    status = "disable"
  }
  redistribute {
    name   = "ospf"
    status = "disable"
  }
  redistribute {
    name   = "static"
    status = "disable"
  }
  redistribute {
    name   = "isis"
    status = "disable"
  }

  redistribute6 {
    name   = "connected"
    status = "disable"
  }
  redistribute6 {
    name   = "rip"
    status = "disable"
  }
  redistribute6 {
    name   = "ospf"
    status = "disable"
  }
  redistribute6 {
    name   = "static"
    status = "disable"
  }
  redistribute6 {
    name   = "isis"
    status = "disable"
  }
}
```

## Argument Reference

The following arguments are supported:

* `as_string` - Router AS number, asplain/asdot/asdot+ format, 0 to disable BGP. *Due to the data type change of API, for other versions of FortiOS, please check variable `as`.*
* `as` - (Required) Router AS number, valid from 1 to 4294967295, 0 to disable BGP. *Due to the data type change of API, for other versions of FortiOS, please check variable `as_string`.*
* `router_id` - Router ID.
* `keepalive_timer` - Frequency to send keep alive requests.
* `holdtime_timer` - Number of seconds to mark peer as dead.
* `always_compare_med` - Enable/disable always compare MED. Valid values: `enable`, `disable`.
* `bestpath_as_path_ignore` - Enable/disable ignore AS path. Valid values: `enable`, `disable`.
* `bestpath_cmp_confed_aspath` - Enable/disable compare federation AS path length. Valid values: `enable`, `disable`.
* `bestpath_cmp_routerid` - Enable/disable compare router ID for identical EBGP paths. Valid values: `enable`, `disable`.
* `bestpath_med_confed` - Enable/disable compare MED among confederation paths. Valid values: `enable`, `disable`.
* `bestpath_med_missing_as_worst` - Enable/disable treat missing MED as least preferred. Valid values: `enable`, `disable`.
* `client_to_client_reflection` - Enable/disable client-to-client route reflection. Valid values: `enable`, `disable`.
* `dampening` - Enable/disable route-flap dampening. Valid values: `enable`, `disable`.
* `deterministic_med` - Enable/disable enforce deterministic comparison of MED. Valid values: `enable`, `disable`.
* `ebgp_multipath` - Enable/disable EBGP multi-path. Valid values: `enable`, `disable`.
* `ibgp_multipath` - Enable/disable IBGP multi-path. Valid values: `enable`, `disable`.
* `enforce_first_as` - Enable/disable enforce first AS for EBGP routes. Valid values: `enable`, `disable`.
* `fast_external_failover` - Enable/disable reset peer BGP session if link goes down. Valid values: `enable`, `disable`.
* `log_neighbour_changes` - Enable logging of BGP neighbour's changes Valid values: `enable`, `disable`.
* `network_import_check` - Enable/disable ensure BGP network route exists in IGP. Valid values: `enable`, `disable`.
* `ignore_optional_capability` - Don't send unknown optional capability notification message Valid values: `enable`, `disable`.
* `additional_path` - Enable/disable selection of BGP IPv4 additional paths. Valid values: `enable`, `disable`.
* `additional_path6` - Enable/disable selection of BGP IPv6 additional paths. Valid values: `enable`, `disable`.
* `additional_path_vpnv4` - Enable/disable selection of BGP VPNv4 additional paths. Valid values: `enable`, `disable`.
* `additional_path_vpnv6` - Enable/disable selection of BGP VPNv6 additional paths. Valid values: `enable`, `disable`.
* `multipath_recursive_distance` - Enable/disable use of recursive distance to select multipath. Valid values: `enable`, `disable`.
* `recursive_next_hop` - Enable/disable recursive resolution of next-hop using BGP route. Valid values: `enable`, `disable`.
* `recursive_inherit_priority` - Enable/disable priority inheritance for recursive resolution. Valid values: `enable`, `disable`.
* `tag_resolve_mode` - Configure tag-match mode. Resolves BGP routes with other routes containing the same tag.
* `cluster_id` - Route reflector cluster ID.
* `confederation_identifier` - Confederation identifier.
* `confederation_peers` - Confederation peers. The structure of `confederation_peers` block is documented below.
* `dampening_route_map` - Criteria for dampening.
* `dampening_reachability_half_life` - Reachability half-life time for penalty (min).
* `dampening_reuse` - Threshold to reuse routes.
* `dampening_suppress` - Threshold to suppress routes.
* `dampening_max_suppress_time` - Maximum minutes a route can be suppressed.
* `dampening_unreachability_half_life` - Unreachability half-life time for penalty (min).
* `default_local_preference` - Default local preference.
* `scan_time` - Background scanner interval (sec), 0 to disable it.
* `distance_external` - Distance for routes external to the AS.
* `distance_internal` - Distance for routes internal to the AS.
* `distance_local` - Distance for routes local to the AS.
* `synchronization` - Enable/disable only advertise routes from iBGP if routes present in an IGP. Valid values: `enable`, `disable`.
* `graceful_restart` - Enable/disable BGP graceful restart capabilities. Valid values: `enable`, `disable`.
* `graceful_restart_time` - Time needed for neighbors to restart (sec).
* `graceful_stalepath_time` - Time to hold stale paths of restarting neighbor (sec).
* `graceful_update_delay` - Route advertisement/selection delay after restart (sec).
* `graceful_end_on_timer` - Enable/disable to exit graceful restart on timer only. Valid values: `enable`, `disable`.
* `additional_path_select` - Number of additional paths to be selected for each IPv4 NLRI.
* `additional_path_select6` - Number of additional paths to be selected for each IPv6 NLRI.
* `additional_path_select_vpnv4` - Number of additional paths to be selected for each VPNv4 NLRI.
* `additional_path_select_vpnv6` - Number of additional paths to be selected for each VPNv6 NLRI.
* `cross_family_conditional_adv` - Enable/disable cross address family conditional advertisement. Valid values: `enable`, `disable`.
* `aggregate_address` - BGP aggregate address table. The structure of `aggregate_address` block is documented below.
* `aggregate_address6` - BGP IPv6 aggregate address table. The structure of `aggregate_address6` block is documented below.
* `neighbor` - BGP neighbor table. The structure of `neighbor` block is documented below.
* `neighbor_group` - BGP neighbor group table. The structure of `neighbor_group` block is documented below.
* `neighbor_range` - BGP neighbor range table. The structure of `neighbor_range` block is documented below.
* `neighbor_range6` - BGP IPv6 neighbor range table. The structure of `neighbor_range6` block is documented below.
* `network` - BGP network table. The structure of `network` block is documented below.
* `network6` - BGP IPv6 network table. The structure of `network6` block is documented below.
* `redistribute` - BGP IPv4 redistribute table. The structure of `redistribute` block is documented below.
* `redistribute6` - BGP IPv6 redistribute table. The structure of `redistribute6` block is documented below.
* `admin_distance` - Administrative distance modifications. The structure of `admin_distance` block is documented below.
* `vrf` - BGP VRF leaking table. The structure of `vrf` block is documented below.
* `vrf6` - BGP IPv6 VRF leaking table. The structure of `vrf6` block is documented below.
* `vrf_leak` - BGP VRF leaking table. The structure of `vrf_leak` block is documented below.
* `vrf_leak6` - BGP IPv6 VRF leaking table. The structure of `vrf_leak6` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `confederation_peers` block supports:

* `peer` - Peer ID.

The `aggregate_address` block supports:

* `id` - ID.
* `prefix` - Aggregate prefix.
* `as_set` - Enable/disable generate AS set path information. Valid values: `enable`, `disable`.
* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `enable`, `disable`.

The `aggregate_address6` block supports:

* `id` - ID.
* `prefix6` - Aggregate IPv6 prefix.
* `as_set` - Enable/disable generate AS set path information. Valid values: `enable`, `disable`.
* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `enable`, `disable`.

The `neighbor` block supports:

* `ip` - IP/IPv6 address of neighbor.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `allowas_in_enable_vpnv4` - Enable/disable to allow my AS in AS path for VPNv4 route. Valid values: `enable`, `disable`.
* `allowas_in_enable_vpnv6` - Enable/disable use of my AS in AS path for VPNv6 route. Valid values: `enable`, `disable`.
* `allowas_in_enable_evpn` - Enable/disable to allow my AS in AS path for L2VPN EVPN route. Valid values: `enable`, `disable`.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `allowas_in_vpnv4` - The maximum number of occurrence of my AS number allowed for VPNv4 route.
* `allowas_in_vpnv6` - The maximum number of occurrence of my AS number allowed for VPNv6 route.
* `allowas_in_evpn` - The maximum number of occurrence of my AS number allowed for L2VPN EVPN route.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
* `attribute_unchanged_vpnv4` - List of attributes that should be unchanged for VPNv4 route. Valid values: `as-path`, `med`, `next-hop`.
* `attribute_unchanged_vpnv6` - List of attributes that should not be changed for VPNv6 route. Valid values: `as-path`, `med`, `next-hop`.
* `activate` - Enable/disable address family IPv4 for this neighbor. Valid values: `enable`, `disable`.
* `activate6` - Enable/disable address family IPv6 for this neighbor. Valid values: `enable`, `disable`.
* `activate_vpnv4` - Enable/disable address family VPNv4 for this neighbor. Valid values: `enable`, `disable`.
* `activate_vpnv6` - Enable/disable address family VPNv6 for this neighbor. Valid values: `enable`, `disable`.
* `activate_evpn` - Enable/disable address family L2VPN EVPN for this neighbor. Valid values: `enable`, `disable`.
* `bfd` - Enable/disable BFD for this neighbor. Valid values: `enable`, `disable`.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart_vpnv4` - Enable/disable advertise VPNv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart_vpnv6` - Enable/disable advertisement of VPNv6 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart_evpn` - Enable/disable advertisement of L2VPN EVPN graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor. Valid values: `enable`, `disable`.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor. Valid values: `enable`, `disable`.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor Valid values: `enable`, `disable`.
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop EBGP neighbors. Valid values: `enable`, `disable`.
* `link_down_failover` - Enable/disable failover upon link down. Valid values: `enable`, `disable`.
* `stale_route` - Enable/disable stale route after neighbor down. Valid values: `enable`, `disable`.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes. Valid values: `enable`, `disable`.
* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes. Valid values: `enable`, `disable`.
* `next_hop_self_vpnv4` - Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor. Valid values: `enable`, `disable`.
* `next_hop_self_vpnv6` - Enable/disable use of outgoing interface's IP address as VPNv6 next-hop for this neighbor. Valid values: `enable`, `disable`.
* `override_capability` - Enable/disable override result of capability negotiation. Valid values: `enable`, `disable`.
* `passive` - Enable/disable sending of open messages to this neighbor. Valid values: `enable`, `disable`.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as_vpnv4` - Enable/disable remove private AS number from VPNv4 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as_vpnv6` - Enable/disable to remove private AS number from VPNv6 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as_evpn` - Enable/disable removing private AS number from L2VPN EVPN outbound updates. Valid values: `enable`, `disable`.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client. Valid values: `enable`, `disable`.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client. Valid values: `enable`, `disable`.
* `route_reflector_client_vpnv4` - Enable/disable VPNv4 AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
* `route_reflector_client_vpnv6` - Enable/disable VPNv6 AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
* `route_reflector_client_evpn` - Enable/disable L2VPN EVPN AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
* `route_server_client` - Enable/disable IPv4 AS route server client. Valid values: `enable`, `disable`.
* `route_server_client6` - Enable/disable IPv6 AS route server client. Valid values: `enable`, `disable`.
* `route_server_client_vpnv4` - Enable/disable VPNv4 AS route server client for this neighbor. Valid values: `enable`, `disable`.
* `route_server_client_vpnv6` - Enable/disable VPNv6 AS route server client for this neighbor. Valid values: `enable`, `disable`.
* `route_server_client_evpn` - Enable/disable L2VPN EVPN AS route server client for this neighbor. Valid values: `enable`, `disable`.
* `rr_attr_allow_change` - Enable/disable allowing change of route attributes when advertising to IPv4 route reflector clients. Valid values: `enable`, `disable`.
* `rr_attr_allow_change6` - Enable/disable allowing change of route attributes when advertising to IPv6 route reflector clients. Valid values: `enable`, `disable`.
* `rr_attr_allow_change_vpnv4` - Enable/disable allowing change of route attributes when advertising to VPNv4 route reflector clients. Valid values: `enable`, `disable`.
* `rr_attr_allow_change_vpnv6` - Enable/disable allowing change of route attributes when advertising to VPNv6 route reflector clients. Valid values: `enable`, `disable`.
* `rr_attr_allow_change_evpn` - Enable/disable allowing change of route attributes when advertising to L2VPN EVPN route reflector clients. Valid values: `enable`, `disable`.
* `shutdown` - Enable/disable shutdown this neighbor. Valid values: `enable`, `disable`.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration_vpnv4` - Enable/disable allow VPNv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration_vpnv6` - Enable/disable VPNv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration_evpn` - Enable/disable L2VPN EVPN inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4. Valid values: `enable`, `disable`.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6. Valid values: `enable`, `disable`.
* `strict_capability_match` - Enable/disable strict capability matching. Valid values: `enable`, `disable`.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_in_vpnv4` - Filter for VPNv4 updates from this neighbor.
* `distribute_list_in_vpnv6` - Filter for VPNv6 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `distribute_list_out_vpnv4` - Filter for VPNv4 updates to this neighbor.
* `distribute_list_out_vpnv6` - Filter for VPNv6 updates to this neighbor.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_in_vpnv4` - BGP filter for VPNv4 inbound routes.
* `filter_list_in_vpnv6` - BGP filter for VPNv6 inbound routes.
* `filter_list_out` - BGP filter for IPv4 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `filter_list_out_vpnv4` - BGP filter for VPNv4 outbound routes.
* `filter_list_out_vpnv6` - BGP filter for VPNv6 outbound routes.
* `interface` - Interface
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `maximum_prefix_vpnv4` - Maximum number of VPNv4 prefixes to accept from this peer.
* `maximum_prefix_vpnv6` - Maximum number of VPNv6 prefixes to accept from this peer.
* `maximum_prefix_evpn` - Maximum number of L2VPN EVPN prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv4` - Maximum VPNv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv6` - Maximum VPNv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_evpn` - Maximum L2VPN EVPN prefix threshold value (1 - 100 percent).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only_vpnv4` - Enable/disable only giving warning message when limit is exceeded for VPNv4 routes. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only_vpnv6` - Enable/disable warning message when limit is exceeded for VPNv6 routes. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only_evpn` - Enable/disable only sending warning message when exceeding limit of L2VPN EVPN routes. Valid values: `enable`, `disable`.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_in_vpnv4` - Inbound filter for VPNv4 updates from this neighbor.
* `prefix_list_in_vpnv6` - Inbound filter for VPNv6 updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `prefix_list_out_vpnv4` - Outbound filter for VPNv4 updates to this neighbor.
* `prefix_list_out_vpnv6` - Outbound filter for VPNv6 updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates. Valid values: `enable`, `disable`.
* `local_as_replace_as` - Replace real AS with local-as in outgoing updates. Valid values: `enable`, `disable`.
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_in_vpnv4` - VPNv4 inbound route map filter.
* `route_map_in_vpnv6` - VPNv6 inbound route map filter.
* `route_map_in_evpn` - L2VPN EVPN inbound route map filter.
* `route_map_out` - IPv4 Outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv4` - VPNv4 outbound route map filter.
* `route_map_out_vpnv6` - VPNv6 outbound route map filter.
* `route_map_out_vpnv4_preferable` - VPNv4 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv6_preferable` - VPNv6 outbound route map filter if this neighbor is preferred.
* `route_map_out_evpn` - L2VPN EVPN outbound route map filter.
* `send_community` - IPv4 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community6` - IPv6 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community_vpnv4` - Send community attribute to neighbor for VPNv4 address family. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community_vpnv6` - Enable/disable sending community attribute to this neighbor for VPNv6 address family. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community_evpn` - Enable/disable sending community attribute to neighbor for L2VPN EVPN address family. Valid values: `standard`, `extended`, `both`, `disable`.
* `keep_alive_timer` - Keep alive timer interval (sec). Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `holdtime_timer` - Interval (sec) before peer considered dead. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `connect_timer` - Interval (sec) for connect timer. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `additional_path` - Enable/disable IPv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `additional_path6` - Enable/disable IPv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `additional_path_vpnv4` - Enable/disable VPNv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `additional_path_vpnv6` - Enable/disable VPNv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv4` - Number of VPNv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv6` - Number of VPNv6 additional paths that can be advertised to this neighbor.
* `password` - Password used in MD5 authentication.
* `auth_options` - Key-chain name for TCP authentication options.
* `conditional_advertise` - Conditional advertisement. The structure of `conditional_advertise` block is documented below.
* `conditional_advertise6` - IPv6 conditional advertisement. The structure of `conditional_advertise6` block is documented below.

The `conditional_advertise` block supports:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - Name of condition route map.
* `condition_type` - Type of condition. Valid values: `exist`, `non-exist`.

The `conditional_advertise6` block supports:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - Name of condition route map.
* `condition_type` - Type of condition. Valid values: `exist`, `non-exist`.

The `neighbor_group` block supports:

* `name` - Neighbor group name.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `allowas_in_enable_vpnv4` - Enable/disable to allow my AS in AS path for VPNv4 route. Valid values: `enable`, `disable`.
* `allowas_in_enable_vpnv6` - Enable/disable use of my AS in AS path for VPNv6 route. Valid values: `enable`, `disable`.
* `allowas_in_enable_evpn` - Enable/disable to allow my AS in AS path for L2VPN EVPN route. Valid values: `enable`, `disable`.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `allowas_in_vpnv4` - The maximum number of occurrence of my AS number allowed for VPNv4 route.
* `allowas_in_vpnv6` - The maximum number of occurrence of my AS number allowed for VPNv6 route.
* `allowas_in_evpn` - The maximum number of occurrence of my AS number allowed for L2VPN EVPN route.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
* `attribute_unchanged_vpnv4` - List of attributes that should be unchanged for VPNv4 route. Valid values: `as-path`, `med`, `next-hop`.
* `attribute_unchanged_vpnv6` - List of attributes that should not be changed for VPNv6 route. Valid values: `as-path`, `med`, `next-hop`.
* `activate` - Enable/disable address family IPv4 for this neighbor. Valid values: `enable`, `disable`.
* `activate6` - Enable/disable address family IPv6 for this neighbor. Valid values: `enable`, `disable`.
* `activate_vpnv4` - Enable/disable address family VPNv4 for this neighbor. Valid values: `enable`, `disable`.
* `activate_vpnv6` - Enable/disable address family VPNv6 for this neighbor. Valid values: `enable`, `disable`.
* `activate_evpn` - Enable/disable address family L2VPN EVPN for this neighbor. Valid values: `enable`, `disable`.
* `bfd` - Enable/disable BFD for this neighbor. Valid values: `enable`, `disable`.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart_vpnv4` - Enable/disable advertise VPNv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart_vpnv6` - Enable/disable advertisement of VPNv6 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart_evpn` - Enable/disable advertisement of L2VPN EVPN graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor. Valid values: `enable`, `disable`.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor. Valid values: `enable`, `disable`.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor Valid values: `enable`, `disable`.
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop EBGP neighbors. Valid values: `enable`, `disable`.
* `link_down_failover` - Enable/disable failover upon link down. Valid values: `enable`, `disable`.
* `stale_route` - Enable/disable stale route after neighbor down. Valid values: `enable`, `disable`.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes. Valid values: `enable`, `disable`.
* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes. Valid values: `enable`, `disable`.
* `next_hop_self_vpnv4` - Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor. Valid values: `enable`, `disable`.
* `next_hop_self_vpnv6` - Enable/disable use of outgoing interface's IP address as VPNv6 next-hop for this neighbor. Valid values: `enable`, `disable`.
* `override_capability` - Enable/disable override result of capability negotiation. Valid values: `enable`, `disable`.
* `passive` - Enable/disable sending of open messages to this neighbor. Valid values: `enable`, `disable`.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as_vpnv4` - Enable/disable remove private AS number from VPNv4 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as_vpnv6` - Enable/disable to remove private AS number from VPNv6 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as_evpn` - Enable/disable removing private AS number from L2VPN EVPN outbound updates. Valid values: `enable`, `disable`.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client. Valid values: `enable`, `disable`.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client. Valid values: `enable`, `disable`.
* `route_reflector_client_vpnv4` - Enable/disable VPNv4 AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
* `route_reflector_client_vpnv6` - Enable/disable VPNv6 AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
* `route_reflector_client_evpn` - Enable/disable L2VPN EVPN AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
* `route_server_client` - Enable/disable IPv4 AS route server client. Valid values: `enable`, `disable`.
* `route_server_client6` - Enable/disable IPv6 AS route server client. Valid values: `enable`, `disable`.
* `route_server_client_vpnv4` - Enable/disable VPNv4 AS route server client for this neighbor. Valid values: `enable`, `disable`.
* `route_server_client_vpnv6` - Enable/disable VPNv6 AS route server client for this neighbor. Valid values: `enable`, `disable`.
* `route_server_client_evpn` - Enable/disable L2VPN EVPN AS route server client for this neighbor. Valid values: `enable`, `disable`.
* `rr_attr_allow_change` - Enable/disable allowing change of route attributes when advertising to IPv4 route reflector clients. Valid values: `enable`, `disable`.
* `rr_attr_allow_change6` - Enable/disable allowing change of route attributes when advertising to IPv6 route reflector clients. Valid values: `enable`, `disable`.
* `rr_attr_allow_change_vpnv4` - Enable/disable allowing change of route attributes when advertising to VPNv4 route reflector clients. Valid values: `enable`, `disable`.
* `rr_attr_allow_change_vpnv6` - Enable/disable allowing change of route attributes when advertising to VPNv6 route reflector clients. Valid values: `enable`, `disable`.
* `rr_attr_allow_change_evpn` - Enable/disable allowing change of route attributes when advertising to L2VPN EVPN route reflector clients. Valid values: `enable`, `disable`.
* `shutdown` - Enable/disable shutdown this neighbor. Valid values: `enable`, `disable`.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration_vpnv4` - Enable/disable allow VPNv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration_vpnv6` - Enable/disable VPNv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration_evpn` - Enable/disable L2VPN EVPN inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4. Valid values: `enable`, `disable`.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6. Valid values: `enable`, `disable`.
* `strict_capability_match` - Enable/disable strict capability matching. Valid values: `enable`, `disable`.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_in_vpnv4` - Filter for VPNv4 updates from this neighbor.
* `distribute_list_in_vpnv6` - Filter for VPNv6 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `distribute_list_out_vpnv4` - Filter for VPNv4 updates to this neighbor.
* `distribute_list_out_vpnv6` - Filter for VPNv6 updates to this neighbor.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_in_vpnv4` - BGP filter for VPNv4 inbound routes.
* `filter_list_in_vpnv6` - BGP filter for VPNv6 inbound routes.
* `filter_list_out` - BGP filter for IPv4 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `filter_list_out_vpnv4` - BGP filter for VPNv4 outbound routes.
* `filter_list_out_vpnv6` - BGP filter for VPNv6 outbound routes.
* `interface` - Interface
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `maximum_prefix_vpnv4` - Maximum number of VPNv4 prefixes to accept from this peer.
* `maximum_prefix_vpnv6` - Maximum number of VPNv6 prefixes to accept from this peer.
* `maximum_prefix_evpn` - Maximum number of L2VPN EVPN prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv4` - Maximum VPNv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv6` - Maximum VPNv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_evpn` - Maximum L2VPN EVPN prefix threshold value (1 - 100 percent).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only_vpnv4` - Enable/disable only giving warning message when limit is exceeded for VPNv4 routes. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only_vpnv6` - Enable/disable warning message when limit is exceeded for VPNv6 routes. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only_evpn` - Enable/disable only sending warning message when exceeding limit of L2VPN EVPN routes. Valid values: `enable`, `disable`.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_in_vpnv4` - Inbound filter for VPNv4 updates from this neighbor.
* `prefix_list_in_vpnv6` - Inbound filter for VPNv6 updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `prefix_list_out_vpnv4` - Outbound filter for VPNv4 updates to this neighbor.
* `prefix_list_out_vpnv6` - Outbound filter for VPNv6 updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `remote_as_filter` - BGP filter for remote AS.
* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates. Valid values: `enable`, `disable`.
* `local_as_replace_as` - Replace real AS with local-as in outgoing updates. Valid values: `enable`, `disable`.
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_in_vpnv4` - VPNv4 inbound route map filter.
* `route_map_in_vpnv6` - VPNv6 inbound route map filter.
* `route_map_in_evpn` - L2VPN EVPN inbound route map filter.
* `route_map_out` - IPv4 Outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv4` - VPNv4 outbound route map filter.
* `route_map_out_vpnv6` - VPNv6 outbound route map filter.
* `route_map_out_vpnv4_preferable` - VPNv4 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv6_preferable` - VPNv6 outbound route map filter if this neighbor is preferred.
* `route_map_out_evpn` - L2VPN EVPN outbound route map filter.
* `send_community` - IPv4 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community6` - IPv6 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community_vpnv4` - Send community attribute to neighbor for VPNv4 address family. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community_vpnv6` - Enable/disable sending community attribute to this neighbor for VPNv6 address family. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community_evpn` - Enable/disable sending community attribute to neighbor for L2VPN EVPN address family. Valid values: `standard`, `extended`, `both`, `disable`.
* `keep_alive_timer` - Keep alive timer interval (sec). Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `holdtime_timer` - Interval (sec) before peer considered dead. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `connect_timer` - Interval (sec) for connect timer. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `additional_path` - Enable/disable IPv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `additional_path6` - Enable/disable IPv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `additional_path_vpnv4` - Enable/disable VPNv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `additional_path_vpnv6` - Enable/disable VPNv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv4` - Number of VPNv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv6` - Number of VPNv6 additional paths that can be advertised to this neighbor.
* `password` - Password used in MD5 authentication.
* `auth_options` - Key-chain name for TCP authentication options.

The `neighbor_range` block supports:

* `id` - Neighbor range ID.
* `prefix` - Neighbor range prefix.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.

The `neighbor_range6` block supports:

* `id` - IPv6 neighbor range ID.
* `prefix6` - IPv6 prefix.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.

The `network` block supports:

* `id` - ID.
* `prefix` - Network prefix.
* `network_import_check` - Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
* `backdoor` - Enable/disable route as backdoor. Valid values: `enable`, `disable`.
* `route_map` - Route map to modify generated route.
* `prefix_name` - Name of firewall address or address group.

The `network6` block supports:

* `id` - ID.
* `prefix6` - Network IPv6 prefix.
* `network_import_check` - Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
* `backdoor` - Enable/disable route as backdoor. Valid values: `enable`, `disable`.
* `route_map` - Route map to modify generated route.

The `redistribute` block supports:

* `name` - Distribute list entry name.
* `status` - Status Valid values: `enable`, `disable`.
* `route_map` - Route map name.

The `redistribute6` block supports:

* `name` - Distribute list entry name.
* `status` - Status Valid values: `enable`, `disable`.
* `route_map` - Route map name.

The `admin_distance` block supports:

* `id` - ID.
* `neighbour_prefix` - Neighbor address prefix.
* `route_list` - Access list of routes to apply new distance to.
* `distance` - Administrative distance to apply (1 - 255).

The `vrf` block supports:

* `vrf` - Origin VRF ID. On FortiOS versions 7.2.0-7.2.3: 0 - 63. On FortiOS versions 7.2.4-7.6.0: 0 - 251. On FortiOS versions >= 7.6.1: 0-511.
* `role` - VRF role. Valid values: `standalone`, `ce`, `pe`.
* `rd` - Route Distinguisher: AA|AA:NN.
* `export_rt` - List of export route target. The structure of `export_rt` block is documented below.
* `import_rt` - List of import route target. The structure of `import_rt` block is documented below.
* `import_route_map` - Import route map.
* `leak_target` - Target VRF table. The structure of `leak_target` block is documented below.

The `export_rt` block supports:

* `route_target` - Attribute: AA|AA:NN.

The `import_rt` block supports:

* `route_target` - Attribute: AA|AA:NN.

The `leak_target` block supports:

* `vrf` - Target VRF ID. On FortiOS versions 7.2.0-7.2.3: 0 - 63. On FortiOS versions 7.2.4-7.6.0: 0 - 251. On FortiOS versions >= 7.6.1: 0-511.
* `route_map` - Route map of VRF leaking.
* `interface` - Interface which is used to leak routes to target VRF.

The `vrf6` block supports:

* `vrf` - Origin VRF ID. On FortiOS versions 7.2.0-7.2.3: 0 - 63. On FortiOS versions 7.2.4-7.6.0: 0 - 251. On FortiOS versions >= 7.6.1: 0-511.
* `role` - VRF role. Valid values: `standalone`, `ce`, `pe`.
* `rd` - Route Distinguisher: AA:NN|A.B.C.D:NN.
* `export_rt` - List of export route target. The structure of `export_rt` block is documented below.
* `import_rt` - List of import route target. The structure of `import_rt` block is documented below.
* `import_route_map` - Import route map.
* `leak_target` - Target VRF table. The structure of `leak_target` block is documented below.

The `export_rt` block supports:

* `route_target` - Attribute: AA:NN|A.B.C.D:NN.

The `import_rt` block supports:

* `route_target` - Attribute: AA:NN|A.B.C.D:NN

The `leak_target` block supports:

* `vrf` - Target VRF ID. On FortiOS versions 7.2.0-7.2.3: 0 - 63. On FortiOS versions 7.2.4-7.6.0: 0 - 251. On FortiOS versions >= 7.6.1: 0-511.
* `route_map` - Route map of VRF leaking.
* `interface` - Interface which is used to leak routes to target VRF.

The `vrf_leak` block supports:

* `vrf` - Origin VRF ID (0 - 31).
* `target` - Target VRF table. The structure of `target` block is documented below.

The `target` block supports:

* `vrf` - Target VRF ID (0 - 31).
* `route_map` - Route map of VRF leaking.
* `interface` - Interface which is used to leak routes to target VRF.

The `vrf_leak6` block supports:

* `vrf` - Origin VRF ID (0 - 31).
* `target` - Target VRF table. The structure of `target` block is documented below.

The `target` block supports:

* `vrf` - Target VRF ID (0 - 31).
* `route_map` - Route map of VRF leaking.
* `interface` - Interface which is used to leak routes to target VRF.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Bgp can be imported using any of these accepted formats:
```
$ terraform import fortios_router_bgp.labelname RouterBgp

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_bgp.labelname RouterBgp
$ unset "FORTIOS_IMPORT_TABLE"
```
