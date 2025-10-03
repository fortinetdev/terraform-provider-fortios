---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_deviceupgradeexemptions"
description: |-
  Configure device upgrade exemptions. Device will stop receiving upgrade notifications on the GUI.
---

# fortios_system_deviceupgradeexemptions
Configure device upgrade exemptions. Device will stop receiving upgrade notifications on the GUI. Applies to FortiOS Version `>= 7.6.4`.

## Argument Reference

The following arguments are supported:

* `fosid` - Device upgrade exemption ID (0 - 65535).
* `fortinet_device` - Fortinet extension device (FortiAP, FortiSwitch, FortiExtender).
* `version` - Highest version of Fortinet firmware to exempt (format in Major.minor.patch, such as 7.6.4).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System DeviceUpgradeExemptions can be imported using any of these accepted formats:
```
$ terraform import fortios_system_deviceupgradeexemptions.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_deviceupgradeexemptions.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
