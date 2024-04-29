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
* `burst` - Enable/disable client rate burst. Valid values: `enable`, `disable`.
* `wmm` - Enable/disable WiFi multi-media (WMM) control. Valid values: `enable`, `disable`.
* `wmm_uapsd` - Enable/disable WMM Unscheduled Automatic Power Save Delivery (U-APSD) power save mode. Valid values: `enable`, `disable`.
* `call_admission_control` - Enable/disable WMM call admission control. Valid values: `enable`, `disable`.
* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones allowed (0 - 60, default = 10).
* `bandwidth_admission_control` - Enable/disable WMM bandwidth admission control. Valid values: `enable`, `disable`.
* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `dscp_wmm_mapping` - Enable/disable Differentiated Services Code Point (DSCP) mapping. Valid values: `enable`, `disable`.
* `dscp_wmm_vo` - DSCP mapping for voice access (default = 48 56). The structure of `dscp_wmm_vo` block is documented below.
* `dscp_wmm_vi` - DSCP mapping for video access (default = 32 40). The structure of `dscp_wmm_vi` block is documented below.
* `dscp_wmm_be` - DSCP mapping for best effort access (default = 0 24). The structure of `dscp_wmm_be` block is documented below.
* `dscp_wmm_bk` - DSCP mapping for background access (default = 8 16). The structure of `dscp_wmm_bk` block is documented below.
* `wmm_dscp_marking` - Enable/disable WMM Differentiated Services Code Point (DSCP) marking. Valid values: `enable`, `disable`.
* `wmm_vo_dscp` - DSCP marking for voice access (default = 48).
* `wmm_vi_dscp` - DSCP marking for video access (default = 32).
* `wmm_be_dscp` - DSCP marking for best effort access (default = 0).
* `wmm_bk_dscp` - DSCP marking for background access (default = 8).
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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
$ terraform import fortios_wirelesscontroller_qosprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_qosprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
