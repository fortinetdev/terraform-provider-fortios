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
* `sn` - Managed-switch serial number.
* `name` - Managed-switch name.
* `description` - Description.
* `switch_profile` - FortiSwitch profile.
* `access_profile` - FortiSwitch access profile.
* `purdue_level` - Purdue Level of this FortiSwitch. Valid values: `1`, `1.5`, `2`, `2.5`, `3`, `3.5`, `4`, `5`, `5.5`.
* `fsw_wan1_peer` - (Required) Fortiswitch WAN1 peer port.
* `fsw_wan1_admin` - FortiSwitch WAN1 admin status; enable to authorize the FortiSwitch as a managed switch. Valid values: `discovered`, `disable`, `enable`.
* `fsw_wan2_peer` - FortiSwitch WAN2 peer port.
* `fsw_wan2_admin` - FortiSwitch WAN2 admin status; enable to authorize the FortiSwitch as a managed switch. Valid values: `discovered`, `disable`, `enable`.
* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection. Valid values: `enable`, `disable`.
* `dhcp_server_access_list` - DHCP snooping server access list. Valid values: `global`, `enable`, `disable`.
* `poe_detection_type` - PoE detection type for FortiSwitch.
* `max_poe_budget` - Max PoE budget for FortiSwitch.
* `poe_lldp_detection` - Enable/disable PoE LLDP detection. Valid values: `enable`, `disable`.
* `directly_connected` - Directly connected FortiSwitch.
* `version` - FortiSwitch version.
* `max_allowed_trunk_members` - FortiSwitch maximum allowed trunk members.
* `pre_provisioned` - Pre-provisioned managed switch.
* `l3_discovered` - Layer 3 management discovered.
* `mgmt_mode` - FortiLink management mode.
* `tunnel_discovered` - SOCKS tunnel management discovered.
* `tdr_supported` - TDR supported.
* `dynamic_capability` - List of features this FortiSwitch supports (not configurable) that is sent to the FortiGate device for subsequent configuration initiated by the FortiGate device.
* `switch_device_tag` - User definable label/tag.
* `switch_dhcp_opt43_key` - DHCP option43 key.
* `mclag_igmp_snooping_aware` - Enable/disable MCLAG IGMP-snooping awareness. Valid values: `enable`, `disable`.
* `dynamically_discovered` - Dynamically discovered FortiSwitch.
* `ptp_status` - Enable/disable PTP profile on this FortiSwitch. Valid values: `disable`, `enable`.
* `ptp_profile` - PTP profile configuration.
* `radius_nas_ip_override` - Use locally defined NAS-IP. Valid values: `disable`, `enable`.
* `radius_nas_ip` - NAS-IP address.
* `route_offload` - Enable/disable route offload on this FortiSwitch. Valid values: `disable`, `enable`.
* `route_offload_mclag` - Enable/disable route offload MCLAG on this FortiSwitch. Valid values: `disable`, `enable`.
* `route_offload_router` - Configure route offload MCLAG IP address. The structure of `route_offload_router` block is documented below.
* `vlan` - Configure VLAN assignment priority. The structure of `vlan` block is documented below.
* `type` - Indication of switch type, physical or virtual. Valid values: `virtual`, `physical`.
* `owner_vdom` - VDOM which owner of port belongs to.
* `flow_identity` - Flow-tracking netflow ipfix switch identity in hex format(00000000-FFFFFFFF default=0).
* `staged_image_version` - Staged image version for FortiSwitch.
* `delayed_restart_trigger` - Delayed restart triggered for this FortiSwitch.
* `firmware_provision` - Enable/disable provisioning of firmware to FortiSwitches on join connection. Valid values: `enable`, `disable`.
* `firmware_provision_version` - Firmware version to provision to this FortiSwitch on bootup (major.minor.build, i.e. 6.2.1234).
* `firmware_provision_latest` - Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
* `ports` - Managed-switch port list. The structure of `ports` block is documented below.
* `ip_source_guard` - IP source guard. The structure of `ip_source_guard` block is documented below.
* `stp_settings` - Configuration method to edit Spanning Tree Protocol (STP) settings used to prevent bridge loops. The structure of `stp_settings` block is documented below.
* `stp_instance` - Configuration method to edit Spanning Tree Protocol (STP) instances. The structure of `stp_instance` block is documented below.
* `override_snmp_sysinfo` - Enable/disable overriding the global SNMP system information. Valid values: `disable`, `enable`.
* `snmp_sysinfo` - Configuration method to edit Simple Network Management Protocol (SNMP) system info. The structure of `snmp_sysinfo` block is documented below.
* `override_snmp_trap_threshold` - Enable/disable overriding the global SNMP trap threshold values. Valid values: `enable`, `disable`.
* `snmp_trap_threshold` - Configuration method to edit Simple Network Management Protocol (SNMP) trap threshold values. The structure of `snmp_trap_threshold` block is documented below.
* `override_snmp_community` - Enable/disable overriding the global SNMP communities. Valid values: `enable`, `disable`.
* `snmp_community` - Configuration method to edit Simple Network Management Protocol (SNMP) communities. The structure of `snmp_community` block is documented below.
* `override_snmp_user` - Enable/disable overriding the global SNMP users. Valid values: `enable`, `disable`.
* `snmp_user` - Configuration method to edit Simple Network Management Protocol (SNMP) users. The structure of `snmp_user` block is documented below.
* `qos_drop_policy` - Set QoS drop-policy. Valid values: `taildrop`, `random-early-detection`.
* `qos_red_probability` - Set QoS RED/WRED drop probability.
* `switch_stp_settings` - Configure spanning tree protocol (STP). The structure of `switch_stp_settings` block is documented below.
* `switch_log` - Configuration method to edit FortiSwitch logging settings (logs are transferred to and inserted into the FortiGate event log). The structure of `switch_log` block is documented below.
* `remote_log` - Configure logging by FortiSwitch device to a remote syslog server. The structure of `remote_log` block is documented below.
* `storm_control` - Configuration method to edit FortiSwitch storm control for measuring traffic activity using data rates to prevent traffic disruption. The structure of `storm_control` block is documented below.
* `mirror` - Configuration method to edit FortiSwitch packet mirror. The structure of `mirror` block is documented below.
* `static_mac` - Configuration method to edit FortiSwitch Static and Sticky MAC. The structure of `static_mac` block is documented below.
* `custom_command` - Configuration method to edit FortiSwitch commands to be pushed to this FortiSwitch device upon rebooting the FortiGate switch controller or the FortiSwitch. The structure of `custom_command` block is documented below.
* `dhcp_snooping_static_client` - Configure FortiSwitch DHCP snooping static clients. The structure of `dhcp_snooping_static_client` block is documented below.
* `igmp_snooping` - Configure FortiSwitch IGMP snooping global settings. The structure of `igmp_snooping` block is documented below.
* `n802_1x_settings` - Configuration method to edit FortiSwitch 802.1X global settings. The structure of `n802_1x_settings` block is documented below.
* `router_vrf` - Configure VRF. The structure of `router_vrf` block is documented below.
* `system_interface` - Configure system interface on FortiSwitch. The structure of `system_interface` block is documented below.
* `router_static` - Configure static routes. The structure of `router_static` block is documented below.
* `system_dhcp_server` - Configure DHCP servers. The structure of `system_dhcp_server` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `route_offload_router` block supports:

