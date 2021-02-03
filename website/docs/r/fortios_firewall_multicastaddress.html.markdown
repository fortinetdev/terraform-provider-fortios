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
* `type` - Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address.
* `subnet` - Broadcast address and subnet.
* `start_ip` - (Required) First IPv4 address (inclusive) in the range for the address.
* `end_ip` - (Required) Final IPv4 address (inclusive) in the range for the address.
* `comment` - Comment.
* `visibility` - Enable/disable visibility of the multicast address on the GUI.
* `associated_interface` - Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_multicastaddress.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
