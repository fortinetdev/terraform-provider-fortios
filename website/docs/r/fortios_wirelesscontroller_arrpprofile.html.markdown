---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_arrpprofile"
description: |-
  Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.
---

# fortios_wirelesscontroller_arrpprofile
Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles. Applies to FortiOS Version `>= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `name` - WiFi ARRP profile name.
* `comment` - Comment.
* `selection_period` - Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
* `monitor_period` - Period in seconds to measure average transmit retries and receive errors (default = 300).
* `weight_managed_ap` - Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
* `weight_rogue_ap` - Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
* `weight_noise_floor` - Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
* `weight_channel_load` - Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
* `weight_spectral_rssi` - Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
* `weight_weather_channel` - Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).
* `weight_dfs_channel` - Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
* `threshold_ap` - Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
* `threshold_noise_floor` - Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
* `threshold_channel_load` - Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
* `threshold_spectral_rssi` - Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
* `threshold_tx_retries` - Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
* `threshold_rx_errors` - Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
* `include_weather_channel` - Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).
* `include_dfs_channel` - Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController ArrpProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_arrpprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
