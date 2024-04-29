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
* `direction` - How this service may be used (source, destination or both). Valid values: `source`, `destination`, `both`.
* `member` - Internet Service group member. The structure of `member` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Internet Service name.
* `id` - Internet Service ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall InternetServiceGroup can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_internetservicegroup.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_internetservicegroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
