---
subcategory: "FortiGate Ztna"
layout: "fortios"
page_title: "FortiOS: fortios_ztna_trafficforwardproxy"
description: |-
  Configure ZTNA traffic forward proxy.
---

# fortios_ztna_trafficforwardproxy
Configure ZTNA traffic forward proxy. Applies to FortiOS Version `>= 7.6.0`.

## Argument Reference

The following arguments are supported:

* `name` - Traffic forward proxy name
* `vip` - Virtual IP name.
* `host` - Virtual or real host name.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `status` - Enable/disable the traffic forward proxy for ZTNA traffic. Valid values: `enable`, `disable`.
* `interface` - interface name
* `port` - Accept incoming traffic on one or more ports (0 - 65535).
* `client_cert` - Enable/disable to request client certificate. Valid values: `disable`, `enable`.
* `user_agent_detect` - Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
* `auth_portal` - Enable/disable authentication portal. Valid values: `disable`, `enable`.
* `auth_virtual_host` - Virtual host for authentication portal.
* `vip6` - Virtual IPv6 name.
* `empty_cert_action` - Action of an empty client certificate. Valid values: `accept`, `block`, `accept-unmanageable`.
* `log_blocked_traffic` - Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
* `svr_pool_multiplex` - Enable/disable server pool multiplexing. Share connected server in HTTP, HTTPS, and web-portal api-gateway. Valid values: `enable`, `disable`.
* `svr_pool_ttl` - Time-to-live in the server pool for idle connections to servers.
* `svr_pool_server_max_request` - Maximum number of requests that servers in server pool handle before disconnecting (default = unlimited).
* `svr_pool_server_max_concurrent_request` - Maximum number of concurrent requests that servers in server pool could handle (default = unlimited).
* `comment` - Comment.
* `h3_support` - Enable/disable HTTP3/QUIC support (default = disable). Valid values: `enable`, `disable`.
* `quic` - QUIC setting. The structure of `quic` block is documented below.
* `ssl_mode` - Apply SSL offloading between the client and the FortiGate (half) or from the client to the FortiGate and from the FortiGate to the server (full). Valid values: `half`, `full`.
* `ssl_certificate` - Name of the certificate to use for SSL handshake. The structure of `ssl_certificate` block is documented below.
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.
* `ssl_algorithm` - Permitted encryption algorithms for SSL sessions according to encryption strength. Valid values: `high`, `medium`, `low`, `custom`.
* `ssl_cipher_suites` - SSL/TLS cipher suites acceptable from a client, ordered by priority. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_server_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high`, `medium`, `low`, `custom`, `client`.
* `ssl_server_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority. The structure of `ssl_server_cipher_suites` block is documented below.
* `ssl_pfs` - Select the cipher suites that can be used for SSL perfect forward secrecy (PFS). Applies to both client and server sessions. Valid values: `require`, `deny`, `allow`.
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `ssl_max_version` - Highest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `ssl_server_min_version` - Lowest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`, `client`.
* `ssl_server_max_version` - Highest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`, `client`.
* `ssl_accept_ffdhe_groups` - Enable/disable FFDHE cipher suite for SSL key exchange. Valid values: `enable`, `disable`.
* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid CBC IV attacks (SSL 3.0 & TLS 1.0 only). May need to be disabled for compatibility with older systems. Valid values: `enable`, `disable`.
* `ssl_client_fallback` - Enable/disable support for preventing Downgrade Attacks on client connections (RFC 7507). Valid values: `disable`, `enable`.
* `ssl_client_renegotiation` - Allow, deny, or require secure renegotiation of client sessions to comply with RFC 5746. Valid values: `allow`, `deny`, `secure`.
* `ssl_client_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the client and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.
* `ssl_client_session_state_timeout` - Number of minutes to keep client to FortiProxy SSL session state.
* `ssl_client_session_state_max` - Maximum number of client to FortiProxy SSL session states to keep.
* `ssl_client_rekey_count` - Maximum length of data in MB before triggering a client rekey (0 = disable).
* `ssl_server_renegotiation` - Enable/disable secure renegotiation to comply with RFC 5746. Valid values: `enable`, `disable`.
* `ssl_server_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the server and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.
* `ssl_server_session_state_timeout` - Number of minutes to keep FortiGate to Server SSL session state.
* `ssl_server_session_state_max` - Maximum number of FortiGate to Server SSL session states to keep.
* `ssl_http_location_conversion` - Enable to replace HTTP with HTTPS in the reply's Location HTTP header field. Valid values: `enable`, `disable`.
* `ssl_http_match_host` - Enable/disable HTTP host matching for location conversion. Valid values: `enable`, `disable`.
* `ssl_hpkp` - Enable/disable including HPKP header in response. Valid values: `disable`, `enable`, `report-only`.
* `ssl_hpkp_primary` - Certificate to generate primary HPKP pin from.
* `ssl_hpkp_backup` - Certificate to generate backup HPKP pin from.
* `ssl_hpkp_age` - Number of seconds the client should honor the HPKP setting.
* `ssl_hpkp_report_uri` - URL to report HPKP violations to.
* `ssl_hpkp_include_subdomains` - Indicate that HPKP header applies to all subdomains. Valid values: `disable`, `enable`.
* `ssl_hsts` - Enable/disable including HSTS header in response. Valid values: `disable`, `enable`.
* `ssl_hsts_age` - Number of seconds the client should honor the HSTS setting.
* `ssl_hsts_include_subdomains` - Indicate that HSTS header applies to all subdomains. Valid values: `disable`, `enable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `quic` block supports:

* `max_idle_timeout` - Maximum idle timeout milliseconds (1 - 60000, default = 30000).
* `max_udp_payload_size` - Maximum UDP payload size in bytes (1200 - 1500, default = 1500).
* `active_connection_id_limit` - Active connection ID limit (1 - 8, default = 2).
* `ack_delay_exponent` - ACK delay exponent (1 - 20, default = 3).
* `max_ack_delay` - Maximum ACK delay in milliseconds (1 - 16383, default = 25).
* `max_datagram_frame_size` - Maximum datagram frame size in bytes (1 - 1500, default = 1500).
* `active_migration` - Enable/disable active migration (default = disable). Valid values: `enable`, `disable`.
* `grease_quic_bit` - Enable/disable grease QUIC bit (default = enable). Valid values: `enable`, `disable`.

The `ssl_certificate` block supports:

* `name` - Certificate list.

The `ssl_cipher_suites` block supports:

* `priority` - SSL/TLS cipher suites priority.
* `cipher` - Cipher suite name. Valid values: `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

The `ssl_server_cipher_suites` block supports:

* `priority` - SSL/TLS cipher suites priority.
* `cipher` - Cipher suite name. Valid values: `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna TrafficForwardProxy can be imported using any of these accepted formats:
```
$ terraform import fortios_ztna_trafficforwardproxy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ztna_trafficforwardproxy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
