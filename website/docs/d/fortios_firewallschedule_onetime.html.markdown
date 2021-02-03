---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_onetime"
description: |-
  Get information on an fortios firewallschedule onetime.
---

# Data Source: fortios_firewallschedule_onetime
Use this data source to get information on an fortios firewallschedule onetime

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallschedule onetime.

## Attribute Reference

The following attributes are exported:

* `name` - Onetime schedule name.
* `start` - Schedule start date and time, format hh:mm yyyy/mm/dd.
* `end` - Schedule end date and time, format hh:mm yyyy/mm/dd.
* `color` - Color of icon on the GUI.
* `expiration_days` - Write an event log message this many days before the schedule expires.

