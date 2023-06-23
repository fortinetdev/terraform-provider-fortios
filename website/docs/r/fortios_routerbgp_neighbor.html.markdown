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
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
* `allowas_in_enable_vpnv4` - Enable/disable to allow my AS in AS path for VPNv4 route. Valid values: `enable`, `disable`.
* `allowas_in_enable_evpn` - Enable/disable to allow my AS in AS path for L2VPN EVPN route. Valid values: `enable`, `disable`.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `allowas_in_vpnv4` - The maximum number of occurrence of my AS number allowed for VPNv4 route.
* `allowas_in_evpn` - The maximum number of occurrence of my AS number allowed for L2VPN EVPN route.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
* `attribute_unchanged_vpnv4` - List of attributes that should be unchanged for VPNv4 route. Valid values: `as-path`, `med`, `next-hop`.
* `activate` - Enable/disable address family IPv4 for this neighbor. Valid values: `enable`, `disable`.
* `activate6` - Enable/disable address family IPv6 for this neighbor. Valid values: `enable`, `disable`.
* `activate_vpnv4` - Enable/disable address family VPNv4 for this neighbor. Valid values: `enable`, `disable`.
* `activate_evpn` - Enable/disable address family L2VPN EVPN for this neighbor. Valid values: `enable`, `disable`.
* `bfd` - Enable/disable BFD for this neighbor. Valid values: `enable`, `disable`.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
* `capability_graceful_restart_vpnv4` - Enable/disable advertise VPNv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
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
* `override_capability` - Enable/disable override result of capability negotiation. Valid values: `enable`, `disable`.
* `passive` - Enable/disable sending of open messages to this neighbor. Valid values: `enable`, `disable`.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as_vpnv4` - Enable/disable remove private AS number from VPNv4 outbound updates. Valid values: `enable`, `disable`.
* `remove_private_as_evpn` - Enable/disable removing private AS number from L2VPN EVPN outbound updates. Valid values: `enable`, `disable`.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client. Valid values: `enable`, `disable`.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client. Valid values: `enable`, `disable`.
* `route_reflector_client_vpnv4` - Enable/disable VPNv4 AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
* `route_reflector_client_evpn` - Enable/disable L2VPN EVPN AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
* `route_server_client` - Enable/disable IPv4 AS route server client. Valid values: `enable`, `disable`.
* `route_server_client6` - Enable/disable IPv6 AS route server client. Valid values: `enable`, `disable`.
* `route_server_client_vpnv4` - Enable/disable VPNv4 AS route server client for this neighbor. Valid values: `enable`, `disable`.
* `route_server_client_evpn` - Enable/disable L2VPN EVPN AS route server client for this neighbor. Valid values: `enable`, `disable`.
* `shutdown` - Enable/disable shutdown this neighbor. Valid values: `enable`, `disable`.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
* `soft_reconfiguration_vpnv4` - Enable/disable allow VPNv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
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
* `maximum_prefix_evpn` - Maximum number of L2VPN EVPN prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_vpnv4` - Maximum VPNv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold_evpn` - Maximum L2VPN EVPN prefix threshold value (1 - 100 percent).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only_vpnv4` - Enable/disable only giving warning message when limit is exceeded for VPNv4 routes. Valid values: `enable`, `disable`.
* `maximum_prefix_warning_only_evpn` - Enable/disable only sending warning message when exceeding limit of L2VPN EVPN routes. Valid values: `enable`, `disable`.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_in_vpnv4` - Inbound filter for VPNv4 updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `prefix_list_out_vpnv4` - Outbound filter for VPNv4 updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates. Valid values: `enable`, `disable`.
* `local_as_replace_as` - Replace real AS with local-as in outgoing updates. Valid values: `enable`, `disable`.
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_in_vpnv4` - VPNv4 inbound route map filter.
* `route_map_in_evpn` - L2VPN EVPN inbound route map filter.
* `route_map_out` - IPv4 Outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv4` - VPNv4 outbound route map filter.
* `route_map_out_vpnv4_preferable` - VPNv4 outbound route map filter if the peer is preferred.
* `route_map_out_evpn` - L2VPN EVPN outbound route map filter.
* `send_community` - IPv4 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community6` - IPv6 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community_vpnv4` - Send community attribute to neighbor for VPNv4 address family. Valid values: `standard`, `extended`, `both`, `disable`.
* `send_community_evpn` - Enable/disable sending community attribute to neighbor for L2VPN EVPN address family. Valid values: `standard`, `extended`, `both`, `disable`.
* `keep_alive_timer` - Keep alive timer interval (sec).
* `holdtime_timer` - Interval (sec) before peer considered dead.
* `connect_timer` - Interval (sec) for connect timer.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight.
* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `additional_path` - Enable/disable IPv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `additional_path6` - Enable/disable IPv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `additional_path_vpnv4` - Enable/disable VPNv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv4` - Number of VPNv4 additional paths that can be advertised to this neighbor.
* `password` - Password used in MD5 authentication.
* `conditional_advertise` - Conditional advertisement. The structure of `conditional_advertise` block is documented below.
* `conditional_advertise6` - IPv6 conditional advertisement. The structure of `conditional_advertise6` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `conditional_advertise` block supports:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - Name of condition route map.
* `condition_type` - Type of condition. Valid values: `exist`, `non-exist`.

The `conditional_advertise6` block supports:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - Name of condition route map.
* `condition_type` - Type of condition. Valid values: `exist`, `non-exist`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ip}}.

## Import

Routerbgp Neighbor can be imported using any of these accepted formats:
```
$ terraform import fortios_routerbgp_neighbor.labelname {{ip}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_routerbgp_neighbor.labelname {{ip}}
$ unset "FORTIOS_IMPORT_TABLE"
```
