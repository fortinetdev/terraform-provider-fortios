---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_ocvpn"
description: |-
  Configure Overlay Controller VPN settings.
---

# fortios_vpn_ocvpn
Configure Overlay Controller VPN settings.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable Overlay Controller cloud assisted VPN. Valid values: `enable`, `disable`.
* `role` - Set device role. Valid values: `spoke`, `primary-hub`, `secondary-hub`.
* `multipath` - Enable/disable multipath redundancy. Valid values: `enable`, `disable`.
* `sdwan` - Enable/disable adding OCVPN tunnels to SDWAN. Valid values: `enable`, `disable`.
* `wan_interface` - FortiGate WAN interfaces to use with OCVPN. The structure of `wan_interface` block is documented below.
* `ip_allocation_block` - Class B subnet reserved for private IP address assignment.
* `poll_interval` - Overlay Controller VPN polling interval.
* `auto_discovery` - Enable/disable auto-discovery shortcuts. Valid values: `enable`, `disable`.
* `eap` - Enable/disable EAP client authentication. Valid values: `enable`, `disable`.
* `eap_users` - EAP authentication user group.
* `nat` - Enable/disable inter-overlay source NAT. Valid values: `enable`, `disable`.
* `overlays` - Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.
* `forticlient_access` - Configure FortiClient settings. The structure of `forticlient_access` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpn_ocvpn.labelname VpnOcvpn
$ unset "FORTIOS_IMPORT_TABLE"
```
