---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_decryptedtrafficmirror"
description: |-
  Configure decrypted traffic mirror.
---

# fortios_firewall_decryptedtrafficmirror
Configure decrypted traffic mirror. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `dstmac` - Set destination MAC address for mirrored traffic.
* `traffic_type` - Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
* `traffic_source` - Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
* `interface` - Decrypted traffic mirror interface The structure of `interface` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `interface` block supports:

* `name` - Decrypted traffic mirror interface.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall DecryptedTrafficMirror can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_decryptedtrafficmirror.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_decryptedtrafficmirror.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
