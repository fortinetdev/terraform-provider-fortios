---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_accesslist6"
description: |-
  Configure IPv6 access lists.
---

# fortios_router_accesslist6
Configure IPv6 access lists.

## Example Usage

```hcl
resource "fortios_router_accesslist6" "trname" {
  comments = "access-list6 test"
  name     = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Comment.
* `rule` - Rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix.
* `prefix6` - IPv6 prefix to define regular filter criteria, such as "any" or subnets.
* `exact_match` - Enable/disable exact prefix match.
* `flags` - Flags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router AccessList6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_accesslist6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
