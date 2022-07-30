---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_addrgrp"
description: |-
  Configure IPv4 address groups.
---

# fortios_firewall_addrgrp
Configure IPv4 address groups.

## Example Usage

```hcl
resource "fortios_firewall_address" "trname1" {
  allow_routing = "disable"
  cache_ttl     = 0
  color         = 0
  end_ip        = "255.0.0.0"
  name          = "1"
  start_ip      = "12.0.0.0"
  subnet        = "12.0.0.0 255.0.0.0"
  type          = "ipmask"
  visibility    = "enable"
  wildcard      = "12.0.0.0 255.0.0.0"
}

resource "fortios_firewall_addrgrp" "trname" {
  allow_routing = "disable"
  color         = 0
  exclude       = "disable"
  name          = "group1"
  visibility    = "enable"

  member {
    name = fortios_firewall_address.trname1.name
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Address group name.
* `type` - Address group type. Valid values: `default`, `folder`.
* `category` - Address group category. Valid values: `default`, `ztna-ems-tag`, `ztna-geo-tag`.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `member` - (Required) Address objects contained within the group. The structure of `member` block is documented below.
* `comment` - Comment.
* `exclude` - Enable/disable address exclusion. Valid values: `enable`, `disable`.
* `exclude_member` - Address exclusion member. The structure of `exclude_member` block is documented below.
* `visibility` - Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
* `color` - Color of icon on the GUI.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `allow_routing` - Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Address name.

The `exclude_member` block supports:

* `name` - Address name.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block supports:

* `name` - Tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Addrgrp can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_addrgrp.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_addrgrp.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
