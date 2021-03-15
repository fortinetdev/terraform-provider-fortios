---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_speedtestschedule"
description: |-
  Speed test schedule for each interface.
---

# fortios_system_speedtestschedule
Speed test schedule for each interface. Applies to FortiOS Version `>= 7.0.0`.

## Argument Reference

The following arguments are supported:

* `interface` - Interface name.
* `status` - Enable/disable scheduled speed test. Valid values: `disable`, `enable`.
* `diffserv` - DSCP used for speed test.
* `server_name` - Speed test server name.
* `schedules` - Schedules for the interface. The structure of `schedules` block is documented below.
* `update_inbandwidth` - Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
* `update_outbandwidth` - Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
* `update_inbandwidth_maximum` - Maximum downloading bandwidth (kbps) to be used in a speed test.
* `update_inbandwidth_minimum` - Minimum downloading bandwidth (kbps) to be considered effective.
* `update_outbandwidth_maximum` - Maximum uploading bandwidth (kbps) to be used in a speed test.
* `update_outbandwidth_minimum` - Minimum uploading bandwidth (kbps) to be considered effective.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `schedules` block supports:

* `name` - Name of a firewall recurring schedule.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{interface}}.

## Import

System SpeedTestSchedule can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_speedtestschedule.labelname {{interface}}
$ unset "FORTIOS_IMPORT_TABLE"
```
