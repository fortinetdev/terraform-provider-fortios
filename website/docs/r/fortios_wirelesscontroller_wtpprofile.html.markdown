---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wtpprofile"
description: |-
  Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.
---

# fortios_wirelesscontroller_wtpprofile
Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.

## Argument Reference

The following arguments are supported:

* `name` - WTP (or FortiAP or AP) profile name.
* `comment` - Comment.
* `platform` - WTP, FortiAP, or AP platform. The structure of `platform` block is documented below.
* `control_message_offload` - Enable/disable CAPWAP control message data channel offload.
* `apcfg_profile` - AP local configuration profile name.
* `ble_profile` - Bluetooth Low Energy profile name.
* `wan_port_mode` - Enable/disable using a WAN port as a LAN port.
* `lan` - WTP LAN port mapping. The structure of `lan` block is documented below.
* `energy_efficient_ethernet` - Enable/disable use of energy efficient Ethernet on WTP.
* `led_state` - Enable/disable use of LEDs on WTP (default = disable).
* `led_schedules` - Recurring firewall schedules for illuminating LEDs on the FortiAP. If led-state is enabled, LEDs will be visible when at least one of the schedules is valid. Separate multiple schedule names with a space. The structure of `led_schedules` block is documented below.
* `dtls_policy` - WTP data channel DTLS policy (default = clear-text).
* `dtls_in_kernel` - Enable/disable data channel DTLS in kernel.
* `max_clients` - Maximum number of stations (STAs) supported by the WTP (default = 0, meaning no client limitation).
* `handoff_rssi` - Minimum received signal strength indicator (RSSI) value for handoff (20 - 30, default = 25).
* `handoff_sta_thresh` - Threshold value for AP handoff.
* `handoff_roaming` - Enable/disable client load balancing during roaming to avoid roaming delay (default = disable).
* `deny_mac_list` - List of MAC addresses that are denied access to this WTP, FortiAP, or AP. The structure of `deny_mac_list` block is documented below.
* `ap_country` - Country in which this WTP, FortiAP or AP will operate (default = NA, automatically use the country configured for the current VDOM).
* `ip_fragment_preventing` - Select how to prevent IP fragmentation for CAPWAP tunneled control and data packets (default = tcp-mss-adjust).
* `tun_mtu_uplink` - Uplink CAPWAP tunnel MTU (0, 576, or 1500 bytes, default = 0).
* `tun_mtu_downlink` - Downlink CAPWAP tunnel MTU (0, 576, or 1500 bytes, default = 0).
* `split_tunneling_acl_path` - Split tunneling ACL path is local/tunnel.
* `split_tunneling_acl_local_ap_subnet` - Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable).
* `split_tunneling_acl` - Split tunneling ACL filter list. The structure of `split_tunneling_acl` block is documented below.
* `allowaccess` - Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space.
* `login_passwd_change` - Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no).
* `login_passwd` - Set the managed WTP, FortiAP, or AP's administrator password.
* `lldp` - Enable/disable Link Layer Discovery Protocol (LLDP) for the WTP, FortiAP, or AP (default = disable).
* `poe_mode` - Set the WTP, FortiAP, or AP's PoE mode.
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `radio_1` - Configuration options for radio 1. The structure of `radio_1` block is documented below.
* `radio_2` - Configuration options for radio 2. The structure of `radio_2` block is documented below.
* `radio_3` - Configuration options for radio 3. The structure of `radio_3` block is documented below.
* `radio_4` - Configuration options for radio 4. The structure of `radio_4` block is documented below.
* `lbs` - Set various location based service (LBS) options. The structure of `lbs` block is documented below.
* `ext_info_enable` - Enable/disable station/VAP/radio extension information.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `platform` block supports:

* `type` - WTP, FortiAP or AP platform type. There are built-in WTP profiles for all supported FortiAP models. You can select a built-in profile and customize it or create a new profile.
* `mode` - Configure operation mode of 5G radios (default = single-5G).
* `ddscan` - Enable/disable use of one radio for dedicated dual-band scanning to detect RF characterization and wireless threat management.

