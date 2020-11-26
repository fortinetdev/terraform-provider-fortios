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
* `visibility` - Enable/disable address group6 visibility in the GUI.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `member` - (Required) Address objects contained within the group. The structure of `member` block is documented below.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_addrgrp6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
