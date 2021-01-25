---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_community"
description: |-
  SNMP community configuration.
---

# fortios_systemsnmp_community
SNMP community configuration.

## Example Usage

```hcl
resource "fortios_systemsnmp_community" "trname" {
  events           = "cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high"
  fosid            = 1
  name             = "fdsie"
  query_v1_port    = 161
  query_v1_status  = "enable"
  query_v2c_port   = 161
  query_v2c_status = "enable"
  status           = "enable"
  trap_v1_lport    = 162
  trap_v1_rport    = 162
  trap_v1_status   = "enable"
  trap_v2c_lport   = 162
  trap_v2c_rport   = 162
  trap_v2c_status  = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `fosid` - (Required) Community ID.
* `name` - (Required) Community name.
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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `hosts` block supports:

* `id` - Host entry ID.
* `source_ip` - Source IPv4 address for SNMP traps.
* `ip` - IPv4 address of the SNMP manager (host).
* `ha_direct` - Enable/disable direct management of HA cluster members.
* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.

The `hosts6` block supports:

* `id` - Host6 entry ID.
* `source_ipv6` - Source IPv6 address for SNMP traps.
* `ipv6` - SNMP manager IPv6 address prefix.
* `ha_direct` - Enable/disable direct management of HA cluster members.
* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SystemSnmp Community can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemsnmp_community.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
