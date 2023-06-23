---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_name"
description: |-
  Configure application signatures.
---

# fortios_application_name
Configure application signatures.

## Argument Reference

The following arguments are supported:

* `name` - Application name.
* `fosid` - Application ID.
* `category` - (Required) Application category ID.
* `sub_category` - Application sub-category ID.
* `popularity` - Application popularity.
* `risk` - Application risk.
* `weight` - Application weight.
* `protocol` - Application protocol.
* `technology` - Application technology.
* `behavior` - Application behavior.
* `vendor` - Application vendor.
* `parameters` - Application parameters. The structure of `parameters` block is documented below.
* `parameter` - Application parameter name.
* `metadata` - Meta data. The structure of `metadata` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `parameters` block supports:

* `name` - Parameter name.

The `metadata` block supports:

* `id` - ID.
* `metaid` - Meta ID.
* `valueid` - Value ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Application Name can be imported using any of these accepted formats:
```
$ terraform import fortios_application_name.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_application_name.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
