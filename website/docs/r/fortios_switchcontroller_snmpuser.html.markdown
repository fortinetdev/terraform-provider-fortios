---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_snmpuser"
description: |-
  Configure FortiSwitch SNMP v3 users globally.
---

# fortios_switchcontroller_snmpuser
Configure FortiSwitch SNMP v3 users globally. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `name` - SNMP user name.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.
* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
* `auth_proto` - Authentication protocol.
* `auth_pwd` - Password for authentication protocol.
* `priv_proto` - Privacy (encryption) protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SnmpUser can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_snmpuser.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_snmpuser.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
