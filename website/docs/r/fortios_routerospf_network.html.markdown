---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_routerospf_network"
description: |-
  OSPF network configuration.
---

# fortios_routerospf_network
OSPF network configuration.

~> The provider supports the definition of Network in Router Ospf `fortios_router_ospf`, and also allows the definition of separate Network resources `fortios_routerospf_network`, but do not use a `fortios_router_ospf` with in-line Network in conjunction with any `fortios_routerospf_network` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `fosid` - Network entry ID.
* `prefix` - Prefix.
* `area` - Attach the network to area.
* `comments` - Comment.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Routerospf Network can be imported using any of these accepted formats:
```
$ terraform import fortios_routerospf_network.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_routerospf_network.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