* `vlan_name` - VLAN name.
* `router_ip` - Router IP address.

The `vlan` block supports:

* `vlan_name` - VLAN name.
* `assignment_priority` - 802.1x Radius (Tunnel-Private-Group-Id) VLANID assign-by-name priority. A smaller value has a higher priority.

The `ports` block supports:

* `port_name` - Switch port name.
* `port_owner` - Switch port name.
* `switch_id` - Switch id.
* `speed` - Switch port speed; default and available settings depend on hardware.
* `speed_mask` - Switch port speed mask.
* `status` - Switch port admin status: up or down. Valid values: `up`, `down`.
* `poe_status` - Enable/disable PoE status. Valid values: `enable`, `disable`.
* `ip_source_guard` - Enable/disable IP source guard. Valid values: `disable`, `enable`.
* `ptp_status` - Enable/disable PTP policy on this FortiSwitch port. Valid values: `disable`, `enable`.
* `ptp_policy` - PTP policy configuration.
* `aggregator_mode` - LACP member select mode. Valid values: `bandwidth`, `count`.
* `flapguard` - Enable/disable flap guard. Valid values: `enable`, `disable`.
* `flap_rate` - Number of stage change events needed within flap-duration.
* `flap_duration` - Period over which flap events are calculated (seconds).
* `flap_timeout` - Flap guard disabling protection (min).
* `rpvst_port` - Enable/disable inter-operability with rapid PVST on this interface. Valid values: `disabled`, `enabled`.
* `poe_pre_standard_detection` - Enable/disable PoE pre-standard detection. Valid values: `enable`, `disable`.
* `port_number` - Port number.
* `port_prefix_type` - Port prefix type.
* `fortilink_port` - FortiLink uplink port.
* `link_status` - Port link status. Valid values: `up`, `down`.
* `poe_capable` - PoE capable.
* `pd_capable` - Powered device capable.
* `stacking_port` - Stacking port.
* `p2p_port` - General peer to peer tunnel port.
* `mclag_icl_port` - MCLAG-ICL port.
* `authenticated_port` - Peer to Peer Authenticated port.
* `restricted_auth_port` - Peer to Peer Restricted Authenticated port.
* `encrypted_port` - Peer to Peer Encrypted port.
* `fiber_port` - Fiber-port.
* `media_type` - Media type.
* `poe_standard` - PoE standard supported.
* `poe_max_power` - PoE maximum power.
* `poe_mode_bt_cabable` - PoE mode IEEE 802.3BT capable.
* `poe_port_mode` - Configure PoE port mode. Valid values: `ieee802-3af`, `ieee802-3at`, `ieee802-3bt`.
* `poe_port_priority` - Configure PoE port priority. Valid values: `critical-priority`, `high-priority`, `low-priority`, `medium-priority`.
* `poe_port_power` - Configure PoE port power. Valid values: `normal`, `perpetual`, `perpetual-fast`.
* `flags` - Port properties flags.
* `virtual_port` - Virtualized switch port.
* `isl_local_trunk_name` - ISL local trunk name.
* `isl_peer_port_name` - ISL peer port name.
* `isl_peer_device_name` - ISL peer device name.
* `isl_peer_device_sn` - ISL peer device serial number.
* `fgt_peer_port_name` - FGT peer port name.
* `fgt_peer_device_name` - FGT peer device name.
* `vlan` - Assign switch ports to a VLAN.
* `allowed_vlans_all` - Enable/disable all defined vlans on this port. Valid values: `enable`, `disable`.
* `allowed_vlans` - Configure switch port tagged vlans The structure of `allowed_vlans` block is documented below.
* `untagged_vlans` - Configure switch port untagged vlans The structure of `untagged_vlans` block is documented below.
* `type` - Interface type: physical or trunk port. Valid values: `physical`, `trunk`.
* `access_mode` - Access mode of the port.
* `matched_dpp_policy` - Matched child policy in the dynamic port policy.
* `matched_dpp_intf_tags` - Matched interface tags in the dynamic port policy.
* `acl_group` - ACL groups on this port. The structure of `acl_group` block is documented below.
* `fortiswitch_acls` - ACLs on this port. The structure of `fortiswitch_acls` block is documented below.
* `dhcp_snooping` - Trusted or untrusted DHCP-snooping interface. Valid values: `untrusted`, `trusted`.
* `dhcp_snoop_option82_trust` - Enable/disable allowance of DHCP with option-82 on untrusted interface. Valid values: `enable`, `disable`.
* `dhcp_snoop_option82_override` - Configure DHCP snooping option 82 override. The structure of `dhcp_snoop_option82_override` block is documented below.
* `arp_inspection_trust` - Trusted or untrusted dynamic ARP inspection. Valid values: `untrusted`, `trusted`.
* `igmp_snooping_flood_reports` - Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled. Valid values: `enable`, `disable`.
* `mcast_snooping_flood_traffic` - Enable/disable flooding of IGMP snooping traffic to this interface. Valid values: `enable`, `disable`.
* `igmp_snooping` - Set IGMP snooping mode for the physical port interface. Valid values: `enable`, `disable`.
* `igmps_flood_reports` - Enable/disable flooding of IGMP reports to this interface when igmp-snooping enabled. Valid values: `enable`, `disable`.
* `igmps_flood_traffic` - Enable/disable flooding of IGMP snooping traffic to this interface. Valid values: `enable`, `disable`.
* `stp_state` - Enable/disable Spanning Tree Protocol (STP) on this interface. Valid values: `enabled`, `disabled`.
* `stp_root_guard` - Enable/disable STP root guard on this interface. Valid values: `enabled`, `disabled`.
* `stp_bpdu_guard` - Enable/disable STP BPDU guard on this interface. Valid values: `enabled`, `disabled`.
* `stp_bpdu_guard_timeout` - BPDU Guard disabling protection (0 - 120 min).
* `edge_port` - Enable/disable this interface as an edge port, bridging connections between workstations and/or computers. Valid values: `enable`, `disable`.
* `discard_mode` - Configure discard mode for port. Valid values: `none`, `all-untagged`, `all-tagged`.
* `packet_sampler` - Enable/disable packet sampling on this interface. Valid values: `enabled`, `disabled`.
* `packet_sample_rate` - Packet sampling rate (0 - 99999 p/sec).
* `sflow_sampler` - Enable/disable sFlow protocol on this interface. Valid values: `enabled`, `disabled`.
* `sflow_sample_rate` - sFlow sampler sample rate (0 - 99999 p/sec).
* `sflow_counter_interval` - sFlow sampling counter polling interval in seconds (0 - 255).
* `sample_direction` - sFlow sample direction. Valid values: `tx`, `rx`, `both`.
* `fec_capable` - FEC capable.
* `fec_state` - State of forward error correction.
* `flow_control` - Flow control direction. Valid values: `disable`, `tx`, `rx`, `both`.
* `pause_meter` - Configure ingress pause metering rate, in kbps (default = 0, disabled).
* `pause_meter_resume` - Resume threshold for resuming traffic on ingress port. Valid values: `75%`, `50%`, `25%`.
* `loop_guard` - Enable/disable loop-guard on this interface, an STP optimization used to prevent network loops. Valid values: `enabled`, `disabled`.
* `loop_guard_timeout` - Loop-guard timeout (0 - 120 min, default = 45).
* `port_policy` - Switch controller dynamic port policy from available options.
* `qos_policy` - Switch controller QoS policy from available options.
* `storm_control_policy` - Switch controller storm control policy from available options.
* `port_security_policy` - Switch controller authentication policy to apply to this managed switch from available options.
* `export_to_pool` - Switch controller export port to pool-list.
* `interface_tags` - Tag(s) associated with the interface for various features including virtual port pool, dynamic port policy. The structure of `interface_tags` block is documented below.
* `export_tags` - Switch controller export tag name. The structure of `export_tags` block is documented below.
* `export_to_pool_flag` - Switch controller export port to pool-list.
* `learning_limit` - Limit the number of dynamic MAC addresses on this Port (1 - 128, 0 = no limit, default).
* `sticky_mac` - Enable or disable sticky-mac on the interface. Valid values: `enable`, `disable`.
* `lldp_status` - LLDP transmit and receive status. Valid values: `disable`, `rx-only`, `tx-only`, `tx-rx`.
* `lldp_profile` - LLDP port TLV profile.
* `export_to` - Export managed-switch port to a tenant VDOM.
* `mac_addr` - Port/Trunk MAC.
* `allow_arp_monitor` - Enable/Disable allow ARP monitor. Valid values: `disable`, `enable`.
* `qnq` - 802.1AD VLANs in the VDom.
* `log_mac_event` - Enable/disable logging for dynamic MAC address events. Valid values: `disable`, `enable`.
* `port_selection_criteria` - Algorithm for aggregate port selection. Valid values: `src-mac`, `dst-mac`, `src-dst-mac`, `src-ip`, `dst-ip`, `src-dst-ip`.
* `description` - Description for port.
* `lacp_speed` - end Link Aggregation Control Protocol (LACP) messages every 30 seconds (slow) or every second (fast). Valid values: `slow`, `fast`.
* `mode` - LACP mode: ignore and do not send control messages, or negotiate 802.3ad aggregation passively or actively. Valid values: `static`, `lacp-passive`, `lacp-active`.
* `bundle` - Enable/disable Link Aggregation Group (LAG) bundling for non-FortiLink interfaces. Valid values: `enable`, `disable`.
* `member_withdrawal_behavior` - Port behavior after it withdraws because of loss of control packets. Valid values: `forward`, `block`.
* `mclag` - Enable/disable multi-chassis link aggregation (MCLAG). Valid values: `enable`, `disable`.
* `min_bundle` - Minimum size of LAG bundle (1 - 24, default = 1)
* `max_bundle` - Maximum size of LAG bundle (1 - 24, default = 24)
* `members` - Aggregated LAG bundle interfaces. The structure of `members` block is documented below.
* `fallback_port` - LACP fallback port.

