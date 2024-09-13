---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_snmpcommunity"
description: |-
  Configure FortiSwitch SNMP v1/v2c communities globally.
---

# fortios_switchcontroller_snmpcommunity
Configure FortiSwitch SNMP v1/v2c communities globally. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `fosid` - SNMP community ID.
* `name` - SNMP community name.
* `status` - Enable/disable this SNMP community. Valid values: `disable`, `enable`.
* `hosts` - Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.
* `query_v2c_port` - SNMP v2c query port (default = 161).
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.
* `trap_v1_lport` - SNMP v2c trap local port (default = 162).
* `trap_v1_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.
* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `events` - SNMP notifications (traps) to send.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `hosts` block supports:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController SnmpCommunity can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_snmpcommunity.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_snmpcommunity.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
