---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnssl_settings"
description: |-
  Get information on fortios vpnssl settings.
---

# Data Source: fortios_vpnssl_settings
Use this data source to get information on fortios vpnssl settings

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable SSL-VPN.
* `reqclientcert` - Enable to require client certificates for all SSL-VPN users.
* `user_peer` - Name of user peer.
* `ssl_max_proto_ver` - SSL maximum protocol version.
* `ssl_min_proto_ver` - SSL minimum protocol version.
* `tlsv1_0` - Enable/disable TLSv1.0.
* `tlsv1_1` - Enable/disable TLSv1.1.
* `tlsv1_2` - Enable/disable TLSv1.2.
* `tlsv1_3` - Enable/disable TLSv1.3.
* `banned_cipher` - Select one or more cipher technologies that cannot be used in SSL-VPN negotiations.
* `ciphersuite` - Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, set ssl-max-proto-ver to tls1-2 or below.
* `ssl_insert_empty_fragment` - Enable/disable insertion of empty fragment.
* `https_redirect` - Enable/disable redirect of port 80 to SSL-VPN port.
* `x_content_type_options` - Add HTTP X-Content-Type-Options header.
* `ssl_client_renegotiation` - Enable to allow client renegotiation by the server if the tunnel goes down.
* `force_two_factor_auth` - Enable to force two-factor authentication for all SSL-VPNs.
* `unsafe_legacy_renegotiation` - Enable/disable unsafe legacy re-negotiation.
* `servercert` - Name of the server certificate to be used for SSL-VPNs.
* `algorithm` - Force the SSL-VPN security level. High allows only high. Medium allows medium and high. Low allows any.
* `tls_groups` - Configure the supported groups for TLS negotiation.
* `idle_timeout` - SSL VPN disconnects if idle for specified time in seconds.
* `auth_timeout` - SSL-VPN authentication timeout (1 - 259200 sec (3 days), 0 for no timeout).
* `login_attempt_limit` - SSL VPN maximum login attempt times before block (0 - 10, default = 2, 0 = no limit).
* `login_block_time` - Time for which a user is blocked from logging in after too many failed login attempts (0 - 86400 sec, default = 60).
* `login_timeout` - SSLVPN maximum login timeout (10 - 180 sec, default = 30).
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
* `browser_language_detection` - Enable/disable overriding the configured system language based on the preferred language of the browser.
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
* `client_sigalgs` - Set signature algorithms related to client authentication. Affects TLS version <= 1.2 only.
* `dual_stack_mode` - Tunnel mode: enable parallel IPv4 and IPv6 tunnel. Web mode: support IPv4 and IPv6 bookmarks in the portal.
* `tunnel_addr_assigned_method` - Method used for assigning address for tunnel.
* `saml_redirect_port` - SAML local redirect port in the machine running FCT (0 - 65535). 0 is to disable redirection on FGT side.
* `web_mode_snat` - Enable/disable use of IP pools defined in firewall policy while using web-mode.
* `ztna_trusted_client` - Enable/disable verification of device certificate for SSLVPN ZTNA session.
* `server_hostname` - Server hostname for HTTPS. When set, will be used for SSL VPN web proxy host header for any redirection.
* `remote_https_cert_check` - Configure how the FortiGate unit checks and responds to the remote HTTPS server's certificate (default = warn-on-error).

The `tunnel_ip_pools` block contains:

* `name` - Address name.

The `tunnel_ipv6_pools` block contains:

* `name` - Address name.

The `source_interface` block contains:

* `name` - Interface name.

The `source_address` block contains:

* `name` - Address name.

The `source_address6` block contains:

* `name` - IPv6 address name.

The `authentication_rule` block contains:

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

The `source_interface` block contains:

* `name` - Interface name.

The `source_address` block contains:

* `name` - Address name.

The `source_address6` block contains:

* `name` - IPv6 address name.

The `users` block contains:

* `name` - User name.

The `groups` block contains:

* `name` - Group name.

