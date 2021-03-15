---
subcategory: "FortiGate ICAP"
layout: "fortios"
page_title: "FortiOS: fortios_icap_server"
description: |-
  Configure ICAP servers.
---

# fortios_icap_server
Configure ICAP servers.

## Example Usage

```hcl
resource "fortios_icap_server" "trname" {
  ip6_address     = "::"
  ip_address      = "1.1.1.1"
  ip_version      = "4"
  max_connections = 100
  name            = "1"
  port            = 22
}
```

## Argument Reference

The following arguments are supported:

* `name` - Server name.
* `ip_version` - IP version. Valid values: `4`, `6`.
* `ip_address` - IPv4 address of the ICAP server.
* `ip6_address` - IPv6 address of the ICAP server.
* `port` - ICAP server port.
* `max_connections` - Maximum number of concurrent connections to ICAP server.
* `secure` - Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
* `ssl_cert` - CA certificate name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Icap Server can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_icap_server.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
