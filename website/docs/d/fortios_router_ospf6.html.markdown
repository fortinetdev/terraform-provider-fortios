---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf6"
description: |-
  Get information on fortios router ospf6.
---

# Data Source: fortios_router_ospf6
Use this data source to get information on fortios router ospf6

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `abr_type` - Area border router type.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `default_information_originate` - Enable/disable generation of default route.
* `log_neighbour_changes` - Enable logging of OSPFv3 neighbour's changes
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type.
* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `router_id` - A.B.C.D, in IPv4 address format.
* `spf_timers` - SPF calculation frequency.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).
* `area` - OSPF6 area configuration. The structure of `area` block is documented below.
* `ospf6_interface` - OSPF6 interface configuration. The structure of `ospf6_interface` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `summary_address` - IPv6 address summary configuration. The structure of `summary_address` block is documented below.

The `area` block contains:

* `id` - Area entry IP address.
* `default_cost` - Summary default cost of stub or NSSA area.
* `nssa_translator_role` - NSSA translator role type.
* `stub_type` - Stub summary setting.
* `type` - Area type setting.
* `nssa_default_information_originate` - Enable/disable originate type 7 default into NSSA area.
* `nssa_default_information_originate_metric` - OSPFv3 default metric.
* `nssa_default_information_originate_metric_type` - OSPFv3 metric type for default routes.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area.
* `authentication` - Authentication mode.
* `key_rollover_interval` - Key roll-over interval.
* `ipsec_auth_alg` - Authentication algorithm.
* `ipsec_enc_alg` - Encryption algorithm.
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
* `range` - OSPF6 area range configuration. The structure of `range` block is documented below.
* `virtual_link` - OSPF6 virtual link configuration. The structure of `virtual_link` block is documented below.

The `ipsec_keys` block contains:

* `spi` - Security Parameters Index.
* `auth_key` - Authentication key.
* `enc_key` - Encryption key.

The `range` block contains:

* `id` - Range entry ID.
* `prefix6` - IPv6 prefix.
* `advertise` - Enable/disable advertise status.

The `virtual_link` block contains:

* `name` - Virtual link entry name.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `peer` - A.B.C.D, peer router ID.
* `authentication` - Authentication mode.
* `key_rollover_interval` - Key roll-over interval.
* `ipsec_auth_alg` - Authentication algorithm.
* `ipsec_enc_alg` - Encryption algorithm.
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `spi` - Security Parameters Index.
* `auth_key` - Authentication key.
* `enc_key` - Encryption key.

The `ospf6_interface` block contains:

* `name` - Interface entry name.
* `area_id` - A.B.C.D, in IPv4 address format.
* `interface` - Configuration interface name.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - priority
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `status` - Enable/disable OSPF6 routing on this interface.
* `network_type` - Network type.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).
* `mtu` - MTU for OSPFv3 packets.
* `mtu_ignore` - Enable/disable ignoring MTU field in DBD packets.
* `authentication` - Authentication mode.
* `key_rollover_interval` - Key roll-over interval.
* `ipsec_auth_alg` - Authentication algorithm.
* `ipsec_enc_alg` - Encryption algorithm.
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
* `neighbor` - OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.

The `ipsec_keys` block contains:

* `spi` - Security Parameters Index.
* `auth_key` - Authentication key.
* `enc_key` - Encryption key.

The `neighbor` block contains:

* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - priority

The `redistribute` block contains:

* `name` - Redistribute name.
* `status` - status
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `metric_type` - Metric type.

The `passive_interface` block contains:

* `name` - Passive interface name.

The `summary_address` block contains:

* `id` - Summary address entry ID.
* `prefix6` - IPv6 prefix.
* `advertise` - Enable/disable advertise status.
* `tag` - Tag value.

