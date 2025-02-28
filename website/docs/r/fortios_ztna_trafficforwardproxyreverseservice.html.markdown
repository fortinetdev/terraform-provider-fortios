---
subcategory: "FortiGate Ztna"
layout: "fortios"
page_title: "FortiOS: fortios_ztna_trafficforwardproxyreverseservice"
description: |-
  Configure ZTNA traffic forward proxy reverse service.
---

# fortios_ztna_trafficforwardproxyreverseservice
Configure ZTNA traffic forward proxy reverse service. Applies to FortiOS Version `7.6.0`.

## Argument Reference

The following arguments are supported:

* `remote_servers` - Connector Remote server The structure of `remote_servers` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `remote_servers` block supports:

* `name` - Remote server name
* `status` - Remote server status. Valid values: `enable`, `disable`.
* `address` - Server adress(IP or FQDN).
* `health_check_interval` - Health check interval in seconds (0 - 600, default = 60, 0 = disable).
* `ssl_max_version` - Highest TLS version acceptable from a server. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `port` - Port number that traffic uses to connect to remote server (0 - 65535;).
* `certificate` - The name of the certificate to use for SSL handshake.
* `trusted_server_ca` - Trusted Server CA certificate used by SSL connection.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ztna TrafficForwardProxyReverseService can be imported using any of these accepted formats:
```
$ terraform import fortios_ztna_trafficforwardproxyreverseservice.labelname ZtnaTrafficForwardProxyReverseService

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ztna_trafficforwardproxyreverseservice.labelname ZtnaTrafficForwardProxyReverseService
$ unset "FORTIOS_IMPORT_TABLE"
```
