---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_group"
description: |-
  Get information on an fortios firewallservice group.
---

# Data Source: fortios_firewallservice_group
Use this data source to get information on an fortios firewallservice group

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallservice group.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Address group name.
* `member` - Service objects contained within the group. The structure of `member` block is documented below.
* `proxy` - Enable/disable web proxy service group.
* `comment` - Comment.
* `color` - Color of icon on the GUI.
* `fabric_object` - Security Fabric global object setting.

The `member` block contains:

* `name` - Address name.

