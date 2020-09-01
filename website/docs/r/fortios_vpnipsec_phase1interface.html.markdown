---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_phase1interface"
description: |-
  Configure VPN remote gateway.
---

# fortios_vpnipsec_phase1interface
Configure VPN remote gateway.

## Example Usage

```hcl
resource "fortios_vpnipsec_phase1interface" "trname2" {
  acct_verify               = "disable"
  add_gw_route              = "disable"
  add_route                 = "enable"
  assign_ip                 = "enable"
  assign_ip_from            = "range"
  authmethod                = "psk"
  auto_discovery_forwarder  = "disable"
  auto_discovery_psk        = "disable"
  auto_discovery_receiver   = "disable"
  auto_discovery_sender     = "disable"
  auto_negotiate            = "enable"
  cert_id_validation        = "enable"
  childless_ike             = "disable"
  client_auto_negotiate     = "disable"
  client_keep_alive         = "disable"
  default_gw                = "0.0.0.0"
  default_gw_priority       = 0
  dhgrp                     = "14 5"
  digital_signature_auth    = "disable"
  distance                  = 15
  dns_mode                  = "manual"
  dpd                       = "on-demand"
  dpd_retrycount            = 3
  dpd_retryinterval         = "20"
  eap                       = "disable"
  eap_identity              = "use-id-payload"
  encap_local_gw4           = "0.0.0.0"
  encap_local_gw6           = "::"
  encap_remote_gw4          = "0.0.0.0"
  encap_remote_gw6          = "::"
  encapsulation             = "none"
  encapsulation_address     = "ike"
  enforce_unique_id         = "disable"
  exchange_interface_ip     = "disable"
  exchange_ip_addr4         = "0.0.0.0"
  exchange_ip_addr6         = "::"
  forticlient_enforcement   = "disable"
  fragmentation             = "enable"
  fragmentation_mtu         = 1200
  group_authentication      = "disable"
  ha_sync_esp_seqno         = "enable"
  idle_timeout              = "disable"
  idle_timeoutinterval      = 15
  ike_version               = "1"
  include_local_lan         = "disable"
  interface                 = "port3"
  ip_version                = "4"
  ipv4_dns_server1          = "0.0.0.0"
  ipv4_dns_server2          = "0.0.0.0"
  ipv4_dns_server3          = "0.0.0.0"
  ipv4_end_ip               = "0.0.0.0"
  ipv4_netmask              = "255.255.255.255"
  ipv4_start_ip             = "0.0.0.0"
  ipv4_wins_server1         = "0.0.0.0"
  ipv4_wins_server2         = "0.0.0.0"
  ipv6_dns_server1          = "::"
  ipv6_dns_server2          = "::"
  ipv6_dns_server3          = "::"
  ipv6_end_ip               = "::"
  ipv6_prefix               = 128
  ipv6_start_ip             = "::"
  keepalive                 = 10
  keylife                   = 86400
  local_gw                  = "0.0.0.0"
  local_gw6                 = "::"
  localid_type              = "auto"
  mesh_selector_type        = "disable"
  mode                      = "main"
  mode_cfg                  = "disable"
  monitor_hold_down_delay   = 0
  monitor_hold_down_time    = "00:00"
  monitor_hold_down_type    = "immediate"
  monitor_hold_down_weekday = "sunday"
  name                      = "t0"
  nattraversal              = "enable"
  negotiate_timeout         = 30
  net_device                = "disable"
  passive_mode              = "disable"
  peertype                  = "any"
  ppk                       = "disable"
  priority                  = 0
  proposal                  = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
  psksecret                 = "eweeeeeeeecee"
  reauth                    = "disable"
  rekey                     = "enable"
  remote_gw                 = "102.2.2.12"
  remote_gw6                = "::"
  rsa_signature_format      = "pkcs1"
  save_password             = "disable"
  send_cert_chain           = "enable"
  signature_hash_alg        = "sha2-512 sha2-384 sha2-256 sha1"
  suite_b                   = "disable"
  tunnel_search             = "selectors"
  type                      = "static"
  unity_support             = "enable"
  wizard_type               = "custom"
  xauthtype                 = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IPsec remote gateway name.
* `type` - Remote gateway type.
* `interface` - (Required) Local physical, aggregate, or VLAN outgoing interface.
* `ip_version` - IP version to use for VPN interface.
* `ike_version` - IKE protocol version.
* `local_gw` - IPv4 address of the local gateway's external interface.
* `local_gw6` - IPv6 address of the local gateway's external interface.
* `remote_gw` - (Required) IPv4 address of the remote gateway's external interface.
* `remote_gw6` - IPv6 address of the remote gateway's external interface.
* `remotegw_ddns` - Domain name of remote gateway (eg. name.DDNS.com).
* `keylife` - Time to wait in seconds before phase 1 encryption key expires.
* `certificate` - The names of up to 4 signed personal certificates.
* `authmethod` - Authentication method.
* `authmethod_remote` - Authentication method (remote side).
* `mode` - The ID protection mode used to establish a secure channel.
* `peertype` - Accept this peer type.
* `peerid` - Accept this peer identity.
* `default_gw` - IPv4 address of default route gateway to use for traffic exiting the interface.
* `default_gw_priority` - Priority for default gateway route. A higher priority number signifies a less preferred route.
* `usrgrp` - User group name for dialup peers.
* `peer` - Accept this peer certificate.
* `peergrp` - Accept this peer certificate group.
* `monitor` - IPsec interface as backup for primary interface.
* `monitor_hold_down_type` - Recovery time method when primary interface re-establishes.
* `monitor_hold_down_delay` - Time to wait in seconds before recovery once primary re-establishes.
* `monitor_hold_down_weekday` - Day of the week to recover once primary re-establishes.
* `monitor_hold_down_time` - Time of day at which to fail back to primary after it re-establishes.
* `net_device` - (Required) Enable/disable kernel device creation.
* `tunnel_search` - Tunnel search method for when the interface is shared.
* `passive_mode` - Enable/disable IPsec passive mode for static tunnels.
* `exchange_interface_ip` - Enable/disable exchange of IPsec interface IP address.
* `exchange_ip_addr4` - IPv4 address to exchange with peers.
* `exchange_ip_addr6` - IPv6 address to exchange with peers
* `mode_cfg` - Enable/disable configuration method.
* `assign_ip` - Enable/disable assignment of IP to IPsec interface via configuration method.
* `assign_ip_from` - Method by which the IP address will be assigned.
* `ipv4_start_ip` - Start of IPv4 range.
* `ipv4_end_ip` - End of IPv4 range.
* `ipv4_netmask` - IPv4 Netmask.
* `dns_mode` - DNS server mode.
* `ipv4_dns_server1` - IPv4 DNS server 1.
* `ipv4_dns_server2` - IPv4 DNS server 2.
* `ipv4_dns_server3` - IPv4 DNS server 3.
* `ipv4_wins_server1` - WINS server 1.
* `ipv4_wins_server2` - WINS server 2.
* `ipv4_exclude_range` - Configuration Method IPv4 exclude ranges.
* `ipv4_split_include` - IPv4 split-include subnets.
* `split_include_service` - Split-include services.
* `ipv4_name` - IPv4 address name.
* `ipv6_start_ip` - Start of IPv6 range.
* `ipv6_end_ip` - End of IPv6 range.
* `ipv6_prefix` - IPv6 prefix.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_dns_server3` - IPv6 DNS server 3.
* `ipv6_exclude_range` - Configuration method IPv6 exclude ranges.
* `ipv6_split_include` - IPv6 split-include subnets.
* `ipv6_name` - IPv6 address name.
* `unity_support` - Enable/disable support for Cisco UNITY Configuration Method extensions.
* `domain` - Instruct unity clients about the default DNS domain.
* `banner` - Message that unity client should display after connecting.
* `include_local_lan` - Enable/disable allow local LAN access on unity clients.
* `ipv4_split_exclude` - IPv4 subnets that should not be sent over the IPsec tunnel.
* `ipv6_split_exclude` - IPv6 subnets that should not be sent over the IPsec tunnel.
* `save_password` - Enable/disable saving XAuth username and password on VPN clients.
* `client_auto_negotiate` - Enable/disable allowing the VPN client to bring up the tunnel when there is no traffic.
* `client_keep_alive` - Enable/disable allowing the VPN client to keep the tunnel up when there is no traffic.
* `backup_gateway` - Instruct unity clients about the backup gateway address(es).
* `proposal` - (Required) Phase1 proposal.
* `add_route` - Enable/disable control addition of a route to peer destination selector.
* `add_gw_route` - Enable/disable automatically add a route to the remote gateway.
* `psksecret` - Pre-shared secret for PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `psksecret_remote` - Pre-shared secret for remote side PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `keepalive` - NAT-T keep alive interval.
* `distance` - Distance for routes added by IKE (1 - 255).
* `priority` - Priority for routes added by IKE (0 - 4294967295).
* `localid` - Local ID.
* `localid_type` - Local ID type.
* `auto_negotiate` - Enable/disable automatic initiation of IKE SA negotiation.
* `negotiate_timeout` - IKE SA negotiation timeout in seconds (1 - 300).
* `fragmentation` - Enable/disable fragment IKE message on re-transmission.
* `dpd` - Dead Peer Detection mode.
* `dpd_retrycount` - Number of DPD retry attempts.
* `dpd_retryinterval` - DPD retry interval.
* `forticlient_enforcement` - Enable/disable FortiClient enforcement.
* `comments` - Comment.
* `send_cert_chain` - Enable/disable sending certificate chain.
* `dhgrp` - DH group.
* `suite_b` - Use Suite-B.
* `eap` - Enable/disable IKEv2 EAP authentication.
* `eap_identity` - IKEv2 EAP peer identity type.
* `acct_verify` - Enable/disable verification of RADIUS accounting record.
* `ppk` - Enable/disable IKEv2 Postquantum Preshared Key (PPK).
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `wizard_type` - GUI VPN Wizard Type.
* `xauthtype` - XAuth type.
* `reauth` - Enable/disable re-authentication upon IKE SA lifetime expiration.
* `authusr` - XAuth user name.
* `authpasswd` - XAuth password (max 35 characters).
* `group_authentication` - Enable/disable IKEv2 IDi group authentication.
* `group_authentication_secret` - Password for IKEv2 IDi group authentication.  (ASCII string or hexadecimal indicated by a leading 0x.)
* `authusrgrp` - Authentication user group.
* `mesh_selector_type` - Add selectors containing subsets of the configuration depending on traffic.
* `idle_timeout` - Enable/disable IPsec tunnel idle timeout.
* `idle_timeoutinterval` - IPsec tunnel idle timeout in minutes (5 - 43200).
* `ha_sync_esp_seqno` - Enable/disable sequence number jump ahead for IPsec HA.
* `auto_discovery_sender` - Enable/disable sending auto-discovery short-cut messages.
* `auto_discovery_receiver` - Enable/disable accepting auto-discovery short-cut messages.
* `auto_discovery_forwarder` - Enable/disable forwarding auto-discovery short-cut messages.
* `auto_discovery_psk` - Enable/disable use of pre-shared secrets for authentication of auto-discovery tunnels.
* `encapsulation` - Enable/disable GRE/VXLAN encapsulation.
* `encapsulation_address` - Source for GRE/VXLAN tunnel address.
* `encap_local_gw4` - Local IPv4 address of GRE/VXLAN tunnel.
* `encap_local_gw6` - Local IPv6 address of GRE/VXLAN tunnel.
* `encap_remote_gw4` - Remote IPv4 address of GRE/VXLAN tunnel.
* `encap_remote_gw6` - Remote IPv6 address of GRE/VXLAN tunnel.
* `vni` - VNI of VXLAN tunnel.
* `nattraversal` - Enable/disable NAT traversal.
* `fragmentation_mtu` - IKE fragmentation MTU (500 - 16000).
* `childless_ike` - Enable/disable childless IKEv2 initiation (RFC 6023).
* `rekey` - Enable/disable phase1 rekey.
* `digital_signature_auth` - Enable/disable IKEv2 Digital Signature Authentication (RFC 7427).
* `signature_hash_alg` - Digital Signature Authentication hash algorithms.
* `rsa_signature_format` - Digital Signature Authentication RSA signature format.
* `enforce_unique_id` - Enable/disable peer ID uniqueness check.
* `cert_id_validation` - Enable/disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945.

The `certificate` block supports:

* `name` - Certificate name.

The `ipv4_exclude_range` block supports:

* `id` - ID.
* `start_ip` - Start of IPv4 exclusive range.
* `end_ip` - End of IPv4 exclusive range.

The `ipv6_exclude_range` block supports:

* `id` - ID.
* `start_ip` - Start of IPv6 exclusive range.
* `end_ip` - End of IPv6 exclusive range.

The `backup_gateway` block supports:

* `address` - Address of backup gateway.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnIpsec Phase1Interface can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnipsec_phase1interface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
