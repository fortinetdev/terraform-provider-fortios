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
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - VXLAN device or interface name. Must be a unique interface name.
* `interface` - Outgoing interface for VXLAN encapsulated traffic.
* `vni` - VXLAN network ID.
* `ip_version` - IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast.
* `remote_ip` - IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip` block is documented below.
* `local_ip` - IPv4 address to use as the source address for egress VXLAN packets.
* `remote_ip6` - IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip6` block is documented below.
* `local_ip6` - IPv6 address to use as the source address for egress VXLAN packets.
* `dstport` - VXLAN destination port (1 - 65535, default = 4789).
* `multicast_ttl` - VXLAN multicast TTL (1-255, default = 0).
* `evpn_id` - EVPN instance.
* `learn_from_traffic` - Enable/disable VXLAN MAC learning from traffic.

The `remote_ip` block contains:

* `ip` - IPv4 address.

The `remote_ip6` block contains:

* `ip6` - IPv6 address.

