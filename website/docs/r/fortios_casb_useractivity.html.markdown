---
subcategory: "FortiGate Casb"
layout: "fortios"
page_title: "FortiOS: fortios_casb_useractivity"
description: |-
  Configure CASB user activity.
---

# fortios_casb_useractivity
Configure CASB user activity. Applies to FortiOS Version `>= 7.4.1`.

## Argument Reference

The following arguments are supported:

* `name` - CASB user activity name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `status` - CASB user activity status. Valid values: `enable`, `disable`.
* `description` - CASB user activity description.
* `type` - CASB user activity type. Valid values: `built-in`, `customized`.
* `casb_name` - CASB user activity signature name.
* `application` - CASB SaaS application name.
* `category` - CASB user activity category. Valid values: `activity-control`, `tenant-control`, `domain-control`, `safe-search-control`, `other`.
* `match_strategy` - CASB user activity match strategy. Valid values: `and`, `or`.
* `match` - CASB user activity match rules. The structure of `match` block is documented below.
* `control_options` - CASB control options. The structure of `control_options` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `match` block supports:

* `id` - CASB user activity match rules ID.
* `strategy` - CASB user activity rules strategy. Valid values: `and`, `or`.
* `rules` - CASB user activity rules. The structure of `rules` block is documented below.

The `rules` block supports:

* `id` - CASB user activity rule ID.
* `type` - CASB user activity rule type. Valid values: `domains`, `host`, `path`, `header`, `header-value`, `method`.
* `domains` - CASB user activity domain list. The structure of `domains` block is documented below.
* `methods` - CASB user activity method list. The structure of `methods` block is documented below.
* `match_pattern` - CASB user activity rule match pattern. Valid values: `simple`, `substr`, `regexp`.
* `match_value` - CASB user activity rule match value.
* `header_name` - CASB user activity rule header name.
* `case_sensitive` - CASB user activity match case sensitive. Valid values: `enable`, `disable`.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `enable`, `disable`.

The `domains` block supports:

* `domain` - Domain list separated by space.

The `methods` block supports:

* `method` - User activity method.

The `control_options` block supports:

* `name` - CASB control option name.
* `status` - CASB control option status. Valid values: `enable`, `disable`.
* `operations` - CASB control option operations. The structure of `operations` block is documented below.

The `operations` block supports:

* `name` - CASB control option operation name.
* `target` - CASB operation target. Valid values: `header`, `path`.
* `action` - CASB operation action. Valid values: `append`, `prepend`, `replace`, `new`, `new-on-not-found`, `delete`.
* `direction` - CASB operation direction. Valid values: `request`.
* `header_name` - CASB operation header name to search.
* `search_pattern` - CASB operation search pattern. Valid values: `simple`, `substr`, `regexp`.
* `search_key` - CASB operation key to search.
* `case_sensitive` - CASB operation search case sensitive. Valid values: `enable`, `disable`.
* `value_from_input` - Enable/disable value from user input. Valid values: `enable`, `disable`.
* `values` - CASB operation new values. The structure of `values` block is documented below.

The `values` block supports:

* `value` - Operation value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb UserActivity can be imported using any of these accepted formats:
```
$ terraform import fortios_casb_useractivity.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_casb_useractivity.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
