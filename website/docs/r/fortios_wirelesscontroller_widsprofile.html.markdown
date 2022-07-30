---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_widsprofile"
description: |-
  Configure wireless intrusion detection system (WIDS) profiles.
---

# fortios_wirelesscontroller_widsprofile
Configure wireless intrusion detection system (WIDS) profiles.

## Argument Reference

The following arguments are supported:

* `name` - WIDS profile name.
* `comment` - Comment.
* `sensor_mode` - Scan WiFi nearby stations (default = disable). Valid values: `disable`, `foreign`, `both`.
* `ap_scan` - Enable/disable rogue AP detection. Valid values: `disable`, `enable`.
* `ap_bgscan_period` - Period of time between background scans (60 - 3600 sec, default = 600).
* `ap_bgscan_intv` - Period of time between scanning two channels (1 - 600 sec, default = 1).
* `ap_bgscan_duration` - Listening time on a scanning channel (10 - 1000 msec, default = 20).
* `ap_bgscan_idle` - Waiting time for channel inactivity before scanning this channel (0 - 1000 msec, default = 0).
* `ap_bgscan_report_intv` - Period of time between background scan reports (15 - 600 sec, default = 30).
* `ap_bgscan_disable_schedules` - Firewall schedules for turning off FortiAP radio background scan. Background scan will be disabled when at least one of the schedules is valid. Separate multiple schedule names with a space. The structure of `ap_bgscan_disable_schedules` block is documented below.
* `ap_bgscan_disable_day` - Optionally turn off scanning for one or more days of the week. Separate the days with a space. By default, no days are set. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `ap_bgscan_disable_start` - Start time, using a 24-hour clock in the format of hh:mm, for disabling background scanning (default = 00:00).
* `ap_bgscan_disable_end` - End time, using a 24-hour clock in the format of hh:mm, for disabling background scanning (default = 00:00).
* `ap_fgscan_report_intv` - Period of time between foreground scan reports (15 - 600 sec, default = 15).
* `ap_scan_passive` - Enable/disable passive scanning. Enable means do not send probe request on any channels (default = disable). Valid values: `enable`, `disable`.
* `ap_scan_threshold` - Minimum signal level/threshold in dBm required for the AP to report detected rogue AP (-95 to -20, default = -90).
* `ap_auto_suppress` - Enable/disable on-wire rogue AP auto-suppression (default = disable). Valid values: `enable`, `disable`.
* `wireless_bridge` - Enable/disable wireless bridge detection (default = disable). Valid values: `enable`, `disable`.
* `deauth_broadcast` - Enable/disable broadcasting de-authentication detection (default = disable). Valid values: `enable`, `disable`.
* `null_ssid_probe_resp` - Enable/disable null SSID probe response detection (default = disable). Valid values: `enable`, `disable`.
* `long_duration_attack` - Enable/disable long duration attack detection based on user configured threshold (default = disable). Valid values: `enable`, `disable`.
* `long_duration_thresh` - Threshold value for long duration attack detection (1000 - 32767 usec, default = 8200).
* `invalid_mac_oui` - Enable/disable invalid MAC OUI detection. Valid values: `enable`, `disable`.
* `weak_wep_iv` - Enable/disable weak WEP IV (Initialization Vector) detection (default = disable). Valid values: `enable`, `disable`.
* `auth_frame_flood` - Enable/disable authentication frame flooding detection (default = disable). Valid values: `enable`, `disable`.
* `auth_flood_time` - Number of seconds after which a station is considered not connected.
* `auth_flood_thresh` - The threshold value for authentication frame flooding.
* `assoc_frame_flood` - Enable/disable association frame flooding detection (default = disable). Valid values: `enable`, `disable`.
* `assoc_flood_time` - Number of seconds after which a station is considered not connected.
* `assoc_flood_thresh` - The threshold value for association frame flooding.
* `spoofed_deauth` - Enable/disable spoofed de-authentication attack detection (default = disable). Valid values: `enable`, `disable`.
* `asleap_attack` - Enable/disable asleap attack detection (default = disable). Valid values: `enable`, `disable`.
* `eapol_start_flood` - Enable/disable EAPOL-Start flooding (to AP) detection (default = disable). Valid values: `enable`, `disable`.
* `eapol_start_thresh` - The threshold value for EAPOL-Start flooding in specified interval.
* `eapol_start_intv` - The detection interval for EAPOL-Start flooding (1 - 3600 sec).
* `eapol_logoff_flood` - Enable/disable EAPOL-Logoff flooding (to AP) detection (default = disable). Valid values: `enable`, `disable`.
* `eapol_logoff_thresh` - The threshold value for EAPOL-Logoff flooding in specified interval.
* `eapol_logoff_intv` - The detection interval for EAPOL-Logoff flooding (1 - 3600 sec).
* `eapol_succ_flood` - Enable/disable EAPOL-Success flooding (to AP) detection (default = disable). Valid values: `enable`, `disable`.
* `eapol_succ_thresh` - The threshold value for EAPOL-Success flooding in specified interval.
* `eapol_succ_intv` - The detection interval for EAPOL-Success flooding (1 - 3600 sec).
* `eapol_fail_flood` - Enable/disable EAPOL-Failure flooding (to AP) detection (default = disable). Valid values: `enable`, `disable`.
* `eapol_fail_thresh` - The threshold value for EAPOL-Failure flooding in specified interval.
* `eapol_fail_intv` - The detection interval for EAPOL-Failure flooding (1 - 3600 sec).
* `eapol_pre_succ_flood` - Enable/disable premature EAPOL-Success flooding (to STA) detection (default = disable). Valid values: `enable`, `disable`.
* `eapol_pre_succ_thresh` - The threshold value for premature EAPOL-Success flooding in specified interval.
* `eapol_pre_succ_intv` - The detection interval for premature EAPOL-Success flooding (1 - 3600 sec).
* `eapol_pre_fail_flood` - Enable/disable premature EAPOL-Failure flooding (to STA) detection (default = disable). Valid values: `enable`, `disable`.
* `eapol_pre_fail_thresh` - The threshold value for premature EAPOL-Failure flooding in specified interval.
* `eapol_pre_fail_intv` - The detection interval for premature EAPOL-Failure flooding (1 - 3600 sec).
* `deauth_unknown_src_thresh` - Threshold value per second to deauth unknown src for DoS attack (0: no limit).
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ap_bgscan_disable_schedules` block supports:

* `name` - Schedule name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController WidsProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_widsprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_widsprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
