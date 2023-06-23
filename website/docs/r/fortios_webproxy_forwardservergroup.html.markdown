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
* `affinity` - Enable/disable affinity, attaching a source-ip's traffic to the assigned forwarding server until the forward-server-affinity-timeout is reached (under web-proxy global). Valid values: `enable`, `disable`.
* `ldb_method` - Load balance method: weighted or least-session.
* `group_down_option` - Action to take when all of the servers in the forward server group are down: block sessions until at least one server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
* `server_list` - Add web forward servers to a list to form a server group. Optionally assign weights to each server. The structure of `server_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `server_list` block supports:

* `name` - Forward server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy ForwardServerGroup can be imported using any of these accepted formats:
```
$ terraform import fortios_webproxy_forwardservergroup.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webproxy_forwardservergroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