The `allowed_vlans` block supports:

* `vlan_name` - VLAN name.

The `untagged_vlans` block supports:

* `vlan_name` - VLAN name.

The `acl_group` block supports:

* `name` - ACL group name.

The `fortiswitch_acls` block supports:

* `id` - ACL ID.

The `dhcp_snoop_option82_override` block supports:

* `vlan_name` - DHCP snooping option 82 VLAN.
* `circuit_id` - Circuit ID string.
* `remote_id` - Remote ID string.

The `interface_tags` block supports:

* `tag_name` - FortiSwitch port tag name when exported to a virtual port pool or matched to dynamic port policy.

The `export_tags` block supports:

* `tag_name` - Switch tag name.

The `members` block supports:

* `member_name` - Interface name from available options.

The `ip_source_guard` block supports:

* `port` - Ingress interface to which source guard is bound.
* `description` - Description.
* `binding_entry` - IP and MAC address configuration. The structure of `binding_entry` block is documented below.

The `binding_entry` block supports:

* `entry_name` - Configure binding pair.
* `ip` - Source IP for this rule.
* `mac` - MAC address for this rule.

The `stp_settings` block supports:

* `local_override` - Enable to configure local STP settings that override global STP settings. Valid values: `enable`, `disable`.
* `name` - Name of local STP settings configuration.
* `status` - Enable/disable STP. Valid values: `enable`, `disable`.
* `revision` - STP revision number (0 - 65535).
* `hello_time` - Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
* `forward_time` - Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
* `max_age` - Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
* `max_hops` - Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
* `pending_timer` - Pending time (1 - 15 sec, default = 4).

