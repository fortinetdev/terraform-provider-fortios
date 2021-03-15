---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_fortiguard"
description: |-
  Configure FortiGuard Web Filter service.
---

# fortios_webfilter_fortiguard
Configure FortiGuard Web Filter service.

## Example Usage

```hcl
resource "fortios_webfilter_fortiguard" "trname" {
  cache_mem_percent         = 2
  cache_mode                = "ttl"
  cache_prefix_match        = "enable"
  close_ports               = "disable"
  ovrd_auth_https           = "enable"
  ovrd_auth_port            = 8008
  ovrd_auth_port_http       = 8008
  ovrd_auth_port_https      = 8010
  ovrd_auth_port_https_flow = 8015
  ovrd_auth_port_warning    = 8020
  warn_auth_https           = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `cache_mode` - Cache entry expiration mode. Valid values: `ttl`, `db-ver`.
* `cache_prefix_match` - Enable/disable prefix matching in the cache. Valid values: `enable`, `disable`.
* `cache_mem_percent` - Maximum percentage of available memory allocated to caching (1 - 15%).
* `ovrd_auth_port_http` - Port to use for FortiGuard Web Filter HTTP override authentication
* `ovrd_auth_port_https` - Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.
* `ovrd_auth_port_https_flow` - Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.
* `ovrd_auth_port_warning` - Port to use for FortiGuard Web Filter Warning override authentication.
* `ovrd_auth_https` - Enable/disable use of HTTPS for override authentication. Valid values: `enable`, `disable`.
* `warn_auth_https` - Enable/disable use of HTTPS for warning and authentication. Valid values: `enable`, `disable`.
* `close_ports` - Close ports used for HTTP/HTTPS override authentication and disable user overrides. Valid values: `enable`, `disable`.
* `request_packet_size_limit` - Limit size of URL request packets sent to FortiGuard server (0 for default).
* `ovrd_auth_port` - Port to use for FortiGuard Web Filter override authentication.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter Fortiguard can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webfilter_fortiguard.labelname WebfilterFortiguard
$ unset "FORTIOS_IMPORT_TABLE"
```
