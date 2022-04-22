---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_exchange"
description: |-
  Configure MS Exchange server entries.
---

# fortios_user_exchange
Configure MS Exchange server entries. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `name` - MS Exchange server entry name.
* `server_name` - MS Exchange server hostname.
* `domain_name` - MS Exchange server fully qualified domain name.
* `username` - User name used to sign in to the server. Must have proper permissions for service.
* `password` - Password for the specified username.
* `ip` - Server IPv4 address.
* `connect_protocol` - Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
* `auth_type` - Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
* `auth_level` - Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
* `http_auth_type` - Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
* `ssl_min_proto_version` - Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
* `auto_discover_kdc` - Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
* `kdc_ip` - KDC IPv4 addresses for Kerberos authentication. The structure of `kdc_ip` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `kdc_ip` block supports:

* `ipv4` - KDC IPv4 addresses for Kerberos authentication.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Exchange can be imported using any of these accepted formats:
```
$ terraform import fortios_user_exchange.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_exchange.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
