---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_managedswitch"
description: |-
  Configure FortiSwitch devices that are managed by this FortiGate.
---

# fortios_switchcontroller_managedswitch
Configure FortiSwitch devices that are managed by this FortiGate.

## Argument Reference

The following arguments are supported:

* `switch_id` - (Required) Managed-switch id.
* `name` - Managed-switch name.
* `description` - Description.
* `switch_profile` - FortiSwitch profile.
* `fsw_wan1_peer` - (Required) Fortiswitch WAN1 peer port.
* `fsw_wan1_admin` - FortiSwitch WAN1 admin status; enable to authorize the FortiSwitch as a managed switch.
* `fsw_wan2_peer` - FortiSwitch WAN2 peer port.
* `fsw_wan2_admin` - FortiSwitch WAN2 admin status; enable to authorize the FortiSwitch as a managed switch.
* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection.
* `directly_connected` - Directly connected FortiSwitch.
* `version` - FortiSwitch version.
* `max_allowed_trunk_members` - FortiSwitch maximum allowed trunk members.
* `pre_provisioned` - Pre-provisioned managed switch.
* `dynamic_capability` - List of features this FortiSwitch supports (not configurable) that is sent to the FortiGate device for subsequent configuration initiated by the FortiGate device.
* `switch_device_tag` - User definable label/tag.
* `dynamically_discovered` - Dynamically discovered FortiSwitch.
* `type` - Indication of switch type, physical or virtual.
* `owner_vdom` - VDOM which owner of port belongs to.
* `staged_image_version` - Staged image version for FortiSwitch.
* `delayed_restart_trigger` - Delayed restart triggered for this FortiSwitch.
* `ports` - Managed-switch port list.
* `stp_settings` - Configuration method to edit Spanning Tree Protocol (STP) settings used to prevent bridge loops.
* `switch_stp_settings` - Configure spanning tree protocol (STP).
* `switch_log` - Configuration method to edit FortiSwitch logging settings (logs are transferred to and inserted into the FortiGate event log).
* `storm_control` - Configuration method to edit FortiSwitch storm control for measuring traffic activity using data rates to prevent traffic disruption.
* `mirror` - Configuration method to edit FortiSwitch packet mirror.
* `static_mac` - Configuration method to edit FortiSwitch Static and Sticky MAC.
* `custom_command` - Configuration method to edit FortiSwitch commands to be pushed to this FortiSwitch device upon rebooting the FortiGate switch controller or the FortiSwitch.
* `igmp_snooping` - Configure FortiSwitch IGMP snooping global settings.
* `n802_1X_settings` - Configuration method to edit FortiSwitch 802.1X global settings.

The `ports` block supports:

* `port_name` - Switch port name.
* `port_owner` - Switch port name.
* `switch_id` - Switch id.
* `speed` - Switch port speed; default and available settings depend on hardware.
* `speed_mask` - Switch port speed mask.
* `status` - Switch port admin status: up or down.
* `poe_status` - Enable/disable PoE status.
* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection.
* `port_number` - Port number.
* `port_prefix_type` - Port prefix type.
* `fortilink_port` - FortiLink uplink port.
* `poe_capable` - PoE capable.
* `stacking_port` - Stacking port.
* `fiber_port` - Fiber-port.
* `flags` - Port properties flags.
* `virtual_port` - Virtualized switch port.
* `isl_local_trunk_name` - ISL local trunk name.
* `isl_peer_port_name` - ISL peer port name.
* `isl_peer_device_name` - ISL peer device name.
* `fgt_peer_port_name` - FGT peer port name.
* `fgt_peer_device_name` - FGT peer device name.
* `vlan` - Assign switch ports to a VLAN.
* `allowed_vlans_all` - Enable/disable all defined vlans on this port.
* `allowed_vlans` - Configure switch port tagged vlans
* `untagged_vlans` - Configure switch port untagged vlans
* `type` - Interface type: physical or trunk port.
* `dhcp_snooping` - Trusted or untrusted DHCP-snooping interface.
* `dhcp_snoop_option82_trust` - Enable/disable allowance of DHCP with option-82 on untrusted interface.
* `arp_inspection_trust` - Trusted or untrusted dynamic ARP inspection.
* `igmp_snooping` - Set IGMP snooping mode for the physical port interface.
* `igmps_flood_reports` - Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled.
* `igmps_flood_traffic` - Enable/disable flooding of IGMP snooping traffic to this interface.
* `stp_state` - Enable/disable Spanning Tree Protocol (STP) on this interface.
* `stp_root_guard` - Enable/disable STP root guard on this interface.
* `stp_bpdu_guard` - Enable/disable STP BPDU guard on this interface.
* `stp_bpdu_guard_timeout` - BPDU Guard disabling protection (0 - 120 min).
* `edge_port` - Enable/disable this interface as an edge port, bridging connections between workstations and/or computers.
* `discard_mode` - Configure discard mode for port.
* `sflow_sampler` - Enable/disable sFlow protocol on this interface.
* `sflow_sample_rate` - sFlow sampler sample rate (0 - 99999 p/sec).
* `sflow_counter_interval` - sFlow sampler counter polling interval (1 - 255 sec).
* `sample_direction` - sFlow sample direction.
* `loop_guard` - Enable/disable loop-guard on this interface, an STP optimization used to prevent network loops.
* `loop_guard_timeout` - Loop-guard timeout (0 - 120 min, default = 45).
* `qos_policy` - Switch controller QoS policy from available options.
* `port_security_policy` - Switch controller authentication policy to apply to this managed switch from available options.
* `export_to_pool` - Switch controller export port to pool-list.
* `export_tags` - Switch controller export tag name.
* `export_to_pool_flag` - Switch controller export port to pool-list.
* `learning_limit` - Limit the number of dynamic MAC addresses on this Port (1 - 128, 0 = no limit, default).
* `sticky_mac` - Enable or disable sticky-mac on the interface.
* `lldp_status` - LLDP transmit and receive status.
* `lldp_profile` - LLDP port TLV profile.
* `export_to` - Export managed-switch port to a tenant VDOM.
* `port_selection_criteria` - Algorithm for aggregate port selection.
* `description` - Description for port.
* `lacp_speed` - end Link Aggregation Control Protocol (LACP) messages every 30 seconds (slow) or every second (fast).
* `mode` - LACP mode: ignore and do not send control messages, or negotiate 802.3ad aggregation passively or actively.
* `bundle` - Enable/disable Link Aggregation Group (LAG) bundling for non-FortiLink interfaces.
* `member_withdrawal_behavior` - Port behavior after it withdraws because of loss of control packets.
* `mclag` - Enable/disable multi-chassis link aggregation (MCLAG).
* `min_bundle` - Minimum size of LAG bundle (1 - 24, default = 1)
* `max_bundle` - Maximum size of LAG bundle (1 - 24, default = 24)
* `members` - Aggregated LAG bundle interfaces.

