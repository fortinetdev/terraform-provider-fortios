---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_routerospf_ospfinterface"
description: |-
  OSPF interface configuration.
---

# fortios_routerospf_ospfinterface
OSPF interface configuration.

~> The provider supports the definition of Ospf-Interface in Router Ospf `fortios_router_ospf`, and also allows the definition of separate Ospf-Interface resources `fortios_routerospf_ospfinterface`, but do not use a `fortios_router_ospf` with in-line Ospf-Interface in conjunction with any `fortios_routerospf_ospfinterface` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `name` - Interface entry name.
* `interface` - Configuration interface name.
* `ip` - IP address.
* `authentication` - Authentication type. Valid values: `none`, `text`, `md5`.
* `authentication_key` - Authentication key.
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
* `database_filter_out` - Enable/disable control of flooding out LSAs. Valid values: `enable`, `disable`.
* `mtu` - MTU for database description packets.
* `mtu_ignore` - Enable/disable ignore MTU. Valid values: `enable`, `disable`.
* `network_type` - Network type. Valid values: `broadcast`, `non-broadcast`, `point-to-point`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.
* `resync_timeout` - Graceful restart neighbor resynchronization timeout.
* `md5_keys` - MD5 key. The structure of `md5_keys` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `md5_keys` block supports:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Routerospf OspfInterface can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_routerospf_ospfinterface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
