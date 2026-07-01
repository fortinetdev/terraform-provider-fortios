---
subcategory: "FortiGate Ztna"
layout: "fortios"
page_title: "FortiOS: fortios_ztna_connectoredge"
description: |-
  Configure ZTNA connector edge.
---

# fortios_ztna_connectoredge
Configure ZTNA connector edge. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `status` - Connector service edge status. Valid values: `enable`, `disable`.
* `interface` - Connector edge interface. The structure of `interface` block is documented below.
* `port` - Connector service edge port(1-65535, default 8443).
* `ssl_min_version` - Lowest TLS version accepted by server. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `ssl_max_version` - Highest TLS version accepted by server. Valid values: `tls-1.1`, `tls-1.2`, `tls-1.3`.
* `server_cert` - Server certificate for mTLS.
* `trusted_client_ca` - CA certificate used for client certificate verification. The structure of `trusted_client_ca` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `interface` block supports:

* `name` - Interface name.

The `trusted_client_ca` block supports:

* `name` - CA certificate list.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ztna ConnectorEdge can be imported using any of these accepted formats:
```
$ terraform import fortios_ztna_connectoredge.labelname ZtnaConnectorEdge

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ztna_connectoredge.labelname ZtnaConnectorEdge
$ unset "FORTIOS_IMPORT_TABLE"
```
