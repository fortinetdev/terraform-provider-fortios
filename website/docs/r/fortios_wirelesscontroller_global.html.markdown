---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_global"
description: |-
  Configure wireless controller global settings.
---

# fortios_wirelesscontroller_global
Configure wireless controller global settings.

## Example Usage

```hcl
resource "fortios_wirelesscontroller_global" "trname" {
  ap_log_server            = "disable"
  ap_log_server_ip         = "0.0.0.0"
  ap_log_server_port       = 0
  control_message_offload  = "ebp-frame aeroscout-tag ap-list sta-list sta-cap-list stats aeroscout-mu"
  data_ethernet_ii         = "disable"
  discovery_mc_addr        = "224.0.1.140"
  fiapp_eth_type           = 5252
  image_download           = "enable"
  ipsec_base_ip            = "169.254.0.1"
  link_aggregation         = "disable"
  max_clients              = 0
  max_retransmit           = 3
  mesh_eth_type            = 8755
  rogue_scan_mac_adjacency = 7
  wtp_share                = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of the wireless controller.
* `location` - Description of the location of the wireless controller.
* `acd_process_count` - Configure the number cw_acd daemons for multi-core CPU support (default = 0).
* `wpad_process_count` - Wpad daemon process count for multi-core CPU support.
* `image_download` - Enable/disable WTP image download at join time. Valid values: `enable`, `disable`.
* `rolling_wtp_upgrade` - Enable/disable rolling WTP upgrade (default = disable). Valid values: `enable`, `disable`.
* `rolling_wtp_upgrade_threshold` - Minimum signal level/threshold in dBm required for the managed WTP to be included in rolling WTP upgrade (-95 to -20, default = -80).
* `max_retransmit` - Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
* `control_message_offload` - Configure CAPWAP control message data channel offload.
* `data_ethernet_ii` - Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = enable). Valid values: `enable`, `disable`.
* `link_aggregation` - Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `enable`, `disable`.
* `mesh_eth_type` - Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
* `fiapp_eth_type` - Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
* `discovery_mc_addr` - Multicast IP address for AP discovery (default = 244.0.1.140).
* `max_clients` - Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
* `rogue_scan_mac_adjacency` - Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
* `ipsec_base_ip` - Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
* `wtp_share` - Enable/disable sharing of WTPs between VDOMs. Valid values: `enable`, `disable`.
* `tunnel_mode` - Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
* `nac_interval` - Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
* `ap_log_server` - Enable/disable configuring FortiGate to redirect wireless event log messages or FortiAPs to send UTM log messages to a syslog server (default = disable). Valid values: `enable`, `disable`.
* `ap_log_server_ip` - IP address that APs or FortiAPs send log messages to.
* `ap_log_server_port` - Port that APs or FortiAPs send log messages to.
* `max_sta_cap` - Maximum number of station cap stored on the controller (default = 0).
* `max_sta_cap_wtp` - Maximum number of station cap's wtp info stored on the controller (1 - 16, default = 8).
* `max_rogue_ap` - Maximum number of rogue APs stored on the controller (default = 0).
* `max_rogue_ap_wtp` - Maximum number of rogue AP's wtp info stored on the controller (1 - 16, default = 16).
* `max_rogue_sta` - Maximum number of rogue stations stored on the controller (default = 0).
* `max_wids_entry` - Maximum number of wids entries stored on the controller (default = 0).
* `max_ble_device` - Maximum number of BLE devices stored on the controller (default = 0).
* `dfs_lab_test` - Enable/disable DFS certificate lab test mode. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Global can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_global.labelname WirelessControllerGlobal

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_global.labelname WirelessControllerGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
