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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `neighbor` block supports:

* `ip` - IPv4 address of the BFD neighbor.
* `interface` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Bfd can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_bfd.labelname RouterBfd
$ unset "FORTIOS_IMPORT_TABLE"
```
