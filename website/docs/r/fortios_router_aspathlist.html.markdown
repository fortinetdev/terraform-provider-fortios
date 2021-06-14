---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_aspathlist"
description: |-
  Configure Autonomous System (AS) path lists.
---

# fortios_router_aspathlist
Configure Autonomous System (AS) path lists.

## Example Usage

```hcl
resource "fortios_router_aspathlist" "trname" {
  name = "aspath1"

  rule {
    action = "deny"
    regexp = "/d+/n"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) AS path list name.
* `rule` - AS path list rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `rule` block supports:

* `id` - ID.
* `action` - Permit or deny route-based operations, based on the route's AS_PATH attribute. Valid values: `deny`, `permit`.
* `regexp` - Regular-expression to match the Border Gateway Protocol (BGP) AS paths.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router AspathList can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_aspathlist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
