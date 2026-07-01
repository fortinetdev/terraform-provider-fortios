---
subcategory: "FortiGate NSX-T"
layout: "fortios"
page_title: "FortiOS: fortios_nsxt_servicechain"
description: |-
  Configure NSX-T service chain.
---

# fortios_nsxt_servicechain
Configure NSX-T service chain. Applies to FortiOS Version `6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.0.14,7.0.15,7.0.16,7.0.17,7.0.18,7.0.19,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.2.7,7.2.8,7.2.9,7.2.10,7.2.11,7.2.12,7.4.0,7.4.1,7.4.2,7.4.3,7.4.4,7.4.5,7.4.6,7.4.7,7.4.8,7.4.9,7.4.11,7.4.12,7.6.0,7.6.1,7.6.2,7.6.3,7.6.4,7.6.5,7.6.6,7.6.7`.

## Argument Reference

The following arguments are supported:

* `fosid` - Chain ID.
* `name` - Chain name.
* `service_index` - Configure service index. The structure of `service_index` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

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
