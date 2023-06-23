---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_acme"
description: |-
  Configure ACME client.
---

# fortios_system_acme
Configure ACME client. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `interface` - Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
* `use_ha_direct` - Enable the use of 'ha-mgmt' interface to connect to the ACME server when 'ha-direct' is enabled in HA configuration Valid values: `enable`, `disable`.
* `source_ip` - Source IPv4 address used to connect to the ACME server.
* `source_ip6` - Source IPv6 address used to connect to the ACME server.
* `accounts` - ACME accounts list. The structure of `accounts` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `interface` block supports:

* `interface_name` - Interface name.

The `accounts` block supports:

* `id` - Account id.
* `status` - Account status.
* `url` - Account url.
* `ca_url` - Account ca_url.
* `email` - Account email.
* `privatekey` - Account Private Key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Acme can be imported using any of these accepted formats:
```
$ terraform import fortios_system_acme.labelname SystemAcme

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_acme.labelname SystemAcme
$ unset "FORTIOS_IMPORT_TABLE"
```
