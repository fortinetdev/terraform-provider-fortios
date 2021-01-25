---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_group"
description: |-
  Configure firewall application groups.
---

# fortios_application_group
Configure firewall application groups.

## Example Usage

```hcl
resource "fortios_application_group" "trname" {
  comment = "group1 test"
  name    = "1"
  type    = "category"
  category {
    id = 2
  }
}
```

## Argument Reference


The following arguments are supported:

* `name` - Application group name.
* `comment` - Comment
* `type` - Application group type.
* `application` - Application ID list. The structure of `application` block is documented below.
* `category` - Application category ID list. The structure of `category` block is documented below.

The `application` block supports:

* `id` - Application IDs.

The `category` block supports:

* `id` - Category IDs.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Application Group can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_application_group.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
