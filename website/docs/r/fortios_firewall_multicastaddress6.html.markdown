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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_multicastaddress6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
