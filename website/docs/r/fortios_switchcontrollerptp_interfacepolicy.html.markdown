---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerptp_interfacepolicy"
description: |-
  PTP interface-policy configuration.
---

# fortios_switchcontrollerptp_interfacepolicy
PTP interface-policy configuration. Applies to FortiOS Version `>= 7.4.1`.

## Argument Reference

The following arguments are supported:

* `name` - Policy name.
* `description` - Description.
* `vlan` - PTP VLAN.
* `vlan_pri` - Configure PTP VLAN priority (0 - 7, default = 4).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerPtp InterfacePolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrollerptp_interfacepolicy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollerptp_interfacepolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
