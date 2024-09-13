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
* `setup_time` - Upgrade preparation start time in UTC (hh:mm yyyy/mm/dd UTC).
* `upgrade_path` - Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.
* `device_type` - Fortinet device type.
* `status` - Current status of the upgrade.
* `failure_reason` - Upgrade failure reason.
* `ha_reboot_controller` - Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
* `next_path_index` - The index of the next image to upgrade to.
* `known_ha_members` - Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled. The structure of `known_ha_members` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `known_ha_members` block supports:

* `serial` - Serial number of HA member


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
