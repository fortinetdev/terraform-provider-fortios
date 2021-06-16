---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_linkmonitor"
description: |-
  Get information on an fortios system linkmonitor.
---

# Data Source: fortios_system_linkmonitor
Use this data source to get information on an fortios system linkmonitor

## Argument Reference

* `name` - (Required) Specify the name of the desired system linkmonitor.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Link monitor name.
* `addr_mode` - Address mode (IPv4 or IPv6).
* `srcintf` - Interface that receives the traffic to be monitored.
* `server` - IP address of the server(s) to be monitored. The structure of `server` block is documented below.
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

The `server` block contains:

* `address` - Server address.

