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
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `community` block supports:

* `id` - Community ID.
* `name` - Community name.
* `status` - Enable/disable this SNMP community. Valid values: `enable`, `disable`.
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
* `hosts` - Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.

The `hosts` block supports:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).

The `user` block supports:

* `name` - SNMP User Name
* `status` - SNMP User Enable Valid values: `enable`, `disable`.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
* `trap_status` - Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
* `auth_proto` - Authentication protocol. Valid values: `md5`, `sha`.
* `auth_pwd` - Password for authentication protocol.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
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
