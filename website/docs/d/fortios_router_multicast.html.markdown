---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicast"
description: |-
  Get information on fortios router multicast.
---

# Data Source: fortios_router_multicast
Use this data source to get information on fortios router multicast

## Argument Reference

No arguments available for the data source.

## Attribute Reference

The following attributes are exported:

* `route_threshold` - Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
* `route_limit` - Maximum number of multicast routes.
* `multicast_routing` - Enable/disable IP multicast routing.
* `pim_sm_global` - PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
* `interface` - PIM interfaces. The structure of `interface` block is documented below.

The `pim_sm_global` block contains:

* `message_interval` - Period of time between sending periodic PIM join/prune messages in seconds (1 - 65535, default = 60).
* `join_prune_holdtime` - Join/prune holdtime (1 - 65535, default = 210).
* `accept_register_list` - Sources allowed to register packets with this Rendezvous Point (RP).
* `accept_source_list` - Sources allowed to send multicast traffic.
* `bsr_candidate` - Enable/disable allowing this router to become a bootstrap router (BSR).
* `bsr_interface` - Interface to advertise as candidate BSR.
* `bsr_priority` - BSR priority (0 - 255, default = 0).
* `bsr_hash` - BSR hash length (0 - 32, default = 10).
* `bsr_allow_quick_refresh` - Enable/disable accept BSR quick refresh packets from neighbors.
* `cisco_register_checksum` - Checksum entire register packet(for old Cisco IOS compatibility).
* `cisco_register_checksum_group` - Cisco register checksum only these groups.
* `cisco_crp_prefix` - Enable/disable making candidate RP compatible with old Cisco IOS.
* `cisco_ignore_rp_set_priority` - Use only hash for RP selection (compatibility with old Cisco IOS).
* `register_rp_reachability` - Enable/disable check RP is reachable before registering packets.
* `register_source` - Override source address in register packets.
* `register_source_interface` - Override with primary interface address.
* `register_source_ip` - Override with local IP address.
* `register_supression` - Period of time to honor register-stop message (1 - 65535 sec, default = 60).
* `null_register_retries` - Maximum retries of null register (1 - 20, default = 1).
* `rp_register_keepalive` - Timeout for RP receiving data on (S,G) tree (1 - 65535 sec, default = 185).
* `spt_threshold` - Enable/disable switching to source specific trees.
* `spt_threshold_group` - Groups allowed to switch to source tree.
* `ssm` - Enable/disable source specific multicast.
* `ssm_range` - Groups allowed to source specific multicast.
* `register_rate_limit` - Limit of packets/sec per source registered through this RP (0 - 65535, default = 0 which means unlimited).
* `rp_address` - Statically configure RP addresses. The structure of `rp_address` block is documented below.

The `rp_address` block contains:

* `id` - ID.
* `ip_address` - RP router address.
* `group` - Groups to use this RP.

The `interface` block contains:

* `name` - Interface name.
* `ttl_threshold` - Minimum TTL of multicast packets that will be forwarded (applied only to new multicast routes) (1 - 255, default = 1).
* `pim_mode` - PIM operation mode.
* `passive` - Enable/disable listening to IGMP but not participating in PIM.
* `bfd` - Enable/disable Protocol Independent Multicast (PIM) Bidirectional Forwarding Detection (BFD).
* `neighbour_filter` - Routers acknowledged as neighbor routers.
* `hello_interval` - Interval between sending PIM hello messages (0 - 65535 sec, default = 30).
* `hello_holdtime` - Time before old neighbor information expires (0 - 65535 sec, default = 105).
* `cisco_exclude_genid` - Exclude GenID from hello packets (compatibility with old Cisco IOS).
* `dr_priority` - DR election priority.
* `propagation_delay` - Delay flooding packets on this interface (100 - 5000 msec, default = 500).
* `state_refresh_interval` - Interval between sending state-refresh packets (1 - 100 sec, default = 60).
* `rp_candidate` - Enable/disable compete to become RP in elections.
* `rp_candidate_group` - Multicast groups managed by this RP.
* `rp_candidate_priority` - Router's priority as RP.
* `rp_candidate_interval` - RP candidate advertisement interval (1 - 16383 sec, default = 60).
* `multicast_flow` - Acceptable source for multicast group.
* `static_group` - Statically set multicast groups to forward out.
* `join_group` - Join multicast groups. The structure of `join_group` block is documented below.
* `igmp` - IGMP configuration options. The structure of `igmp` block is documented below.

The `join_group` block contains:

* `address` - Multicast group IP address.

The `igmp` block contains:

* `access_group` - Groups IGMP hosts are allowed to join.
* `version` - Maximum version of IGMP to support.
* `immediate_leave_group` - Groups to drop membership for immediately after receiving IGMPv2 leave.
* `last_member_query_interval` - Timeout between IGMPv2 leave and removing group (1 - 65535 msec, default = 1000).
* `last_member_query_count` - Number of group specific queries before removing group (2 - 7, default = 2).
* `query_max_response_time` - Maximum time to wait for a IGMP query response (1 - 25 sec, default = 10).
* `query_interval` - Interval between queries to IGMP hosts (1 - 65535 sec, default = 125).
* `query_timeout` - Timeout between queries before becoming querier for network (60 - 900, default = 255).
* `router_alert_check` - Enable/disable require IGMP packets contain router alert option.

