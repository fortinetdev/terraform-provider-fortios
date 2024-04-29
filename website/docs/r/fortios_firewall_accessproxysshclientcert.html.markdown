---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_accessproxysshclientcert"
description: |-
  Configure Access Proxy SSH client certificate.
---

# fortios_firewall_accessproxysshclientcert
Configure Access Proxy SSH client certificate. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - SSH client certificate name.
* `source_address` - Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address. Valid values: `enable`, `disable`.
* `permit_x11_forwarding` - Enable/disable appending permit-x11-forwarding certificate extension. Valid values: `enable`, `disable`.
* `permit_agent_forwarding` - Enable/disable appending permit-agent-forwarding certificate extension. Valid values: `enable`, `disable`.
* `permit_port_forwarding` - Enable/disable appending permit-port-forwarding certificate extension. Valid values: `enable`, `disable`.
* `permit_pty` - Enable/disable appending permit-pty certificate extension. Valid values: `enable`, `disable`.
* `permit_user_rc` - Enable/disable appending permit-user-rc certificate extension. Valid values: `enable`, `disable`.
* `cert_extension` - Configure certificate extension for user certificate. The structure of `cert_extension` block is documented below.
* `auth_ca` - Name of the SSH server public key authentication CA.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `cert_extension` block supports:

* `name` - Name of certificate extension.
* `critical` - Critical option. Valid values: `no`, `yes`.
* `type` - Type of certificate extension. Valid values: `fixed`, `user`.
* `data` - Data of certificate extension.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall AccessProxySshClientCert can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_accessproxysshclientcert.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_accessproxysshclientcert.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
