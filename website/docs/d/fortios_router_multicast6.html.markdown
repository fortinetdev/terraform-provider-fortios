---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicast6"
description: |-
  Get information on fortios router multicast6.
---

# Data Source: fortios_router_multicast6
Use this data source to get information on fortios router multicast6

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `multicast_routing` - Enable/disable IPv6 multicast routing.
* `multicast_pmtu` - Enable/disable PMTU for IPv6 multicast.
* `interface` - Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
* `pim_sm_global` - PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.

The `interface` block contains:

* `name` - Interface name.
* `hello_interval` - Interval between sending PIM hello messages  (1 - 65535 sec, default = 30)..
* `hello_holdtime` - Time before old neighbour information expires (1 - 65535 sec, default = 105).

The `pim_sm_global` block contains:

* `register_rate_limit` - Limit of packets/sec per source registered through this RP (0 means unlimited).
* `rp_address` - Statically configured RP addresses. The structure of `rp_address` block is documented below.

The `rp_address` block contains:

* `id` - ID of the entry.
* `ip6_address` - RP router IPv6 address.

