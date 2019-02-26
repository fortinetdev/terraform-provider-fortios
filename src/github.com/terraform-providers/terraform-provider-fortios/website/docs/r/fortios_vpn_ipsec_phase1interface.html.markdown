---
layout: "fortios"
page_title: "FortiOS: fortios_vpn_ipsec_phase1interface"
sidebar_current: "docs-fortios-resource-vpn-ipsec-phase1interface"
description: |-
  Provides a resource to use phase1-interface to define a phase 1 definition for a route-based (interface mode) IPsec VPN tunnel that generates authentication and encryption keys automatically.
---

# fortios_vpn_ipsec_phase1interface

Provides a resource to use phase1-interface to define a phase 1 definition for a route-based (interface mode) IPsec VPN tunnel that generates authentication and encryption keys automatically.

## Example Usage
fortios_vpn_ipsec_phase1interface needs to be set with fortios_vpn_ipsec_phase2interface. See section fortios_vpn_ipsec_phase2interface.

## Argument Reference
The following arguments are supported:
* `name` - (Required) IPsec remote gateway name.
* `type` - (Required) Remote gateway type.
* `interface` - (Required) Local physical, aggregate, or VLAN outgoing interface.
* `peertype` - Accept this peer type.
* `proposal` - Phase1 proposal.
* `comments` - Comment.
* `wizard_type` - GUI VPN Wizard Type.
* `remote_gw` - (Required) IPv4 address of the remote gateway's external interface.
* `psksecret` - (Required) Pre-shared secret for PSK authentication.
* `certificate` - Names of signed personal certificates.
* `peerid` - Accept this peer identity.
* `peer` - Accept this peer certificate.
* `peergrp` - Accept this peer certificate group.
* `ipv4_split_include` - IPv4 split-include subnets.
* `split_include_service` - Split-include services.
* `ipv4_split_exclude` - IPv4 subnets that should not be sent over the IPsec tunnel.

## Attributes Reference
The following attributes are exported:
* `id` - The ID of the phase1-interface item.
* `name` - IPsec remote gateway name.
* `type` - Remote gateway type.
* `interface` - Local physical, aggregate, or VLAN outgoing interface.
* `peertype` - Accept this peer type.
* `proposal` - Phase1 proposal.
* `comments` - Comment.
* `wizard_type` - GUI VPN Wizard Type.
* `remote_gw` - IPv4 address of the remote gateway's external interface.
* `psksecret` - Pre-shared secret for PSK authentication.
* `certificate` - Names of signed personal certificates.
* `peerid` - Accept this peer identity.
* `peer` - Accept this peer certificate.
* `peergrp` - Accept this peer certificate group.
* `ipv4_split_include` - IPv4 split-include subnets.
* `split_include_service` - Split-include services.
* `ipv4_split_exclude` - IPv4 subnets that should not be sent over the IPsec tunnel.

