---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp"
description: |-
  Get information on fortios router bgp.
---

# Data Source: fortios_router_bgp
Use this data source to get information on fortios router bgp

## Example Usage

```hcl
data "fortios_router_bgp" sample1 {
}

output output1 {
  value = data.fortios_router_bgp.sample1.neighbor
}
```

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `as` - Router AS number, valid from 1 to 4294967295, 0 to disable BGP.
* `router_id` - Router ID.
* `keepalive_timer` - Frequency to send keep alive requests.
* `holdtime_timer` - Number of seconds to mark peer as dead.
* `always_compare_med` - Enable/disable always compare MED.
* `bestpath_as_path_ignore` - Enable/disable ignore AS path.
* `bestpath_cmp_confed_aspath` - Enable/disable compare federation AS path length.
* `bestpath_cmp_routerid` - Enable/disable compare router ID for identical EBGP paths.
* `bestpath_med_confed` - Enable/disable compare MED among confederation paths.
* `bestpath_med_missing_as_worst` - Enable/disable treat missing MED as least preferred.
* `client_to_client_reflection` - Enable/disable client-to-client route reflection.
* `dampening` - Enable/disable route-flap dampening.
* `deterministic_med` - Enable/disable enforce deterministic comparison of MED.
* `ebgp_multipath` - Enable/disable EBGP multi-path.
* `ibgp_multipath` - Enable/disable IBGP multi-path.
* `enforce_first_as` - Enable/disable enforce first AS for EBGP routes.
* `fast_external_failover` - Enable/disable reset peer BGP session if link goes down.
* `log_neighbour_changes` - Enable logging of BGP neighbour's changes
* `network_import_check` - Enable/disable ensure BGP network route exists in IGP.
* `ignore_optional_capability` - Don't send unknown optional capability notification message
* `additional_path` - Enable/disable selection of BGP IPv4 additional paths.
* `additional_path6` - Enable/disable selection of BGP IPv6 additional paths.
* `additional_path_vpnv4` - Enable/disable selection of BGP VPNv4 additional paths.
* `multipath_recursive_distance` - Enable/disable use of recursive distance to select multipath.
* `recursive_next_hop` - Enable/disable recursive resolution of next-hop using BGP route.
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
* `synchronization` - Enable/disable only advertise routes from iBGP if routes present in an IGP.
* `graceful_restart` - Enable/disable BGP graceful restart capabilities.
* `graceful_restart_time` - Time needed for neighbors to restart (sec).
* `graceful_stalepath_time` - Time to hold stale paths of restarting neighbor (sec).
* `graceful_update_delay` - Route advertisement/selection delay after restart (sec).
* `graceful_end_on_timer` - Enable/disable to exit graceful restart on timer only.
* `additional_path_select` - Number of additional paths to be selected for each IPv4 NLRI.
* `additional_path_select6` - Number of additional paths to be selected for each IPv6 NLRI.
* `additional_path_select_vpnv4` - Number of additional paths to be selected for each VPNv4 NLRI.
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

The `confederation_peers` block contains:

* `peer` - Peer ID.

The `aggregate_address` block contains:

* `id` - ID.
* `prefix` - Aggregate prefix.
* `as_set` - Enable/disable generate AS set path information.
* `summary_only` - Enable/disable filter more specific routes from updates.

The `aggregate_address6` block contains:

* `id` - ID.
* `prefix6` - Aggregate IPv6 prefix.
* `as_set` - Enable/disable generate AS set path information.
* `summary_only` - Enable/disable filter more specific routes from updates.

The `neighbor` block contains:

