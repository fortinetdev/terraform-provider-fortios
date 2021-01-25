---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_aspathlist"
description: |-
  Get information on an fortios router aspathlist.
---

# Data Source: fortios_router_aspathlist
Use this data source to get information on an fortios router aspathlist

## Argument Reference

* `name` - (Required) Specify the name of the desired router aspathlist.

## Attribute Reference

The following attributes are exported:

* `name` - AS path list name.
* `rule` - AS path list rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - ID.
* `action` - Permit or deny route-based operations, based on the route's AS_PATH attribute.
* `regexp` - Regular-expression to match the Border Gateway Protocol (BGP) AS paths.

