---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_peergrp"
description: |-
  Configure peer groups.
---

# fortios_user_peergrp
Configure peer groups.

## Example Usage

```hcl
resource "fortios_user_peer" "trname2" {
  ca                  = "EC-ACC"
  cn_type             = "string"
  ldap_mode           = "password"
  mandatory_ca_verify = "enable"
  name                = "u2"
  two_factor          = "disable"
}

resource "fortios_user_peergrp" "trname" {
  name = "g1"

  member {
    name = fortios_user_peer.trname2.name
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Peer group name.
* `member` - Peer group members. The structure of `member` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Peer group member name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Peergrp can be imported using any of these accepted formats:
```
$ terraform import fortios_user_peergrp.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_peergrp.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