The `stp_instance` block supports:

* `id` - Instance ID.
* `priority` - Priority. Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`, `61440`.

The `snmp_sysinfo` block supports:

* `status` - Enable/disable SNMP. Valid values: `disable`, `enable`.
* `engine_id` - Local SNMP engine ID string (max 24 char).
* `description` - System description.
* `contact_info` - Contact information.
* `location` - System location.

The `snmp_trap_threshold` block supports:

* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.

The `snmp_community` block supports:

* `id` - SNMP community ID.
* `name` - SNMP community name.
* `status` - Enable/disable this SNMP community. Valid values: `disable`, `enable`.
* `hosts` - Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.
* `query_v2c_port` - SNMP v2c query port (default = 161).
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.
* `trap_v1_lport` - SNMP v2c trap local port (default = 162).
* `trap_v1_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.
* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `events` - SNMP notifications (traps) to send.

The `hosts` block supports:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).

The `snmp_user` block supports:

* `name` - SNMP user name.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.
* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
* `auth_proto` - Authentication protocol.
* `auth_pwd` - Password for authentication protocol.
* `priv_proto` - Privacy (encryption) protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.

The `switch_stp_settings` block supports:

* `status` - Enable/disable STP. Valid values: `enable`, `disable`.

The `switch_log` block supports:

