---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_snmptrapthreshold"
description: |-
  Configure FortiSwitch SNMP trap threshold values globally.
---

# fortios_switchcontroller_snmptrapthreshold
Configure FortiSwitch SNMP trap threshold values globally.

## Argument Reference

The following arguments are supported:

* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController SnmpTrapThreshold can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_snmptrapthreshold.labelname SwitchControllerSnmpTrapThreshold
$ unset "FORTIOS_IMPORT_TABLE"
```
