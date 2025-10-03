---
subcategory: "FortiGate Extension-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extensioncontroller_extenderprofile"
description: |-
  FortiExtender extender profile configuration.
---

# fortios_extensioncontroller_extenderprofile
FortiExtender extender profile configuration. Applies to FortiOS Version `>= 7.2.1`.

## Argument Reference

The following arguments are supported:

* `name` - FortiExtender profile name.
* `fosid` - ID.
* `model` - Model.
* `extension` - Extension option. Valid values: `wan-extension`, `lan-extension`.
* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
* `login_password` - Set the managed extender's administrator password.
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `cellular` - FortiExtender cellular configuration. The structure of `cellular` block is documented below.
* `lan_extension` - FortiExtender lan extension configuration. The structure of `lan_extension` block is documented below.
* `wifi` - FortiExtender wifi configuration. The structure of `wifi` block is documented below.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `cellular` block supports:

* `dataplan` - Dataplan names. The structure of `dataplan` block is documented below.
* `controller_report` - FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
* `sms_notification` - FortiExtender cellular SMS notification configuration. The structure of `sms_notification` block is documented below.
* `modem1` - Configuration options for modem 1. The structure of `modem1` block is documented below.
* `modem2` - Configuration options for modem 2. The structure of `modem2` block is documented below.

The `dataplan` block supports:

* `name` - Dataplan name.

The `controller_report` block supports:

* `status` - FortiExtender controller report status. Valid values: `disable`, `enable`.
* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.

The `sms_notification` block supports:

* `status` - FortiExtender SMS notification status. Valid values: `disable`, `enable`.
* `alert` - SMS alert list. The structure of `alert` block is documented below.
* `receiver` - SMS notification receiver list. The structure of `receiver` block is documented below.

The `alert` block supports:

* `system_reboot` - Display string when system rebooted.
* `data_exhausted` - Display string when data exhausted.
* `session_disconnect` - Display string when session disconnected.
* `low_signal_strength` - Display string when signal strength is low.
* `os_image_fallback` - Display string when falling back to a previous OS image.
* `mode_switch` - Display string when mode is switched.
* `fgt_backup_mode_switch` - Display string when FortiGate backup mode switched.

The `receiver` block supports:

* `name` - FortiExtender SMS notification receiver name.
* `status` - SMS notification receiver status. Valid values: `disable`, `enable`.
* `phone_number` - Receiver phone number. Format: [+][country code][area code][local phone number]. For example, +16501234567.
* `alert` - Alert multi-options. Valid values: `system-reboot`, `data-exhausted`, `session-disconnect`, `low-signal-strength`, `mode-switch`, `os-image-fallback`, `fgt-backup-mode-switch`.

The `modem1` block supports:

