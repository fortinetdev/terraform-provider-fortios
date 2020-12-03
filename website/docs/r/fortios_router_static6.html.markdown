---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_static6"
description: |-
  Configure IPv6 static routing tables.
---

# fortios_router_static6
Configure IPv6 static routing tables.

## Example Usage

```hcl
resource "fortios_router_static6" "trname" {
  bfd              = "disable"
  blackhole        = "disable"
  device           = "port3"
  devindex         = 5
  distance         = 10
  dst              = "2001:db8::/32"
  gateway          = "::"
  priority         = 32
  seq_num          = 1
  status           = "enable"
  virtual_wan_link = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `seq_num` - (Required) Sequence number.
* `status` - Enable/disable this static route.
* `dst` - (Required) Destination IPv6 prefix.
* `gateway` - IPv6 address of the gateway.
* `device` - (Required) Gateway out interface or tunnel.
* `devindex` - Device index (0 - 4294967295).
* `distance` - Administrative distance (1 - 255).
* `priority` - Administrative priority (0 - 4294967295).
* `comment` - Optional comments.
* `blackhole` - Enable/disable black hole.
* `virtual_wan_link` - Enable/disable egress through the virtual-wan-link.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Static6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_static6.labelname {{seq_num}}
$ unset "FORTIOS_IMPORT_TABLE"
```
