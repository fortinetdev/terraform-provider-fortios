---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_prefixlist6"
description: |-
  Get information on an fortios router prefixlist6.
---

# Data Source: fortios_router_prefixlist6
Use this data source to get information on an fortios router prefixlist6

## Argument Reference

* `name` - (Required) Specify the name of the desired router prefixlist6.

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comments` - Comment.
* `rule` - IPv6 prefix list rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - Rule ID.
* `action` - Permit or deny packets that match this rule.
* `prefix6` - IPv6 prefix to define regular filter criteria, such as "any" or subnets.
* `ge` - Minimum prefix length to be matched (0 - 128).
* `le` - Maximum prefix length to be matched (0 - 128).
* `flags` - Flags.

