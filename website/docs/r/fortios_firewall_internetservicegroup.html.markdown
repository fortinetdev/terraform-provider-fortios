---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicegroup"
description: |-
  Configure group of Internet Service.
---

# fortios_firewall_internetservicegroup
Configure group of Internet Service.

## Example Usage

```hcl
resource "fortios_firewall_internetservicegroup" "trname" {
  direction = "both"
  name      = "1"

  member {
    id = 65641
  }
  member {
    id = 65646
  }
  member {
    id = 196747
  }
  member {
    id = 327781
  }
  member {
    id = 327786
  }
  member {
    id = 327791
  }
  member {
    id = 327839
  }
}
```

## Argument Reference


The following arguments are supported:

* `name` - Internet Service group name.
* `comment` - Comment.
* `direction` - How this service may be used (source, destination or both).
* `member` - Internet Service group member. The structure of `member` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `member` block supports:

* `id` - Internet Service ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall InternetServiceGroup can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetservicegroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
