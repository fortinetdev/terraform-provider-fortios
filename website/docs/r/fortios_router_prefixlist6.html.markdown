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

The `rule` block supports:

* `id` - Rule ID.
* `action` - Permit or deny packets that match this rule.
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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_prefixlist6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