The `lan` block supports:

* `port_mode` - LAN port mode.
* `port_ssid` - Bridge LAN port to SSID.
* `port1_mode` - LAN port 1 mode.
* `port1_ssid` - Bridge LAN port 1 to SSID.
* `port2_mode` - LAN port 2 mode.
* `port2_ssid` - Bridge LAN port 2 to SSID.
* `port3_mode` - LAN port 3 mode.
* `port3_ssid` - Bridge LAN port 3 to SSID.
* `port4_mode` - LAN port 4 mode.
* `port4_ssid` - Bridge LAN port 4 to SSID.
* `port5_mode` - LAN port 5 mode.
* `port5_ssid` - Bridge LAN port 5 to SSID.
* `port6_mode` - LAN port 6 mode.
* `port6_ssid` - Bridge LAN port 6 to SSID.
* `port7_mode` - LAN port 7 mode.
* `port7_ssid` - Bridge LAN port 7 to SSID.
* `port8_mode` - LAN port 8 mode.
* `port8_ssid` - Bridge LAN port 8 to SSID.
* `port_esl_mode` - ESL port mode.
* `port_esl_ssid` - Bridge ESL port to SSID.

The `led_schedules` block supports:

* `name` - LED schedule name.

The `deny_mac_list` block supports:

* `id` - ID.
* `mac` - A WiFi device with this MAC address is denied access to this WTP, FortiAP or AP.

The `split_tunneling_acl` block supports:

* `id` - ID.
* `dest_ip` - Destination IP and mask for the split-tunneling subnet.

The `radio_1` block supports:

* `radio_id` - radio-id
* `mode` - Mode of radio 1. Radio 1 can be disabled, configured as an access point, a rogue AP monitor, or a sniffer.
* `band` - WiFi band that Radio 1 operates on.
* `band_5g_type` - WiFi 5G band type.
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable).
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).
* `airtime_fairness` - Enable/disable airtime fairness (default = disable).
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc.
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.
* `channel_bonding` - Channel bandwidth: 80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = disable).
* `auto_power_high` - Automatic transmit power high limit in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_low` - Automatic transmission power low limit in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `power_level` - Radio power level as a percentage of the maximum transmit power (0 - 100, default = 100).
* `dtim` - DTIM interval. The frequency to transmit Delivery Traffic Indication Message (or Map) (DTIM) messages (1 - 255, default = 1). Set higher to save client battery life.
* `beacon_interval` - Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable).
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable).
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable).
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable).
* `channel_utilization` - Enable/disable measuring channel utilization.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `vap_all` - Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable).
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).

The `vaps` block supports:

* `name` - Virtual Access Point (VAP) name.

The `channel` block supports:

* `chan` - Channel number.

The `radio_2` block supports:

* `radio_id` - radio-id
* `mode` - Mode of radio 2. Radio 2 can be disabled, configured as an access point, a rogue AP monitor, or a sniffer.
* `band` - WiFi band that Radio 2 operates on.
* `band_5g_type` - WiFi 5G band type.
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable).
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).
* `airtime_fairness` - Enable/disable airtime fairness (default = disable).
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc.
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.
* `channel_bonding` - Channel bandwidth: 80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = disable).
* `auto_power_high` - Automatic transmit power high limit in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_low` - Automatic transmission power low limit in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `power_level` - Radio power level as a percentage of the maximum transmit power (0 - 100, default = 100).
* `dtim` - DTIM interval. The frequency to transmit Delivery Traffic Indication Message (or Map) (DTIM) messages (1 - 255, default = 1). Set higher to save client battery life.
* `beacon_interval` - Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable).
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable).
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable).
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable).
* `channel_utilization` - Enable/disable measuring channel utilization.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `vap_all` - Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable).
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).

The `vaps` block supports:

* `name` - Virtual Access Point (VAP) name.

