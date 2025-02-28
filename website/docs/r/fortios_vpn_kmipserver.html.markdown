---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_kmipserver"
description: |-
  KMIP server entry configuration.
---

# fortios_vpn_kmipserver
KMIP server entry configuration. Applies to FortiOS Version `>= 7.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - KMIP server entry name.
* `server_list` - KMIP server list. The structure of `server_list` block is documented below.
* `username` - User name to use for connectivity to the KMIP server.
* `password` - Password to use for connectivity to the KMIP server.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `server_identity_check` - Enable/disable KMIP server identity check (verify server FQDN/IP address against the server certificate). Valid values: `enable`, `disable`.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `source_ip` - FortiGate IP address to be used for communication with the KMIP server.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `server_list` block supports:

* `id` - ID
* `status` - Enable/disable KMIP server. Valid values: `enable`, `disable`.
* `server` - KMIP server FQDN or IP address.
* `port` - KMIP server port.
* `cert` - Client certificate to use for connectivity to the KMIP server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn KmipServer can be imported using any of these accepted formats:
```
$ terraform import fortios_vpn_kmipserver.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpn_kmipserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