* `local_override` - Enable to configure local logging settings that override global logging settings. Valid values: `enable`, `disable`.
* `status` - Enable/disable adding FortiSwitch logs to the FortiGate event log. Valid values: `enable`, `disable`.
* `severity` - Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

The `remote_log` block supports:

* `name` - Remote log name.
* `status` - Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
* `server` - IPv4 address of the remote syslog server.
* `port` - Remote syslog server listening port.
* `severity` - Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `csv` - Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
* `facility` - Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.

The `storm_control` block supports:

* `local_override` - Enable to override global FortiSwitch storm control settings for this FortiSwitch. Valid values: `enable`, `disable`.
* `rate` - Rate in packets per second at which storm control drops excess traffic, default=500. On FortiOS versions 6.2.0-7.2.12: 1 - 10000000. On FortiOS versions >= 7.4.0: 0-10000000, drop-all=0.
* `burst_size_level` - Increase level to handle bursty traffic (0 - 4, default = 0).
* `unknown_unicast` - Enable/disable storm control to drop unknown unicast traffic. Valid values: `enable`, `disable`.
* `unknown_multicast` - Enable/disable storm control to drop unknown multicast traffic. Valid values: `enable`, `disable`.
* `broadcast` - Enable/disable storm control to drop broadcast traffic. Valid values: `enable`, `disable`.

