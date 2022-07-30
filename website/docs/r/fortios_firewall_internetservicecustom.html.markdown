---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicecustom"
description: |-
  Configure custom Internet Services.
---

# fortios_firewall_internetservicecustom
Configure custom Internet Services.

## Argument Reference

The following arguments are supported:

* `name` - Internet Service name.
* `reputation` - Reputation level of the custom Internet Service.
* `comment` - Comment.
* `entry` - Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entry` block supports:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the custom entry. The structure of `port_range` block is documented below.
* `dst` - Destination address or address group name. The structure of `dst` block is documented below.

The `port_range` block supports:

* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).

The `dst` block supports:

* `name` - Select the destination address or address group object from available options.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall InternetServiceCustom can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_internetservicecustom.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_internetservicecustom.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
