---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_label"
description: |-
  Configure labels used by DLP blocking.
---

# fortios_dlp_label
Configure labels used by DLP blocking. Applies to FortiOS Version `>= 7.6.3`.

## Argument Reference

The following arguments are supported:

* `name` - Name of table containing the label.
* `type` - Label type.
* `mpip_type` - MPIP label type. Valid values: `remote`, `local`.
* `connector` - Name of SDN connector.
* `comment` - Optional comments.
* `entries` - DLP label entries. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `id` - ID.
* `fortidata_label_name` - Name of FortiData label
* `mpip_label_name` - Name of MPIP label.
* `guid` - MPIP label guid.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Label can be imported using any of these accepted formats:
```
$ terraform import fortios_dlp_label.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_dlp_label.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
