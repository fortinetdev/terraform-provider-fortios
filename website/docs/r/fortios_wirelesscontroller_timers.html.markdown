---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_timers"
description: |-
  Configure CAPWAP timers.
---

# fortios_wirelesscontroller_timers
Configure CAPWAP timers.

## Argument Reference

The following arguments are supported:

* `echo_interval` - Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
* `discovery_interval` - Time between discovery requests (2 - 180 sec, default = 5).
* `client_idle_timeout` - Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
* `rogue_ap_log` - Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
* `fake_ap_log` - Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
* `darrp_optimize` - Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
* `darrp_day` - Weekday on which to run DARRP optimization.
* `darrp_time` - Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrp_time` block is documented below.
* `sta_stats_interval` - Time between running client (station) reports (1 - 255 sec, default = 1).
* `vap_stats_interval` - Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
* `radio_stats_interval` - Time between running radio reports (1 - 255 sec, default = 15).
* `sta_capability_interval` - Time between running station capability reports (1 - 255 sec, default = 30).
* `sta_locate_timer` - Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
* `ipsec_intf_cleanup` - Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
* `ble_scan_report_intv` - Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
* `drma_interval` - Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `darrp_time` block supports:

* `time` - Time.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Timers can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_timers.labelname WirelessControllerTimers
$ unset "FORTIOS_IMPORT_TABLE"
```
