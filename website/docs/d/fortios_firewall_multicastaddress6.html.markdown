---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastaddress6"
description: |-
  Get information on an fortios firewall multicastaddress6.
---

# Data Source: fortios_firewall_multicastaddress6
Use this data source to get information on an fortios firewall multicastaddress6

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall multicastaddress6.

## Attribute Reference

The following attributes are exported:

* `name` - IPv6 multicast address name.
* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `comment` - Comment.
* `visibility` - Enable/disable visibility of the IPv6 multicast address on the GUI.
* `color` - Color of icon on the GUI.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

