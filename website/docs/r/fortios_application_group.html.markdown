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
* `application` - Application ID list.
* `category` - Application category ID list.

The `application` block supports:

* `id` - Application IDs.

The `category` block supports:

* `id` - Category IDs.


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
