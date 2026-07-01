---
subcategory: "FortiGate Ztna"
layout: "fortios"
page_title: "FortiOS: fortios_ztna_serviceconnector"
description: |-
  Configure ZTNA service connector.
---

# fortios_ztna_serviceconnector
Configure ZTNA service connector. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - Service-connector name
* `connection_mode` - Connection mode. Valid values: `forward`, `reverse`.
* `forward_address` - service-connector address(IP or FQDN).
* `forward_port` - Port number that forward traffic uses to connect to
* `forward_destination_cn` - CN for forward server.
* `certificate` - The name of the certificate to use for SSL handshake.
* `trusted_ca` - Trusted CA certificate used by SSL inspection.
* `encryption` - Enable/disable Encryption (default = disable). Valid values: `enable`, `disable`.
* `ssl_max_version` - Highest SSL/TLS version acceptable from a server. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a server. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `url_map` - URL pattern to match.
* `relay_dev_info` - Enable/disable device info relay. Valid values: `enable`, `disable`.
* `relay_user_info` - Enable/disable user info relay. Valid values: `enable`, `disable`.
* `log` - Enable/disable logging of traffic. Valid values: `enable`, `disable`.
* `health_check_interval` - health check interval(0-600 seconds, default 60, 0 means disable).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna ServiceConnector can be imported using any of these accepted formats:
```
$ terraform import fortios_ztna_serviceconnector.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ztna_serviceconnector.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
