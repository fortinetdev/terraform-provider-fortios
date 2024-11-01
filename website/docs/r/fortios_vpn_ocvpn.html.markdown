---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_ocvpn"
description: |-
  Configure Overlay Controller VPN settings.
---

# fortios_vpn_ocvpn
Configure Overlay Controller VPN settings. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.0.14,7.0.15,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.2.7,7.2.8,7.2.9,7.2.10`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable Overlay Controller cloud assisted VPN. Valid values: `enable`, `disable`.
* `role` - Set device role. Valid values: `spoke`, `primary-hub`, `secondary-hub`.
* `multipath` - Enable/disable multipath redundancy. Valid values: `enable`, `disable`.
* `sdwan` - Enable/disable adding OCVPN tunnels to SDWAN. Valid values: `enable`, `disable`.
* `sdwan_zone` - Set SD-WAN zone.
* `wan_interface` - FortiGate WAN interfaces to use with OCVPN. The structure of `wan_interface` block is documented below.
* `ip_allocation_block` - Class B subnet reserved for private IP address assignment.
* `poll_interval` - Overlay Controller VPN polling interval.
* `auto_discovery` - Enable/disable auto-discovery shortcuts. Valid values: `enable`, `disable`.
* `auto_discovery_shortcut_mode` - Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.
* `eap` - Enable/disable EAP client authentication. Valid values: `enable`, `disable`.
* `eap_users` - EAP authentication user group.
* `nat` - Enable/disable inter-overlay source NAT. Valid values: `enable`, `disable`.
* `overlays` - Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.
* `forticlient_access` - Configure FortiClient settings. The structure of `forticlient_access` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `wan_interface` block supports:

* `name` - Interface name.

The `overlays` block supports:

* `overlay_name` - Overlay name.
* `inter_overlay` - Allow or deny traffic from other overlays. Valid values: `allow`, `deny`.
* `id` - ID.
* `name` - Overlay name.
* `assign_ip` - Enable/disable client address assignment. Valid values: `enable`, `disable`.
* `ipv4_start_ip` - Start of client IPv4 range.
* `ipv4_end_ip` - End of client IPv4 range.
* `subnets` - Internal subnets to register with OCVPN service. The structure of `subnets` block is documented below.

The `subnets` block supports:

* `id` - ID.
* `type` - Subnet type. Valid values: `subnet`, `interface`.
* `subnet` - IPv4 address and subnet mask.
* `interface` - LAN interface.

The `forticlient_access` block supports:

* `status` - Enable/disable FortiClient to access OCVPN networks. Valid values: `enable`, `disable`.
* `psksecret` - Pre-shared secret for FortiClient PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `auth_groups` - FortiClient user authentication groups. The structure of `auth_groups` block is documented below.

The `auth_groups` block supports:

* `name` - Group name.
* `auth_group` - Authentication user group for FortiClient access.
* `overlays` - OCVPN overlays to allow access to. The structure of `overlays` block is documented below.

The `overlays` block supports:

* `overlay_name` - Overlay name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn Ocvpn can be imported using any of these accepted formats:
```
$ terraform import fortios_vpn_ocvpn.labelname VpnOcvpn

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpn_ocvpn.labelname VpnOcvpn
$ unset "FORTIOS_IMPORT_TABLE"
```
