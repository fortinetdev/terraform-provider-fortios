---
subcategory: "FortiGate ICAP"
layout: "fortios"
page_title: "FortiOS: fortios_icap_servergroup"
description: |-
  Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing.
---

# fortios_icap_servergroup
Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing. Applies to FortiOS Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
* `ldb_method` - Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.
* `server_list` - Add ICAP servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `server_list` block supports:

* `name` - ICAP server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Icap ServerGroup can be imported using any of these accepted formats:
```
$ terraform import fortios_icap_servergroup.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_icap_servergroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