* `ip` - IP/IPv6 address of neighbor.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `allowas_in_vpnv4` - The maximum number of occurrence of my AS number allowed for VPNv4 route.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged.
* `attribute_unchanged_vpnv4` - List of attributes that should be unchanged for VPNv4 route.
* `activate` - Enable/disable address family IPv4 for this neighbor.
* `activate6` - Enable/disable address family IPv6 for this neighbor.
* `activate_vpnv4` - Enable/disable address family VPNv4 for this neighbor.
* `bfd` - Enable/disable BFD for this neighbor.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor.
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor.
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor.
* `capability_graceful_restart_vpnv4` - Enable/disable advertise VPNv4 graceful restart capability to this neighbor.
* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop EBGP neighbors.
* `link_down_failover` - Enable/disable failover upon link down.
* `stale_route` - Enable/disable stale route after neighbor down.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor.
* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.
* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.
* `next_hop_self_vpnv4` - Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor.
* `override_capability` - Enable/disable override result of capability negotiation.
* `passive` - Enable/disable sending of open messages to this neighbor.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates.
* `remove_private_as_vpnv4` - Enable/disable remove private AS number from VPNv4 outbound updates.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client.
* `route_reflector_client_vpnv4` - Enable/disable VPNv4 AS route reflector client for this neighbor.
* `route_server_client` - Enable/disable IPv4 AS route server client.
* `route_server_client6` - Enable/disable IPv6 AS route server client.
* `route_server_client_vpnv4` - Enable/disable VPNv4 AS route server client for this neighbor.
* `shutdown` - Enable/disable shutdown this neighbor.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration.
* `soft_reconfiguration_vpnv4` - Enable/disable allow VPNv4 inbound soft reconfiguration.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6.
* `strict_capability_match` - Enable/disable strict capability matching.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_in_vpnv4` - Filter for VPNv4 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `distribute_list_out_vpnv4` - Filter for VPNv4 updates to this neighbor.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_out` - BGP filter for IPv4 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `interface` - Interface
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `maximum_prefix_vpnv4` - Maximum number of VPNv4 prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv4` - Maximum VPNv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded.
* `maximum_prefix_warning_only_vpnv4` - Enable/disable only giving warning message when limit is exceeded for VPNv4 routes.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_in_vpnv4` - Inbound filter for VPNv4 updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `prefix_list_out_vpnv4` - Outbound filter for VPNv4 updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates.
* `local_as_replace_as` - Replace real AS with local-as in outgoing updates.
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_in_vpnv4` - VPNv4 inbound route map filter.
* `route_map_out` - IPv4 Outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv4` - VPNv4 outbound route map filter.
* `route_map_out_vpnv4_preferable` - VPNv4 outbound route map filter if the peer is preferred.
* `send_community` - IPv4 Send community attribute to neighbor.
* `send_community6` - IPv6 Send community attribute to neighbor.
* `send_community_vpnv4` - Send community attribute to neighbor for VPNv4 address family.
* `keep_alive_timer` - Keep alive timer interval (sec).
* `holdtime_timer` - Interval (sec) before peer considered dead.
* `connect_timer` - Interval (sec) for connect timer.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight.
* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `additional_path` - Enable/disable IPv4 additional-path capability.
* `additional_path6` - Enable/disable IPv6 additional-path capability.
* `additional_path_vpnv4` - Enable/disable VPNv4 additional-path capability.
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv4` - Number of VPNv4 additional paths that can be advertised to this neighbor.
* `password` - Password used in MD5 authentication.
* `conditional_advertise` - Conditional advertisement. The structure of `conditional_advertise` block is documented below.
* `conditional_advertise6` - IPv6 conditional advertisement. The structure of `conditional_advertise6` block is documented below.

The `conditional_advertise` block contains:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - Name of condition route map.
* `condition_type` - Type of condition.

The `conditional_advertise6` block contains:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - Name of condition route map.
* `condition_type` - Type of condition.

The `neighbor_group` block contains:

* `name` - Neighbor group name.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `allowas_in_vpnv4` - The maximum number of occurrence of my AS number allowed for VPNv4 route.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged.
* `attribute_unchanged_vpnv4` - List of attributes that should be unchanged for VPNv4 route.
* `activate` - Enable/disable address family IPv4 for this neighbor.
* `activate6` - Enable/disable address family IPv6 for this neighbor.
* `activate_vpnv4` - Enable/disable address family VPNv4 for this neighbor.
* `bfd` - Enable/disable BFD for this neighbor.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor.
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor.
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor.
* `capability_graceful_restart_vpnv4` - Enable/disable advertise VPNv4 graceful restart capability to this neighbor.
* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop EBGP neighbors.
* `link_down_failover` - Enable/disable failover upon link down.
* `stale_route` - Enable/disable stale route after neighbor down.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor.
* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.
* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.
* `next_hop_self_vpnv4` - Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor.
* `override_capability` - Enable/disable override result of capability negotiation.
* `passive` - Enable/disable sending of open messages to this neighbor.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates.
* `remove_private_as_vpnv4` - Enable/disable remove private AS number from VPNv4 outbound updates.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client.
* `route_reflector_client_vpnv4` - Enable/disable VPNv4 AS route reflector client for this neighbor.
* `route_server_client` - Enable/disable IPv4 AS route server client.
* `route_server_client6` - Enable/disable IPv6 AS route server client.
* `route_server_client_vpnv4` - Enable/disable VPNv4 AS route server client for this neighbor.
* `shutdown` - Enable/disable shutdown this neighbor.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration.
* `soft_reconfiguration_vpnv4` - Enable/disable allow VPNv4 inbound soft reconfiguration.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6.
* `strict_capability_match` - Enable/disable strict capability matching.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_in_vpnv4` - Filter for VPNv4 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `distribute_list_out_vpnv4` - Filter for VPNv4 updates to this neighbor.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_out` - BGP filter for IPv4 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `interface` - Interface
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `maximum_prefix_vpnv4` - Maximum number of VPNv4 prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv4` - Maximum VPNv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded.
* `maximum_prefix_warning_only_vpnv4` - Enable/disable only giving warning message when limit is exceeded for VPNv4 routes.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_in_vpnv4` - Inbound filter for VPNv4 updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `prefix_list_out_vpnv4` - Outbound filter for VPNv4 updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates.
* `local_as_replace_as` - Replace real AS with local-as in outgoing updates.
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_in_vpnv4` - VPNv4 inbound route map filter.
* `route_map_out` - IPv4 Outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv4` - VPNv4 outbound route map filter.
* `route_map_out_vpnv4_preferable` - VPNv4 outbound route map filter if the peer is preferred.
* `send_community` - IPv4 Send community attribute to neighbor.
* `send_community6` - IPv6 Send community attribute to neighbor.
* `send_community_vpnv4` - Send community attribute to neighbor for VPNv4 address family.
* `keep_alive_timer` - Keep alive timer interval (sec).
* `holdtime_timer` - Interval (sec) before peer considered dead.
* `connect_timer` - Interval (sec) for connect timer.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight.
* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `additional_path` - Enable/disable IPv4 additional-path capability.
* `additional_path6` - Enable/disable IPv6 additional-path capability.
* `additional_path_vpnv4` - Enable/disable VPNv4 additional-path capability.
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv4` - Number of VPNv4 additional paths that can be advertised to this neighbor.

