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
* `addr_type` - Address type of the remote ICAP server: IPv4, IPv6 or FQDN. Valid values: `ip4`, `ip6`, `fqdn`.
* `ip_version` - IP version. Valid values: `4`, `6`.
* `ip_address` - IPv4 address of the ICAP server.
* `ip6_address` - IPv6 address of the ICAP server.
* `fqdn` - ICAP remote server Fully Qualified Domain Name (FQDN).
* `port` - ICAP server port.
* `max_connections` - Maximum number of concurrent connections to ICAP server. Must not be less than wad-worker-count.
* `secure` - Enable/disable secure connection to ICAP server. Valid values: `enable`, `disable`.
* `ssl_cert` - CA certificate name.
* `healthcheck` - Enable/disable ICAP remote server health checking. Attempts to connect to the remote ICAP server to verify that the server is operating normally. Valid values: `disable`, `enable`.
* `healthcheck_service` - ICAP Service name to use for health checks.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Icap Server can be imported using any of these accepted formats:
```
$ terraform import fortios_icap_server.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_icap_server.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
