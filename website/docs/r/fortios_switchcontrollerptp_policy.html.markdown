---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerptp_policy"
description: |-
  PTP policy configuration.
---

# fortios_switchcontrollerptp_policy
PTP policy configuration. Applies to FortiOS Version `>= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `name` - Policy name.
* `status` - Enable/disable PTP policy. Valid values: `disable`, `enable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerPtp Policy can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrollerptp_policy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollerptp_policy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
