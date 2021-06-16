---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_snmpcommunity"
description: |-
  Configure FortiSwitch SNMP v1/v2c communities globally.
---

# fortios_switchcontroller_snmpcommunity
Configure FortiSwitch SNMP v1/v2c communities globally.

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
* `events` - SNMP notifications (traps) to send. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `ent-conf-change`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_snmpcommunity.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
