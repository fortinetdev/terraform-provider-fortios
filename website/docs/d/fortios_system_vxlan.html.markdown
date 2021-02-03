---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vxlan"
description: |-
  Get information on an fortios system vxlan.
---

# Data Source: fortios_system_vxlan
Use this data source to get information on an fortios system vxlan

## Argument Reference

* `name` - (Required) Specify the name of the desired system vxlan.

## Attribute Reference

The following attributes are exported:

* `name` - VXLAN device or interface name. Must be a unique interface name.
* `interface` - Outgoing interface for VXLAN encapsulated traffic.
* `vni` - VXLAN network ID.
* `ip_version` - IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast.
* `remote_ip` - IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip` block is documented below.
* `remote_ip6` - IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip6` block is documented below.
* `dstport` - VXLAN destination port (1 - 65535, default = 4789).
* `multicast_ttl` - VXLAN multicast TTL (1-255, default = 0).

The `remote_ip` block contains:

* `ip` - IPv4 address.

The `remote_ip6` block contains:

* `ip6` - IPv6 address.

