---
subcategory: "FortiGate Sctp-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_sctpfilter_profile"
description: |-
  Configure SCTP filter profiles.
---

# fortios_sctpfilter_profile
Configure SCTP filter profiles. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `comment` - Comment.
* `ppid_filters` - PPID filters list. The structure of `ppid_filters` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ppid_filters` block supports:

* `id` - ID.
* `ppid` - Payload protocol identifier.
* `action` - Action taken when PPID is matched. Valid values: `pass`, `reset`, `replace`.
* `comment` - Comment.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SctpFilter Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_sctpfilter_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_sctpfilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
