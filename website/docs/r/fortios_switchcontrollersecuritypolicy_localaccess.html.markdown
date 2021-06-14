---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollersecuritypolicy_localaccess"
description: |-
  Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch.
---

# fortios_switchcontrollersecuritypolicy_localaccess
Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch.

## Argument Reference

The following arguments are supported:

* `name` - Policy name.
* `mgmt_allowaccess` - Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
* `internal_allowaccess` - Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerSecurityPolicy LocalAccess can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollersecuritypolicy_localaccess.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
