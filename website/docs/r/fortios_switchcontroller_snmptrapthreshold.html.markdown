---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_snmptrapthreshold"
description: |-
  Configure FortiSwitch SNMP trap threshold values globally.
---

# fortios_switchcontroller_snmptrapthreshold
Configure FortiSwitch SNMP trap threshold values globally. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController SnmpTrapThreshold can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_snmptrapthreshold.labelname SwitchControllerSnmpTrapThreshold

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_snmptrapthreshold.labelname SwitchControllerSnmpTrapThreshold
$ unset "FORTIOS_IMPORT_TABLE"
```
