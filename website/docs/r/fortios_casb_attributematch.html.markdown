---
subcategory: "FortiGate Casb"
layout: "fortios"
page_title: "FortiOS: fortios_casb_attributematch"
description: |-
  Configure CASB SaaS application.
---

# fortios_casb_attributematch
Configure CASB SaaS application. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `name` - CASB tenant match name.
* `application` - CASB tenant application name.
* `match_strategy` - CASB tenant match strategy. Valid values: `and`, `or`.
* `match` - CASB tenant match rules. The structure of `match` block is documented below.
* `attribute` - CASB tenant match rules. The structure of `attribute` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `match` block supports:

* `id` - CASB attribute match rule ID.
* `rule_strategy` - CASB attribute match rule strategy. Valid values: `and`, `or`.
* `rule` - CASB attribute match rule. The structure of `rule` block is documented below.

The `rule` block supports:

* `id` - CASB attribute rule ID.
* `attribute` - CASB attribute match name.
* `match_pattern` - CASB attribute match pattern. Valid values: `simple`, `substr`, `regexp`.
* `match_value` - CASB attribute match value.
* `case_sensitive` - CASB attribute match case sensitive. Valid values: `enable`, `disable`.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `enable`, `disable`.

The `attribute` block supports:

* `name` - CASB attribute match name.
* `match_pattern` - CASB attribute match pattern. Valid values: `simple`, `substr`, `regexp`.
* `match_value` - CASB attribute match value.
* `case_sensitive` - CASB attribute match case sensitive. Valid values: `enable`, `disable`.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb AttributeMatch can be imported using any of these accepted formats:
```
$ terraform import fortios_casb_attributematch.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_casb_attributematch.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
