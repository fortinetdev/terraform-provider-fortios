---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bfd"
description: |-
  Configure BFD.
---

# fortios_router_bfd
Configure BFD.

## Argument Reference

The following arguments are supported:

* `neighbor` - neighbor The structure of `neighbor` block is documented below.
* `multihop_template` - BFD multi-hop template table. The structure of `multihop_template` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `neighbor` block supports:

* `ip` - IPv4 address of the BFD neighbor.
* `interface` - Interface name.

The `multihop_template` block supports:

* `id` - ID.
* `src` - Source prefix.
* `dst` - Destination prefix.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval (milliseconds).
* `bfd_required_min_rx` - BFD required minimal receive interval (milliseconds).
* `bfd_detect_mult` - BFD detection multiplier.
* `auth_mode` - Authentication mode. Valid values: `none`, `md5`.
* `md5_key` - MD5 key of key ID 1.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Bfd can be imported using any of these accepted formats:
```
$ terraform import fortios_router_bfd.labelname RouterBfd

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_bfd.labelname RouterBfd
$ unset "FORTIOS_IMPORT_TABLE"
```
