---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_system"
description: |-
  Configure system-wide switch controller settings.
---

# fortios_switchcontroller_system
Configure system-wide switch controller settings.

## Argument Reference

The following arguments are supported:

* `parallel_process_override` - Enable/disable parallel process override. Valid values: `disable`, `enable`.
* `parallel_process` - Maximum number of parallel processes (1 - 300, default = 1).
* `data_sync_interval` - Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
* `iot_weight_threshold` - MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
* `iot_scan_interval` - IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
* `iot_holdoff` - MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
* `iot_mac_idle` - MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
* `nac_periodic_interval` - Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController System can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_system.labelname SwitchControllerSystem
$ unset "FORTIOS_IMPORT_TABLE"
```
