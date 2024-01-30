---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_exactdatamatch"
description: |-
  Configure exact-data-match template used by DLP scan.
---

# fortios_dlp_exactdatamatch
Configure exact-data-match template used by DLP scan. Applies to FortiOS Version `>= 7.4.2`.

## Argument Reference

The following arguments are supported:

* `name` - Name of table containing the exact-data-match template.
* `optional` - Number of optional columns need to match.
* `data` - External resource for exact data match.
* `columns` - DLP exact-data-match column types. The structure of `columns` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `columns` block supports:

* `index` - Column index.
* `type` - Data-type for this column.
* `optional` - Enable/disable optional match. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp ExactDataMatch can be imported using any of these accepted formats:
```
$ terraform import fortios_dlp_exactdatamatch.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_dlp_exactdatamatch.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
