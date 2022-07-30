---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_routerospf6_ospf6interface"
description: |-
  OSPF6 interface configuration.
---

# fortios_routerospf6_ospf6interface
OSPF6 interface configuration.

~> The provider supports the definition of Ospf6-Interface in Router Ospf6 `fortios_router_ospf6`, and also allows the definition of separate Ospf6-Interface resources `fortios_routerospf6_ospf6interface`, but do not use a `fortios_router_ospf6` with in-line Ospf6-Interface in conjunction with any `fortios_routerospf6_ospf6interface` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `name` - Interface entry name.
* `area_id` - A.B.C.D, in IPv4 address format.
* `interface` - Configuration interface name.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - priority
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `status` - Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
* `network_type` - Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
* `mtu` - MTU for OSPFv3 packets.
* `mtu_ignore` - Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
* `authentication` - Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
* `key_rollover_interval` - Key roll-over interval.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
* `neighbor` - OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ipsec_keys` block supports:

* `spi` - Security Parameters Index.
* `auth_key` - Authentication key.
* `enc_key` - Encryption key.

The `neighbor` block supports:

* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - priority


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Routerospf6 Ospf6Interface can be imported using any of these accepted formats:
```
$ terraform import fortios_routerospf6_ospf6interface.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_routerospf6_ospf6interface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
