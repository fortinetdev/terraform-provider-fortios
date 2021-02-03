---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_static"
description: |-
  Configure IPv4 static routing tables.
---

# fortios_router_static
Configure IPv4 static routing tables.

## Example Usage

```hcl
resource "fortios_router_static" "trname" {
  bfd                 = "disable"
  blackhole           = "disable"
  device              = "port4"
  distance            = 10
  dst                 = "1.0.0.0 255.240.0.0"
  dynamic_gateway     = "disable"
  gateway             = "0.0.0.0"
  internet_service    = 0
  link_monitor_exempt = "disable"
  priority            = 22
  seq_num             = 1
  src                 = "0.0.0.0 0.0.0.0"
  status              = "enable"
  virtual_wan_link    = "disable"
  vrf                 = 0
  weight              = 2
}
```

## Argument Reference

The following arguments are supported:

* `seq_num` - Sequence number.
* `status` - Enable/disable this static route.
* `dst` - (Required) Destination IP and mask for this route.
* `src` - Source prefix for this route.
* `gateway` - Gateway IP for this route.
* `distance` - Administrative distance (1 - 255).
* `weight` - Administrative weight (0 - 255).
* `priority` - Administrative priority (0 - 4294967295).
* `device` - Gateway out interface or tunnel.
* `comment` - Optional comments.
* `blackhole` - Enable/disable black hole.
* `dynamic_gateway` - Enable use of dynamic gateway retrieved from a DHCP or PPP server.
* `sdwan` - Enable/disable egress through SD-WAN.
* `virtual_wan_link` - Enable/disable egress through the virtual-wan-link.
* `dstaddr` - Name of firewall address or address group.
* `internet_service` - Application ID in the Internet service database.
* `internet_service_custom` - Application name in the Internet service custom database.
* `link_monitor_exempt` - Enable/disable withdrawing this route when link monitor or health check is down.
* `vrf` - Virtual Routing Forwarding ID.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Static can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_static.labelname {{seq_num}}
$ unset "FORTIOS_IMPORT_TABLE"
```
