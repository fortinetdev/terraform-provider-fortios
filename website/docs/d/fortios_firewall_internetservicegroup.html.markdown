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
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Internet Service group name.
* `comment` - Comment.
* `direction` - How this service may be used (source, destination or both).
* `member` - Internet Service group member. The structure of `member` block is documented below.

The `member` block contains:

* `name` - Internet Service name.
* `id` - Internet Service ID.

