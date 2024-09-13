---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_community"
description: |-
  Get information on an fortios systemsnmp community.
---

# Data Source: fortios_systemsnmp_community
Use this data source to get information on an fortios systemsnmp community

## Argument Reference

* `fosid` - (Required) Specify the fosid of the desired systemsnmp community.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fosid` - Community ID.
* `name` - Community name.
* `status` - Enable/disable this SNMP community.
* `hosts` - Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
* `hosts6` - Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
* `query_v1_status` - Enable/disable SNMP v1 queries.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries.
* `query_v2c_port` - SNMP v2c query port (default = 161).
* `trap_v1_status` - Enable/disable SNMP v1 traps.
* `trap_v1_lport` - SNMP v1 trap local port (default = 162).
* `trap_v1_rport` - SNMP v1 trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps.
* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `events` - SNMP trap events.
* `mib_view` - SNMP access control MIB view.
* `vdoms` - SNMP access control VDOMs. The structure of `vdoms` block is documented below.

The `hosts` block contains:

* `id` - Host entry ID.
* `source_ip` - Source IPv4 address for SNMP traps.
* `ip` - IPv4 address of the SNMP manager (host).
* `ha_direct` - Enable/disable direct management of HA cluster members.
* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

The `hosts6` block contains:

* `id` - Host6 entry ID.
* `source_ipv6` - Source IPv6 address for SNMP traps.
* `ipv6` - SNMP manager IPv6 address prefix.
* `ha_direct` - Enable/disable direct management of HA cluster members.
* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

The `vdoms` block contains:

* `name` - VDOM name

