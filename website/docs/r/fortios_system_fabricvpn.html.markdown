---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fabricvpn"
description: |-
  Setup for self orchestrated fabric auto discovery VPN.
---

# fortios_system_fabricvpn
Setup for self orchestrated fabric auto discovery VPN. Applies to FortiOS Version `>= 7.2.4`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable Fabric VPN. Valid values: `enable`, `disable`.
* `sync_mode` - Setting synchronised by fabric or manual. Valid values: `enable`, `disable`.
* `branch_name` - Branch name.
* `policy_rule` - Policy creation rule. Valid values: `health-check`, `manual`, `auto`.
* `vpn_role` - Fabric VPN role. Valid values: `hub`, `spoke`.
* `overlays` - Local overlay interfaces table. The structure of `overlays` block is documented below.
* `advertised_subnets` - Local advertised subnets. The structure of `advertised_subnets` block is documented below.
* `loopback_address_block` - IPv4 address and subnet mask for hub's loopback address, syntax: X.X.X.X/24.
* `loopback_interface` - Loopback interface.
* `loopback_advertised_subnet` - Loopback advertised subnet reference.
* `psksecret` - Pre-shared secret for ADVPN.
* `bgp_as` - BGP Router AS number, valid from 1 to 4294967295.
* `sdwan_zone` - Reference to created SD-WAN zone.
* `health_checks` - Underlying health checks.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `overlays` block supports:

* `name` - Overlay name.
* `overlay_tunnel_block` - IPv4 address and subnet mask for the overlay tunnel , syntax: X.X.X.X/24.
* `remote_gw` - IP address of the hub gateway (Set by hub).
* `interface` - Underlying interface name.
* `bgp_neighbor` - Underlying BGP neighbor entry.
* `overlay_policy` - The overlay policy to allow ADVPN thru traffic.
* `bgp_network` - Underlying BGP network.
* `route_policy` - Underlying router policy.
* `bgp_neighbor_group` - Underlying BGP neighbor group entry.
* `bgp_neighbor_range` - Underlying BGP neighbor range entry.
* `ipsec_phase1` - IPsec interface.
* `sdwan_member` - Reference to SD-WAN member entry.

The `advertised_subnets` block supports:

* `id` - ID.
* `prefix` - Network prefix.
* `access` - Access policy direction. Valid values: `inbound`, `bidirectional`.
* `bgp_network` - Underlying BGP network.
* `firewall_address` - Underlying firewall address.
* `policies` - Underlying policies.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FabricVpn can be imported using any of these accepted formats:
```
$ terraform import fortios_system_fabricvpn.labelname SystemFabricVpn

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_fabricvpn.labelname SystemFabricVpn
$ unset "FORTIOS_IMPORT_TABLE"
```
