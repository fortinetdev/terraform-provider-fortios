---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_macsyncsettings"
description: |-
  Configure global MAC synchronization settings.
---

# fortios_switchcontroller_macsyncsettings
Configure global MAC synchronization settings.

## Argument Reference

The following arguments are supported:

* `mac_sync_interval` - Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController MacSyncSettings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_macsyncsettings.labelname SwitchControllerMacSyncSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
