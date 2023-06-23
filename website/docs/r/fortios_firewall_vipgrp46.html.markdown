---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vipgrp46"
description: |-
  Configure IPv4 to IPv6 virtual IP groups.
---

# fortios_firewall_vipgrp46
Configure IPv4 to IPv6 virtual IP groups. Applies to FortiOS Version `<= 7.0.0`.

## Example Usage

```hcl
resource "fortios_firewall_vip46" "trname1" {
  arp_reply   = "enable"
  color       = 0
  extip       = "10.202.1.100"
  extport     = "0-65535"
  fosid       = 0
  ldb_method  = "static"
  mappedip    = "2001:1:1:2::100"
  mappedport  = "0-65535"
  name        = "vip46s2"
  portforward = "disable"
  protocol    = "tcp"
  type        = "static-nat"
}

resource "fortios_firewall_vipgrp46" "trname" {
  color = 0
  name  = "vipgrp46"

  member {
    name = fortios_firewall_vip46.trname1.name
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - VIP46 group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `member` - (Required) Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - VIP46 name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vipgrp46 can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_vipgrp46.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_vipgrp46.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
