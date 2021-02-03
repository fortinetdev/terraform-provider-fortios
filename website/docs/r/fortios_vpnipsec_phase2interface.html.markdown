---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_phase2interface"
description: |-
  Configure VPN autokey tunnel.
---

# fortios_vpnipsec_phase2interface
Configure VPN autokey tunnel.

## Example Usage

```hcl
resource "fortios_vpnipsec_phase1interface" "trname3" {
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
  name                      = "absct1"
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
  remote_gw                 = "2.22.2.2"
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

resource "fortios_vpnipsec_phase2interface" "trname2" {
  add_route                = "phase1"
  auto_discovery_forwarder = "phase1"
  auto_discovery_sender    = "phase1"
  auto_negotiate           = "disable"
  dhcp_ipsec               = "disable"
  dhgrp                    = "14 5"
  dst_addr_type            = "subnet"
  dst_end_ip6              = "::"
  dst_port                 = 0
  dst_subnet               = "0.0.0.0 0.0.0.0"
  encapsulation            = "tunnel-mode"
  keepalive                = "disable"
  keylife_type             = "seconds"
  keylifekbs               = 5120
  keylifeseconds           = 43200
  l2tp                     = "disable"
  name                     = "abscp2"
  pfs                      = "enable"
  phase1name               = fortios_vpnipsec_phase1interface.trname3.name
  proposal                 = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
  protocol                 = 0
  replay                   = "enable"
  route_overlap            = "use-new"
  single_source            = "disable"
  src_addr_type            = "subnet"
  src_end_ip6              = "::"
  src_port                 = 0
  src_subnet               = "0.0.0.0 0.0.0.0"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IPsec tunnel name.
* `phase1name` - (Required) Phase 1 determines the options required for phase 2.
* `dhcp_ipsec` - Enable/disable DHCP-IPsec.
* `proposal` - (Required) Phase2 proposal.
* `pfs` - Enable/disable PFS feature.
* `ipv4_df` - Enable/disable setting and resetting of IPv4 'Don't Fragment' bit.
* `dhgrp` - Phase2 DH group.
* `replay` - Enable/disable replay detection.
* `keepalive` - Enable/disable keep alive.
* `auto_negotiate` - Enable/disable IPsec SA auto-negotiation.
* `add_route` - Enable/disable automatic route addition.
* `auto_discovery_sender` - Enable/disable sending short-cut messages.
* `auto_discovery_forwarder` - Enable/disable forwarding short-cut messages.
* `keylifeseconds` - Phase2 key life in time in seconds (120 - 172800).
* `keylifekbs` - Phase2 key life in number of bytes of traffic (5120 - 4294967295).
* `keylife_type` - Keylife type.
* `single_source` - Enable/disable single source IP restriction.
* `route_overlap` - Action for overlapping routes.
* `encapsulation` - ESP encapsulation mode.
* `l2tp` - Enable/disable L2TP over IPsec.
* `comments` - Comment.
* `initiator_ts_narrow` - Enable/disable traffic selector narrowing for IKEv2 initiator.
* `diffserv` - Enable/disable applying DSCP value to the IPsec tunnel outer IP header.
* `diffservcode` - DSCP value to be applied to the IPsec tunnel outer IP header.
* `protocol` - Quick mode protocol selector (1 - 255 or 0 for all).
* `src_name` - Local proxy ID name.
* `src_name6` - Local proxy ID name.
* `src_addr_type` - Local proxy ID type.
* `src_start_ip` - Local proxy ID start.
* `src_start_ip6` - Local proxy ID IPv6 start.
* `src_end_ip` - Local proxy ID end.
* `src_end_ip6` - Local proxy ID IPv6 end.
* `src_subnet` - Local proxy ID subnet.
* `src_subnet6` - Local proxy ID IPv6 subnet.
* `src_port` - Quick mode source port (1 - 65535 or 0 for all).
* `dst_name` - Remote proxy ID name.
* `dst_name6` - Remote proxy ID name.
* `dst_addr_type` - Remote proxy ID type.
* `dst_start_ip` - Remote proxy ID IPv4 start.
* `dst_start_ip6` - Remote proxy ID IPv6 start.
* `dst_end_ip` - Remote proxy ID IPv4 end.
* `dst_end_ip6` - Remote proxy ID IPv6 end.
* `dst_subnet` - Remote proxy ID IPv4 subnet.
* `dst_subnet6` - Remote proxy ID IPv6 subnet.
* `dst_port` - Quick mode destination port (1 - 65535 or 0 for all).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnIpsec Phase2Interface can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnipsec_phase2interface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
