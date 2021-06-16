---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_group"
description: |-
  Get information on an fortios firewallschedule group.
---

# Data Source: fortios_firewallschedule_group
Use this data source to get information on an fortios firewallschedule group

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallschedule group.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Schedule group name.
* `member` - Schedules added to the schedule group. The structure of `member` block is documented below.
* `color` - Color of icon on the GUI.

The `member` block contains:

* `name` - Schedule name.

