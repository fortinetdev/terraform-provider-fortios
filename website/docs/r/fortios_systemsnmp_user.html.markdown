---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_user"
description: |-
  SNMP user configuration.
---

# fortios_systemsnmp_user
SNMP user configuration.

## Example Usage

```hcl
resource "fortios_systemsnmp_user" "trname" {
  auth_proto     = "sha"
  events         = "cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high"
  ha_direct      = "disable"
  name           = "1"
  priv_proto     = "aes"
  queries        = "disable"
  query_port     = 161
  security_level = "no-auth-no-priv"
  source_ip      = "0.0.0.0"
  source_ipv6    = "::"
  status         = "disable"
  trap_lport     = 162
  trap_rport     = 162
  trap_status    = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) SNMP user name.
* `status` - Enable/disable this SNMP user. Valid values: `enable`, `disable`.
* `trap_status` - Enable/disable traps for this SNMP user. Valid values: `enable`, `disable`.
* `trap_lport` - SNMPv3 local trap port (default = 162).
* `trap_rport` - SNMPv3 trap remote port (default = 162).
* `queries` - Enable/disable SNMP queries for this user. Valid values: `enable`, `disable`.
* `query_port` - SNMPv3 query port (default = 161).
* `notify_hosts` - SNMP managers to send notifications (traps) to.
* `notify_hosts6` - IPv6 SNMP managers to send notifications (traps) to.
* `source_ip` - Source IP for SNMP trap.
* `source_ipv6` - Source IPv6 for SNMP trap.
* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `enable`, `disable`.
* `events` - SNMP notifications (traps) to send.
* `mib_view` - SNMP access control MIB view.
* `vdoms` - SNMP access control VDOMs. The structure of `vdoms` block is documented below.
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
* `auth_proto` - Authentication protocol.
* `auth_pwd` - Password for authentication protocol.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.
* `priv_pwd` - Password for privacy (encryption) protocol.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `vdoms` block supports:

* `name` - VDOM name


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemSnmp User can be imported using any of these accepted formats:
```
$ terraform import fortios_systemsnmp_user.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_systemsnmp_user.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
