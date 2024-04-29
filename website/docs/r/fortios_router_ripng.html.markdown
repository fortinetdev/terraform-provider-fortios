---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ripng"
description: |-
  Configure RIPng.
---

# fortios_router_ripng
Configure RIPng.

## Example Usage

```hcl
resource "fortios_router_ripng" "trname" {
  default_information_originate = "disable"
  default_metric                = 1
  garbage_timer                 = 120
  max_out_metric                = 0
  timeout_timer                 = 180
  update_timer                  = 30

  redistribute {
    metric = 10
    name   = "connected"
    status = "disable"
  }
  redistribute {
    metric = 10
    name   = "static"
    status = "disable"
  }
  redistribute {
    metric = 10
    name   = "ospf"
    status = "disable"
  }
  redistribute {
    metric = 10
    name   = "bgp"
    status = "disable"
  }
  redistribute {
    metric = 10
    name   = "isis"
    status = "disable"
  }
}
```

## Argument Reference

The following arguments are supported:

* `default_information_originate` - Enable/disable generation of default route. Valid values: `enable`, `disable`.
* `default_metric` - Default metric.
* `max_out_metric` - Maximum metric allowed to output(0 means 'not set').
* `distance` - distance The structure of `distance` block is documented below.
* `distribute_list` - Distribute list. The structure of `distribute_list` block is documented below.
* `neighbor` - neighbor The structure of `neighbor` block is documented below.
* `network` - Network. The structure of `network` block is documented below.
* `aggregate_address` - Aggregate address. The structure of `aggregate_address` block is documented below.
* `offset_list` - Offset list. The structure of `offset_list` block is documented below.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `update_timer` - Update timer.
* `timeout_timer` - Timeout timer.
* `garbage_timer` - Garbage timer.
* `interface` - RIPng interface configuration. The structure of `interface` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `distance` block supports:

* `id` - Distance ID.
* `distance` - Distance (1 - 255).
* `prefix6` - Distance prefix6.
* `access_list6` - Access list for route destination.

The `distribute_list` block supports:

* `id` - Distribute list ID.
* `status` - status Valid values: `enable`, `disable`.
* `direction` - Distribute list direction. Valid values: `in`, `out`.
* `listname` - Distribute access/prefix list name.
* `interface` - Distribute list interface name.

The `neighbor` block supports:

* `id` - Neighbor entry ID.
* `ip6` - IPv6 link-local address.
* `interface` - Interface name.

The `network` block supports:

* `id` - Network entry ID.
* `prefix` - Network IPv6 link-local prefix.

The `aggregate_address` block supports:

* `id` - Aggregate address entry ID.
* `prefix6` - Aggregate address prefix.

The `offset_list` block supports:

* `id` - Offset-list ID.
* `status` - status Valid values: `enable`, `disable`.
* `direction` - Offset list direction. Valid values: `in`, `out`.
* `access_list6` - IPv6 access list name.
* `offset` - offset
* `interface` - Interface name.

The `passive_interface` block supports:

* `name` - Passive interface name.

The `redistribute` block supports:

* `name` - Redistribute name.
* `status` - status Valid values: `enable`, `disable`.
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.

The `interface` block supports:

* `name` - Interface name.
* `split_horizon_status` - Enable/disable split horizon. Valid values: `enable`, `disable`.
* `split_horizon` - Enable/disable split horizon. Valid values: `poisoned`, `regular`.
* `flags` - Flags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ripng can be imported using any of these accepted formats:
```
$ terraform import fortios_router_ripng.labelname RouterRipng

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_ripng.labelname RouterRipng
$ unset "FORTIOS_IMPORT_TABLE"
```
