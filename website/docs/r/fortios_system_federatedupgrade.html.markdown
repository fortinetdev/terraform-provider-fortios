---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_federatedupgrade"
description: |-
  Coordinate federated upgrades within the Security Fabric.
---

# fortios_system_federatedupgrade
Coordinate federated upgrades within the Security Fabric. Applies to FortiOS Version `>= 7.0.0`.

## Argument Reference

The following arguments are supported:

* `status` - Current status of the upgrade.
* `failure_reason` - Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`.
* `failure_device` - Serial number of the node to include.
* `upgrade_id` - Unique identifier for this upgrade.
* `next_path_index` - The index of the next image to upgrade to.
* `node_list` - Nodes which will be included in the upgrade. The structure of `node_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `node_list` block supports:

* `serial` - Serial number of the node to include.
* `timing` - Whether the upgrade should be run immediately, or at a scheduled time. Valid values: `immediate`, `scheduled`.
* `time` - Scheduled time for the upgrade. Format hh:mm yyyy/mm/dd UTC.
* `setup_time` - When the upgrade was configured. Format hh:mm yyyy/mm/dd UTC.
* `upgrade_path` - Image IDs to upgrade through.
* `device_type` - What type of device this node represents. Valid values: `fortigate`, `fortiswitch`, `fortiap`.
* `coordinating_fortigate` - The serial of the FortiGate that controls this device


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FederatedUpgrade can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_federatedupgrade.labelname SystemFederatedUpgrade
$ unset "FORTIOS_IMPORT_TABLE"
```
