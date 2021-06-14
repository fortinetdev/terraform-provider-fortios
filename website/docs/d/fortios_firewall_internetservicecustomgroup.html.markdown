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
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Custom Internet Service group name.
* `comment` - Comment.
* `member` - Custom Internet Service group members. The structure of `member` block is documented below.

The `member` block contains:

* `name` - Group member name.

