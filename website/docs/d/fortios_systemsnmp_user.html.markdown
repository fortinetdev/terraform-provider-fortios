---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_user"
description: |-
  Get information on an fortios systemsnmp user.
---

# Data Source: fortios_systemsnmp_user
Use this data source to get information on an fortios systemsnmp user

## Argument Reference

* `name` - (Required) Specify the name of the desired systemsnmp user.

## Attribute Reference

The following attributes are exported:

* `name` - SNMP user name.
* `status` - Enable/disable this SNMP user.
* `trap_status` - Enable/disable traps for this SNMP user.
* `trap_lport` - SNMPv3 local trap port (default = 162).
* `trap_rport` - SNMPv3 trap remote port (default = 162).
* `queries` - Enable/disable SNMP queries for this user.
* `query_port` - SNMPv3 query port (default = 161).
* `notify_hosts` - SNMP managers to send notifications (traps) to.
* `notify_hosts6` - IPv6 SNMP managers to send notifications (traps) to.
* `source_ip` - Source IP for SNMP trap.
* `source_ipv6` - Source IPv6 for SNMP trap.
* `ha_direct` - Enable/disable direct management of HA cluster members.
* `events` - SNMP notifications (traps) to send.
* `security_level` - Security level for message authentication and encryption.
* `auth_proto` - Authentication protocol.
* `auth_pwd` - Password for authentication protocol.
* `priv_proto` - Privacy (encryption) protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.

