---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_fastfallback"
description: |-
  Proxy destination connection fast-fallback.
---

# fortios_webproxy_fastfallback
Proxy destination connection fast-fallback. Applies to FortiOS Version `>= 7.4.1`.

## Argument Reference

The following arguments are supported:

* `name` - Configure a name for the fast-fallback entry.
* `status` - Enable/disable the fast-fallback entry. Valid values: `enable`, `disable`.
* `connection_mode` - Connection mode for multiple destinations. Valid values: `sequentially`, `simultaneously`.
* `protocol` - Connection protocols for multiple destinations. Valid values: `IPv4-first`, `IPv6-first`, `IPv4-only`, `IPv6-only`.
* `connection_timeout` - Number of milliseconds to wait before starting another connection (200 - 1800000, default = 200). For sequential connection-mode only.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy FastFallback can be imported using any of these accepted formats:
```
$ terraform import fortios_webproxy_fastfallback.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webproxy_fastfallback.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
