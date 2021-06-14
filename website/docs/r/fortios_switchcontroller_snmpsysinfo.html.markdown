---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_snmpsysinfo"
description: |-
  Configure FortiSwitch SNMP system information globally.
---

# fortios_switchcontroller_snmpsysinfo
Configure FortiSwitch SNMP system information globally.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable SNMP. Valid values: `disable`, `enable`.
* `engine_id` - Local SNMP engine ID string (max 24 char).
* `description` - System description.
* `contact_info` - Contact information.
* `location` - System location.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController SnmpSysinfo can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_snmpsysinfo.labelname SwitchControllerSnmpSysinfo
$ unset "FORTIOS_IMPORT_TABLE"
```
