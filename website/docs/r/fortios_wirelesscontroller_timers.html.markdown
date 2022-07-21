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
* `client_idle_rehome_timeout` - Time after which a client is considered idle and disconnected from the home controller (2 - 3600 sec, default = 20, 0 for no timeout).
* `auth_timeout` - Time after which a client is considered failed in RADIUS authentication and times out (5 - 30 sec, default = 5).
* `rogue_ap_log` - Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
* `fake_ap_log` - Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
* `rogue_ap_cleanup` - Time period in minutes to keep rogue AP after it is gone (default = 0).
* `darrp_optimize` - Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
* `darrp_day` - Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
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
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `darrp_time` block supports:

* `time` - Time.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Timers can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_timers.labelname WirelessControllerTimers

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_timers.labelname WirelessControllerTimers
$ unset "FORTIOS_IMPORT_TABLE"
```
