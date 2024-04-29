---
subcategory: "FortiGate Videofilter"
layout: "fortios"
page_title: "FortiOS: fortios_videofilter_keyword"
description: |-
  Configure video filter keywords.
---

# fortios_videofilter_keyword
Configure video filter keywords. Applies to FortiOS Version `>= 7.4.2`.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `name` - Name.
* `comment` - Comment.
* `match` - Keyword matching logic. Valid values: `or`, `and`.
* `word` - List of keywords. The structure of `word` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `word` block supports:

* `name` - Name.
* `comment` - Comment.
* `pattern_type` - Pattern type. Valid values: `wildcard`, `regex`.
* `status` - Enable(consider)/disable(ignore) this keyword. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Videofilter Keyword can be imported using any of these accepted formats:
```
$ terraform import fortios_videofilter_keyword.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_videofilter_keyword.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
