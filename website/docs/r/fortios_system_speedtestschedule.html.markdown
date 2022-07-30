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
* `dynamic_server` - Enable/disable dynamic server option. Valid values: `disable`, `enable`.
* `update_inbandwidth` - Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable`, `enable`.
* `update_outbandwidth` - Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable`, `enable`.
* `update_inbandwidth_maximum` - Maximum downloading bandwidth (kbps) to be used in a speed test.
* `update_inbandwidth_minimum` - Minimum downloading bandwidth (kbps) to be considered effective.
* `update_outbandwidth_maximum` - Maximum uploading bandwidth (kbps) to be used in a speed test.
* `update_outbandwidth_minimum` - Minimum uploading bandwidth (kbps) to be considered effective.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `schedules` block supports:

* `name` - Name of a firewall recurring schedule.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{interface}}.

## Import

System SpeedTestSchedule can be imported using any of these accepted formats:
```
$ terraform import fortios_system_speedtestschedule.labelname {{interface}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_speedtestschedule.labelname {{interface}}
$ unset "FORTIOS_IMPORT_TABLE"
```
