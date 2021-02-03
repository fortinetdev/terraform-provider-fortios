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

## Attribute Reference

The following attributes are exported:

* `name` - Schedule group name.
* `member` - Schedules added to the schedule group. The structure of `member` block is documented below.
* `color` - Color of icon on the GUI.

The `member` block contains:

* `name` - Schedule name.

