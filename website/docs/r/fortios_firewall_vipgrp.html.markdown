---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vipgrp"
description: |-
  Configure IPv4 virtual IP groups.
---

# fortios_firewall_vipgrp
Configure IPv4 virtual IP groups.

## Example Usage

```hcl
resource "fortios_firewall_vip" "trname1" {
  extintf = "any"
  extport = "0-65535"
  extip   = "2.0.0.1-2.0.0.4"
  name    = "vips2"
  mappedip {
    range = "3.0.0.0-3.0.0.3"
  }
}

resource "fortios_firewall_vipgrp" "trname" {
  color     = 0
  interface = "any"
  name      = "vipgrp1"

  member {
    name = fortios_firewall_vip.trname1.name
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - VIP group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `interface` - (Required) interface
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `member` - (Required) Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.

The `member` block supports:

* `name` - VIP name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vipgrp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_vipgrp.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
