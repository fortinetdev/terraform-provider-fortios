---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicegroup"
description: |-
  Get information on an fortios firewall internetservicegroup.
---

# Data Source: fortios_firewall_internetservicegroup
Use this data source to get information on an fortios firewall internetservicegroup

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall internetservicegroup.

## Attribute Reference

The following attributes are exported:

* `name` - Internet Service group name.
* `comment` - Comment.
* `direction` - How this service may be used (source, destination or both).
* `member` - Internet Service group member. The structure of `member` block is documented below.

The `member` block contains:

* `name` - Internet Service name.
* `id` - Internet Service ID.

