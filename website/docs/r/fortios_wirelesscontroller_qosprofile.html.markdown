---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_qosprofile"
description: |-
  Configure WiFi quality of service (QoS) profiles.
---

# fortios_wirelesscontroller_qosprofile
Configure WiFi quality of service (QoS) profiles.

## Argument Reference


The following arguments are supported:

* `name` - WiFi QoS profile name.
* `comment` - Comment.
* `uplink` - Maximum uplink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `downlink` - Maximum downlink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `uplink_sta` - Maximum uplink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `downlink_sta` - Maximum downlink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).
* `burst` - Enable/disable client rate burst.
* `wmm` - Enable/disable WiFi multi-media (WMM) control.
* `wmm_uapsd` - Enable/disable WMM Unscheduled Automatic Power Save Delivery (U-APSD) power save mode.
* `call_admission_control` - Enable/disable WMM call admission control.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones allowed (0 - 60, default = 10).
* `bandwidth_admission_control` - Enable/disable WMM bandwidth admission control.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `dscp_wmm_mapping` - Enable/disable Differentiated Services Code Point (DSCP) mapping.
* `dscp_wmm_vo` - DSCP mapping for voice access (default = 48 56). The structure of `dscp_wmm_vo` block is documented below.
* `dscp_wmm_vi` - DSCP mapping for video access (default = 32 40). The structure of `dscp_wmm_vi` block is documented below.
* `dscp_wmm_be` - DSCP mapping for best effort access (default = 0 24). The structure of `dscp_wmm_be` block is documented below.
* `dscp_wmm_bk` - DSCP mapping for background access (default = 8 16). The structure of `dscp_wmm_bk` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dscp_wmm_vo` block supports:

* `id` - DSCP WMM mapping numbers (0 - 63).

The `dscp_wmm_vi` block supports:

* `id` - DSCP WMM mapping numbers (0 - 63).

The `dscp_wmm_be` block supports:

* `id` - DSCP WMM mapping numbers (0 - 63).

The `dscp_wmm_bk` block supports:

* `id` - DSCP WMM mapping numbers (0 - 63).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController QosProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_qosprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
