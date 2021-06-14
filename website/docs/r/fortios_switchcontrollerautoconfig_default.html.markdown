---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerautoconfig_default"
description: |-
  Configure default auto-config QoS policy for FortiSwitch.
---

# fortios_switchcontrollerautoconfig_default
Configure default auto-config QoS policy for FortiSwitch.

## Argument Reference

The following arguments are supported:

* `fgt_policy` - Default FortiLink auto-config policy.
* `isl_policy` - Default ISL auto-config policy.
* `icl_policy` - Default ICL auto-config policy.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchControllerAutoConfig Default can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerautoconfig_default.labelname SwitchControllerAutoConfigDefault
$ unset "FORTIOS_IMPORT_TABLE"
```
