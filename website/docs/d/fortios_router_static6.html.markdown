---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_static6"
description: |-
  Get information on an fortios router static6.
---

# Data Source: fortios_router_static6
Use this data source to get information on an fortios router static6

## Argument Reference

* `seq_num` - (Required) Specify the seq_num of the desired router static6.

## Attribute Reference

The following attributes are exported:

* `seq_num` - Sequence number.
* `status` - Enable/disable this static route.
* `dst` - Destination IPv6 prefix.
* `gateway` - IPv6 address of the gateway.
* `device` - Gateway out interface or tunnel.
* `devindex` - Device index (0 - 4294967295).
* `distance` - Administrative distance (1 - 255).
* `priority` - Administrative priority (0 - 4294967295).
* `comment` - Optional comments.
* `blackhole` - Enable/disable black hole.
* `sdwan` - Enable/disable egress through the SD-WAN.
* `virtual_wan_link` - Enable/disable egress through the virtual-wan-link.
* `link_monitor_exempt` - Enable/disable withdrawal of this static route when link monitor or health check is down.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).

