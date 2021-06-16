---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_routerospf_neighbor"
description: |-
  OSPF neighbor configuration are used when OSPF runs on non-broadcast media
---

# fortios_routerospf_neighbor
OSPF neighbor configuration are used when OSPF runs on non-broadcast media

~> The provider supports the definition of Neighbor in Router Ospf `fortios_router_ospf`, and also allows the definition of separate Neighbor resources `fortios_routerospf_neighbor`, but do not use a `fortios_router_ospf` with in-line Neighbor in conjunction with any `fortios_routerospf_neighbor` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `fosid` - Neighbor entry ID.
* `ip` - Interface IP address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - Priority.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Routerospf Neighbor can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_routerospf_neighbor.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
