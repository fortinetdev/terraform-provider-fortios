---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_snmp"
description: |-
  Configure SNMP.
---

# fortios_wirelesscontroller_snmp
Configure SNMP.

## Argument Reference

The following arguments are supported:

* `engine_id` - AC SNMP engineId string (maximum 24 characters).
* `contact_info` - Contact Information.
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_high_mem_threshold` - Memory usage when trap is sent.
* `community` - SNMP Community Configuration. The structure of `community` block is documented below.
* `user` - SNMP User Configuration. The structure of `user` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `community` block supports:

* `id` - Community ID.
* `name` - Community name.
* `status` - Enable/disable this SNMP community.
* `query_v1_status` - Enable/disable SNMP v1 queries.
* `query_v2c_status` - Enable/disable SNMP v2c queries.
* `trap_v1_status` - Enable/disable SNMP v1 traps.
* `trap_v2c_status` - Enable/disable SNMP v2c traps.
* `hosts` - Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.

The `hosts` block supports:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).

The `user` block supports:

* `name` - SNMP User Name
* `status` - SNMP User Enable
* `queries` - Enable/disable SNMP queries for this user.
* `trap_status` - Enable/disable traps for this SNMP user.
* `security_level` - Security level for message authentication and encryption.
* `auth_proto` - Authentication protocol.
* `auth_pwd` - Password for authentication protocol.
* `priv_proto` - Privacy (encryption) protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.
* `notify_hosts` - Configure SNMP User Notify Hosts.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Snmp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_snmp.labelname WirelessControllerSnmp
$ unset "FORTIOS_IMPORT_TABLE"
```
