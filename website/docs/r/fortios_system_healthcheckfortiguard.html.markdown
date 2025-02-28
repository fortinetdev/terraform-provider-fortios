---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_healthcheckfortiguard"
description: |-
  SD-WAN status checking or health checking. Identify a server predefine by FortiGuard and determine how SD-WAN verifies that FGT can communicate with it.
---

# fortios_system_healthcheckfortiguard
SD-WAN status checking or health checking. Identify a server predefine by FortiGuard and determine how SD-WAN verifies that FGT can communicate with it. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `name` - Status check or predefined health-check targets name.
* `server` - Status check or predefined health-check domain name.
* `obsolete` - Indicates whether Health Check service can be used.
* `protocol` - Protocol name. Valid values: `ping`, `tcp-echo`, `udp-echo`, `http`, `https`, `twamp`, `dns`, `tcp-connect`, `ftp`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System HealthCheckFortiguard can be imported using any of these accepted formats:
```
$ terraform import fortios_system_healthcheckfortiguard.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_healthcheckfortiguard.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
