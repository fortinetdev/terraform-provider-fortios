---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_hsprofile"
description: |-
  Configure hotspot profile.
---

# fortios_wirelesscontrollerhotspot20_hsprofile
Configure hotspot profile.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Hotspot profile name.
* `access_network_type` - Access network type.
* `access_network_internet` - Enable/disable connectivity to the Internet.
* `access_network_asra` - Enable/disable additional step required for access (ASRA).
* `access_network_esr` - Enable/disable emergency services reachable (ESR).
* `access_network_uesa` - Enable/disable unauthenticated emergency service accessible (UESA).
* `venue_group` - Venue group.
* `venue_type` - Venue type.
* `hessid` - Homogeneous extended service set identifier (HESSID).
* `proxy_arp` - Enable/disable Proxy ARP.
* `l2tif` - Enable/disable Layer 2 traffic inspection and filtering.
* `pame_bi` - Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI).
* `anqp_domain_id` - ANQP Domain ID (0-65535).
* `domain_name` - Domain name.
* `osu_ssid` - Online sign up (OSU) SSID.
* `gas_comeback_delay` - GAS comeback delay (0 or 100 - 4000 milliseconds, default = 500).
* `gas_fragmentation_limit` - GAS fragmentation limit (512 - 4096, default = 1024).
* `dgaf` - Enable/disable downstream group-addressed forwarding (DGAF).
* `deauth_request_timeout` - Deauthentication request timeout (in seconds).
* `wnm_sleep_mode` - Enable/disable wireless network management (WNM) sleep mode.
* `bss_transition` - Enable/disable basic service set (BSS) transition Support.
* `venue_name` - Venue name.
* `roaming_consortium` - Roaming consortium list name.
* `nai_realm` - NAI realm list name.
* `oper_friendly_name` - Operator friendly name.
* `osu_provider` - Manually selected list of OSU provider(s). The structure of `osu_provider` block is documented below.
* `wan_metrics` - WAN metric name.
* `network_auth` - Network authentication name.
* `n3gpp_plmn` - 3GPP PLMN name.
* `conn_cap` - Connection capability name.
* `qos_map` - QoS MAP set ID.
* `ip_addr_type` - IP address type name.

The `osu_provider` block supports:

* `name` - OSU provider name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 HsProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_hsprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
