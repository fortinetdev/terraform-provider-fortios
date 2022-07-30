---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_region"
description: |-
  Define region table.
---

# fortios_firewall_region
Define region table. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - Region ID.
* `name` - Region name.
* `city` - City ID list. The structure of `city` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `city` block supports:

* `id` - City ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall Region can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_region.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_region.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
