---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_log"
description: |-
  Configure wireless controller event log filters.
---

# fortios_wirelesscontroller_log
Configure wireless controller event log filters. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable wireless event logging. Valid values: `enable`, `disable`.
* `addrgrp_log` - Lowest severity level to log address group message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `ble_log` - Lowest severity level to log BLE detection message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `clb_log` - Lowest severity level to log client load balancing message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `dhcp_starv_log` - Lowest severity level to log DHCP starvation event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `led_sched_log` - Lowest severity level to log LED schedule event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `radio_event_log` - Lowest severity level to log radio event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `rogue_event_log` - Lowest severity level to log rogue AP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `sta_event_log` - Lowest severity level to log station event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `sta_locate_log` - Lowest severity level to log station locate message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `wids_log` - Lowest severity level to log WIDS message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `wtp_event_log` - Lowest severity level to log WTP event message. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Log can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_log.labelname WirelessControllerLog

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_log.labelname WirelessControllerLog
$ unset "FORTIOS_IMPORT_TABLE"
```
