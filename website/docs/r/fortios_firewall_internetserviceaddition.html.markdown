---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceaddition"
description: |-
  Configure Internet Services Addition.
---

# fortios_firewall_internetserviceaddition
Configure Internet Services Addition. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `fosid` - Internet Service ID in the Internet Service database.
* `comment` - Comment.
* `entry` - Entries added to the Internet Service addition database. The structure of `entry` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entry` block supports:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the custom entry. The structure of `port_range` block is documented below.

The `port_range` block supports:

* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceAddition can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_internetserviceaddition.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_internetserviceaddition.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
