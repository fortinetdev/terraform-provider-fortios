---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ripng"
description: |-
  Get information on fortios router ripng.
---

# Data Source: fortios_router_ripng
Use this data source to get information on fortios router ripng

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `default_information_originate` - Enable/disable generation of default route.
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

The `distance` block contains:

* `id` - Distance ID.
* `distance` - Distance (1 - 255).
* `prefix6` - Distance prefix6.
* `access_list6` - Access list for route destination.

The `distribute_list` block contains:

* `id` - Distribute list ID.
* `status` - status
* `direction` - Distribute list direction.
* `listname` - Distribute access/prefix list name.
* `interface` - Distribute list interface name.

The `neighbor` block contains:

* `id` - Neighbor entry ID.
* `ip6` - IPv6 link-local address.
* `interface` - Interface name.

The `network` block contains:

* `id` - Network entry ID.
* `prefix` - Network IPv6 link-local prefix.

The `aggregate_address` block contains:

* `id` - Aggregate address entry ID.
* `prefix6` - Aggregate address prefix.

The `offset_list` block contains:

* `id` - Offset-list ID.
* `status` - status
* `direction` - Offset list direction.
* `access_list6` - IPv6 access list name.
* `offset` - offset
* `interface` - Interface name.

The `passive_interface` block contains:

* `name` - Passive interface name.

The `redistribute` block contains:

* `name` - Redistribute name.
* `status` - status
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.

The `interface` block contains:

* `name` - Interface name.
* `split_horizon_status` - Enable/disable split horizon.
* `split_horizon` - Enable/disable split horizon.
* `flags` - Flags.

