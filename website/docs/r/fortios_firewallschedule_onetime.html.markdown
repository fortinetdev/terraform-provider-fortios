---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_onetime"
description: |-
  Onetime schedule configuration.
---

# fortios_firewallschedule_onetime
Onetime schedule configuration.

## Example Usage

```hcl
resource "fortios_firewallschedule_onetime" "trname" {
  color           = 0
  end             = "00:00 2020/12/12"
  expiration_days = 2
  name            = "onetime1"
  start           = "00:00 2010/12/12"
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) Onetime schedule name.
* `start` - (Required) Schedule start date and time, format hh:mm yyyy/mm/dd.
* `end` - (Required) Schedule end date and time, format hh:mm yyyy/mm/dd.
* `color` - Color of icon on the GUI.
* `expiration_days` - Write an event log message this many days before the schedule expires.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallSchedule Onetime can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallschedule_onetime.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
