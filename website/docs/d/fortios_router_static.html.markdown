---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_static"
description: |-
  Get information on an fortios router static.
---

# Data Source: fortios_router_static
Use this data source to get information on an fortios router static

## Example Usage

```hcl
data "fortios_router_static" sample1 {
  seq_num = 1
}

output output1 {
  value = data.fortios_router_static.sample1
}
```

## Argument Reference

* `seq_num` - (Required) Specify the seq_num of the desired router static.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `seq_num` - Sequence number.
* `status` - Enable/disable this static route.
* `dst` - Destination IP and mask for this route.
* `src` - Source prefix for this route.
* `gateway` - Gateway IP for this route.
* `preferred_source` - Preferred source IP for this route.
* `distance` - Administrative distance (1 - 255).
* `weight` - Administrative weight (0 - 255).
* `priority` - Administrative priority (0 - 4294967295).
* `device` - Gateway out interface or tunnel.
* `comment` - Optional comments.
* `blackhole` - Enable/disable black hole.
* `dynamic_gateway` - Enable use of dynamic gateway retrieved from a DHCP or PPP server.
* `sdwan_zone` - Choose SD-WAN Zone. The structure of `sdwan_zone` block is documented below.
* `sdwan` - Enable/disable egress through SD-WAN.
* `virtual_wan_link` - Enable/disable egress through the virtual-wan-link.
* `dstaddr` - Name of firewall address or address group.
* `internet_service` - Application ID in the Internet service database.
* `internet_service_custom` - Application name in the Internet service custom database.
* `link_monitor_exempt` - Enable/disable withdrawing this route when link monitor or health check is down.
* `tag` - Route tag.
* `vrf` - Virtual Routing Forwarding ID.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).

The `sdwan_zone` block contains:

* `name` - SD-WAN zone name.

