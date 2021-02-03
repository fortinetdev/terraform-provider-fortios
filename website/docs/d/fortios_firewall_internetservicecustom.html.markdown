---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicecustom"
description: |-
  Get information on an fortios firewall internetservicecustom.
---

# Data Source: fortios_firewall_internetservicecustom
Use this data source to get information on an fortios firewall internetservicecustom

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall internetservicecustom.

## Attribute Reference

The following attributes are exported:

* `name` - Internet Service name.
* `reputation` - Reputation level of the custom Internet Service.
* `comment` - Comment.
* `entry` - Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.

The `entry` block contains:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the custom entry. The structure of `port_range` block is documented below.
* `dst` - Destination address or address group name. The structure of `dst` block is documented below.

The `port_range` block contains:

* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).

The `dst` block contains:

* `name` - Select the destination address or address group object from available options.

