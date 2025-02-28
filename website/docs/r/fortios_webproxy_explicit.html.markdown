---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_explicit"
description: |-
  Configure explicit Web proxy settings.
---

# fortios_webproxy_explicit
Configure explicit Web proxy settings.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable the explicit Web proxy for HTTP and HTTPS session. Valid values: `enable`, `disable`.
* `secure_web_proxy` - Enable/disable/require the secure web proxy for HTTP and HTTPS session. Valid values: `disable`, `enable`, `secure`.
* `ftp_over_http` - Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
* `socks` - Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
* `http_incoming_port` - Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
* `http_connection_mode` - HTTP connection mode (default = static). Valid values: `static`, `multiplex`, `serverpool`.
* `https_incoming_port` - Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
* `secure_web_proxy_cert` - Name of certificates for secure web proxy. The structure of `secure_web_proxy_cert` block is documented below.
* `client_cert` - Enable/disable to request client certificate. Valid values: `disable`, `enable`.
* `user_agent_detect` - Enable/disable to detect device type by HTTP user-agent if no client certificate provided. Valid values: `disable`, `enable`.
* `empty_cert_action` - Action of an empty client certificate. Valid values: `accept`, `block`, `accept-unmanageable`.
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
* `ftp_incoming_port` - Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
* `socks_incoming_port` - Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
* `incoming_ip` - Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
* `outgoing_ip` - Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `ipv6_status` - Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
* `incoming_ip6` - Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
* `outgoing_ip6` - Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
* `strict_guest` - Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
* `pref_dns_result` - Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4).
* `unknown_http_version` - Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
* `realm` - Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
* `sec_default_action` - Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
* `https_replacement_message` - Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
* `message_upon_server_error` - Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
* `pac_file_server_status` - Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
* `pac_file_url` - PAC file access URL.
* `pac_file_server_port` - Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
* `pac_file_through_https` - Enable/disable to get Proxy Auto-Configuration (PAC) through HTTPS. Valid values: `enable`, `disable`.
* `pac_file_name` - Pac file name.
* `pac_file_data` - PAC file contents enclosed in quotes (maximum of 256K bytes).
* `pac_policy` - PAC policies. The structure of `pac_policy` block is documented below.
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
* `trace_auth_no_rsp` - Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `secure_web_proxy_cert` block supports:

* `name` - Certificate list.

The `pac_policy` block supports:

* `policyid` - Policy ID.
* `status` - Enable/disable policy. Valid values: `enable`, `disable`.
* `srcaddr` - Source address objects. The structure of `srcaddr` block is documented below.
* `srcaddr6` - Source address6 objects. The structure of `srcaddr6` block is documented below.
* `dstaddr` - Destination address objects. The structure of `dstaddr` block is documented below.
* `pac_file_name` - Pac file name.
* `pac_file_data` - PAC file contents enclosed in quotes (maximum of 256K bytes).
* `comments` - Optional comments.

The `srcaddr` block supports:

* `name` - Address name.

The `srcaddr6` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WebProxy Explicit can be imported using any of these accepted formats:
```
$ terraform import fortios_webproxy_explicit.labelname WebProxyExplicit

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webproxy_explicit.labelname WebProxyExplicit
$ unset "FORTIOS_IMPORT_TABLE"
```
