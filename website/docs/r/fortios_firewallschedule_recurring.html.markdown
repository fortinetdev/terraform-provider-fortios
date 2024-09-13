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
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `start` - (Required) Time of day to start the schedule, format hh:mm.
* `end` - (Required) Time of day to end the schedule, format hh:mm.
* `day` - One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
* `color` - Color of icon on the GUI.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallSchedule Recurring can be imported using any of these accepted formats:
```
$ terraform import fortios_firewallschedule_recurring.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallschedule_recurring.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
