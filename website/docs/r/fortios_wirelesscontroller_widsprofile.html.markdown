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
* `sensor_mode` - Scan nearby WiFi stations (default = disable). Valid values: `disable`, `foreign`, `both`.
* `ap_scan` - Enable/disable rogue AP detection. Valid values: `disable`, `enable`.
* `ap_scan_channel_list_2g_5g` - Selected ap scan channel list for 2.4G and 5G bands. The structure of `ap_scan_channel_list_2g_5g` block is documented below.
* `ap_scan_channel_list_6g` - Selected ap scan channel list for 6G band. The structure of `ap_scan_channel_list_6g` block is documented below.
* `ap_bgscan_period` - Period between background scans (default = 600). On FortiOS versions 6.2.0-6.2.6: 60 - 3600 sec. On FortiOS versions 6.4.0-7.0.1: 10 - 3600 sec.
* `ap_bgscan_intv` - Period between successive channel scans (1 - 600 sec). On FortiOS versions 6.2.0-7.0.1: default = 1. On FortiOS versions >= 7.0.2: default = 3.
* `ap_bgscan_duration` - Listen time on scanning a channel (10 - 1000 msec). On FortiOS versions 6.2.0-7.0.1: default = 20. On FortiOS versions >= 7.0.2: default = 30.
* `ap_bgscan_idle` - Wait time for channel inactivity before scanning this channel (0 - 1000 msec). On FortiOS versions 6.2.0-7.0.1: default = 0. On FortiOS versions >= 7.0.2: default = 20.
* `ap_bgscan_report_intv` - Period between background scan reports (15 - 600 sec, default = 30).
* `ap_bgscan_disable_schedules` - Firewall schedules for turning off FortiAP radio background scan. Background scan will be disabled when at least one of the schedules is valid. Separate multiple schedule names with a space. The structure of `ap_bgscan_disable_schedules` block is documented below.
* `ap_bgscan_disable_day` - Optionally turn off scanning for one or more days of the week. Separate the days with a space. By default, no days are set. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `ap_bgscan_disable_start` - Start time, using a 24-hour clock in the format of hh:mm, for disabling background scanning (default = 00:00).
* `ap_bgscan_disable_end` - End time, using a 24-hour clock in the format of hh:mm, for disabling background scanning (default = 00:00).
* `ap_fgscan_report_intv` - Period between foreground scan reports (15 - 600 sec, default = 15).
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
* `reassoc_flood` - Enable/disable reassociation flood detection (default = disable). Valid values: `enable`, `disable`.
* `reassoc_flood_time` - Detection Window Period.
* `reassoc_flood_thresh` - The threshold value for reassociation flood.
* `probe_flood` - Enable/disable probe flood detection (default = disable). Valid values: `enable`, `disable`.
* `probe_flood_time` - Detection Window Period.
* `probe_flood_thresh` - The threshold value for probe flood.
* `bcn_flood` - Enable/disable bcn flood detection (default = disable). Valid values: `enable`, `disable`.
* `bcn_flood_time` - Detection Window Period.
* `bcn_flood_thresh` - The threshold value for bcn flood.
* `rts_flood` - Enable/disable rts flood detection (default = disable). Valid values: `enable`, `disable`.
* `rts_flood_time` - Detection Window Period.
* `rts_flood_thresh` - The threshold value for rts flood.
* `cts_flood` - Enable/disable cts flood detection (default = disable). Valid values: `enable`, `disable`.
* `cts_flood_time` - Detection Window Period.
* `cts_flood_thresh` - The threshold value for cts flood.
* `client_flood` - Enable/disable client flood detection (default = disable). Valid values: `enable`, `disable`.
* `client_flood_time` - Detection Window Period.
* `client_flood_thresh` - The threshold value for client flood.
* `block_ack_flood` - Enable/disable block_ack flood detection (default = disable). Valid values: `enable`, `disable`.
* `block_ack_flood_time` - Detection Window Period.
* `block_ack_flood_thresh` - The threshold value for block_ack flood.
* `pspoll_flood` - Enable/disable pspoll flood detection (default = disable). Valid values: `enable`, `disable`.
* `pspoll_flood_time` - Detection Window Period.
* `pspoll_flood_thresh` - The threshold value for pspoll flood.
* `netstumbler` - Enable/disable netstumbler detection (default = disable). Valid values: `enable`, `disable`.
* `netstumbler_time` - Detection Window Period.
* `netstumbler_thresh` - The threshold value for netstumbler.
* `wellenreiter` - Enable/disable wellenreiter detection (default = disable). Valid values: `enable`, `disable`.
* `wellenreiter_time` - Detection Window Period.
* `wellenreiter_thresh` - The threshold value for wellenreiter.
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
* `windows_bridge` - Enable/disable windows bridge detection (default = disable). Valid values: `enable`, `disable`.
* `disassoc_broadcast` - Enable/disable broadcast dis-association detection (default = disable). Valid values: `enable`, `disable`.
* `ap_spoofing` - Enable/disable AP spoofing detection (default = disable). Valid values: `enable`, `disable`.
* `chan_based_mitm` - Enable/disable channel based mitm detection (default = disable). Valid values: `enable`, `disable`.
* `adhoc_valid_ssid` - Enable/disable adhoc using valid SSID detection (default = disable). Valid values: `enable`, `disable`.
* `adhoc_network` - Enable/disable adhoc network detection (default = disable). Valid values: `enable`, `disable`.
* `eapol_key_overflow` - Enable/disable overflow EAPOL key detection (default = disable). Valid values: `enable`, `disable`.
* `ap_impersonation` - Enable/disable AP impersonation detection (default = disable). Valid values: `enable`, `disable`.
* `invalid_addr_combination` - Enable/disable invalid address combination detection (default = disable). Valid values: `enable`, `disable`.
* `beacon_wrong_channel` - Enable/disable beacon wrong channel detection (default = disable). Valid values: `enable`, `disable`.
* `ht_greenfield` - Enable/disable HT greenfield detection (default = disable). Valid values: `enable`, `disable`.
* `overflow_ie` - Enable/disable overflow IE detection (default = disable). Valid values: `enable`, `disable`.
* `malformed_ht_ie` - Enable/disable malformed HT IE detection (default = disable). Valid values: `enable`, `disable`.
* `malformed_auth` - Enable/disable malformed auth frame detection (default = disable). Valid values: `enable`, `disable`.
* `malformed_association` - Enable/disable malformed association request detection (default = disable). Valid values: `enable`, `disable`.
* `ht_40mhz_intolerance` - Enable/disable HT 40 MHz intolerance detection (default = disable). Valid values: `enable`, `disable`.
* `valid_ssid_misuse` - Enable/disable valid SSID misuse detection (default = disable). Valid values: `enable`, `disable`.
* `valid_client_misassociation` - Enable/disable valid client misassociation detection (default = disable). Valid values: `enable`, `disable`.
* `hotspotter_attack` - Enable/disable hotspotter attack detection (default = disable). Valid values: `enable`, `disable`.
* `pwsave_dos_attack` - Enable/disable power save DOS attack detection (default = disable). Valid values: `enable`, `disable`.
* `omerta_attack` - Enable/disable omerta attack detection (default = disable). Valid values: `enable`, `disable`.
* `disconnect_station` - Enable/disable disconnect station detection (default = disable). Valid values: `enable`, `disable`.
* `unencrypted_valid` - Enable/disable unencrypted valid detection (default = disable). Valid values: `enable`, `disable`.
* `fata_jack` - Enable/disable FATA-Jack detection (default = disable). Valid values: `enable`, `disable`.
* `risky_encryption` - Enable/disable Risky Encryption detection (default = disable). Valid values: `enable`, `disable`.
* `fuzzed_beacon` - Enable/disable fuzzed beacon detection (default = disable). Valid values: `enable`, `disable`.
* `fuzzed_probe_request` - Enable/disable fuzzed probe request detection (default = disable). Valid values: `enable`, `disable`.
* `fuzzed_probe_response` - Enable/disable fuzzed probe response detection (default = disable). Valid values: `enable`, `disable`.
* `air_jack` - Enable/disable AirJack detection (default = disable). Valid values: `enable`, `disable`.
* `wpa_ft_attack` - Enable/disable WPA FT attack detection (default = disable). Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ap_scan_channel_list_2g_5g` block supports:

* `chan` - Channel number.

The `ap_scan_channel_list_6g` block supports:

* `chan` - Channel 6g number.

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
