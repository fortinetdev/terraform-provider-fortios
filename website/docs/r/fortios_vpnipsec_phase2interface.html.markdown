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
* `dhcp_ipsec` - Enable/disable DHCP-IPsec. Valid values: `enable`, `disable`.
* `proposal` - (Required) Phase2 proposal. Valid values: `null-md5`, `null-sha1`, `null-sha256`, `null-sha384`, `null-sha512`, `des-null`, `des-md5`, `des-sha1`, `des-sha256`, `des-sha384`, `des-sha512`, `3des-null`, `3des-md5`, `3des-sha1`, `3des-sha256`, `3des-sha384`, `3des-sha512`, `aes128-null`, `aes128-md5`, `aes128-sha1`, `aes128-sha256`, `aes128-sha384`, `aes128-sha512`, `aes128gcm`, `aes192-null`, `aes192-md5`, `aes192-sha1`, `aes192-sha256`, `aes192-sha384`, `aes192-sha512`, `aes256-null`, `aes256-md5`, `aes256-sha1`, `aes256-sha256`, `aes256-sha384`, `aes256-sha512`, `aes256gcm`, `chacha20poly1305`, `aria128-null`, `aria128-md5`, `aria128-sha1`, `aria128-sha256`, `aria128-sha384`, `aria128-sha512`, `aria192-null`, `aria192-md5`, `aria192-sha1`, `aria192-sha256`, `aria192-sha384`, `aria192-sha512`, `aria256-null`, `aria256-md5`, `aria256-sha1`, `aria256-sha256`, `aria256-sha384`, `aria256-sha512`, `seed-null`, `seed-md5`, `seed-sha1`, `seed-sha256`, `seed-sha384`, `seed-sha512`.
* `pfs` - Enable/disable PFS feature. Valid values: `enable`, `disable`.
* `ipv4_df` - Enable/disable setting and resetting of IPv4 'Don't Fragment' bit. Valid values: `enable`, `disable`.
* `dhgrp` - Phase2 DH group. Valid values: `1`, `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `27`, `28`, `29`, `30`, `31`, `32`.
* `replay` - Enable/disable replay detection. Valid values: `enable`, `disable`.
* `keepalive` - Enable/disable keep alive. Valid values: `enable`, `disable`.
* `auto_negotiate` - Enable/disable IPsec SA auto-negotiation. Valid values: `enable`, `disable`.
* `add_route` - Enable/disable automatic route addition. Valid values: `phase1`, `enable`, `disable`.
* `inbound_dscp_copy` - Enable/disable copying of the DSCP field in the ESP header to the inner IP header. Valid values: `phase1`, `enable`, `disable`.
* `auto_discovery_sender` - Enable/disable sending short-cut messages. Valid values: `phase1`, `enable`, `disable`.
* `auto_discovery_forwarder` - Enable/disable forwarding short-cut messages. Valid values: `phase1`, `enable`, `disable`.
* `keylifeseconds` - Phase2 key life in time in seconds (120 - 172800).
* `keylifekbs` - Phase2 key life in number of bytes of traffic (5120 - 4294967295).
* `keylife_type` - Keylife type. Valid values: `seconds`, `kbs`, `both`.
* `single_source` - Enable/disable single source IP restriction. Valid values: `enable`, `disable`.
* `route_overlap` - Action for overlapping routes. Valid values: `use-old`, `use-new`, `allow`.
* `encapsulation` - ESP encapsulation mode. Valid values: `tunnel-mode`, `transport-mode`.
* `l2tp` - Enable/disable L2TP over IPsec. Valid values: `enable`, `disable`.
* `comments` - Comment.
* `initiator_ts_narrow` - Enable/disable traffic selector narrowing for IKEv2 initiator. Valid values: `enable`, `disable`.
* `diffserv` - Enable/disable applying DSCP value to the IPsec tunnel outer IP header. Valid values: `enable`, `disable`.
* `diffservcode` - DSCP value to be applied to the IPsec tunnel outer IP header.
* `protocol` - Quick mode protocol selector (1 - 255 or 0 for all).
* `src_name` - Local proxy ID name.
* `src_name6` - Local proxy ID name.
* `src_addr_type` - Local proxy ID type. Valid values: `subnet`, `range`, `ip`, `name`, `subnet6`, `range6`, `ip6`, `name6`.
* `src_start_ip` - Local proxy ID start.
* `src_start_ip6` - Local proxy ID IPv6 start.
* `src_end_ip` - Local proxy ID end.
* `src_end_ip6` - Local proxy ID IPv6 end.
* `src_subnet` - Local proxy ID subnet.
* `src_subnet6` - Local proxy ID IPv6 subnet.
* `src_port` - Quick mode source port (1 - 65535 or 0 for all).
* `dst_name` - Remote proxy ID name.
* `dst_name6` - Remote proxy ID name.
* `dst_addr_type` - Remote proxy ID type. Valid values: `subnet`, `range`, `ip`, `name`, `subnet6`, `range6`, `ip6`, `name6`.
* `dst_start_ip` - Remote proxy ID IPv4 start.
* `dst_start_ip6` - Remote proxy ID IPv6 start.
* `dst_end_ip` - Remote proxy ID IPv4 end.
* `dst_end_ip6` - Remote proxy ID IPv6 end.
* `dst_subnet` - Remote proxy ID IPv4 subnet.
* `dst_subnet6` - Remote proxy ID IPv6 subnet.
* `dst_port` - Quick mode destination port (1 - 65535 or 0 for all).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnIpsec Phase2Interface can be imported using any of these accepted formats:
```
$ terraform import fortios_vpnipsec_phase2interface.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnipsec_phase2interface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
