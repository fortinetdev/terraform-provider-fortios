---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_evpn"
description: |-
  Configure EVPN instance.
---

# fortios_system_evpn
Configure EVPN instance. Applies to FortiOS Version `>= 7.4.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `rd` - Route Distinguisher: AA|AA:NN|A.B.C.D:NN.
* `import_rt` - List of import route targets. The structure of `import_rt` block is documented below.
* `export_rt` - List of export route targets. The structure of `export_rt` block is documented below.
* `ip_local_learning` - Enable/disable IP address local learning. Valid values: `enable`, `disable`.
* `arp_suppression` - Enable/disable ARP suppression. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `import_rt` block supports:

* `route_target` - Route target: AA|AA:NN.

The `export_rt` block supports:

* `route_target` - Route target: AA|AA:NN.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Evpn can be imported using any of these accepted formats:
```
$ terraform import fortios_system_evpn.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_evpn.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
