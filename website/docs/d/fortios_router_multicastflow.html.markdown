---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicastflow"
description: |-
  Get information on an fortios router multicastflow.
---

# Data Source: fortios_router_multicastflow
Use this data source to get information on an fortios router multicastflow

## Argument Reference

* `name` - (Required) Specify the name of the desired router multicastflow.

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comments` - Comment.
* `flows` - Multicast-flow entries. The structure of `flows` block is documented below.

The `flows` block contains:

* `id` - Flow ID.
* `group_addr` - Multicast group IP address.
* `source_addr` - Multicast source IP address.

