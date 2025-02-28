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
* `nat_session_keep_alive` - Maximal time in seconds between control requests sent by the managed WTP, AP, or FortiAP (0 - 255 sec, default = 0).
* `discovery_interval` - Time between discovery requests (2 - 180 sec, default = 5).
* `client_idle_timeout` - Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
* `client_idle_rehome_timeout` - Time after which a client is considered idle and disconnected from the home controller (2 - 3600 sec, default = 20, 0 for no timeout).
* `auth_timeout` - Time after which a client is considered failed in RADIUS authentication and times out (5 - 30 sec, default = 5).
* `rogue_ap_log` - Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
* `fake_ap_log` - Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
* `sta_cap_cleanup` - Time period in minutes to keep station capability data after it is gone (default = 0).
* `rogue_ap_cleanup` - Time period in minutes to keep rogue AP after it is gone (default = 0).
* `rogue_sta_cleanup` - Time period in minutes to keep rogue station after it is gone (default = 0).
* `wids_entry_cleanup` - Time period in minutes to keep wids entry after it is gone (default = 0).
* `ble_device_cleanup` - Time period in minutes to keep BLE device after it is gone (default = 60).
* `darrp_optimize` - Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 1800).
* `darrp_day` - Weekday on which to run DARRP optimization. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `darrp_time` - Time at which DARRP optimizations run (you can add up to 8 times). The structure of `darrp_time` block is documented below.
* `sta_stats_interval` - Time between running client (station) reports (1 - 255 sec). On FortiOS versions 6.2.0-7.4.1: default = 1. On FortiOS versions >= 7.4.2: default = 10.
* `vap_stats_interval` - Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).
* `radio_stats_interval` - Time between running radio reports (1 - 255 sec, default = 15).
* `sta_capability_interval` - Time between running station capability reports (1 - 255 sec, default = 30).
* `sta_locate_timer` - Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
* `ipsec_intf_cleanup` - Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
* `ble_scan_report_intv` - Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
* `drma_interval` - Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
* `ap_reboot_wait_interval1` - Time in minutes to wait before AP reboots when there is no controller detected (5 - 65535, default = 0, 0 for no reboot).
* `ap_reboot_wait_time` - Time to reboot the AP when there is no controller detected and standalone SSIDs are pushed to the AP in the previous session, format hh:mm.
* `ap_reboot_wait_interval2` - Time in minutes to wait before AP reboots when there is no controller detected and standalone SSIDs are pushed to the AP in the previous session (5 - 65535, default = 0, 0 for no reboot).
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
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
