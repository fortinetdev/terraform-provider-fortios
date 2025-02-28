---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_linkmonitor"
description: |-
  Configure Link Health Monitor.
---

# fortios_system_linkmonitor
Configure Link Health Monitor.

## Example Usage

```hcl
resource "fortios_system_linkmonitor" "trname" {
  addr_mode                = "ipv4"
  failtime                 = 5
  gateway_ip               = "2.2.2.2"
  gateway_ip6              = "::"
  ha_priority              = 1
  http_agent               = "Chrome/ Safari/"
  http_get                 = "/"
  interval                 = 1
  name                     = "1"
  packet_size              = 64
  port                     = 80
  protocol                 = "ping"
  recoverytime             = 5
  security_mode            = "none"
  source_ip                = "0.0.0.0"
  source_ip6               = "::"
  srcintf                  = "port4"
  status                   = "enable"
  update_cascade_interface = "enable"
  update_static_route      = "enable"
  server {
    address = "3.3.3.3"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Link monitor name.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
* `srcintf` - Interface that receives the traffic to be monitored.
* `server_config` - Mode of server configuration. Valid values: `default`, `individual`.
* `server_type` - Server type (static or dynamic). Valid values: `static`, `dynamic`.
* `server` - (Required) IP address of the server(s) to be monitored. The structure of `server` block is documented below.
* `protocol` - Protocols used to monitor the server.
* `port` - Port number of the traffic to be used to monitor the server.
* `gateway_ip` - Gateway IP address used to probe the server.
* `gateway_ip6` - Gateway IPv6 address used to probe the server.
* `route` - Subnet to monitor. The structure of `route` block is documented below.
* `source_ip` - Source IP address used in packet to the server.
* `source_ip6` - Source IPv6 address used in packet to the server.
* `http_get` - If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_match` - String that you expect to see in the HTTP-GET requests of the traffic to be monitored.
* `interval` - Detection interval. On FortiOS versions 6.2.0: 1 - 3600 sec, default = 5. On FortiOS versions 6.2.4-7.0.10, 7.2.0-7.2.4: 500 - 3600 * 1000 msec, default = 500. On FortiOS versions 7.0.11-7.0.17, >= 7.2.6: 20 - 3600 * 1000 msec, default = 500.
* `probe_timeout` - Time to wait before a probe packet is considered lost (default = 500). On FortiOS versions 6.2.4-7.0.10, 7.2.0-7.2.4: 500 - 5000 msec. On FortiOS versions 7.0.11-7.0.17, >= 7.2.6: 20 - 5000 msec.
* `failtime` - Number of retry attempts before the server is considered down (default = 5). On FortiOS versions 6.2.0-7.0.5: 1 - 10. On FortiOS versions >= 7.0.6: 1 - 3600.
* `recoverytime` - Number of successful responses received before server is considered recovered (default = 5). On FortiOS versions 6.2.0-7.0.5: 1 - 10. On FortiOS versions >= 7.0.6: 1 - 3600.
* `probe_count` - Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
* `security_mode` - Twamp controller security mode. Valid values: `none`, `authentication`.
* `password` - Twamp controller password in authentication mode
* `packet_size` - Packet size of a TWAMP test session.
* `ha_priority` - HA election priority (1 - 50).
* `fail_weight` - Threshold weight to trigger link failure alert.
* `update_cascade_interface` - Enable/disable update cascade interface. Valid values: `enable`, `disable`.
* `update_static_route` - Enable/disable updating the static route. Valid values: `enable`, `disable`.
* `update_policy_route` - Enable/disable updating the policy route. Valid values: `enable`, `disable`.
* `status` - Enable/disable this link monitor. Valid values: `enable`, `disable`.
* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `class_id` - Traffic class ID.
* `service_detection` - Only use monitor to read quality values. If enabled, static routes and cascade interfaces will not be updated. Valid values: `enable`, `disable`.
* `server_list` - Servers for link-monitor to monitor. The structure of `server_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `server` block supports:

* `address` - Server address.

The `route` block supports:

* `subnet` - IP and netmask (x.x.x.x/y).

The `server_list` block supports:

* `id` - Server ID.
* `dst` - IP address of the server to be monitored.
* `protocol` - Protocols used to monitor the server.
* `port` - Port number of the traffic to be used to monitor the server.
* `weight` - Weight of the monitor to this dst (0 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System LinkMonitor can be imported using any of these accepted formats:
```
$ terraform import fortios_system_linkmonitor.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_linkmonitor.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
