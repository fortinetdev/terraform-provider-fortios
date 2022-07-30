---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_group"
description: |-
  Schedule group configuration.
---

# fortios_firewallschedule_group
Schedule group configuration.

## Example Usage

```hcl
resource "fortios_firewallschedule_recurring" "trname1" {
  color = 0
  day   = "sunday"
  end   = "00:00"
  name  = "recurring2"
  start = "00:00"
}

resource "fortios_firewallschedule_group" "trname" {
  color = 0
  name  = "groupg1"

  member {
    name = fortios_firewallschedule_recurring.trname1.name
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Schedule group name.
* `member` - (Required) Schedules added to the schedule group. The structure of `member` block is documented below.
* `color` - Color of icon on the GUI.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Schedule name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallSchedule Group can be imported using any of these accepted formats:
```
$ terraform import fortios_firewallschedule_group.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallschedule_group.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
