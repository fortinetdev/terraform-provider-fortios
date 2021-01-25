---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_recurring"
description: |-
  Recurring schedule configuration.
---

# fortios_firewallschedule_recurring
Recurring schedule configuration.

## Example Usage

```hcl
resource "fortios_firewallschedule_recurring" "trname" {
  color = 0
  day   = "sunday"
  end   = "00:00"
  name  = "recurring1"
  start = "00:00"
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) Recurring schedule name.
* `start` - (Required) Time of day to start the schedule, format hh:mm.
* `end` - (Required) Time of day to end the schedule, format hh:mm.
* `day` - One or more days of the week on which the schedule is valid. Separate the names of the days with a space.
* `color` - Color of icon on the GUI.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallSchedule Recurring can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallschedule_recurring.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
