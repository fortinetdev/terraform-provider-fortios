---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_npupost"
description: |-
  Configure NPU attributes after interface initialization.
---

# fortios_system_npupost
Configure NPU attributes after interface initialization. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `port_npu_map` - Configure port to NPU group list. The structure of `port_npu_map` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `port_npu_map` block supports:

* `interface` - Set NPU interface port for NPU group mapping.
* `npu_group` - Mapping NPU group list. The structure of `npu_group` block is documented below.

The `npu_group` block supports:

* `group_name` - NPU group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System NpuPost can be imported using any of these accepted formats:
```
$ terraform import fortios_system_npupost.labelname SystemNpuPost

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_npupost.labelname SystemNpuPost
$ unset "FORTIOS_IMPORT_TABLE"
```
