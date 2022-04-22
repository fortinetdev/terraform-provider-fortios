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
* `prefix` - (Required) Network prefix.
* `network_import_check` - Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
* `backdoor` - Enable/disable route as backdoor. Valid values: `enable`, `disable`.
* `route_map` - Route map to modify generated route.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Routerbgp Network can be imported using any of these accepted formats:
```
$ terraform import fortios_routerbgp_network.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_routerbgp_network.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
