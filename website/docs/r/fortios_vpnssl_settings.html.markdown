---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnssl_settings"
description: |-
  Configure SSL VPN.
---

# fortios_vpnssl_settings
Configure SSL VPN.

## Example Usage

```hcl
resource "fortios_vpnssl_settings" "trname" {
  login_attempt_limit = 2
  login_block_time    = 60
  login_timeout       = 30
  port                = 443
  servercert          = "self-sign"
}
```

## Argument Reference

The following arguments are supported:

* `reqclientcert` - Enable to require client certificates for all SSL-VPN users.
* `user_peer` - Name of user peer.
* `ssl_max_proto_ver` - SSL maximum protocol version.
* `ssl_min_proto_ver` - SSL minimum protocol version.
* `tlsv1_0` - Enable/disable TLSv1.0.
* `tlsv1_1` - Enable/disable TLSv1.1.
* `tlsv1_2` - Enable/disable TLSv1.2.
* `tlsv1_3` - Enable/disable TLSv1.3.
* `banned_cipher` - Select one or more cipher technologies that cannot be used in SSL-VPN negotiations.
* `ssl_insert_empty_fragment` - Enable/disable insertion of empty fragment.
* `https_redirect` - Enable/disable redirect of port 80 to SSL-VPN port.
* `x_content_type_options` - Add HTTP X-Content-Type-Options header.
* `ssl_client_renegotiation` - Enable to allow client renegotiation by the server if the tunnel goes down.
* `force_two_factor_auth` - Enable to force two-factor authentication for all SSL-VPNs.
* `unsafe_legacy_renegotiation` - Enable/disable unsafe legacy re-negotiation.
* `servercert` - Name of the server certificate to be used for SSL-VPNs.
* `algorithm` - Force the SSL-VPN security level. High allows only high. Medium allows medium and high. Low allows any.
* `idle_timeout` - SSL VPN disconnects if idle for specified time in seconds.
* `auth_timeout` - SSL-VPN authentication timeout (1 - 259200 sec (3 days), 0 for no timeout).
* `login_attempt_limit` - SSL VPN maximum login attempt times before block (0 - 10, default = 2, 0 = no limit).
* `login_block_time` - Time for which a user is blocked from logging in after too many failed login attempts (0 - 86400 sec, default = 60).
* `login_timeout` - SSLVPN maximum login timeout (10 - 180 sec, default = 30).
* `dtls_hello_timeout` - SSLVPN maximum DTLS hello timeout (10 - 60 sec, default = 10).
* `tunnel_ip_pools` - Names of the IPv4 IP Pool firewall objects that define the IP addresses reserved for remote clients. The structure of `tunnel_ip_pools` block is documented below.
* `tunnel_ipv6_pools` - Names of the IPv6 IP Pool firewall objects that define the IP addresses reserved for remote clients. The structure of `tunnel_ipv6_pools` block is documented below.
* `dns_suffix` - DNS suffix used for SSL-VPN clients.
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `wins_server1` - WINS server 1.
* `wins_server2` - WINS server 2.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_wins_server1` - IPv6 WINS server 1.
* `ipv6_wins_server2` - IPv6 WINS server 2.
* `route_source_interface` - Enable to allow SSL-VPN sessions to bypass routing and bind to the incoming interface.
* `url_obscuration` - Enable to obscure the host name of the URL of the web browser display.
* `http_compression` - Enable to allow HTTP compression over SSL-VPN tunnels.
* `http_only_cookie` - Enable/disable SSL-VPN support for HttpOnly cookies.
* `deflate_compression_level` - Compression level (0~9).
* `deflate_min_data_size` - Minimum amount of data that triggers compression (200 - 65535 bytes).
* `port` - SSL-VPN access port (1 - 65535).
* `port_precedence` - Enable means that if SSL-VPN connections are allowed on an interface admin GUI connections are blocked on that interface.
* `auto_tunnel_static_route` - Enable to auto-create static routes for the SSL-VPN tunnel IP addresses.
* `header_x_forwarded_for` - Forward the same, add, or remove HTTP header.
* `source_interface` - SSL VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.
* `source_address` - Source address of incoming traffic. The structure of `source_address` block is documented below.
* `source_address_negate` - Enable/disable negated source address match.
* `source_address6` - IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.
* `source_address6_negate` - Enable/disable negated source IPv6 address match.
* `default_portal` - Default SSL VPN portal.
* `authentication_rule` - Authentication rule for SSL VPN. The structure of `authentication_rule` block is documented below.
* `dtls_tunnel` - Enable DTLS to prevent eavesdropping, tampering, or message forgery.
* `dtls_max_proto_ver` - DTLS maximum protocol version.
* `dtls_min_proto_ver` - DTLS minimum protocol version.
* `check_referer` - Enable/disable verification of referer field in HTTP request header.
* `http_request_header_timeout` - SSL-VPN session is disconnected if an HTTP request header is not received within this time (1 - 60 sec, default = 20).
* `http_request_body_timeout` - SSL-VPN session is disconnected if an HTTP request body is not received within this time (1 - 60 sec, default = 20).
* `auth_session_check_source_ip` - Enable/disable checking of source IP for authentication session.
* `tunnel_connect_without_reauth` - Enable/disable tunnel connection without re-authorization if previous connection dropped.
* `tunnel_user_session_timeout` - Time out value to clean up user session after tunnel connection is dropped (1 - 255 sec, default=30).
* `hsts_include_subdomains` - Add HSTS includeSubDomains response header.
* `transform_backward_slashes` - Transform backward slashes to forward slashes in URLs.
* `encode_2f_sequence` - Encode \2F sequence to forward slash in URLs.
* `encrypt_and_store_password` - Encrypt and store user passwords for SSL-VPN web sessions.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `tunnel_ip_pools` block supports:

* `name` - Address name.

The `tunnel_ipv6_pools` block supports:

* `name` - Address name.

The `source_interface` block supports:

* `name` - Interface name.

The `source_address` block supports:

* `name` - Address name.

The `source_address6` block supports:

* `name` - IPv6 address name.

The `authentication_rule` block supports:

* `id` - ID (0 - 4294967295).
* `source_interface` - SSL VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.
* `source_address` - Source address of incoming traffic. The structure of `source_address` block is documented below.
* `source_address_negate` - Enable/disable negated source address match.
* `source_address6` - IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.
* `source_address6_negate` - Enable/disable negated source IPv6 address match.
* `users` - User name. The structure of `users` block is documented below.
* `groups` - User groups. The structure of `groups` block is documented below.
* `portal` - SSL VPN portal.
* `realm` - SSL VPN realm.
* `client_cert` - Enable/disable SSL VPN client certificate restrictive.
* `user_peer` - Name of user peer.
* `cipher` - SSL VPN cipher strength.
* `auth` - SSL VPN authentication method restriction.

The `source_interface` block supports:

* `name` - Interface name.

The `source_address` block supports:

* `name` - Address name.

The `source_address6` block supports:

* `name` - IPv6 address name.

The `users` block supports:

* `name` - User name.

The `groups` block supports:

* `name` - Group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

VpnSsl Settings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnssl_settings.labelname VpnSslSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
