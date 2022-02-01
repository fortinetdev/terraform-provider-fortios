---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf"
description: |-
  Get information on fortios router ospf.
---

# Data Source: fortios_router_ospf
Use this data source to get information on fortios router ospf

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `abr_type` - Area border router type.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `distance_external` - Administrative external distance.
* `distance_inter_area` - Administrative inter-area distance.
* `distance_intra_area` - Administrative intra-area distance.
* `database_overflow` - Enable/disable database overflow.
* `database_overflow_max_lsas` - Database overflow maximum LSAs.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `default_information_originate` - Enable/disable generation of default route.
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type.
* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `distance` - Distance of the route.
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility.
* `router_id` - Router ID.
* `spf_timers` - SPF calculation frequency.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `log_neighbour_changes` - Enable logging of OSPF neighbour's changes
* `distribute_list_in` - Filter incoming routes.
* `distribute_route_map_in` - Filter incoming external routes by route-map.
* `restart_mode` - OSPF restart mode (graceful or LLS).
* `restart_period` - Graceful restart period.
* `area` - OSPF area configuration. The structure of `area` block is documented below.
* `ospf_interface` - OSPF interface configuration. The structure of `ospf_interface` block is documented below.
* `network` - OSPF network configuration. The structure of `network` block is documented below.
* `neighbor` - OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `summary_address` - IP address summary configuration. The structure of `summary_address` block is documented below.
* `distribute_list` - Distribute list configuration. The structure of `distribute_list` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.

The `area` block contains:

* `id` - Area entry IP address.
* `shortcut` - Enable/disable shortcut option.
* `authentication` - Authentication type.
* `default_cost` - Summary default cost of stub or NSSA area.
* `nssa_translator_role` - NSSA translator role type.
* `stub_type` - Stub summary setting.
* `type` - Area type setting.
* `nssa_default_information_originate` - Redistribute, advertise, or do not originate Type-7 default route into NSSA area.
* `nssa_default_information_originate_metric` - OSPF default metric.
* `nssa_default_information_originate_metric_type` - OSPF metric type for default routes.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area.
* `comments` - Comment.
* `range` - OSPF area range configuration. The structure of `range` block is documented below.
* `virtual_link` - OSPF virtual link configuration. The structure of `virtual_link` block is documented below.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.

The `range` block contains:

* `id` - Range entry ID.
* `prefix` - Prefix.
* `advertise` - Enable/disable advertise status.
* `substitute` - Substitute prefix.
* `substitute_status` - Enable/disable substitute status.

The `virtual_link` block contains:

* `name` - Virtual link entry name.
* `authentication` - Authentication type.
* `authentication_key` - Authentication key.
* `keychain` - Message-digest key-chain name.
* `md5_key` - MD5 key.
* `md5_keychain` - Authentication MD5 key-chain name.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `peer` - Peer IP.
* `md5_keys` - MD5 key. The structure of `md5_keys` block is documented below.

The `md5_keys` block contains:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.

The `filter_list` block contains:

* `id` - Filter list entry ID.
* `list` - Access-list or prefix-list name.
* `direction` - Direction.

The `ospf_interface` block contains:

* `name` - Interface entry name.
* `comments` - Comment.
* `interface` - Configuration interface name.
* `ip` - IP address.
* `authentication` - Authentication type.
* `authentication_key` - Authentication key.
* `keychain` - Message-digest key-chain name.
* `md5_key` - MD5 key.
* `md5_keychain` - Authentication MD5 key-chain name.
* `prefix_length` - Prefix length.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - Priority.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `hello_multiplier` - Number of hello packets within dead interval.
* `database_filter_out` - Enable/disable control of flooding out LSAs.
* `mtu` - MTU for database description packets.
* `mtu_ignore` - Enable/disable ignore MTU.
* `network_type` - Network type.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `status` - Enable/disable status.
* `resync_timeout` - Graceful restart neighbor resynchronization timeout.
* `md5_keys` - MD5 key. The structure of `md5_keys` block is documented below.

The `md5_keys` block contains:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.

The `network` block contains:

* `id` - Network entry ID.
* `prefix` - Prefix.
* `area` - Attach the network to area.
* `comments` - Comment.

The `neighbor` block contains:

* `id` - Neighbor entry ID.
* `ip` - Interface IP address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - Priority.

The `passive_interface` block contains:

* `name` - Passive interface name.

The `summary_address` block contains:

* `id` - Summary address entry ID.
* `prefix` - Prefix.
* `tag` - Tag value.
* `advertise` - Enable/disable advertise status.

The `distribute_list` block contains:

* `id` - Distribute list entry ID.
* `access_list` - Access list name.
* `protocol` - Protocol type.

The `redistribute` block contains:

* `name` - Redistribute name.
* `status` - status
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `metric_type` - Metric type.
* `tag` - Tag value.

