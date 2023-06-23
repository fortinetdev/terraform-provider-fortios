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
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Onetime schedule name.
* `start` - Schedule start date and time, format hh:mm yyyy/mm/dd.
* `start_utc` - Schedule start date and time, in epoch format.
* `end` - Schedule end date and time, format hh:mm yyyy/mm/dd.
* `end_utc` - Schedule end date and time, in epoch format.
* `color` - Color of icon on the GUI.
* `expiration_days` - Write an event log message this many days before the schedule expires.
* `fabric_object` - Security Fabric global object setting.