The `neighbor_range` block contains:

* `id` - Neighbor range ID.
* `prefix` - Neighbor range prefix.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.

The `neighbor_range6` block contains:

* `id` - IPv6 neighbor range ID.
* `prefix6` - IPv6 prefix.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.

The `network` block contains:

* `id` - ID.
* `prefix` - Network prefix.
* `network_import_check` - Configure insurance of BGP network route existence in IGP.
* `backdoor` - Enable/disable route as backdoor.
* `route_map` - Route map to modify generated route.

The `network6` block contains:

* `id` - ID.
* `prefix6` - Network IPv6 prefix.
* `network_import_check` - Configure insurance of BGP network route existence in IGP.
* `backdoor` - Enable/disable route as backdoor.
* `route_map` - Route map to modify generated route.

The `redistribute` block contains:

* `name` - Distribute list entry name.
* `status` - Status
* `route_map` - Route map name.

The `redistribute6` block contains:

* `name` - Distribute list entry name.
* `status` - Status
* `route_map` - Route map name.

The `admin_distance` block contains:

* `id` - ID.
* `neighbour_prefix` - Neighbor address prefix.
* `route_list` - Access list of routes to apply new distance to.
* `distance` - Administrative distance to apply (1 - 255).

The `vrf` block contains:

* `vrf` - Origin VRF ID (0 - 63).
* `role` - VRF role.
* `rd` - Route Distinguisher: AA|AA:NN.
* `export_rt` - List of export route target. The structure of `export_rt` block is documented below.
* `import_rt` - List of import route target. The structure of `import_rt` block is documented below.
* `import_route_map` - Import route map.
* `leak_target` - Target VRF table. The structure of `leak_target` block is documented below.

The `export_rt` block contains:

* `route_target` - Attribute: AA|AA:NN.

The `import_rt` block contains:

* `route_target` - Attribute: AA|AA:NN.

The `leak_target` block contains:

* `vrf` - Target VRF ID (0 - 63).
* `route_map` - Route map of VRF leaking.
* `interface` - Interface which is used to leak routes to target VRF.

The `vrf6` block contains:

* `vrf` - Origin VRF ID (0 - 63).
* `leak_target` - Target VRF table. The structure of `leak_target` block is documented below.

The `leak_target` block contains:

* `vrf` - Target VRF ID (0 - 63).
* `route_map` - Route map of VRF leaking.
* `interface` - Interface which is used to leak routes to target VRF.

The `vrf_leak` block contains:

* `vrf` - Origin VRF ID <0 - 31>.
* `target` - Target VRF table. The structure of `target` block is documented below.

The `target` block contains:

* `vrf` - Target VRF ID <0 - 31>.
* `route_map` - Route map of VRF leaking.
* `interface` - Interface which is used to leak routes to target VRF.

The `vrf_leak6` block contains:

* `vrf` - Origin VRF ID <0 - 31>.
* `target` - Target VRF table. The structure of `target` block is documented below.

The `target` block contains:

* `vrf` - Target VRF ID <0 - 31>.
* `route_map` - Route map of VRF leaking.
* `interface` - Interface which is used to leak routes to target VRF.

