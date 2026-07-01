---
subcategory: "FortiGate Ztna"
layout: "fortios"
page_title: "FortiOS: fortios_ztna_reverseconnector"
description: |-
  Configure ZTNA Reverse-Connector.
---

# fortios_ztna_reverseconnector
Configure ZTNA Reverse-Connector. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `name` - Reverse-Connector name
* `status` - Reverse-Connector status. Valid values: `enable`, `disable`.
* `address` - Connector service edge adress(IP or FQDN).
* `port` - Port number that traffic uses to connect to connector service edge. On FortiOS versions 7.6.1-7.6.7: 0 - 65535. On FortiOS versions >= 8.0.0: 1 - 65535.
* `default_incoming_vip` - Default Incoming Virtual IP name.
* `source_ip` - FortiGate IPv4 address to be used for ZTNA reverse-connector connection.
* `source_ip_interface` - Source interface to be used for ZTNA reverse-connector connection.
* `health_check_interval` - Health check interval in seconds (0 - 600, default = 60, 0 = disable).
* `ssl_min_version` - Lowest SSL/TLS version acceptable from a server. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `ssl_max_version` - Highest TLS version acceptable from a server. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `certificate` - The name of the certificate to use for SSL handshake.
* `trusted_server_ca` - Trusted Server CA certificate used by SSL connection.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna ReverseConnector can be imported using any of these accepted formats:
```
$ terraform import fortios_ztna_reverseconnector.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ztna_reverseconnector.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
