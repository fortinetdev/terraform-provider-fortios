---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_prefixlist"
description: |-
  Get information on an fortios router prefixlist.
---

# Data Source: fortios_router_prefixlist
Use this data source to get information on an fortios router prefixlist

## Argument Reference

* `name` - (Required) Specify the name of the desired router prefixlist.

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comments` - Comment.
* `rule` - IPv4 prefix list rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix.
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.
* `ge` - Minimum prefix length to be matched (0 - 32).
* `le` - Maximum prefix length to be matched (0 - 32).
* `flags` - Flags.

