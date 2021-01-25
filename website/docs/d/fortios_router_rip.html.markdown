---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip"
description: |-
  Get information on fortios router rip.
---

# Data Source: fortios_router_rip
Use this data source to get information on fortios router rip

## Argument Reference

No arguments available for the data source.

## Attribute Reference

The following attributes are exported:

* `default_information_originate` - Enable/disable generation of default route.
* `default_metric` - Default metric.
* `max_out_metric` - Maximum metric allowed to output(0 means 'not set').
* `recv_buffer_size` - Receiving buffer size.
* `distance` - distance The structure of `distance` block is documented below.
* `distribute_list` - Distribute list. The structure of `distribute_list` block is documented below.
* `neighbor` - neighbor The structure of `neighbor` block is documented below.
* `network` - network The structure of `network` block is documented below.
* `offset_list` - Offset list. The structure of `offset_list` block is documented below.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `update_timer` - Update timer in seconds.
* `timeout_timer` - Timeout timer in seconds.
* `garbage_timer` - Garbage timer in seconds.
* `version` - RIP version.
* `interface` - RIP interface configuration. The structure of `interface` block is documented below.

The `distance` block contains:

* `id` - Distance ID.
* `prefix` - Distance prefix.
* `distance` - Distance (1 - 255).
* `access_list` - Access list for route destination.

The `distribute_list` block contains:

* `id` - Distribute list ID.
* `status` - status
* `direction` - Distribute list direction.
* `listname` - Distribute access/prefix list name.
* `interface` - Distribute list interface name.

The `neighbor` block contains:

* `id` - Neighbor entry ID.
* `ip` - IP address.

The `network` block contains:

* `id` - Network entry ID.
* `prefix` - Network prefix.

The `offset_list` block contains:

* `id` - Offset-list ID.
* `status` - status
* `direction` - Offset list direction.
* `access_list` - Access list name.
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
* `auth_keychain` - Authentication key-chain name.
* `auth_mode` - Authentication mode.
* `auth_string` - Authentication string/password.
* `receive_version` - Receive version.
* `send_version` - Send version.
* `send_version2_broadcast` - Enable/disable broadcast version 1 compatible packets.
* `split_horizon_status` - Enable/disable split horizon.
* `split_horizon` - Enable/disable split horizon.
* `flags` - flags

