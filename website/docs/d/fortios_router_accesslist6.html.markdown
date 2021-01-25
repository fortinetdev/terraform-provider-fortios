---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_accesslist6"
description: |-
  Get information on an fortios router accesslist6.
---

# Data Source: fortios_router_accesslist6
Use this data source to get information on an fortios router accesslist6

## Argument Reference

* `name` - (Required) Specify the name of the desired router accesslist6.

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comments` - Comment.
* `rule` - Rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix.
* `prefix6` - IPv6 prefix to define regular filter criteria, such as "any" or subnets.
* `exact_match` - Enable/disable exact prefix match.
* `flags` - Flags.

