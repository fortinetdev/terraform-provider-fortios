---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_addrgrp6"
description: |-
  Configure IPv6 address groups.
---

# fortios_firewall_addrgrp6
Configure IPv6 address groups.

## Example Usage

```hcl
resource "fortios_firewall_address6" "trname1" {
  cache_ttl  = 0
  color      = 0
  end_ip     = "::"
  host       = ""
  host_type  = "any"
  ip6        = "fdff:ffff::/120"
  name       = "sewec1"
  start_ip   = ""
  type       = "ipprefix"
  visibility = "enable"
}

resource "fortios_firewall_addrgrp6" "trname" {
  color      = 0
  name       = "grp61"
  visibility = "enable"

  member {
    name = fortios_firewall_address6.trname1.name
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - IPv6 address group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address group6 visibility in the GUI. Valid values: `enable`, `disable`.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `member` - (Required) Address objects contained within the group. The structure of `member` block is documented below.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Address6/addrgrp6 name.

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

Firewall Addrgrp6 can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_addrgrp6.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_addrgrp6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