* `redundant_mode` - FortiExtender mode. Valid values: `disable`, `enable`.
* `redundant_intf` - Redundant interface.
* `conn_status` - Connection status.
* `default_sim` - Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.
* `gps` - FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.
* `sim1_pin` - SIM #1 PIN status. Valid values: `disable`, `enable`.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable`, `enable`.
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin_code` - SIM #2 PIN password.
* `preferred_carrier` - Preferred carrier.
* `auto_switch` - FortiExtender auto switch configuration. The structure of `auto_switch` block is documented below.
* `multiple_pdn` - Multiple-PDN enable/disable. Valid values: `disable`, `enable`.
* `pdn1_dataplan` - PDN1-dataplan.
* `pdn2_dataplan` - PDN2-dataplan.
* `pdn3_dataplan` - PDN3-dataplan.
* `pdn4_dataplan` - PDN4-dataplan.

The `auto_switch` block supports:

* `disconnect` - Auto switch by disconnect. Valid values: `disable`, `enable`.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `disconnect_period` - Automatically switch based on disconnect period.
* `signal` - Automatically switch based on signal strength. Valid values: `disable`, `enable`.
* `dataplan` - Automatically switch based on data usage. Valid values: `disable`, `enable`.
* `switch_back` - Auto switch with switch back multi-options. Valid values: `time`, `timer`.
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).

The `modem2` block supports:

* `redundant_mode` - FortiExtender mode. Valid values: `disable`, `enable`.
* `redundant_intf` - Redundant interface.
* `conn_status` - Connection status.
* `default_sim` - Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.
* `gps` - FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.
* `sim1_pin` - SIM #1 PIN status. Valid values: `disable`, `enable`.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable`, `enable`.
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin_code` - SIM #2 PIN password.
* `preferred_carrier` - Preferred carrier.
* `auto_switch` - FortiExtender auto switch configuration. The structure of `auto_switch` block is documented below.
* `multiple_pdn` - Multiple-PDN enable/disable. Valid values: `disable`, `enable`.
* `pdn1_dataplan` - PDN1-dataplan.
* `pdn2_dataplan` - PDN2-dataplan.
* `pdn3_dataplan` - PDN3-dataplan.
* `pdn4_dataplan` - PDN4-dataplan.

The `auto_switch` block supports:

* `disconnect` - Auto switch by disconnect. Valid values: `disable`, `enable`.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `disconnect_period` - Automatically switch based on disconnect period.
* `signal` - Automatically switch based on signal strength. Valid values: `disable`, `enable`.
* `dataplan` - Automatically switch based on data usage. Valid values: `disable`, `enable`.
* `switch_back` - Auto switch with switch back multi-options. Valid values: `time`, `timer`.
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).

The `lan_extension` block supports:

* `link_loadbalance` - LAN extension link load balance strategy. Valid values: `activebackup`, `loadbalance`.
* `ipsec_tunnel` - IPsec tunnel name.
* `backhaul_interface` - IPsec phase1 interface.
* `backhaul_ip` - IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
* `backhaul` - LAN extension backhaul tunnel configuration. The structure of `backhaul` block is documented below.
* `downlinks` - Config FortiExtender downlink interface for LAN extension. The structure of `downlinks` block is documented below.
* `traffic_split_services` - Config FortiExtender traffic split interface for LAN extension. The structure of `traffic_split_services` block is documented below.

The `backhaul` block supports:

* `name` - FortiExtender LAN extension backhaul name.
* `port` - FortiExtender uplink port. Valid values: `wan`, `lte1`, `lte2`, `port1`, `port2`, `port3`, `port4`, `port5`, `sfp`.
* `role` - FortiExtender uplink port. Valid values: `primary`, `secondary`.
* `weight` - WRR weight parameter.

The `downlinks` block supports:

* `name` - FortiExtender LAN extension downlink config entry name.
* `type` - FortiExtender LAN extension downlink type [port/vap]. Valid values: `port`, `vap`.
* `port` - FortiExtender LAN extension downlink port.
* `vap` - FortiExtender LAN extension downlink vap.
* `pvid` - FortiExtender LAN extension downlink PVID (1 - 4089).
* `vids` - FortiExtender LAN extension downlink VIDs. The structure of `vids` block is documented below.

The `vids` block supports:

* `vid` - Please enter VID numbers (1 - 4089) with space separated. Up to 50 VIDs are accepted.

The `traffic_split_services` block supports:

* `name` - FortiExtender LAN extension tunnel split entry name.
* `vsdb` - Select vsdb [enable/disable]. Valid values: `disable`, `enable`.
* `address` - Address selection.
* `service` - Service selection.

The `wifi` block supports:

* `country` - Country in which this FEX will operate (default = NA).
* `radio_1` - Radio-1 config for Wi-Fi 2.4GHz The structure of `radio_1` block is documented below.
* `radio_2` - Radio-2 config for Wi-Fi 5GHz The structure of `radio_2` block is documented below.

The `radio_1` block supports:

* `mode` - Wi-Fi radio mode AP(LAN mode) / Client(WAN mode). Valid values: `AP`, `Client`.
* `band` - Wi-Fi band selection 2.4GHz / 5GHz. Valid values: `2.4GHz`.
* `status` - Enable/disable Wi-Fi radio. Valid values: `disable`, `enable`.
* `operating_standard` - Wi-Fi operating standard. Valid values: `auto`, `11A-N-AC-AX`, `11A-N-AC`, `11A-N`, `11A`, `11N-AC-AX`, `11AC-AX`, `11AC`, `11N-AC`, `11B-G-N-AX`, `11B-G-N`, `11B-G`, `11B`, `11G-N-AX`, `11N-AX`, `11AX`, `11G-N`, `11N`, `11G`.
* `guard_interval` - Wi-Fi guard interval. Valid values: `auto`, `400ns`, `800ns`.
* `channel` - Wi-Fi channels. Valid values: `CH1`, `CH2`, `CH3`, `CH4`, `CH5`, `CH6`, `CH7`, `CH8`, `CH9`, `CH10`, `CH11`.
* `bandwidth` - Wi-Fi channel bandwidth. Valid values: `auto`, `20MHz`, `40MHz`, `80MHz`.
* `power_level` - Wi-Fi power level in percent (0 - 100, 0 = auto, default = 100).
* `beacon_interval` - Wi-Fi beacon interval in miliseconds (100 - 3500, default = 100).
* `n80211d` - Enable/disable Wi-Fi 802.11d. Valid values: `disable`, `enable`.
* `max_clients` - Maximum number of Wi-Fi radio clients (0 - 512, 0 = unlimited, default = 0).
* `extension_channel` - Wi-Fi extension channel. Valid values: `auto`, `higher`, `lower`.
* `bss_color_mode` - Wi-Fi 802.11AX BSS color mode. Valid values: `auto`, `static`.
* `bss_color` - Wi-Fi 802.11AX BSS color value (0 - 63, 0 = disable, default = 0).
* `lan_ext_vap` - Wi-Fi LAN-Extention VAP. Select only one VAP.
* `local_vaps` - Wi-Fi local VAP. Select up to three VAPs. The structure of `local_vaps` block is documented below.

The `local_vaps` block supports:

* `name` - Wi-Fi local VAP name.

The `radio_2` block supports:

* `mode` - Wi-Fi radio mode AP(LAN mode) / Client(WAN mode). Valid values: `AP`, `Client`.
* `band` - Wi-Fi band selection 2.4GHz / 5GHz. Valid values: `5GHz`.
* `status` - Enable/disable Wi-Fi radio. Valid values: `disable`, `enable`.
* `operating_standard` - Wi-Fi operating standard. Valid values: `auto`, `11A-N-AC-AX`, `11A-N-AC`, `11A-N`, `11A`, `11N-AC-AX`, `11AC-AX`, `11AC`, `11N-AC`, `11B-G-N-AX`, `11B-G-N`, `11B-G`, `11B`, `11G-N-AX`, `11N-AX`, `11AX`, `11G-N`, `11N`, `11G`.
* `guard_interval` - Wi-Fi guard interval. Valid values: `auto`, `400ns`, `800ns`.
* `channel` - Wi-Fi channels. Valid values: `CH36`, `CH40`, `CH44`, `CH48`, `CH52`, `CH56`, `CH60`, `CH64`, `CH100`, `CH104`, `CH108`, `CH112`, `CH116`, `CH120`, `CH124`, `CH128`, `CH132`, `CH136`, `CH140`, `CH144`, `CH149`, `CH153`, `CH157`, `CH161`, `CH165`.
* `bandwidth` - Wi-Fi channel bandwidth. Valid values: `auto`, `20MHz`, `40MHz`, `80MHz`.
* `power_level` - Wi-Fi power level in percent (0 - 100, 0 = auto, default = 100).
* `beacon_interval` - Wi-Fi beacon interval in miliseconds (100 - 3500, default = 100).
* `n80211d` - Enable/disable Wi-Fi 802.11d. Valid values: `disable`, `enable`.
* `max_clients` - Maximum number of Wi-Fi radio clients (0 - 512, 0 = unlimited, default = 0).
* `extension_channel` - Wi-Fi extension channel. Valid values: `auto`, `higher`, `lower`.
* `bss_color_mode` - Wi-Fi 802.11AX BSS color mode. Valid values: `auto`, `static`.
* `bss_color` - Wi-Fi 802.11AX BSS color value (0 - 63, 0 = disable, default = 0).
* `lan_ext_vap` - Wi-Fi LAN-Extention VAP. Select only one VAP.
* `local_vaps` - Wi-Fi local VAP. Select up to three VAPs. The structure of `local_vaps` block is documented below.

The `local_vaps` block supports:

* `name` - Wi-Fi local VAP name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController ExtenderProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_extensioncontroller_extenderprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_extensioncontroller_extenderprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
