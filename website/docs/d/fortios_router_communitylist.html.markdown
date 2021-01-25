---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_communitylist"
description: |-
  Get information on an fortios router communitylist.
---

# Data Source: fortios_router_communitylist
Use this data source to get information on an fortios router communitylist

## Argument Reference

* `name` - (Required) Specify the name of the desired router communitylist.

## Attribute Reference

The following attributes are exported:

* `name` - Community list name.
* `type` - Community list type (standard or expanded).
* `rule` - Community list rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - ID.
* `action` - Permit or deny route-based operations, based on the route's COMMUNITY attribute.
* `regexp` - Ordered list of COMMUNITY attributes as a regular expression.
* `match` - Community specifications for matching a reserved community.

