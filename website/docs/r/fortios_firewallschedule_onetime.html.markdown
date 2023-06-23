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
* `start_utc` - Schedule start date and time, in epoch format.
* `end` - (Required) Schedule end date and time, format hh:mm yyyy/mm/dd.
* `end_utc` - Schedule end date and time, in epoch format.
* `color` - Color of icon on the GUI.
* `expiration_days` - Write an event log message this many days before the schedule expires.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallSchedule Onetime can be imported using any of these accepted formats:
```
$ terraform import fortios_firewallschedule_onetime.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallschedule_onetime.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
