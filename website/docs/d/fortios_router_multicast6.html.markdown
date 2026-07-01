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
* `pim_sm_global_vrf` - per-VRF PIM sparse-mode global settings. The structure of `pim_sm_global_vrf` block is documented below.

The `interface` block contains:

* `name` - Interface name.
* `hello_interval` - Interval between sending PIM hello messages  (1 - 65535 sec, default = 30)..
* `hello_holdtime` - Time before old neighbour information expires (1 - 65535 sec, default = 105).
* `rp_candidate` - Enable/disable compete to become RP in elections.
* `rp_candidate_group` - Multicast groups managed by this RP.
* `rp_candidate_priority` - Router's priority as RP.
* `rp_candidate_interval` - RP candidate advertisement interval (1 - 16383 sec, default = 60).
* `static_group` - Statically set IPv6 multicast groups to forward out.

The `pim_sm_global` block contains:

* `bsr_candidate` - Enable/disable allowing this router to become a bootstrap router (BSR).
* `bsr_interface` - Interface to advertise as candidate BSR.
* `bsr_priority` - BSR priority (0 - 255, default = 0).
* `bsr_hash` - BSR hash length (0 - 128, default = 126).
* `bsr_allow_quick_refresh` - Enable/disable accept BSR quick refresh packets from neighbors.
* `cisco_crp_prefix` - Enable/disable making candidate RP compatible with old Cisco IOS.
* `register_rate_limit` - Limit of packets/sec per source registered through this RP (0 means unlimited).
* `cisco_ignore_rp_set_priority` - Use only hash for RP selection (compatibility with old Cisco IOS).
* `spt_threshold` - Enable/disable switching to source specific trees.
* `spt_threshold_group` - Groups allowed to switch to source tree.
* `pim_use_sdwan` - Enable/disable use of SDWAN when checking RPF neighbor and sending of REG packet.
* `rp_address` - Statically configured RP addresses. The structure of `rp_address` block is documented below.

The `rp_address` block contains:

* `id` - ID of the entry.
* `ip6_address` - RP router IPv6 address.
* `group` - Groups to use this RP.

The `pim_sm_global_vrf` block contains:

* `vrf` - VRF ID.
* `bsr_candidate` - Enable/disable allowing this router to become a bootstrap router (BSR).
* `bsr_interface` - Interface to advertise as candidate BSR.
* `bsr_priority` - BSR priority (0 - 255, default = 0).
* `bsr_hash` - BSR hash length (0 - 128, default = 126).
* `bsr_allow_quick_refresh` - Enable/disable accept BSR quick refresh packets from neighbors.
* `cisco_crp_prefix` - Enable/disable making candidate RP compatible with old Cisco IOS.
* `rp_address` - Statically configured RP addresses. The structure of `rp_address` block is documented below.

The `rp_address` block contains:

* `id` - ID of the entry.
* `ip6_address` - RP router IPv6 address.
* `group` - Groups to use this RP.

