---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_stormcontrolpolicy"
description: |-
  Configure FortiSwitch storm control policy to be applied on managed-switch ports.
---

# fortios_switchcontroller_stormcontrolpolicy
Configure FortiSwitch storm control policy to be applied on managed-switch ports.

## Argument Reference

The following arguments are supported:

* `name` - Storm control policy name.
* `description` - Description of the storm control policy.
* `storm_control_mode` - Set Storm control mode.
* `rate` - Threshold rate in packets per second at which storm traffic is controlled in override mode (default=500, 0 to drop all).
* `unknown_unicast` - Enable/disable storm control to drop/allow unknown unicast traffic in override mode.
* `unknown_multicast` - Enable/disable storm control to drop/allow unknown multicast traffic in override mode.
* `broadcast` - Enable/disable storm control to drop/allow broadcast traffic in override mode.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController StormControlPolicy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_stormcontrolpolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
