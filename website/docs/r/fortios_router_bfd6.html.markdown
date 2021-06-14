---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bfd6"
description: |-
  Configure IPv6 BFD.
---

# fortios_router_bfd6
Configure IPv6 BFD.

## Argument Reference

The following arguments are supported:

* `neighbor` - Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `neighbor` block supports:

* `ip6_address` - IPv6 address of the BFD neighbor.
* `interface` - Interface to the BFD neighbor.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Bfd6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_bfd6.labelname RouterBfd6
$ unset "FORTIOS_IMPORT_TABLE"
```
