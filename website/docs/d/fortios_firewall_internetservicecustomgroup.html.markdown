---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicecustomgroup"
description: |-
  Get information on an fortios firewall internetservicecustomgroup.
---

# Data Source: fortios_firewall_internetservicecustomgroup
Use this data source to get information on an fortios firewall internetservicecustomgroup

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall internetservicecustomgroup.

## Attribute Reference

The following attributes are exported:

* `name` - Custom Internet Service group name.
* `comment` - Comment.
* `member` - Custom Internet Service group members. The structure of `member` block is documented below.

The `member` block contains:

* `name` - Group member name.

