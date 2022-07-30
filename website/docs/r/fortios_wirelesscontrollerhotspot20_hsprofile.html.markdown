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

* `name` - Hotspot profile name.
* `release` - Hotspot 2.0 Release number (1, 2, 3, default = 2).
* `access_network_type` - Access network type. Valid values: `private-network`, `private-network-with-guest-access`, `chargeable-public-network`, `free-public-network`, `personal-device-network`, `emergency-services-only-network`, `test-or-experimental`, `wildcard`.
* `access_network_internet` - Enable/disable connectivity to the Internet. Valid values: `enable`, `disable`.
* `access_network_asra` - Enable/disable additional step required for access (ASRA). Valid values: `enable`, `disable`.
* `access_network_esr` - Enable/disable emergency services reachable (ESR). Valid values: `enable`, `disable`.
* `access_network_uesa` - Enable/disable unauthenticated emergency service accessible (UESA). Valid values: `enable`, `disable`.
* `venue_group` - Venue group. Valid values: `unspecified`, `assembly`, `business`, `educational`, `factory`, `institutional`, `mercantile`, `residential`, `storage`, `utility`, `vehicular`, `outdoor`.
* `venue_type` - Venue type. Valid values: `unspecified`, `arena`, `stadium`, `passenger-terminal`, `amphitheater`, `amusement-park`, `place-of-worship`, `convention-center`, `library`, `museum`, `restaurant`, `theater`, `bar`, `coffee-shop`, `zoo-or-aquarium`, `emergency-center`, `doctor-office`, `bank`, `fire-station`, `police-station`, `post-office`, `professional-office`, `research-facility`, `attorney-office`, `primary-school`, `secondary-school`, `university-or-college`, `factory`, `hospital`, `long-term-care-facility`, `rehab-center`, `group-home`, `prison-or-jail`, `retail-store`, `grocery-market`, `auto-service-station`, `shopping-mall`, `gas-station`, `private`, `hotel-or-motel`, `dormitory`, `boarding-house`, `automobile`, `airplane`, `bus`, `ferry`, `ship-or-boat`, `train`, `motor-bike`, `muni-mesh-network`, `city-park`, `rest-area`, `traffic-control`, `bus-stop`, `kiosk`.
* `hessid` - Homogeneous extended service set identifier (HESSID).
* `proxy_arp` - Enable/disable Proxy ARP. Valid values: `enable`, `disable`.
* `l2tif` - Enable/disable Layer 2 traffic inspection and filtering. Valid values: `enable`, `disable`.
* `pame_bi` - Enable/disable Pre-Association Message Exchange BSSID Independent (PAME-BI). Valid values: `disable`, `enable`.
* `anqp_domain_id` - ANQP Domain ID (0-65535).
* `domain_name` - Domain name.
* `osu_ssid` - Online sign up (OSU) SSID.
* `gas_comeback_delay` - GAS comeback delay (0 or 100 - 4000 milliseconds, default = 500).
* `gas_fragmentation_limit` - GAS fragmentation limit (512 - 4096, default = 1024).
* `dgaf` - Enable/disable downstream group-addressed forwarding (DGAF). Valid values: `enable`, `disable`.
* `deauth_request_timeout` - Deauthentication request timeout (in seconds).
* `wnm_sleep_mode` - Enable/disable wireless network management (WNM) sleep mode. Valid values: `enable`, `disable`.
* `bss_transition` - Enable/disable basic service set (BSS) transition Support. Valid values: `enable`, `disable`.
* `venue_name` - Venue name.
* `venue_url` - Venue name.
* `roaming_consortium` - Roaming consortium list name.
* `nai_realm` - NAI realm list name.
* `oper_friendly_name` - Operator friendly name.
* `oper_icon` - Operator icon.
* `advice_of_charge` - Advice of charge.
* `osu_provider_nai` - OSU Provider NAI.
* `terms_and_conditions` - Terms and conditions.
* `osu_provider` - Manually selected list of OSU provider(s). The structure of `osu_provider` block is documented below.
* `wan_metrics` - WAN metric name.
* `network_auth` - Network authentication name.
* `n3gpp_plmn` - 3GPP PLMN name.
* `conn_cap` - Connection capability name.
* `qos_map` - QoS MAP set ID.
* `ip_addr_type` - IP address type name.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `osu_provider` block supports:

* `name` - OSU provider name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 HsProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_hsprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_hsprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
