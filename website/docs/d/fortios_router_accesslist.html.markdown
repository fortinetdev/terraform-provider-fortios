---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_accesslist"
description: |-
  Get information on an fortios router accesslist.
---

# Data Source: fortios_router_accesslist
Use this data source to get information on an fortios router accesslist

## Argument Reference

* `name` - (Required) Specify the name of the desired router accesslist.

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comments` - Comment.
* `rule` - Rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix.
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.
* `wildcard` - Wildcard to define Cisco-style wildcard filter criteria.
* `exact_match` - Enable/disable exact match.
* `flags` - Flags.

