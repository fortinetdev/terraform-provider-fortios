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
* `type` - Remote gateway type. Valid values: `static`, `dynamic`, `ddns`.
* `interface` - (Required) Local physical, aggregate, or VLAN outgoing interface.
* `ip_version` - IP version to use for VPN interface. Valid values: `4`, `6`.
* `ike_version` - IKE protocol version. Valid values: `1`, `2`.
* `local_gw` - IPv4 address of the local gateway's external interface.
* `local_gw6` - IPv6 address of the local gateway's external interface.
* `remote_gw` - IPv4 address of the remote gateway's external interface.
* `remote_gw6` - IPv6 address of the remote gateway's external interface.
* `remotegw_ddns` - Domain name of remote gateway (eg. name.DDNS.com).
* `keylife` - Time to wait in seconds before phase 1 encryption key expires.
* `certificate` - The names of up to 4 signed personal certificates. The structure of `certificate` block is documented below.
* `authmethod` - Authentication method. Valid values: `psk`, `signature`.
* `authmethod_remote` - Authentication method (remote side). Valid values: `psk`, `signature`.
* `mode` - The ID protection mode used to establish a secure channel. Valid values: `aggressive`, `main`.
* `peertype` - Accept this peer type. Valid values: `any`, `one`, `dialup`, `peer`, `peergrp`.
* `peerid` - Accept this peer identity.
* `default_gw` - IPv4 address of default route gateway to use for traffic exiting the interface.
* `default_gw_priority` - Priority for default gateway route. A higher priority number signifies a less preferred route.
* `usrgrp` - User group name for dialup peers.
* `peer` - Accept this peer certificate.
* `peergrp` - Accept this peer certificate group.
* `monitor` - IPsec interface as backup for primary interface.
* `monitor_hold_down_type` - Recovery time method when primary interface re-establishes. Valid values: `immediate`, `delay`, `time`.
* `monitor_hold_down_delay` - Time to wait in seconds before recovery once primary re-establishes.
* `monitor_hold_down_weekday` - Day of the week to recover once primary re-establishes. Valid values: `everyday`, `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `monitor_hold_down_time` - Time of day at which to fail back to primary after it re-establishes.
* `net_device` - Enable/disable kernel device creation. Valid values: `enable`, `disable`.
* `tunnel_search` - Tunnel search method for when the interface is shared. Valid values: `selectors`, `nexthop`.
* `passive_mode` - Enable/disable IPsec passive mode for static tunnels. Valid values: `enable`, `disable`.
* `exchange_interface_ip` - Enable/disable exchange of IPsec interface IP address. Valid values: `enable`, `disable`.
* `exchange_ip_addr4` - IPv4 address to exchange with peers.
* `exchange_ip_addr6` - IPv6 address to exchange with peers
* `aggregate_member` - Enable/disable use as an aggregate member. Valid values: `enable`, `disable`.
* `aggregate_weight` - Link weight for aggregate.
* `mode_cfg` - Enable/disable configuration method. Valid values: `disable`, `enable`.
* `mode_cfg_allow_client_selector` - Enable/disable mode-cfg client to use custom phase2 selectors. Valid values: `disable`, `enable`.
* `assign_ip` - Enable/disable assignment of IP to IPsec interface via configuration method. Valid values: `disable`, `enable`.
* `assign_ip_from` - Method by which the IP address will be assigned. Valid values: `range`, `usrgrp`, `dhcp`, `name`.
* `ipv4_start_ip` - Start of IPv4 range.
* `ipv4_end_ip` - End of IPv4 range.
* `ipv4_netmask` - IPv4 Netmask.
* `dhcp_ra_giaddr` - Relay agent gateway IP address to use in the giaddr field of DHCP requests.
* `dhcp6_ra_linkaddr` - Relay agent IPv6 link address to use in DHCP6 requests.
* `dns_mode` - DNS server mode. Valid values: `manual`, `auto`.
* `ipv4_dns_server1` - IPv4 DNS server 1.
* `ipv4_dns_server2` - IPv4 DNS server 2.
* `ipv4_dns_server3` - IPv4 DNS server 3.
* `ipv4_wins_server1` - WINS server 1.
* `ipv4_wins_server2` - WINS server 2.
* `ipv4_exclude_range` - Configuration Method IPv4 exclude ranges. The structure of `ipv4_exclude_range` block is documented below.
* `ipv4_split_include` - IPv4 split-include subnets.
* `split_include_service` - Split-include services.
* `ipv4_name` - IPv4 address name.
* `ipv6_start_ip` - Start of IPv6 range.
* `ipv6_end_ip` - End of IPv6 range.
* `ipv6_prefix` - IPv6 prefix.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_dns_server3` - IPv6 DNS server 3.
* `ipv6_exclude_range` - Configuration method IPv6 exclude ranges. The structure of `ipv6_exclude_range` block is documented below.
* `ipv6_split_include` - IPv6 split-include subnets.
* `ipv6_name` - IPv6 address name.
* `ip_delay_interval` - IP address reuse delay interval in seconds (0 - 28800).
* `unity_support` - Enable/disable support for Cisco UNITY Configuration Method extensions. Valid values: `disable`, `enable`.
* `domain` - Instruct unity clients about the default DNS domain.
* `banner` - Message that unity client should display after connecting.
* `include_local_lan` - Enable/disable allow local LAN access on unity clients. Valid values: `disable`, `enable`.
* `ipv4_split_exclude` - IPv4 subnets that should not be sent over the IPsec tunnel.
* `ipv6_split_exclude` - IPv6 subnets that should not be sent over the IPsec tunnel.
* `save_password` - Enable/disable saving XAuth username and password on VPN clients. Valid values: `disable`, `enable`.
* `client_auto_negotiate` - Enable/disable allowing the VPN client to bring up the tunnel when there is no traffic. Valid values: `disable`, `enable`.
* `client_keep_alive` - Enable/disable allowing the VPN client to keep the tunnel up when there is no traffic. Valid values: `disable`, `enable`.
* `backup_gateway` - Instruct unity clients about the backup gateway address(es). The structure of `backup_gateway` block is documented below.
* `proposal` - (Required) Phase1 proposal. Valid values: `des-md5`, `des-sha1`, `des-sha256`, `des-sha384`, `des-sha512`, `3des-md5`, `3des-sha1`, `3des-sha256`, `3des-sha384`, `3des-sha512`, `aes128-md5`, `aes128-sha1`, `aes128-sha256`, `aes128-sha384`, `aes128-sha512`, `aes128gcm-prfsha1`, `aes128gcm-prfsha256`, `aes128gcm-prfsha384`, `aes128gcm-prfsha512`, `aes192-md5`, `aes192-sha1`, `aes192-sha256`, `aes192-sha384`, `aes192-sha512`, `aes256-md5`, `aes256-sha1`, `aes256-sha256`, `aes256-sha384`, `aes256-sha512`, `aes256gcm-prfsha1`, `aes256gcm-prfsha256`, `aes256gcm-prfsha384`, `aes256gcm-prfsha512`, `chacha20poly1305-prfsha1`, `chacha20poly1305-prfsha256`, `chacha20poly1305-prfsha384`, `chacha20poly1305-prfsha512`, `aria128-md5`, `aria128-sha1`, `aria128-sha256`, `aria128-sha384`, `aria128-sha512`, `aria192-md5`, `aria192-sha1`, `aria192-sha256`, `aria192-sha384`, `aria192-sha512`, `aria256-md5`, `aria256-sha1`, `aria256-sha256`, `aria256-sha384`, `aria256-sha512`, `seed-md5`, `seed-sha1`, `seed-sha256`, `seed-sha384`, `seed-sha512`.
* `add_route` - Enable/disable control addition of a route to peer destination selector. Valid values: `disable`, `enable`.
* `add_gw_route` - Enable/disable automatically add a route to the remote gateway. Valid values: `enable`, `disable`.
* `psksecret` - Pre-shared secret for PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `psksecret_remote` - Pre-shared secret for remote side PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `keepalive` - NAT-T keep alive interval.
* `distance` - Distance for routes added by IKE (1 - 255).
* `priority` - Priority for routes added by IKE (0 - 4294967295).
* `localid` - Local ID.
* `localid_type` - Local ID type. Valid values: `auto`, `fqdn`, `user-fqdn`, `keyid`, `address`, `asn1dn`.
* `auto_negotiate` - Enable/disable automatic initiation of IKE SA negotiation. Valid values: `enable`, `disable`.
* `negotiate_timeout` - IKE SA negotiation timeout in seconds (1 - 300).
* `fragmentation` - Enable/disable fragment IKE message on re-transmission. Valid values: `enable`, `disable`.
* `ip_fragmentation` - Determine whether IP packets are fragmented before or after IPsec encapsulation. Valid values: `pre-encapsulation`, `post-encapsulation`.
* `dpd` - Dead Peer Detection mode. Valid values: `disable`, `on-idle`, `on-demand`.
* `dpd_retrycount` - Number of DPD retry attempts.
* `dpd_retryinterval` - DPD retry interval.
* `forticlient_enforcement` - Enable/disable FortiClient enforcement. Valid values: `enable`, `disable`.
* `comments` - Comment.
* `npu_offload` - Enable/disable offloading NPU. Valid values: `enable`, `disable`.
* `send_cert_chain` - Enable/disable sending certificate chain. Valid values: `enable`, `disable`.
* `dhgrp` - DH group. Valid values: `1`, `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `27`, `28`, `29`, `30`, `31`, `32`.
* `suite_b` - Use Suite-B. Valid values: `disable`, `suite-b-gcm-128`, `suite-b-gcm-256`.
* `eap` - Enable/disable IKEv2 EAP authentication. Valid values: `enable`, `disable`.
* `eap_identity` - IKEv2 EAP peer identity type. Valid values: `use-id-payload`, `send-request`.
* `eap_exclude_peergrp` - Peer group excluded from EAP authentication.
* `acct_verify` - Enable/disable verification of RADIUS accounting record. Valid values: `enable`, `disable`.
* `ppk` - Enable/disable IKEv2 Postquantum Preshared Key (PPK). Valid values: `disable`, `allow`, `require`.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `wizard_type` - GUI VPN Wizard Type.
* `xauthtype` - XAuth type. Valid values: `disable`, `client`, `pap`, `chap`, `auto`.
* `reauth` - Enable/disable re-authentication upon IKE SA lifetime expiration. Valid values: `disable`, `enable`.
* `authusr` - XAuth user name.
* `authpasswd` - XAuth password (max 35 characters).
* `group_authentication` - Enable/disable IKEv2 IDi group authentication. Valid values: `enable`, `disable`.
* `group_authentication_secret` - Password for IKEv2 IDi group authentication.  (ASCII string or hexadecimal indicated by a leading 0x.)
* `authusrgrp` - Authentication user group.
* `mesh_selector_type` - Add selectors containing subsets of the configuration depending on traffic. Valid values: `disable`, `subnet`, `host`.
* `idle_timeout` - Enable/disable IPsec tunnel idle timeout. Valid values: `enable`, `disable`.
* `idle_timeoutinterval` - IPsec tunnel idle timeout in minutes (5 - 43200).
* `ha_sync_esp_seqno` - Enable/disable sequence number jump ahead for IPsec HA. Valid values: `enable`, `disable`.
* `inbound_dscp_copy` - Enable/disable copy the dscp in the ESP header to the inner IP Header. Valid values: `enable`, `disable`.
* `auto_discovery_sender` - Enable/disable sending auto-discovery short-cut messages. Valid values: `enable`, `disable`.
* `auto_discovery_receiver` - Enable/disable accepting auto-discovery short-cut messages. Valid values: `enable`, `disable`.
* `auto_discovery_forwarder` - Enable/disable forwarding auto-discovery short-cut messages. Valid values: `enable`, `disable`.
* `auto_discovery_psk` - Enable/disable use of pre-shared secrets for authentication of auto-discovery tunnels. Valid values: `enable`, `disable`.
* `auto_discovery_shortcuts` - Control deletion of child short-cut tunnels when the parent tunnel goes down. Valid values: `independent`, `dependent`.
* `auto_discovery_offer_interval` - Interval between shortcut offer messages in seconds (1 - 300, default = 5).
* `encapsulation` - Enable/disable GRE/VXLAN encapsulation.
* `encapsulation_address` - Source for GRE/VXLAN tunnel address. Valid values: `ike`, `ipv4`, `ipv6`.
* `encap_local_gw4` - Local IPv4 address of GRE/VXLAN tunnel.
* `encap_local_gw6` - Local IPv6 address of GRE/VXLAN tunnel.
* `encap_remote_gw4` - Remote IPv4 address of GRE/VXLAN tunnel.
* `encap_remote_gw6` - Remote IPv6 address of GRE/VXLAN tunnel.
* `vni` - VNI of VXLAN tunnel.
* `nattraversal` - Enable/disable NAT traversal. Valid values: `enable`, `disable`, `forced`.
* `esn` - Extended sequence number (ESN) negotiation. Valid values: `require`, `allow`, `disable`.
* `fragmentation_mtu` - IKE fragmentation MTU (500 - 16000).
* `childless_ike` - Enable/disable childless IKEv2 initiation (RFC 6023). Valid values: `enable`, `disable`.
* `rekey` - Enable/disable phase1 rekey. Valid values: `enable`, `disable`.
* `digital_signature_auth` - Enable/disable IKEv2 Digital Signature Authentication (RFC 7427). Valid values: `enable`, `disable`.
* `signature_hash_alg` - Digital Signature Authentication hash algorithms. Valid values: `sha1`, `sha2-256`, `sha2-384`, `sha2-512`.
* `rsa_signature_format` - Digital Signature Authentication RSA signature format. Valid values: `pkcs1`, `pss`.
* `enforce_unique_id` - Enable/disable peer ID uniqueness check. Valid values: `disable`, `keep-new`, `keep-old`.
* `cert_id_validation` - Enable/disable cross validation of peer ID and the identity in the peer's certificate as specified in RFC 4945. Valid values: `enable`, `disable`.
* `fec_egress` - Enable/disable Forward Error Correction for egress IPsec traffic. Valid values: `enable`, `disable`.
* `fec_send_timeout` - Timeout in milliseconds before sending Forward Error Correction packets (1 - 1000).
* `fec_base` - Number of base Forward Error Correction packets (1 - 100).
* `fec_codec` - ipsec fec encoding/decoding algorithm (0: reed-solomon, 1: xor).
* `fec_redundant` - Number of redundant Forward Error Correction packets (1 - 100).
* `fec_ingress` - Enable/disable Forward Error Correction for ingress IPsec traffic. Valid values: `enable`, `disable`.
* `fec_receive_timeout` - Timeout in milliseconds before dropping Forward Error Correction packets (1 - 10000).
* `fec_health_check` - SD-WAN health check.
* `fec_mapping_profile` - Forward Error Correction (FEC) mapping profile.
* `network_overlay` - Enable/disable network overlays. Valid values: `disable`, `enable`.
* `network_id` - VPN gateway network ID.
* `loopback_asymroute` - Enable/disable asymmetric routing for IKE traffic on loopback interface. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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
$ terraform import fortios_vpnipsec_phase1interface.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnipsec_phase1interface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
