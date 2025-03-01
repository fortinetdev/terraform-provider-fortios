---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_forticlient"
description: |-
  Configure FortiClient policy realm.
---

# fortios_vpnipsec_forticlient
Configure FortiClient policy realm. Applies to FortiOS Version `<= 7.6.0`.

## Example Usage

```hcl
# fortios_vpnipsec_phase1interface.trname2:
resource "fortios_vpnipsec_phase1interface" "trname4" {
  acct_verify               = "disable"
  add_gw_route              = "disable"
  add_route                 = "enable"
  assign_ip                 = "enable"
  assign_ip_from            = "range"
  authmethod                = "psk"
  authusrgrp                = "Guest-group"
  auto_discovery_forwarder  = "disable"
  auto_discovery_psk        = "disable"
  auto_discovery_receiver   = "disable"
  auto_discovery_sender     = "disable"
  auto_negotiate            = "enable"
  cert_id_validation        = "enable"
  childless_ike             = "disable"
  client_auto_negotiate     = "disable"
  client_keep_alive         = "disable"
  comments                  = "VPN: Dialup_IPsec (Created by VPN wizard)"
  default_gw                = "0.0.0.0"
  default_gw_priority       = 0
  dhgrp                     = "14 5"
  digital_signature_auth    = "disable"
  distance                  = 15
  dns_mode                  = "auto"
  dpd                       = "on-idle"
  dpd_retrycount            = 3
  dpd_retryinterval         = "60"
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
  interface                 = "port4"
  ip_version                = "4"
  ipv4_dns_server1          = "0.0.0.0"
  ipv4_dns_server2          = "0.0.0.0"
  ipv4_dns_server3          = "0.0.0.0"
  ipv4_end_ip               = "10.10.10.10"
  ipv4_netmask              = "255.255.255.192"
  ipv4_split_include        = "FIREWALL_AUTH_PORTAL_ADDRESS"
  ipv4_start_ip             = "10.10.10.1"
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
  mode                      = "aggressive"
  mode_cfg                  = "enable"
  monitor_hold_down_delay   = 0
  monitor_hold_down_time    = "00:00"
  monitor_hold_down_type    = "immediate"
  monitor_hold_down_weekday = "sunday"
  name                      = "dia1"
  nattraversal              = "enable"
  negotiate_timeout         = 30
  net_device                = "enable"
  passive_mode              = "disable"
  peertype                  = "any"
  psksecret                 = "NCIEW32930293203932"
  ppk                       = "disable"
  priority                  = 0
  proposal                  = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
  reauth                    = "disable"
  rekey                     = "enable"
  remote_gw                 = "0.0.0.0"
  remote_gw6                = "::"
  rsa_signature_format      = "pkcs1"
  save_password             = "enable"
  send_cert_chain           = "enable"
  signature_hash_alg        = "sha2-512 sha2-384 sha2-256 sha1"
  suite_b                   = "disable"
  tunnel_search             = "selectors"
  type                      = "dynamic"
  unity_support             = "enable"
  wizard_type               = "dialup-forticlient"
  xauthtype                 = "auto"
}

# fortios_vpnipsec_phase2interface.trname1:
resource "fortios_vpnipsec_phase2interface" "trname3" {
  add_route                = "phase1"
  auto_discovery_forwarder = "phase1"
  auto_discovery_sender    = "phase1"
  auto_negotiate           = "disable"
  dhcp_ipsec               = "disable"
  dhgrp                    = "14 5"
  dst_addr_type            = "subnet"
  dst_end_ip               = "0.0.0.0"
  dst_end_ip6              = "::"
  dst_port                 = 0
  dst_start_ip             = "0.0.0.0"
  dst_start_ip6            = "::"
  dst_subnet               = "0.0.0.0 0.0.0.0"
  dst_subnet6              = "::/0"
  encapsulation            = "tunnel-mode"
  keepalive                = "disable"
  keylife_type             = "seconds"
  keylifekbs               = 5120
  keylifeseconds           = 43200
  l2tp                     = "disable"
  name                     = "dia2"
  pfs                      = "enable"
  phase1name               = fortios_vpnipsec_phase1interface.trname4.name
  proposal                 = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
  protocol                 = 0
  replay                   = "enable"
  route_overlap            = "use-new"
  single_source            = "disable"
  src_addr_type            = "subnet"
  src_end_ip               = "0.0.0.0"
  src_end_ip6              = "::"
  src_port                 = 0
  src_start_ip             = "0.0.0.0"
  src_start_ip6            = "::"
  src_subnet               = "0.0.0.0 0.0.0.0"
  src_subnet6              = "::/0"
}

resource "fortios_vpnipsec_forticlient" "trname" {
  phase2name    = fortios_vpnipsec_phase2interface.trname3.name
  realm         = "1"
  status        = "enable"
  usergroupname = "Guest-group"
}
```

## Argument Reference

The following arguments are supported:

* `realm` - FortiClient realm name.
* `usergroupname` - (Required) User group name for FortiClient users.
* `phase2name` - (Required) Phase 2 tunnel name that you defined in the FortiClient dialup configuration.
* `status` - Enable/disable this FortiClient configuration. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{realm}}.

## Import

VpnIpsec Forticlient can be imported using any of these accepted formats:
```
$ terraform import fortios_vpnipsec_forticlient.labelname {{realm}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnipsec_forticlient.labelname {{realm}}
$ unset "FORTIOS_IMPORT_TABLE"
```
