---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_dictionary"
description: |-
  Configure dictionaries used by DLP blocking.
---

# fortios_dlp_dictionary
Configure dictionaries used by DLP blocking. Applies to FortiOS Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `name` - Name of table containing the dictionary.
* `match_type` - Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.
* `match_around` - Enable/disable match-around support. Valid values: `enable`, `disable`.
* `comment` - Optional comments.
* `entries` - DLP dictionary entries. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `id` - ID.
* `type` - Pattern type to match.
* `pattern` - Pattern to match.
* `ignore_case` - Enable/disable ignore case. Valid values: `enable`, `disable`.
* `repeat` - Enable/disable repeat match. Valid values: `enable`, `disable`.
* `status` - Enable/disable this pattern. Valid values: `enable`, `disable`.
* `comment` - Optional comments.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Dictionary can be imported using any of these accepted formats:
```
$ terraform import fortios_dlp_dictionary.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_dlp_dictionary.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
