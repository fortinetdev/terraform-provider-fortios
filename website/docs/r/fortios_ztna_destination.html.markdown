---
subcategory: "FortiGate Ztna"
layout: "fortios"
page_title: "FortiOS: fortios_ztna_destination"
description: |-
  Configure ZTNA destination.
---

# fortios_ztna_destination
Configure ZTNA destination. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - Destination name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `type` - ZTNA destination type. Valid values: `on-premise`, `saas`.
* `address` - Address or address group of the ZTNA destination.
* `mappedport` - Port for communicating with the real server.
* `protocol` - Protocol type based on IANA numbers. Valid values: `TCP`, `UDP`, `ALL`.
* `saas_application` - SaaS application controlled by this ZTNA destination. The structure of `saas_application` block is documented below.
* `conn_type` - Connection type. Valid values: `traffic-forwarding`, `ssh`.
* `ssh_client_cert` - Configure access-proxy SSH client certificate profile.
* `ssh_host_key_validation` - Enable/disable SSH host key validation. Valid values: `disable`, `enable`.
* `ssh_host_key` - Configure host keys (one or more may be configured). The structure of `ssh_host_key` block is documented below.
* `external_auth` - Enable/disable use of external browser as user-agent for SAML user authentication. Valid values: `enable`, `disable`.
* `tunnel_encryption` - Tunnel encryption. Valid values: `enable`, `disable`.
* `domain` - Wildcard domain name of the ZTNA destination.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `saas_application` block supports:

* `name` - SaaS application name.

The `ssh_host_key` block supports:

* `name` - Host key name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna Destination can be imported using any of these accepted formats:
```
$ terraform import fortios_ztna_destination.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ztna_destination.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
