---
subcategory: "FortiGate Casb"
layout: "fortios"
page_title: "FortiOS: fortios_casb_saasapplication"
description: |-
  Configure CASB SaaS application.
---

# fortios_casb_saasapplication
Configure CASB SaaS application. Applies to FortiOS Version `>= 7.4.1`.

## Argument Reference

The following arguments are supported:

* `name` - SaaS application name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `status` - Enable/disable setting. Valid values: `enable`, `disable`.
* `type` - SaaS application type. Valid values: `built-in`, `customized`.
* `casb_name` - SaaS application signature name.
* `description` - SaaS application description.
* `domains` - SaaS application domain list. The structure of `domains` block is documented below.
* `output_attributes` - SaaS application output attributes. The structure of `output_attributes` block is documented below.
* `input_attributes` - SaaS application input attributes. The structure of `input_attributes` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `domains` block supports:

* `domain` - Domain list separated by space.

The `output_attributes` block supports:

* `name` - CASB attribute name.
* `attr_type` - CASB attribute type. Valid values: `tenant`.
* `description` - CASB attribute description.
* `type` - CASB attribute format type. Valid values: `string`, `string-list`, `integer`, `integer-list`, `boolean`.
* `required` - CASB attribute required. Valid values: `enable`, `disable`.

The `input_attributes` block supports:

* `name` - CASB attribute name.
* `attr_type` - CASB attribute type. Valid values: `tenant`.
* `description` - CASB attribute description.
* `type` - CASB attribute format type. Valid values: `string`, `string-list`, `integer`, `integer-list`, `boolean`.
* `required` - CASB attribute required. Valid values: `enable`, `disable`.
* `default` - CASB attribute default value. Valid values: `string`, `string-list`.
* `fallback_input` - CASB attribute legacy input. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb SaasApplication can be imported using any of these accepted formats:
```
$ terraform import fortios_casb_saasapplication.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_casb_saasapplication.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
