---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_prefixlist6"
description: |-
  Configure IPv6 prefix lists.
---

# fortios_router_prefixlist6
Configure IPv6 prefix lists.

## Example Usage

```hcl
resource "fortios_router_prefixlist6" "trname" {
  name = "pr2"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Comment.
* `rule` - IPv6 prefix list rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `rule` block supports:

* `id` - Rule ID.
* `action` - Permit or deny packets that match this rule. Valid values: `permit`, `deny`.
* `prefix6` - IPv6 prefix to define regular filter criteria, such as "any" or subnets.
* `ge` - Minimum prefix length to be matched (0 - 128).
* `le` - Maximum prefix length to be matched (0 - 128).
* `flags` - Flags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router PrefixList6 can be imported using any of these accepted formats:
```
$ terraform import fortios_router_prefixlist6.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_prefixlist6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
