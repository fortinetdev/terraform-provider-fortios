---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wagprofile"
description: |-
  Configure wireless access gateway (WAG) profiles used for tunnels on AP.
---

# fortios_wirelesscontroller_wagprofile
Configure wireless access gateway (WAG) profiles used for tunnels on AP.

## Argument Reference

The following arguments are supported:

* `name` - Tunnel profile name.
* `comment` - Comment.
* `tunnel_type` - Tunnel type.
* `wag_ip` - IP Address of the wireless access gateway.
* `wag_port` - UDP port of the wireless access gateway.
* `ping_interval` - Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
* `ping_number` - Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
* `return_packet_timeout` - Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
* `dhcp_ip_addr` - IP address of the monitoring DHCP request packet sent through the tunnel.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController WagProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_wagprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
