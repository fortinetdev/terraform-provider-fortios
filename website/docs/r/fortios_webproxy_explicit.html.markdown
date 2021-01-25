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

* `status` - Enable/disable the explicit Web proxy for HTTP and HTTPS session.
* `ftp_over_http` - Enable to proxy FTP-over-HTTP sessions sent from a web browser.
* `socks` - Enable/disable the SOCKS proxy.
* `http_incoming_port` - Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
* `https_incoming_port` - Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
* `ftp_incoming_port` - Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
* `socks_incoming_port` - Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
* `incoming_ip` - Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
* `outgoing_ip` - Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
* `ipv6_status` - Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command.
* `incoming_ip6` - Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
* `outgoing_ip6` - Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
* `strict_guest` - Enable/disable strict guest user checking by the explicit web proxy.
* `pref_dns_result` - Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4).
* `unknown_http_version` - Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
* `realm` - Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
* `sec_default_action` - Accept or deny explicit web proxy sessions when no web proxy firewall policy exists.
* `https_replacement_message` - Enable/disable sending the client a replacement message for HTTPS requests.
* `message_upon_server_error` - Enable/disable displaying a replacement message when a server error is detected.
* `pac_file_server_status` - Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile.
* `pac_file_url` - PAC file access URL.
* `pac_file_server_port` - Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
* `pac_file_name` - Pac file name.
* `pac_file_data` - PAC file contents enclosed in quotes (maximum of 256K bytes).
* `pac_policy` - PAC policies. The structure of `pac_policy` block is documented below.
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low.
* `trace_auth_no_rsp` - Enable/disable logging timed-out authentication requests.

The `pac_policy` block supports:

* `policyid` - Policy ID.
* `status` - Enable/disable policy.
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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WebProxy Explicit can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webproxy_explicit.labelname WebProxyExplicit
$ unset "FORTIOS_IMPORT_TABLE"
```
