---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_routerbgp_network"
description: |-
  BGP network table.
---

# fortios_routerbgp_network
BGP network table.

~> The provider supports the definition of Network in Router Bgp `fortios_router_bgp`, and also allows the definition of separate Network resources `fortios_routerbgp_network`, but do not use a `fortios_router_bgp` with in-line Network in conjunction with any `fortios_routerbgp_network` resources, otherwise conflicts and overwrite will occur.


## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `prefix` - (Required) Prefix of network.
* `backdoor` - Enable/disable route as backdoor.
* `route_map` - Route map to modify generated route.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Routerbgp Network can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_routerbgp_network.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
