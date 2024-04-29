---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastaddress6"
description: |-
  Configure IPv6 multicast address.
---

# fortios_firewall_multicastaddress6
Configure IPv6 multicast address.

## Example Usage

```hcl
resource "fortios_firewall_multicastaddress6" "trname" {
  color      = 0
  ip6        = "ff02::1:ff0e:8c6c/128"
  name       = "1"
  visibility = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IPv6 multicast address name.
* `ip6` - (Required) IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `comment` - Comment.
* `visibility` - Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable`, `disable`.
* `color` - Color of icon on the GUI.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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

Firewall MulticastAddress6 can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_multicastaddress6.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_multicastaddress6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
