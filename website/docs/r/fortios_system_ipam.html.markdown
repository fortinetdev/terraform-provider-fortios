---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipam"
description: |-
  Configure IP address management services.
---

# fortios_system_ipam
Configure IP address management services. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable IP address management services. Valid values: `enable`, `disable`.
* `server_type` - Configure the type of IPAM server to use.
* `automatic_conflict_resolution` - Enable/disable automatic conflict resolution. Valid values: `disable`, `enable`.
* `manage_lan_addresses` - Enable/disable default management of LAN interface addresses. Valid values: `disable`, `enable`.
* `manage_lan_extension_addresses` - Enable/disable default management of FortiExtender LAN extension interface addresses. Valid values: `disable`, `enable`.
* `manage_ssid_addresses` - Enable/disable default management of FortiAP SSID addresses. Valid values: `disable`, `enable`.
* `pools` - Configure IPAM pools. The structure of `pools` block is documented below.
* `rules` - Configure IPAM allocation rules. The structure of `rules` block is documented below.
* `pool_subnet` - Configure IPAM pool subnet, Class A - Class B subnet.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `pools` block supports:

* `name` - IPAM pool name.
* `description` - Description.
* `subnet` - Configure IPAM pool subnet, Class A - Class B subnet.

The `rules` block supports:

* `name` - IPAM rule name.
* `description` - Description.
* `device` - Configure serial number or wildcard of Fortigate to match. The structure of `device` block is documented below.
* `interface` - Configure name or wildcard of interface to match. The structure of `interface` block is documented below.
* `role` - Configure role of interface to match. Valid values: `any`, `lan`, `wan`, `dmz`, `undefined`.
* `pool` - Configure name of IPAM pool to use. The structure of `pool` block is documented below.
* `dhcp` - Enable/disable DHCP server for matching IPAM interfaces. Valid values: `enable`, `disable`.

The `device` block supports:

* `name` - Fortigate serial number or wildcard.

The `interface` block supports:

* `name` - Interface name or wildcard.

The `pool` block supports:

* `name` - Ipam pool name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ipam can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ipam.labelname SystemIpam

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ipam.labelname SystemIpam
$ unset "FORTIOS_IMPORT_TABLE"
```
