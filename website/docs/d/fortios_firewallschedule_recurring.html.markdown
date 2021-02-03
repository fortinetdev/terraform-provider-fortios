---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_recurring"
description: |-
  Get information on an fortios firewallschedule recurring.
---

# Data Source: fortios_firewallschedule_recurring
Use this data source to get information on an fortios firewallschedule recurring

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallschedule recurring.

## Attribute Reference

The following attributes are exported:

* `name` - Recurring schedule name.
* `start` - Time of day to start the schedule, format hh:mm.
* `end` - Time of day to end the schedule, format hh:mm.
* `day` - One or more days of the week on which the schedule is valid. Separate the names of the days with a space.
* `color` - Color of icon on the GUI.

