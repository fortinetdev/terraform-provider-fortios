---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip"
description: |-
  Configure RIP.
---

# fortios_router_rip
Configure RIP.

## Example Usage

```hcl
resource "fortios_router_rip" "trname" {
  default_information_originate = "disable"
  default_metric                = 1
  garbage_timer                 = 120
  max_out_metric                = 0
  recv_buffer_size              = 655360
  timeout_timer                 = 180
  update_timer                  = 30
  version                       = "2"

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

* `default_information_originate` - Enable/disable generation of default route.
* `default_metric` - Default metric.
* `max_out_metric` - Maximum metric allowed to output(0 means 'not set').
* `recv_buffer_size` - Receiving buffer size.
* `distance` - distance
* `distribute_list` - Distribute list.
* `neighbor` - neighbor
* `network` - network
* `offset_list` - Offset list.
* `passive_interface` - Passive interface configuration.
* `redistribute` - Redistribute configuration.
* `update_timer` - Update timer in seconds.
* `timeout_timer` - Timeout timer in seconds.
* `garbage_timer` - Garbage timer in seconds.
* `version` - RIP version.
* `interface` - RIP interface configuration.

The `distance` block supports:

* `id` - Distance ID.
* `prefix` - Distance prefix.
* `distance` - Distance (1 - 255).
* `access_list` - Access list for route destination.

The `distribute_list` block supports:

* `id` - Distribute list ID.
* `status` - status
* `direction` - Distribute list direction.
* `listname` - Distribute access/prefix list name.
* `interface` - Distribute list interface name.

The `neighbor` block supports:

* `id` - Neighbor entry ID.
* `ip` - IP address.

The `network` block supports:

* `id` - Network entry ID.
* `prefix` - Network prefix.

The `offset_list` block supports:

* `id` - Offset-list ID.
* `status` - status
* `direction` - Offset list direction.
* `access_list` - Access list name.
* `offset` - offset
* `interface` - Interface name.

The `passive_interface` block supports:

* `name` - Passive interface name.

The `redistribute` block supports:

* `name` - Redistribute name.
* `status` - status
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.

The `interface` block supports:

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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Rip can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_rip.labelname RouterRip
$ unset "FORTIOS_IMPORT_TABLE"
```
