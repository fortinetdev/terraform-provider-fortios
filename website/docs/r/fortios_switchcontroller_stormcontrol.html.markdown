---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_stormcontrol"
description: |-
  Configure FortiSwitch storm control.
---

# fortios_switchcontroller_stormcontrol
Configure FortiSwitch storm control.

## Argument Reference

The following arguments are supported:

* `rate` - Rate in packets per second at which storm traffic is controlled (1 - 10000000, default = 500). Storm control drops excess traffic data rates beyond this threshold.
* `unknown_unicast` - Enable/disable storm control to drop unknown unicast traffic. Valid values: `enable`, `disable`.
* `unknown_multicast` - Enable/disable storm control to drop unknown multicast traffic. Valid values: `enable`, `disable`.
* `broadcast` - Enable/disable storm control to drop broadcast traffic. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController StormControl can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_stormcontrol.labelname SwitchControllerStormControl
$ unset "FORTIOS_IMPORT_TABLE"
```
