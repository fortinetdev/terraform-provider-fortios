---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_communitylist"
description: |-
  Configure community lists.
---

# fortios_router_communitylist
Configure community lists.

## Example Usage

```hcl
resource "fortios_router_communitylist" "trname" {
  name = "1"
  type = "standard"

  rule {
    action = "deny"
    match  = "123:234 345:456"
    regexp = "123:234 345:456"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Community list name.
* `type` - (Required) Community list type (standard or expanded). Valid values: `standard`, `expanded`.
* `rule` - Community list rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `id` - ID.
* `action` - Permit or deny route-based operations, based on the route's COMMUNITY attribute. Valid values: `deny`, `permit`.
* `regexp` - Ordered list of COMMUNITY attributes as a regular expression.
* `match` - Community specifications for matching a reserved community.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router CommunityList can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_communitylist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
