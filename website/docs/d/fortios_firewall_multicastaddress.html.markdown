---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastaddress"
description: |-
  Get information on an fortios firewall multicastaddress.
---

# Data Source: fortios_firewall_multicastaddress
Use this data source to get information on an fortios firewall multicastaddress

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall multicastaddress.

## Attribute Reference

The following attributes are exported:

* `name` - Multicast address name.
* `type` - Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address.
* `subnet` - Broadcast address and subnet.
* `start_ip` - First IPv4 address (inclusive) in the range for the address.
* `end_ip` - Final IPv4 address (inclusive) in the range for the address.
* `comment` - Comment.
* `visibility` - Enable/disable visibility of the multicast address on the GUI.
* `associated_interface` - Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

