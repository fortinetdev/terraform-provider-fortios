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
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `seq_num` - Sequence number.
* `status` - Enable/disable this static route.
* `dst` - Destination IPv6 prefix.
* `gateway` - IPv6 address of the gateway.
* `device` - Gateway out interface or tunnel.
* `devindex` - Device index (0 - 4294967295).
* `distance` - Administrative distance (1 - 255).
* `weight` - Administrative weight (0 - 255).
* `priority` - Administrative priority (0 - 4294967295).
* `comment` - Optional comments.
* `blackhole` - Enable/disable black hole.
* `dynamic_gateway` - Enable use of dynamic gateway retrieved from Router Advertisement (RA).
* `sdwan_zone` - Choose SD-WAN Zone. The structure of `sdwan_zone` block is documented below.
* `dstaddr` - Name of firewall address or address group.
* `sdwan` - Enable/disable egress through the SD-WAN.
* `virtual_wan_link` - Enable/disable egress through the virtual-wan-link.
* `link_monitor_exempt` - Enable/disable withdrawal of this static route when link monitor or health check is down.
* `vrf` - Virtual Routing Forwarding ID.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).

The `sdwan_zone` block contains:

* `name` - SD-WAN zone name.