The `channel` block supports:

* `chan` - Channel number.

The `radio_3` block supports:

* `mode` - Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, or a sniffer.
* `band` - WiFi band that Radio 3 operates on.
* `band_5g_type` - WiFi 5G band type.
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable).
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).
* `airtime_fairness` - Enable/disable airtime fairness (default = disable).
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc.
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `power_level` - Radio power level as a percentage of the maximum transmit power (0 - 100, default = 100).
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `beacon_interval` - Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable).
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable).
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable).
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable).
* `channel_utilization` - Enable/disable measuring channel utilization.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `vap_all` - Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable).
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).

The `vaps` block supports:

* `name` - Virtual Access Point (VAP) name.

The `channel` block supports:

* `chan` - Channel number.

The `radio_4` block supports:

* `mode` - Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, or a sniffer.
* `band` - WiFi band that Radio 3 operates on.
* `band_5g_type` - WiFi 5G band type.
* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable).
* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low).
* `airtime_fairness` - Enable/disable airtime fairness (default = disable).
* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc.
* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default.
* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable).
* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable).
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns.
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence.
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable).
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `power_level` - Radio power level as a percentage of the maximum transmit power (0 - 100, default = 100).
* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `beacon_interval` - Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable).
* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable).
* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable).
* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable).
* `channel_utilization` - Enable/disable measuring channel utilization.
* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance.
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable).
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable).
* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable).
* `vap_all` - Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable).
* `vaps` - Manually selected list of Virtual Access Points (VAPs). The structure of `vaps` block is documented below.
* `channel` - Selected list of wireless radio channels. The structure of `channel` block is documented below.
* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).

The `vaps` block supports:

* `name` - Virtual Access Point (VAP) name.

The `channel` block supports:

* `chan` - Channel number.

The `lbs` block supports:

* `ekahau_blink_mode` - Enable/disable Ekahua blink mode (also called AiRISTA Flow Blink Mode) to find the location of devices connected to a wireless LAN (default = disable).
* `ekahau_tag` - WiFi frame MAC address or WiFi Tag.
* `erc_server_ip` - IP address of Ekahua RTLS Controller (ERC).
* `erc_server_port` - Ekahua RTLS Controller (ERC) UDP listening port.
* `aeroscout` - Enable/disable AeroScout Real Time Location Service (RTLS) support.
* `aeroscout_server_ip` - IP address of AeroScout server.
* `aeroscout_server_port` - AeroScout server UDP listening port.
* `aeroscout_mu` - Enable/disable AeroScout support.
* `aeroscout_ap_mac` - Use BSSID or board MAC address as AP MAC address in the Aeroscout AP message.
* `aeroscout_mmu_report` - Enable/disable MU compounded report.
* `aeroscout_mu_factor` - AeroScout Mobile Unit (MU) mode dilution factor (default = 20).
* `aeroscout_mu_timeout` - AeroScout MU mode timeout (0 - 65535 sec, default = 5).
* `fortipresence` - Enable/disable FortiPresence to monitor the location and activity of WiFi clients even if they don't connect to this WiFi network (default = disable).
* `fortipresence_server` - FortiPresence server IP address.
* `fortipresence_port` - FortiPresence server UDP listening port (default = 3000).
* `fortipresence_secret` - FortiPresence secret password (max. 16 characters).
* `fortipresence_project` - FortiPresence project name (max. 16 characters, default = fortipresence).
* `fortipresence_frequency` - FortiPresence report transmit frequency (5 - 65535 sec, default = 30).
* `fortipresence_rogue` - Enable/disable FortiPresence finding and reporting rogue APs.
* `fortipresence_unassoc` - Enable/disable FortiPresence finding and reporting unassociated stations.
* `fortipresence_ble` - Enable/disable FortiPresence finding and reporting BLE devices.
* `station_locate` - Enable/disable client station locating services for all clients, whether associated or not (default = disable).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController WtpProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_wtpprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
