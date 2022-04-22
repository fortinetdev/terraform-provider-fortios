---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_macsyncsettings"
description: |-
  Configure global MAC synchronization settings.
---

# fortios_switchcontroller_macsyncsettings
Configure global MAC synchronization settings. Applies to FortiOS Version `<= 6.2.0`.

## Argument Reference

The following arguments are supported:

* `mac_sync_interval` - Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController MacSyncSettings can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_macsyncsettings.labelname SwitchControllerMacSyncSettings

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_macsyncsettings.labelname SwitchControllerMacSyncSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
