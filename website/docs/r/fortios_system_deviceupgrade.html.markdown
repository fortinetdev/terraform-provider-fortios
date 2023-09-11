---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_deviceupgrade"
description: |-
  Independent upgrades for managed devices.
---

# fortios_system_deviceupgrade
Independent upgrades for managed devices. Applies to FortiOS Version `>= 7.2.4`.

## Argument Reference

The following arguments are supported:

* `serial` - Serial number of the node to include.
* `timing` - Run immediately or at a scheduled time. Valid values: `immediate`, `scheduled`.
* `maximum_minutes` - Maximum number of minutes to allow for immediate upgrade preparation.
* `time` - Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
* `setup_time` - Upgrade configuration time in UTC (hh:mm yyyy/mm/dd UTC).
* `upgrade_path` - Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.
* `device_type` - Fortinet device type. Valid values: `fortiswitch`, `fortiap`, `fortiextender`.
* `status` - Current status of the upgrade. Valid values: `disabled`, `initialized`, `downloading`, `device-disconnected`, `ready`, `coordinating`, `staging`, `final-check`, `upgrade-devices`, `cancelled`, `confirmed`, `done`, `failed`.
* `failure_reason` - Upgrade failure reason.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial}}.

## Import

System DeviceUpgrade can be imported using any of these accepted formats:
```
$ terraform import fortios_system_deviceupgrade.labelname {{serial}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_deviceupgrade.labelname {{serial}}
$ unset "FORTIOS_IMPORT_TABLE"
```
