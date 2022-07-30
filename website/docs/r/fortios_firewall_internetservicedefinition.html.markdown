---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicedefinition"
description: |-
  Internet Service definition.
---

# fortios_firewall_internetservicedefinition
Internet Service definition.

## Argument Reference

The following arguments are supported:

* `fosid` - Internet Service application list ID.
* `entry` - Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entry` block supports:

* `seq_num` - Entry sequence number.
* `category_id` - Internet Service category ID.
* `name` - Internet Service name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the definition entry. The structure of `port_range` block is documented below.
* `port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535). 0 means undefined.

The `port_range` block supports:

* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (1 to 65535).
* `end_port` - Ending TCP/UDP/SCTP destination port (1 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceDefinition can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_internetservicedefinition.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_internetservicedefinition.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
