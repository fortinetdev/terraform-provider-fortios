---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_prefixlist"
description: |-
  Configure IPv4 prefix lists.
---

# fortios_router_prefixlist
Configure IPv4 prefix lists.

## Example Usage

```hcl
resource "fortios_router_prefixlist" "trname" {
  name = "pr1"
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Comment.
* `rule` - IPv4 prefix list rule. The structure of `rule` block is documented below.

The `rule` block supports:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix.
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.
* `ge` - Minimum prefix length to be matched (0 - 32).
* `le` - Maximum prefix length to be matched (0 - 32).
* `flags` - Flags.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router PrefixList can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_prefixlist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
