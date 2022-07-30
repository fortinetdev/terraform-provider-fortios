---
subcategory: "FortiGate NSX-T"
layout: "fortios"
page_title: "FortiOS: fortios_nsxt_servicechain"
description: |-
  Configure NSX-T service chain.
---

# fortios_nsxt_servicechain
Configure NSX-T service chain. Applies to FortiOS Version `>= 7.0.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - Chain ID.
* `name` - Chain name.
* `service_index` - Configure service index. The structure of `service_index` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `service_index` block supports:

* `id` - Service index.
* `reverse_index` - Reverse service index.
* `name` - Index name.
* `vd` - VDOM name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Nsxt ServiceChain can be imported using any of these accepted formats:
```
$ terraform import fortios_nsxt_servicechain.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_nsxt_servicechain.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
