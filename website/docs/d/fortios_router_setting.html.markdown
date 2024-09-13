---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_setting"
description: |-
  Get information on fortios router setting.
---

# Data Source: fortios_router_setting
Use this data source to get information on fortios router setting

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `show_filter` - Prefix-list as filter for showing routes.
* `hostname` - Hostname for this virtual domain router.
* `kernel_route_distance` - Administrative distance for routes learned from kernel (0 - 255).
* `ospf_debug_lsa_flags` - ospf_debug_lsa_flags
* `ospf_debug_nfsm_flags` - ospf_debug_nfsm_flags
* `ospf_debug_packet_flags` - ospf_debug_packet_flags
* `ospf_debug_events_flags` - ospf_debug_events_flags
* `ospf_debug_route_flags` - ospf_debug_route_flags
* `ospf_debug_ifsm_flags` - ospf_debug_ifsm_flags
* `ospf_debug_nsm_flags` - ospf_debug_nsm_flags
* `rip_debug_flags` - rip_debug_flags
* `bgp_debug_flags` - bgp_debug_flags
* `igmp_debug_flags` - igmp_debug_flags
* `pimdm_debug_flags` - pimdm_debug_flags
* `pimsm_debug_simple_flags` - pimsm_debug_simple_flags
* `pimsm_debug_timer_flags` - pimsm_debug_timer_flags
* `pimsm_debug_joinprune_flags` - pimsm_debug_joinprune_flags
* `imi_debug_flags` - imi_debug_flags
* `isis_debug_flags` - isis_debug_flags
* `ospf6_debug_lsa_flags` - ospf6_debug_lsa_flags
* `ospf6_debug_nfsm_flags` - ospf6_debug_nfsm_flags
* `ospf6_debug_packet_flags` - ospf6_debug_packet_flags
* `ospf6_debug_events_flags` - ospf6_debug_events_flags
* `ospf6_debug_route_flags` - ospf6_debug_route_flags
* `ospf6_debug_ifsm_flags` - ospf6_debug_ifsm_flags
* `ospf6_debug_nsm_flags` - ospf6_debug_nsm_flags
* `ripng_debug_flags` - ripng_debug_flags

