---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollersecuritypolicy_captiveportal"
description: |-
  Names of VLANs that use captive portal authentication.
---

# fortios_switchcontrollersecuritypolicy_captiveportal
Names of VLANs that use captive portal authentication. Applies to FortiOS Version `<= 6.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - Policy name.
* `vlan` - Names of VLANs that use captive portal authentication.
* `policy_type` - Policy type. Valid values: `captive-portal`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerSecurityPolicy CaptivePortal can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollersecuritypolicy_captiveportal.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
