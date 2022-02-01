---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vip6"
description: |-
  Configure virtual IP for IPv6.
---

# fortios_firewall_vip6
Configure virtual IP for IPv6.

## Example Usage

```hcl
resource "fortios_firewall_vip6" "trname" {
  arp_reply                        = "enable"
  color                            = 0
  extip                            = "2001:1:1:12::100"
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
  mappedip                         = "2001:1:1:12::200"
  mappedport                       = "0-65535"
  max_embryonic_connections        = 1000
  name                             = "vip6s1"
  outlook_web_access               = "disable"
  persistence                      = "none"
  portforward                      = "disable"
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
}
```

## Argument Reference

The following arguments are supported:

* `name` - Virtual ip6 name.
* `fosid` - Custom defined ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `comment` - Comment.
* `type` - Configure a static NAT or server load balance VIP.
* `src_filter` - Source IP6 filter (x:x:x:x:x:x:x:x/x). Separate addresses with spaces. The structure of `src_filter` block is documented below.
* `extip` - (Required) IP address or address range on the external interface that you want to map to an address or address range on the destination network.
* `mappedip` - (Required) Mapped IP address range in the format startIP-endIP.
* `nat_source_vip` - Enable to perform SNAT on traffic from mappedip to the extip for all egress interfaces. Valid values: `disable`, `enable`.
* `arp_reply` - Enable to respond to ARP requests for this virtual IP address. Enabled by default. Valid values: `disable`, `enable`.
* `portforward` - Enable port forwarding. Valid values: `disable`, `enable`.
* `protocol` - Protocol to use when forwarding packets. Valid values: `tcp`, `udp`, `sctp`.
* `extport` - Incoming port number range that you want to map to a port number range on the destination network.
* `mappedport` - Port number range on the destination network to which the external port number range is mapped.
* `color` - Color of icon on the GUI.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`, `http-host`.
* `server_type` - Protocol to be load balanced by the virtual server (also called the server load balance virtual IP). Valid values: `http`, `https`, `imaps`, `pop3s`, `smtps`, `ssl`, `tcp`, `udp`, `ip`.
* `http_redirect` - Enable/disable redirection of HTTP to HTTPS Valid values: `enable`, `disable`.
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none`, `http-cookie`, `ssl-session-id`.
* `nat66` - Enable/disable DNAT66. Valid values: `disable`, `enable`.
* `nat64` - Enable/disable DNAT64. Valid values: `disable`, `enable`.
* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
* `realservers` - Select the real servers that this server load balancing VIP will distribute traffic to. The structure of `realservers` block is documented below.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable`, `enable`.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 seconds. 0 = no time limit.
* `http_cookie_share` - Control sharing of cookies across virtual servers. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable`, `same-ip`.
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable`, `enable`.
* `http_multiplex` - Enable/disable HTTP multiplexing. Valid values: `enable`, `disable`.
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
* `ssl_server_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the server and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.
* `ssl_server_session_state_timeout` - Number of minutes to keep FortiGate to Server SSL session state.
* `ssl_server_session_state_max` - Maximum number of FortiGate to Server SSL session states to keep.
* `ssl_http_location_conversion` - Enable to replace HTTP with HTTPS in the reply's Location HTTP header field. Valid values: `enable`, `disable`.
* `ssl_http_match_host` - Enable/disable HTTP host matching for location conversion. Valid values: `enable`, `disable`.
* `ssl_hpkp` - Enable/disable including HPKP header in response. Valid values: `disable`, `enable`, `report-only`.
* `ssl_hpkp_primary` - Certificate to generate primary HPKP pin from.
* `ssl_hpkp_backup` - Certificate to generate backup HPKP pin from.
* `ssl_hpkp_age` - Number of minutes the web browser should keep HPKP.
* `ssl_hpkp_report_uri` - URL to report HPKP violations to.
* `ssl_hpkp_include_subdomains` - Indicate that HPKP header applies to all subdomains. Valid values: `disable`, `enable`.
* `ssl_hsts` - Enable/disable including HSTS header in response. Valid values: `disable`, `enable`.
* `ssl_hsts_age` - Number of seconds the client should honour the HSTS setting.
* `ssl_hsts_include_subdomains` - Indicate that HSTS header applies to all subdomains. Valid values: `disable`, `enable`.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status. The structure of `monitor` block is documented below.
* `max_embryonic_connections` - Maximum number of incomplete connections.
* `embedded_ipv4_address` - Enable/disable embedded IPv4 address. Valid values: `disable`, `enable`.
* `ipv4_mappedip` - Start-mapped-IPv4-address [-end mapped-IPv4-address].
* `ipv4_mappedport` - IPv4 port number range on the destination network to which the external port number range is mapped.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `src_filter` block supports:

* `range` - Source-filter range.

The `realservers` block supports:

* `id` - Real server ID.
* `ip` - IPv6 address of the real server.
* `port` - Port for communicating with the real server. Required if port forwarding is enabled.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `holddown_interval` - Time in seconds that the health check monitor continues to monitor an unresponsive server that should be active.
* `healthcheck` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`, `vip`.
* `http_host` - HTTP server domain name in HTTP header.
* `max_connections` - Max number of active connections that can directed to the real server. When reached, sessions are sent to other real servers.
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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vip6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_vip6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
