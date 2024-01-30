---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vip"
description: |-
  Configure virtual IP for IPv4.
---

# fortios_firewall_vip
Configure virtual IP for IPv4.

## Example Usage

```hcl
resource "fortios_firewall_vip" "trname" {
  arp_reply                        = "enable"
  color                            = 0
  dns_mapping_ttl                  = 0
  extintf                          = "any"
  extip                            = "1.0.0.1-1.0.0.2"
  extport                          = "0-65535"
  fosid                            = 0
  http_cookie_age                  = 60
  http_cookie_domain_from_host     = "disable"
  http_cookie_generation           = 0
  http_cookie_share                = "same-ip"
  http_ip_header                   = "disable"
  http_multiplex                   = "disable"
  https_cookie_secure              = "disable"
  ldb_method                       = "static"
  mappedport                       = "0-65535"
  max_embryonic_connections        = 1000
  name                             = "vips1"
  nat_source_vip                   = "disable"
  outlook_web_access               = "disable"
  persistence                      = "none"
  portforward                      = "disable"
  portmapping_type                 = "1-to-1"
  protocol                         = "tcp"
  ssl_algorithm                    = "high"
  ssl_client_fallback              = "enable"
  ssl_client_renegotiation         = "secure"
  ssl_client_session_state_max     = 1000
  ssl_client_session_state_timeout = 30
  ssl_client_session_state_type    = "both"
  ssl_dh_bits                      = "2048"
  ssl_hpkp                         = "disable"
  ssl_hpkp_age                     = 5184000
  ssl_hpkp_include_subdomains      = "disable"
  ssl_hsts                         = "disable"
  ssl_hsts_age                     = 5184000
  ssl_hsts_include_subdomains      = "disable"
  ssl_http_location_conversion     = "disable"
  ssl_http_match_host              = "enable"
  ssl_max_version                  = "tls-1.2"
  ssl_min_version                  = "tls-1.1"
  ssl_mode                         = "half"
  ssl_pfs                          = "require"
  ssl_send_empty_frags             = "enable"
  ssl_server_algorithm             = "client"
  ssl_server_max_version           = "client"
  ssl_server_min_version           = "client"
  ssl_server_session_state_max     = 100
  ssl_server_session_state_timeout = 60
  ssl_server_session_state_type    = "both"
  type                             = "static-nat"
  weblogic_server                  = "disable"
  websphere_server                 = "disable"

  mappedip {
    range = "3.0.0.0-3.0.0.1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Virtual IP name.
* `fosid` - Custom defined ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `comment` - Comment.
* `type` - Configure a static NAT, load balance, server load balance, DNS translation, or FQDN VIP.
* `dns_mapping_ttl` - DNS mapping TTL (Set to zero to use TTL in DNS response, default = 0).
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`, `http-host`.
* `src_filter` - Source address filter. Each address must be either an IP/subnet (x.x.x.x/n) or a range (x.x.x.x-y.y.y.y). Separate addresses with spaces. The structure of `src_filter` block is documented below.
* `service` - Service name. The structure of `service` block is documented below.
* `extip` - IP address or address range on the external interface that you want to map to an address or address range on the destination network.
* `extaddr` - External FQDN address name. The structure of `extaddr` block is documented below.
* `h2_support` - Enable/disable HTTP2 support (default = enable). Valid values: `enable`, `disable`.
* `h3_support` - Enable/disable HTTP3/QUIC support (default = disable). Valid values: `enable`, `disable`.
* `quic` - QUIC setting. The structure of `quic` block is documented below.
* `nat44` - Enable/disable NAT44. Valid values: `disable`, `enable`.
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.
* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
* `mappedip` - IP address or address range on the destination network to which the external IP address is mapped. The structure of `mappedip` block is documented below.
* `mapped_addr` - Mapped FQDN address name.
* `extintf` - Interface connected to the source network that receives the packets that will be forwarded to the destination network.
* `arp_reply` - Enable to respond to ARP requests for this virtual IP address. Enabled by default. Valid values: `disable`, `enable`.
* `server_type` - Protocol to be load balanced by the virtual server (also called the server load balance virtual IP). Valid values: `http`, `https`, `imaps`, `pop3s`, `smtps`, `ssl`, `tcp`, `udp`, `ip`.
* `http_redirect` - Enable/disable redirection of HTTP to HTTPS Valid values: `enable`, `disable`.
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none`, `http-cookie`, `ssl-session-id`.
* `nat_source_vip` - Enable/disable forcing the source NAT mapped IP to the external IP for all traffic. Valid values: `disable`, `enable`.
* `portforward` - Enable/disable port forwarding. Valid values: `disable`, `enable`.
* `status` - Enable/disable VIP. Valid values: `disable`, `enable`.
* `protocol` - Protocol to use when forwarding packets. Valid values: `tcp`, `udp`, `sctp`, `icmp`.
* `extport` - Incoming port number range that you want to map to a port number range on the destination network.
* `mappedport` - Port number range on the destination network to which the external port number range is mapped.
* `gratuitous_arp_interval` - Enable to have the VIP send gratuitous ARPs. 0=disabled. Set from 5 up to 8640000 seconds to enable.
* `srcintf_filter` - Interfaces to which the VIP applies. Separate the names with spaces. The structure of `srcintf_filter` block is documented below.
* `portmapping_type` - Port mapping type. Valid values: `1-to-1`, `m-to-n`.
* `realservers` - Select the real servers that this server load balancing VIP will distribute traffic to. The structure of `realservers` block is documented below.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable`, `enable`.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 seconds. 0 = no time limit.
* `http_cookie_share` - Control sharing of cookies across virtual servers. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable`, `same-ip`.
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable`, `enable`.
* `http_multiplex` - Enable/disable HTTP multiplexing. Valid values: `enable`, `disable`.
* `http_multiplex_ttl` - Time-to-live for idle connections to servers.
* `http_multiplex_max_request` - Maximum number of requests that a multiplex server can handle before disconnecting sessions (default = unlimited).
* `http_multiplex_max_concurrent_request` - Maximum number of concurrent requests that a multiplex server can handle (default = unlimited).
* `http_supported_max_version` - Maximum supported HTTP versions. default = HTTP2 Valid values: `http1`, `http2`.
* `http_ip_header` - For HTTP multiplexing, enable to add the original client IP address in the XForwarded-For HTTP header. Valid values: `enable`, `disable`.
* `http_ip_header_name` - For HTTP multiplexing, enter a custom HTTPS header name. The original client IP address is added to this header. If empty, X-Forwarded-For is used.
* `outlook_web_access` - Enable to add the Front-End-Https header for Microsoft Outlook Web Access. Valid values: `disable`, `enable`.
* `weblogic_server` - Enable to add an HTTP header to indicate SSL offloading for a WebLogic server. Valid values: `disable`, `enable`.
* `websphere_server` - Enable to add an HTTP header to indicate SSL offloading for a WebSphere server. Valid values: `disable`, `enable`.
* `ssl_mode` - Apply SSL offloading between the client and the FortiGate (half) or from the client to the FortiGate and from the FortiGate to the server (full). Valid values: `half`, `full`.
* `ssl_certificate` - The name of the SSL certificate to use for SSL acceleration.
* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.
* `ssl_algorithm` - Permitted encryption algorithms for SSL sessions according to encryption strength. Valid values: `high`, `medium`, `low`, `custom`.
* `ssl_cipher_suites` - SSL/TLS cipher suites acceptable from a client, ordered by priority. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_server_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high`, `medium`, `low`, `custom`, `client`.
* `ssl_server_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority. The structure of `ssl_server_cipher_suites` block is documented below.
* `ssl_pfs` - Select the cipher suites that can be used for SSL perfect forward secrecy (PFS). Applies to both client and server sessions. Valid values: `require`, `deny`, `allow`.
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a client.
* `ssl_max_version` - Highest SSL/TLS version acceptable from a client.
* `ssl_server_min_version` - Lowest SSL/TLS version acceptable from a server. Use the client setting by default.
* `ssl_server_max_version` - Highest SSL/TLS version acceptable from a server. Use the client setting by default.
* `ssl_accept_ffdhe_groups` - Enable/disable FFDHE cipher suite for SSL key exchange. Valid values: `enable`, `disable`.
* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid CBC IV attacks (SSL 3.0 & TLS 1.0 only). May need to be disabled for compatibility with older systems. Valid values: `enable`, `disable`.
* `ssl_client_fallback` - Enable/disable support for preventing Downgrade Attacks on client connections (RFC 7507). Valid values: `disable`, `enable`.
* `ssl_client_renegotiation` - Allow, deny, or require secure renegotiation of client sessions to comply with RFC 5746. Valid values: `allow`, `deny`, `secure`.
* `ssl_client_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the client and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.
* `ssl_client_session_state_timeout` - Number of minutes to keep client to FortiGate SSL session state.
* `ssl_client_session_state_max` - Maximum number of client to FortiGate SSL session states to keep.
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
* `ssl_hpkp_age` - Number of seconds the client should honour the HPKP setting.
* `ssl_hpkp_report_uri` - URL to report HPKP violations to.
* `ssl_hpkp_include_subdomains` - Indicate that HPKP header applies to all subdomains. Valid values: `disable`, `enable`.
* `ssl_hsts` - Enable/disable including HSTS header in response. Valid values: `disable`, `enable`.
* `ssl_hsts_age` - Number of seconds the client should honour the HSTS setting.
* `ssl_hsts_include_subdomains` - Indicate that HSTS header applies to all subdomains. Valid values: `disable`, `enable`.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status. The structure of `monitor` block is documented below.
* `max_embryonic_connections` - Maximum number of incomplete connections.
* `color` - Color of icon on the GUI.
* `ipv6_mappedip` - Start-mapped-IPv6-address [-end mapped-IPv6-address].
* `ipv6_mappedport` - IPv6 port number range on the destination network to which the external port number range is mapped.
* `one_click_gslb_server` - Enable/disable one click GSLB server integration with FortiGSLB. Valid values: `disable`, `enable`.
* `gslb_hostname` - Hostname to use within the configured FortiGSLB domain.
* `gslb_domain_name` - Domain to use when integrating with FortiGSLB.
* `gslb_public_ips` - Publicly accessible IP addresses for the FortiGSLB service. The structure of `gslb_public_ips` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `src_filter` block supports:

* `range` - Source-filter range.

The `service` block supports:

* `name` - Service name.

The `extaddr` block supports:

* `name` - Address name.

The `quic` block supports:

* `max_idle_timeout` - Maximum idle timeout milliseconds (1 - 60000, default = 30000).
* `max_udp_payload_size` - Maximum UDP payload size in bytes (1200 - 1500, default = 1500).
* `active_connection_id_limit` - Active connection ID limit (1 - 8, default = 2).
* `ack_delay_exponent` - ACK delay exponent (1 - 20, default = 3).
* `max_ack_delay` - Maximum ACK delay in milliseconds (1 - 16383, default = 25).
* `max_datagram_frame_size` - Maximum datagram frame size in bytes (1 - 1500, default = 1500).
* `active_migration` - Enable/disable active migration (default = disable). Valid values: `enable`, `disable`.
* `grease_quic_bit` - Enable/disable grease QUIC bit (default = enable). Valid values: `enable`, `disable`.

The `mappedip` block supports:

* `range` - Mapped IP range.

The `srcintf_filter` block supports:

* `interface_name` - Interface name.

The `realservers` block supports:

* `id` - Real server ID.
* `type` - Type of address. Valid values: `ip`, `address`.
* `address` - Dynamic address of the real server.
* `ip` - IP address of the real server.
* `port` - Port for communicating with the real server. Required if port forwarding is enabled.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `holddown_interval` - Time in seconds that the health check monitor continues to monitor and unresponsive server that should be active.
* `healthcheck` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`, `vip`.
* `http_host` - HTTP server domain name in HTTP header.
* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `enable`, `disable`.
* `max_connections` - Max number of active connections that can be directed to the real server. When reached, sessions are sent to other real servers.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `client_ip` - Only clients in this IP range can connect to this real server.

The `ssl_cipher_suites` block supports:

* `priority` - SSL/TLS cipher suites priority.
* `cipher` - Cipher suite name.
* `versions` - SSL/TLS versions that the cipher suite can be used with.

The `ssl_server_cipher_suites` block supports:

* `priority` - SSL/TLS cipher suites priority.
* `cipher` - Cipher suite name.
* `versions` - SSL/TLS versions that the cipher suite can be used with.

The `monitor` block supports:

* `name` - Health monitor name.

The `gslb_public_ips` block supports:

* `index` - Index of this public IP setting.
* `ip` - The publicly accessible IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vip can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_vip.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_vip.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
