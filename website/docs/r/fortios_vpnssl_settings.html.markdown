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

* `status` - Enable/disable SSL-VPN. Valid values: `enable`, `disable`.
* `reqclientcert` - Enable to require client certificates for all SSL-VPN users. Valid values: `enable`, `disable`.
* `user_peer` - Name of user peer.
* `ssl_max_proto_ver` - SSL maximum protocol version. Valid values: `tls1-0`, `tls1-1`, `tls1-2`, `tls1-3`.
* `ssl_min_proto_ver` - SSL minimum protocol version. Valid values: `tls1-0`, `tls1-1`, `tls1-2`, `tls1-3`.
* `tlsv1_0` - Enable/disable TLSv1.0. Valid values: `enable`, `disable`.
* `tlsv1_1` - Enable/disable TLSv1.1. Valid values: `enable`, `disable`.
* `tlsv1_2` - Enable/disable TLSv1.2. Valid values: `enable`, `disable`.
* `tlsv1_3` - Enable/disable TLSv1.3. Valid values: `enable`, `disable`.
* `banned_cipher` - Select one or more cipher technologies that cannot be used in SSL-VPN negotiations.
* `ciphersuite` - Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, set ssl-max-proto-ver to tls1-2 or below. Valid values: `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-AES-128-CCM-SHA256`, `TLS-AES-128-CCM-8-SHA256`.
* `ssl_insert_empty_fragment` - Enable/disable insertion of empty fragment. Valid values: `enable`, `disable`.
* `https_redirect` - Enable/disable redirect of port 80 to SSL-VPN port. Valid values: `enable`, `disable`.
* `x_content_type_options` - Add HTTP X-Content-Type-Options header. Valid values: `enable`, `disable`.
* `ssl_client_renegotiation` - Enable to allow client renegotiation by the server if the tunnel goes down. Valid values: `disable`, `enable`.
* `force_two_factor_auth` - Enable to force two-factor authentication for all SSL-VPNs. Valid values: `enable`, `disable`.
* `unsafe_legacy_renegotiation` - Enable/disable unsafe legacy re-negotiation. Valid values: `enable`, `disable`.
* `servercert` - Name of the server certificate to be used for SSL-VPNs.
* `algorithm` - Force the SSL-VPN security level. High allows only high. Medium allows medium and high. Low allows any. Valid values: `high`, `medium`, `default`, `low`.
* `tls_groups` - Configure the supported groups for TLS negotiation. Valid values: `P-521`, `P-384`, `P-256`, `ML-KEM512`, `ML-KEM768`, `ML-KEM1024`, `P-384-MLKEM1024`, `P-256-MLKEM768`, `X25519-MLKEM768`, `X448`, `X25519`, `FFDHE2048`, `FFDHE3072`, `FFDHE4096`, `FFDHE6144`, `FFDHE8192`.
* `idle_timeout` - SSL VPN disconnects if idle for specified time in seconds.
* `auth_timeout` - Agentless VPN authentication timeout (1 - 259200 sec (3 days), 0 for no timeout).
* `login_attempt_limit` - Agentless VPN maximum login attempt times before block (0 - 10, default = 2, 0 = no limit).
* `login_block_time` - Time for which a user is blocked from logging in after too many failed login attempts (0 - 86400 sec, default = 60).
* `login_timeout` - Agentless VPN maximum login timeout (10 - 180 sec, default = 30).
* `dtls_hello_timeout` - SSLVPN maximum DTLS hello timeout (10 - 60 sec, default = 10).
* `dtls_heartbeat_idle_timeout` - Idle timeout before DTLS heartbeat is sent.
* `dtls_heartbeat_interval` - Interval between DTLS heartbeat.
* `dtls_heartbeat_fail_count` - Number of missing heartbeats before the connection is considered dropped.
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
* `route_source_interface` - Enable to allow SSL-VPN sessions to bypass routing and bind to the incoming interface. Valid values: `enable`, `disable`.
* `url_obscuration` - Enable to obscure the host name of the URL of the web browser display. Valid values: `enable`, `disable`.
* `http_compression` - Enable to allow HTTP compression over SSL-VPN tunnels. Valid values: `enable`, `disable`.
* `http_only_cookie` - Enable/disable SSL-VPN support for HttpOnly cookies. Valid values: `enable`, `disable`.
* `deflate_compression_level` - Compression level (0~9).
* `deflate_min_data_size` - Minimum amount of data that triggers compression (200 - 65535 bytes).
* `port` - Agentless VPN access port (1 - 65535).
* `port_precedence` - Enable means that if SSL-VPN connections are allowed on an interface admin GUI connections are blocked on that interface. Valid values: `enable`, `disable`.
* `auto_tunnel_static_route` - Enable to auto-create static routes for the SSL-VPN tunnel IP addresses. Valid values: `enable`, `disable`.
* `header_x_forwarded_for` - Forward the same, add, or remove HTTP header. Valid values: `pass`, `add`, `remove`.
* `source_interface` - SSL VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.
* `source_address` - Source address of incoming traffic. The structure of `source_address` block is documented below.
* `source_address_negate` - Enable/disable negated source address match. Valid values: `enable`, `disable`.
* `source_address6` - IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.
* `source_address6_negate` - Enable/disable negated source IPv6 address match. Valid values: `enable`, `disable`.
* `default_portal` - Default SSL VPN portal.
* `authentication_rule` - Authentication rule for SSL VPN. The structure of `authentication_rule` block is documented below.
* `browser_language_detection` - Enable/disable overriding the configured system language based on the preferred language of the browser. Valid values: `enable`, `disable`.
* `dtls_tunnel` - Enable DTLS to prevent eavesdropping, tampering, or message forgery. Valid values: `enable`, `disable`.
* `dtls_max_proto_ver` - DTLS maximum protocol version. Valid values: `dtls1-0`, `dtls1-2`.
* `dtls_min_proto_ver` - DTLS minimum protocol version. Valid values: `dtls1-0`, `dtls1-2`.
* `check_referer` - Enable/disable verification of referer field in HTTP request header. Valid values: `enable`, `disable`.
* `http_request_header_timeout` - Agentless VPN session is disconnected if an HTTP request header is not received within this time (1 - 60 sec, default = 20).
* `http_request_body_timeout` - Agentless VPN session is disconnected if an HTTP request body is not received within this time (1 - 60 sec, default = 20).
* `auth_session_check_source_ip` - Enable/disable checking of source IP for authentication session. Valid values: `enable`, `disable`.
* `tunnel_connect_without_reauth` - Enable/disable tunnel connection without re-authorization if previous connection dropped. Valid values: `enable`, `disable`.
* `tunnel_user_session_timeout` - Number of seconds after which user sessions are cleaned up after tunnel connection is dropped (default = 30). On FortiOS versions 6.2.0-7.4.3: 1 - 255 sec. On FortiOS versions 7.4.4-7.6.2: 1 - 86400 sec.
* `hsts_include_subdomains` - Add HSTS includeSubDomains response header. Valid values: `enable`, `disable`.
* `transform_backward_slashes` - Transform backward slashes to forward slashes in URLs. Valid values: `enable`, `disable`.
* `encode_2f_sequence` - Encode \2F sequence to forward slash in URLs. Valid values: `enable`, `disable`.
* `encrypt_and_store_password` - Encrypt and store user passwords for SSL-VPN web sessions. Valid values: `enable`, `disable`.
* `client_sigalgs` - Set signature algorithms related to client authentication. Affects TLS version <= 1.2 only. Valid values: `no-rsa-pss`, `all`.
* `dual_stack_mode` - Tunnel mode: enable parallel IPv4 and IPv6 tunnel. Web mode: support IPv4 and IPv6 bookmarks in the portal. Valid values: `enable`, `disable`.
* `tunnel_addr_assigned_method` - Method used for assigning address for tunnel. Valid values: `first-available`, `round-robin`.
* `saml_redirect_port` - SAML local redirect port in the machine running FortiClient (0 - 65535). 0 is to disable redirection on FGT side.
* `web_mode_snat` - Enable/disable use of IP pools defined in firewall policy while using web-mode. Valid values: `enable`, `disable`.
* `ztna_trusted_client` - Enable/disable verification of device certificate for SSLVPN ZTNA session. Valid values: `enable`, `disable`.
* `server_hostname` - Server hostname for HTTPS. When set, will be used for SSL VPN web proxy host header for any redirection.
* `remote_https_cert_check` - Configure how the FortiGate unit checks and responds to the remote HTTPS server's certificate (default = warn-on-error). Valid values: `no-check`, `warn-on-error`, `reject-on-error`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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
* `source_address_negate` - Enable/disable negated source address match. Valid values: `enable`, `disable`.
* `source_address6` - IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.
* `source_address6_negate` - Enable/disable negated source IPv6 address match. Valid values: `enable`, `disable`.
* `users` - User name. The structure of `users` block is documented below.
* `groups` - User groups. The structure of `groups` block is documented below.
* `portal` - SSL VPN portal.
* `realm` - SSL VPN realm.
* `client_cert` - Enable/disable SSL VPN client certificate restrictive. Valid values: `enable`, `disable`.
* `user_peer` - Name of user peer.
* `cipher` - SSL VPN cipher strength. Valid values: `any`, `high`, `medium`.
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
$ terraform import fortios_vpnssl_settings.labelname VpnSslSettings

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnssl_settings.labelname VpnSslSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
