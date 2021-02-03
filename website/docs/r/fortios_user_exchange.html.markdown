---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_exchange"
description: |-
  Configure MS Exchange server entries.
---

# fortios_user_exchange
Configure MS Exchange server entries.

## Argument Reference

The following arguments are supported:

* `name` - MS Exchange server entry name.
* `server_name` - MS Exchange server hostname.
* `domain_name` - MS Exchange server fully qualified domain name.
* `username` - User name used to sign in to the server. Must have proper permissions for service.
* `password` - Password for the specified username.
* `ip` - Server IPv4 address.
* `connect_protocol` - Connection protocol used to connect to MS Exchange service.
* `auth_type` - Authentication security type used for the RPC protocol layer.
* `auth_level` - Authentication security level used for the RPC protocol layer.
* `http_auth_type` - Authentication security type used for the HTTP transport.
* `ssl_min_proto_version` - Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting).
* `auto_discover_kdc` - Enable/disable automatic discovery of KDC IP addresses.
* `kdc_ip` - KDC IPv4 addresses for Kerberos authentication. The structure of `kdc_ip` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `kdc_ip` block supports:

* `ipv4` - KDC IPv4 addresses for Kerberos authentication.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Exchange can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_exchange.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
