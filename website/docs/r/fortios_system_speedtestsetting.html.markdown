---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_speedtestsetting"
description: |-
  Configure speed test setting.
---

# fortios_system_speedtestsetting
Configure speed test setting. Applies to FortiOS Version `7.2.6,7.4.1`.

## Argument Reference

The following arguments are supported:

* `latency_threshold` - Speed test latency threshold in milliseconds (0 - 2000, default = 60) for the Auto mode. If the latency exceeds this threshold, the speed test will use the UDP protocol; otherwise, it will use the TCP protocol.
* `multiple_tcp_stream` - Number of parallel client streams (1 - 64, default = 4) for the TCP protocol to run during the speed test.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SpeedTestSetting can be imported using any of these accepted formats:
```
$ terraform import fortios_system_speedtestsetting.labelname SystemSpeedTestSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_speedtestsetting.labelname SystemSpeedTestSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