The `allowed_vlans` block supports:

* `vlan_name` - VLAN name.

The `untagged_vlans` block supports:

* `vlan_name` - VLAN name.

The `export_tags` block supports:

* `tag_name` - Switch tag name.

The `members` block supports:

* `member_name` - Interface name from available options.

The `stp_settings` block supports:

* `local_override` - Enable to configure local STP settings that override global STP settings.
* `name` - Name of local STP settings configuration.
* `status` - Enable/disable STP.
* `revision` - STP revision number (0 - 65535).
* `hello_time` - Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
* `forward_time` - Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
* `max_age` - Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
* `max_hops` - Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
* `pending_timer` - Pending time (1 - 15 sec, default = 4).

The `switch_stp_settings` block supports:

* `status` - Enable/disable STP.

The `switch_log` block supports:

* `local_override` - Enable to configure local logging settings that override global logging settings.
* `status` - Enable/disable adding FortiSwitch logs to the FortiGate event log.
* `severity` - Severity of FortiSwitch logs that are added to the FortiGate event log.

The `storm_control` block supports:

* `local_override` - Enable to override global FortiSwitch storm control settings for this FortiSwitch.
* `rate` - Rate in packets per second at which storm traffic is controlled (1 - 10000000, default = 500). Storm control drops excess traffic data rates beyond this threshold.
* `unknown_unicast` - Enable/disable storm control to drop unknown unicast traffic.
* `unknown_multicast` - Enable/disable storm control to drop unknown multicast traffic.
* `broadcast` - Enable/disable storm control to drop broadcast traffic.

The `mirror` block supports:

* `name` - Mirror name.
* `status` - Active/inactive mirror configuration.
* `switching_packet` - Enable/disable switching functionality when mirroring.
* `dst` - Destination port.
* `src_ingress` - Source ingress interfaces.
* `src_egress` - Source egress interfaces.

The `src_ingress` block supports:

* `name` - Interface name.

The `src_egress` block supports:

* `name` - Interface name.

The `static_mac` block supports:

* `id` - Id
* `type` - Type.
* `vlan` - Vlan.
* `mac` - MAC address.
* `interface` - Interface name.
* `description` - Description.

The `custom_command` block supports:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Names of commands to be pushed to this FortiSwitch device, as configured under config switch-controller custom-command.

The `igmp_snooping` block supports:

* `local_override` - Enable/disable overriding the global IGMP snooping configuration.
* `aging_time` - Maximum time to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
* `flood_unknown_multicast` - Enable/disable unknown multicast flooding.

The `n802_1X_settings` block supports:

* `local_override` - Enable to override global 802.1X settings on individual FortiSwitches.
* `link_down_auth` - Authentication state to set if a link is down.
* `reauth_period` - Reauthentication time interval (1 - 1440 min, default = 60, 0 = disable).
* `max_reauth_attempt` - Maximum number of authentication attempts (0 - 15, default = 3).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{switch_id}}.

## Import

SwitchController ManagedSwitch can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_managedswitch.labelname {{switch_id}}
$ unset "FORTIOS_IMPORT_TABLE"
```