The `mirror` block supports:

* `name` - Mirror name.
* `status` - Active/inactive mirror configuration. Valid values: `active`, `inactive`.
* `switching_packet` - Enable/disable switching functionality when mirroring. Valid values: `enable`, `disable`.
* `dst` - Destination port.
* `src_ingress` - Source ingress interfaces. The structure of `src_ingress` block is documented below.
* `src_egress` - Source egress interfaces. The structure of `src_egress` block is documented below.

The `src_ingress` block supports:

* `name` - Interface name.

The `src_egress` block supports:

* `name` - Interface name.

The `static_mac` block supports:

* `id` - Id
* `type` - Type. Valid values: `static`, `sticky`.
* `vlan` - Vlan.
* `mac` - MAC address.
* `interface` - Interface name.
* `description` - Description.

The `custom_command` block supports:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Names of commands to be pushed to this FortiSwitch device, as configured under config switch-controller custom-command.

The `dhcp_snooping_static_client` block supports:

* `name` - Client name.
* `vlan` - VLAN name.
* `ip` - Client static IP address.
* `mac` - Client MAC address.
* `port` - Interface name.

The `igmp_snooping` block supports:

* `local_override` - Enable/disable overriding the global IGMP snooping configuration. Valid values: `enable`, `disable`.
* `aging_time` - Maximum time to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
* `flood_unknown_multicast` - Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
* `vlans` - Configure IGMP snooping VLAN. The structure of `vlans` block is documented below.

The `vlans` block supports:

* `vlan_name` - List of FortiSwitch VLANs.
* `proxy` - IGMP snooping proxy for the VLAN interface. Valid values: `disable`, `enable`, `global`.
* `querier` - Enable/disable IGMP snooping querier for the VLAN interface. Valid values: `disable`, `enable`.
* `querier_addr` - IGMP snooping querier address.
* `version` - IGMP snooping querier version.

The `n802_1x_settings` block supports:

