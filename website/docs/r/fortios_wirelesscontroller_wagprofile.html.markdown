---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wagprofile"
description: |-
  Configure wireless access gateway (WAG) profiles used for tunnels on AP.
---

# fortios_wirelesscontroller_wagprofile
Configure wireless access gateway (WAG) profiles used for tunnels on AP. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `name` - Tunnel profile name.
* `comment` - Comment.
* `tunnel_type` - Tunnel type. Valid values: `l2tpv3`, `gre`.
* `wag_ip` - IP Address of the wireless access gateway.
* `wag_port` - UDP port of the wireless access gateway.
* `ping_interval` - Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
* `ping_number` - Number of the tunnel mointoring echo packets (1 - 65535, default = 5).
* `return_packet_timeout` - Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
* `dhcp_ip_addr` - IP address of the monitoring DHCP request packet sent through the tunnel.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController WagProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_wagprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_wagprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
