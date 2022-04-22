---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerautoconfig_policy"
description: |-
  Configure FortiSwitch Auto-Config QoS policy.
---

# fortios_switchcontrollerautoconfig_policy
Configure FortiSwitch Auto-Config QoS policy.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Auto-Config QoS policy name
* `qos_policy` - Auto-Config QoS policy.
* `storm_control_policy` - Auto-Config storm control policy.
* `poe_status` - Enable/disable PoE status. Valid values: `enable`, `disable`.
* `igmp_flood_report` - Enable/disable IGMP flood report. Valid values: `enable`, `disable`.
* `igmp_flood_traffic` - Enable/disable IGMP flood traffic. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerAutoConfig Policy can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrollerautoconfig_policy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollerautoconfig_policy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