* `local_override` - Enable to override global 802.1X settings on individual FortiSwitches. Valid values: `enable`, `disable`.
* `link_down_auth` - Authentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
* `reauth_period` - Reauthentication time interval (1 - 1440 min, default = 60, 0 = disable).
* `max_reauth_attempt` - Maximum number of authentication attempts (0 - 15, default = 3).
* `tx_period` - 802.1X Tx period (seconds, default=30).
* `mab_reauth` - Enable or disable MAB reauthentication settings. Valid values: `disable`, `enable`.
* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen). Valid values: `colon`, `hyphen`, `none`, `single-hyphen`.
* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen). Valid values: `colon`, `hyphen`, `none`, `single-hyphen`.
* `mac_calling_station_delimiter` - MAC calling station delimiter (default = hyphen). Valid values: `colon`, `hyphen`, `none`, `single-hyphen`.
* `mac_called_station_delimiter` - MAC called station delimiter (default = hyphen). Valid values: `colon`, `hyphen`, `none`, `single-hyphen`.
* `mac_case` - MAC case (default = lowercase). Valid values: `lowercase`, `uppercase`.

The `router_vrf` block supports:

* `name` - VRF entry name.
* `switch_id` - Switch ID.
* `vrfid` - VRF ID.

The `system_interface` block supports:

* `name` - Interface name.
* `switch_id` - Switch ID.
* `mode` - Interface addressing mode. Valid values: `static`, `dhcp`.
* `ip` - IP and mask for this interface.
* `status` - Enable/disable interface status. Valid values: `disable`, `enable`.
* `allowaccess` - Permitted types of management access to this interface. Valid values: `ping`, `https`, `http`, `ssh`, `snmp`, `telnet`, `radius-acct`.
* `vlan` - VLAN name.
* `type` - Interface type. Valid values: `vlan`, `physical`.
* `interface` - Interface name.
* `vrf` - VRF for this route.

The `router_static` block supports:

* `id` - Entry sequence number.
* `switch_id` - Switch ID.
* `blackhole` - Enable/disable blackhole on this route. Valid values: `disable`, `enable`.
* `comment` - Comment.
* `device` - Gateway out interface.
* `distance` - Administrative distance for the route (1 - 255, default = 10).
* `dst` - Destination ip and mask for this route.
* `dynamic_gateway` - Enable/disable dynamic gateway. Valid values: `disable`, `enable`.
* `gateway` - Gateway ip for this route.
* `status` - Enable/disable route status. Valid values: `disable`, `enable`.
* `vrf` - VRF for this route.

The `system_dhcp_server` block supports:

* `id` - ID.
* `switch_id` - Switch ID.
* `status` - Enable/disable this DHCP configuration. Valid values: `disable`, `enable`.
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `dns_service` - Options for assigning DNS servers to DHCP clients. Valid values: `local`, `default`, `specify`.
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients. Valid values: `local`, `default`, `specify`.
* `ntp_server1` - NTP server 1.
* `ntp_server2` - NTP server 2.
* `ntp_server3` - NTP server 3.
* `default_gateway` - Default gateway IP address assigned by the DHCP server.
* `netmask` - Netmask assigned by the DHCP server.
* `interface` - DHCP server can assign IP configurations to clients connected to this interface.
* `ip_range` - DHCP IP range configuration. The structure of `ip_range` block is documented below.
* `options` - DHCP options. The structure of `options` block is documented below.

The `ip_range` block supports:

* `id` - ID.
* `start_ip` - Start of IP range.
* `end_ip` - End of IP range.

The `options` block supports:

* `id` - ID.
* `code` - DHCP option code.
* `type` - DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.
* `value` - DHCP option value.
* `ip` - DHCP option IPs.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{switch_id}}.

## Import

SwitchController ManagedSwitch can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_managedswitch.labelname {{switch_id}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_managedswitch.labelname {{switch_id}}
$ unset "FORTIOS_IMPORT_TABLE"
```
