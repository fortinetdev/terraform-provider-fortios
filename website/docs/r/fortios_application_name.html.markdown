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
* `parameter` - Application parameter name.
* `metadata` - Meta data. The structure of `metadata` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_application_name.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
