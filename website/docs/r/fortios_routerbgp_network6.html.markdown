---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_routerbgp_network6"
description: |-
  BGP IPv6 network table.
---

# fortios_routerbgp_network6
BGP IPv6 network table.

~> The provider supports the definition of Network6 in Router Bgp `fortios_router_bgp`, and also allows the definition of separate Network6 resources `fortios_routerbgp_network6`, but do not use a `fortios_router_bgp` with in-line Network6 in conjunction with any `fortios_routerbgp_network6` resources, otherwise conflicts and overwrite will occur.



## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `prefix6` - Network IPv6 prefix.
* `backdoor` - Enable/disable route as backdoor. Valid values: `enable`, `disable`.
* `route_map` - Route map to modify generated route.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Routerbgp Network6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_routerbgp_network6.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
