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
* `addr_mode` - Address mode (IPv4 or IPv6).
* `srcintf` - Interface that receives the traffic to be monitored.
* `server` - (Required) IP address of the server(s) to be monitored. The structure of `server` block is documented below.
* `protocol` - Protocols used to monitor the server.
* `port` - Port number of the traffic to be used to monitor the server.
* `gateway_ip` - Gateway IP address used to probe the server.
* `gateway_ip6` - Gateway IPv6 address used to probe the server.
* `source_ip` - Source IP address used in packet to the server.
* `source_ip6` - Source IPv6 address used in packet to the server.
* `http_get` - If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_match` - String that you expect to see in the HTTP-GET requests of the traffic to be monitored.
* `interval` - Detection interval (1 - 3600 sec, default = 5).
* `probe_timeout` - Time to wait before a probe packet is considered lost (500 - 5000 msec, default = 500).
* `failtime` - Number of retry attempts before the server is considered down (1 - 10, default = 5)
* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 10, default = 5).
* `probe_count` - Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
* `security_mode` - Twamp controller security mode.
* `password` - Twamp controller password in authentication mode
* `packet_size` - Packet size of a twamp test session,
* `ha_priority` - HA election priority (1 - 50).
* `update_cascade_interface` - Enable/disable update cascade interface.
* `update_static_route` - Enable/disable updating the static route.
* `status` - Enable/disable this link monitor.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server` block supports:

* `address` - Server address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System LinkMonitor can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_linkmonitor.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
