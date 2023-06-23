---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_pcpserver"
description: |-
  Configure PCP server information.
---

# fortios_system_pcpserver
Configure PCP server information. Applies to FortiOS Version `>= 7.4.0`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable PCP server. Valid values: `enable`, `disable`.
* `pools` - Configure PCP pools. The structure of `pools` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `pools` block supports:

* `name` - PCP pool name.
* `description` - Description.
* `id` - ID.
* `client_subnet` - Subnets from which PCP requests are accepted. The structure of `client_subnet` block is documented below.
* `ext_intf` - External interface name.
* `arp_reply` - Enable to respond to ARP requests for external IP (default = enable). Valid values: `disable`, `enable`.
* `extip` - IP address or address range on the external interface that you want to map to an address on the internal network.
* `extport` - Incoming port number range that you want to map to a port number on the internal network.
* `minimal_lifetime` - Minimal lifetime of a PCP mapping in seconds (60 - 300, default = 120).
* `maximal_lifetime` - Maximal lifetime of a PCP mapping in seconds (3600 - 604800, default = 86400).
* `client_mapping_limit` - Mapping limit per client (0 - 65535, default = 0, 0 = unlimited).
* `mapping_filter_limit` - Filter limit per mapping (0 - 5, default = 1).
* `allow_opcode` - Allowed PCP opcode. Valid values: `map`, `peer`, `announce`.
* `third_party` - Allow/disallow third party option. Valid values: `allow`, `disallow`.
* `third_party_subnet` - Subnets from which third party requests are accepted. The structure of `third_party_subnet` block is documented below.
* `multicast_announcement` - Enable/disable multicast announcements. Valid values: `enable`, `disable`.
* `announcement_count` - Number of multicast announcements.
* `intl_intf` - Internal interface name. The structure of `intl_intf` block is documented below.
* `recycle_delay` - Minimum delay (in seconds) the PCP Server will wait before recycling mappings that have expired (0 - 3600, default = 0).

The `client_subnet` block supports:

* `subnet` - Client subnets.

The `third_party_subnet` block supports:

* `subnet` - Third party subnets.

The `intl_intf` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System PcpServer can be imported using any of these accepted formats:
```
$ terraform import fortios_system_pcpserver.labelname SystemPcpServer

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_pcpserver.labelname SystemPcpServer
$ unset "FORTIOS_IMPORT_TABLE"
```
