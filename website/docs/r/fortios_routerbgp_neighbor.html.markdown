---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_routerbgp_neighbor"
description: |-
  BGP neighbor table.
---

# fortios_routerbgp_neighbor
BGP neighbor table.

~> The provider supports the definition of Neighbor in Router Bgp `fortios_router_bgp`, and also allows the definition of separate Neighbor resources `fortios_routerbgp_neighbor`, but do not use a `fortios_router_bgp` with in-line Neighbor in conjunction with any `fortios_routerbgp_neighbor` resources, otherwise conflicts and overwrite will occur.


## Argument Reference

The following arguments are supported:

* `ip` - (Required) IP/IPv6 address of neighbor.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged.
* `activate` - Enable/disable address family IPv4 for this neighbor.
* `activate6` - Enable/disable address family IPv6 for this neighbor.
* `bfd` - Enable/disable BFD for this neighbor.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor.
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor.
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor.
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
* `override_capability` - Enable/disable override result of capability negotiation.
* `passive` - Enable/disable sending of open messages to this neighbor.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client.
* `route_server_client` - Enable/disable IPv4 AS route server client.
* `route_server_client6` - Enable/disable IPv6 AS route server client.
* `shutdown` - Enable/disable shutdown this neighbor.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6.
* `strict_capability_match` - Enable/disable strict capability matching.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_out` - BGP filter for IPv4 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `interface` - Interface
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates.
* `local_as_replace_as` - Replace real AS with local-as in outgoing updates.
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_out` - IPv4 Outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `send_community` - IPv4 Send community attribute to neighbor.
* `send_community6` - IPv6 Send community attribute to neighbor.
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
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `password` - Password used in MD5 authentication.
* `conditional_advertise` - Conditional advertisement. The structure of `conditional_advertise` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `conditional_advertise` block supports:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - Name of condition route map.
* `condition_type` - Type of condition.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ip}}.

## Import

Routerbgp Neighbor can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_routerbgp_neighbor.labelname {{ip}}
$ unset "FORTIOS_IMPORT_TABLE"
```
