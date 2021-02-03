---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_log"
description: |-
  Configure wireless controller event log filters.
---

# fortios_wirelesscontroller_log
Configure wireless controller event log filters.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable wireless event logging.
* `addrgrp_log` - Lowest severity level to log address group message.
* `ble_log` - Lowest severity level to log BLE detection message.
* `clb_log` - Lowest severity level to log client load balancing message.
* `dhcp_starv_log` - Lowest severity level to log DHCP starvation event message.
* `led_sched_log` - Lowest severity level to log LED schedule event message.
* `radio_event_log` - Lowest severity level to log radio event message.
* `rogue_event_log` - Lowest severity level to log rogue AP event message.
* `sta_event_log` - Lowest severity level to log station event message.
* `sta_locate_log` - Lowest severity level to log station locate message.
* `wids_log` - Lowest severity level to log WIDS message.
* `wtp_event_log` - Lowest severity level to log WTP event message.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Log can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_log.labelname WirelessControllerLog
$ unset "FORTIOS_IMPORT_TABLE"
```
