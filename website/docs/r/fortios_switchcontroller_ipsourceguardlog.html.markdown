---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_ipsourceguardlog"
description: |-
  Configure FortiSwitch ip source guard log.
---

# fortios_switchcontroller_ipsourceguardlog
Configure FortiSwitch ip source guard log. Applies to FortiOS Version `>= 7.6.4`.

## Argument Reference

The following arguments are supported:

* `log_violations` - Enable/Disable log violations for IP source guard logging. Valid values: `enable`, `disable`.
* `violation_timer` - IP source gurad log violation timer in seconds (0 - 1500, default = 0).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController IpSourceGuardLog can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_ipsourceguardlog.labelname SwitchControllerIpSourceGuardLog

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_ipsourceguardlog.labelname SwitchControllerIpSourceGuardLog
$ unset "FORTIOS_IMPORT_TABLE"
```
