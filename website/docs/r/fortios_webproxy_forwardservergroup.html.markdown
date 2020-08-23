---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_forwardservergroup"
description: |-
  Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.
---

# fortios_webproxy_forwardservergroup
Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.

## Example Usage

```hcl
resource "fortios_webproxy_forwardserver" "trname1" {
  addr_type          = "fqdn"
  healthcheck        = "disable"
  ip                 = "0.0.0.0"
  monitor            = "http://www.google.com"
  name               = "test1"
  port               = 1128
  server_down_option = "block"
}

resource "fortios_webproxy_forwardservergroup" "trname1" {
  affinity          = "disable"
  group_down_option = "block"
  ldb_method        = "weighted"
  name              = "sg1"

  server_list {
    name   = fortios_webproxy_forwardserver.trname1.name
    weight = 12
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Configure a forward server group consisting one or multiple forward servers. Supports failover and load balancing.
* `affinity` - Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global).
* `ldb_method` - Load balance method: weighted or least-session.
* `group_down_option` - Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination.
* `server_list` - Add web forward servers to a list to form a server group. Optionally assign weights to each server.

The `server_list` block supports:

* `name` - Forward server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy ForwardServerGroup can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webproxy_forwardservergroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
