---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_qkd"
description: |-
  Configure Quantum Key Distribution servers
---

# fortios_vpn_qkd
Configure Quantum Key Distribution servers Applies to FortiOS Version `>= 7.4.2`.

## Argument Reference

The following arguments are supported:

* `name` - Quantum Key Distribution configuration name.
* `server` - IPv4, IPv6 or DNS address of the KME.
* `port` - Port to connect to on the KME.
* `fosid` - Quantum Key Distribution ID assigned by the KME.
* `peer` - Authenticate Quantum Key Device's certificate with the peer/peergrp.
* `certificate` - Names of up to 4 certificates to offer to the KME. The structure of `certificate` block is documented below.
* `comment` - Comment.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `certificate` block supports:

* `name` - Certificate name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn Qkd can be imported using any of these accepted formats:
```
$ terraform import fortios_vpn_qkd.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpn_qkd.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
