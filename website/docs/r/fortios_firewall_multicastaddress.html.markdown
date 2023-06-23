---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastaddress"
description: |-
  Configure multicast addresses.
---

# fortios_firewall_multicastaddress
Configure multicast addresses.

## Example Usage

```hcl
resource "fortios_firewall_multicastaddress" "trname" {
  color      = 0
  end_ip     = "224.0.0.22"
  name       = "1"
  start_ip   = "224.0.0.11"
  subnet     = "224.0.0.11 224.0.0.22"
  type       = "multicastrange"
  visibility = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Multicast address name.
* `type` - Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange`, `broadcastmask`.
* `subnet` - Broadcast address and subnet.
* `start_ip` - First IPv4 address (inclusive) in the range for the address.
* `end_ip` - Final IPv4 address (inclusive) in the range for the address.
* `comment` - Comment.
* `visibility` - Enable/disable visibility of the multicast address on the GUI. Valid values: `enable`, `disable`.
* `associated_interface` - Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
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

Firewall MulticastAddress can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_multicastaddress.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_multicastaddress.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
