---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_accessproxy6"
description: |-
  Configure IPv6 access proxy.
---

# fortios_firewall_accessproxy6
Configure IPv6 access proxy. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - Access Proxy name.
* `vip` - Virtual IP name.
* `client_cert` - Enable/disable to request client certificate. Valid values: `disable`, `enable`.
* `user_agent_detect` - Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
* `auth_portal` - Enable/disable authentication portal. Valid values: `disable`, `enable`.
* `auth_virtual_host` - Virtual host for authentication portal.
* `empty_cert_action` - Action of an empty client certificate.
* `log_blocked_traffic` - Enable/disable logging of blocked traffic. Valid values: `enable`, `disable`.
* `add_vhost_domain_to_dnsdb` - Enable/disable adding vhost/domain to dnsdb for ztna dox tunnel. Valid values: `enable`, `disable`.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `api_gateway` - Set IPv4 API Gateway. The structure of `api_gateway` block is documented below.
* `api_gateway6` - Set IPv6 API Gateway. The structure of `api_gateway6` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `api_gateway` block supports:

* `id` - API Gateway ID.
* `url_map` - URL pattern to match.
* `service` - Service.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static`, `round-robin`, `weighted`, `first-alive`, `http-host`.
* `virtual_host` - Virtual host.
* `url_map_type` - Type of url-map. Valid values: `sub-string`, `wildcard`, `regex`.
* `realservers` - Select the real servers that this Access Proxy will distribute traffic to. The structure of `realservers` block is documented below.
* `application` - SaaS application controlled by this Access Proxy. The structure of `application` block is documented below.
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none`, `http-cookie`.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable`, `enable`.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_share` - Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable`, `same-ip`.
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable`, `enable`.
* `saml_server` - SAML service provider configuration for VIP authentication.
* `saml_redirect` - Enable/disable SAML redirection after successful authentication. Valid values: `disable`, `enable`.
* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.
* `ssl_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high`, `medium`, `low`.
* `ssl_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a server. Valid values: `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `ssl_max_version` - Highest SSL/TLS version acceptable from a server. Valid values: `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `ssl_vpn_web_portal` - SSL-VPN web portal.

The `realservers` block supports:

* `id` - Real server ID.
* `addr_type` - Type of address. Valid values: `ip`, `fqdn`.
* `address` - Address or address group of the real server.
* `ip` - IP address of the real server.
* `domain` - Wildcard domain name of the real server.
* `port` - Port for communicating with the real server.
* `mappedport` - Port for communicating with the real server.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.
* `type` - TCP forwarding server type. Valid values: `tcp-forwarding`, `ssh`.
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `http_host` - HTTP server domain name in HTTP header.
* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`.
* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping`, `http`, `tcp-connect`.
* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `enable`, `disable`.
* `ssh_client_cert` - Set access-proxy SSH client certificate profile.
* `ssh_host_key_validation` - Enable/disable SSH real server host key validation. Valid values: `disable`, `enable`.
* `ssh_host_key` - One or more server host key. The structure of `ssh_host_key` block is documented below.

The `ssh_host_key` block supports:

* `name` - Server host key name.

The `application` block supports:

* `name` - SaaS application name.

The `ssl_cipher_suites` block supports:

* `priority` - SSL/TLS cipher suites priority.
* `cipher` - Cipher suite name. Valid values: `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

The `api_gateway6` block supports:

* `id` - API Gateway ID.
* `url_map` - URL pattern to match.
* `service` - Service.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static`, `round-robin`, `weighted`, `first-alive`, `http-host`.
* `virtual_host` - Virtual host.
* `url_map_type` - Type of url-map. Valid values: `sub-string`, `wildcard`, `regex`.
* `realservers` - Select the real servers that this Access Proxy will distribute traffic to. The structure of `realservers` block is documented below.
* `application` - SaaS application controlled by this Access Proxy. The structure of `application` block is documented below.
* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none`, `http-cookie`.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable`, `enable`.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `http_cookie_share` - Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable`, `same-ip`.
* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable`, `enable`.
* `saml_server` - SAML service provider configuration for VIP authentication.
* `saml_redirect` - Enable/disable SAML redirection after successful authentication. Valid values: `disable`, `enable`.
* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.
* `ssl_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high`, `medium`, `low`.
* `ssl_cipher_suites` - SSL/TLS cipher suites to offer to a server, ordered by priority. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a server. Valid values: `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `ssl_max_version` - Highest SSL/TLS version acceptable from a server. Valid values: `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `ssl_vpn_web_portal` - SSL-VPN web portal.

The `realservers` block supports:

* `id` - Real server ID.
* `addr_type` - Type of address. Valid values: `ip`, `fqdn`.
* `address` - Address or address group of the real server.
* `ip` - IPv6 address of the real server.
* `domain` - Wildcard domain name of the real server.
* `port` - Port for communicating with the real server.
* `mappedport` - Port for communicating with the real server.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.
* `type` - TCP forwarding server type. Valid values: `tcp-forwarding`, `ssh`.
* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
* `http_host` - HTTP server domain name in HTTP header.
* `health_check` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`.
* `health_check_proto` - Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping`, `http`, `tcp-connect`.
* `holddown_interval` - Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `enable`, `disable`.
* `ssh_client_cert` - Set access-proxy SSH client certificate profile.
* `ssh_host_key_validation` - Enable/disable SSH real server host key validation. Valid values: `disable`, `enable`.
* `ssh_host_key` - One or more server host key. The structure of `ssh_host_key` block is documented below.

The `ssh_host_key` block supports:

* `name` - Server host key name.

The `application` block supports:

* `name` - SaaS application name.

The `ssl_cipher_suites` block supports:

* `priority` - SSL/TLS cipher suites priority.
* `cipher` - Cipher suite name. Valid values: `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall AccessProxy6 can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_accessproxy6.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_accessproxy6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
