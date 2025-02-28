---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_interface"
description: |-
  Configure interfaces.
---

# fortios_system_interface
Configure interfaces.

## Example Usage

```hcl
resource "fortios_system_interface" "trname" {
  algorithm    = "L4"
  defaultgw    = "enable"
  distance     = 5
  ip           = "0.0.0.0 0.0.0.0"
  mtu          = 1500
  mtu_override = "disable"
  name         = "port3"
  type         = "physical"
  vdom         = "root"
  mode         = "dhcp"
  snmp_index   = 3
  description  = "Created by Terraform Provider for FortiOS"
  ipv6 {
    nd_mode = "basic"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `vdom` - (Required) Interface is in this virtual domain (VDOM).
* `vrf` - Virtual Routing Forwarding ID.
* `cli_conn_status` - CLI connection status.
* `fortilink` - Enable FortiLink to dedicate this interface to manage other Fortinet devices. Valid values: `enable`, `disable`.
* `switch_controller_source_ip` - Source IP address used in FortiLink over L3 connections. Valid values: `outbound`, `fixed`.
* `mode` - Addressing mode (static, DHCP, PPPoE). Valid values: `static`, `dhcp`, `pppoe`.
* `client_options` - DHCP client options. The structure of `client_options` block is documented below.
* `distance` - Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
* `priority` - Priority of learned routes.
* `dhcp_relay_interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `dhcp_relay_interface` - Specify outgoing interface to reach server.
* `dhcp_relay_vrf_select` - VRF ID used for connection to server. Set to -1 means unset this variable. CLI output may have different value on different FortiOS version.
* `dhcp_broadcast_flag` - Enable/disable setting of the broadcast flag in messages sent by the DHCP client (default = enable). Valid values: `disable`, `enable`.
* `dhcp_relay_service` - Enable/disable allowing this interface to act as a DHCP relay. Valid values: `disable`, `enable`.
* `dhcp_relay_ip` - DHCP relay IP address.
* `dhcp_relay_source_ip` - IP address used by the DHCP relay as its source IP.
* `dhcp_relay_circuit_id` - DHCP relay circuit ID.
* `dhcp_relay_link_selection` - DHCP relay link selection.
* `dhcp_relay_request_all_server` - Enable/disable sending DHCP request to all servers. Valid values: `disable`, `enable`.
* `dhcp_relay_allow_no_end_option` - Enable/disable relaying DHCP messages with no end option. Valid values: `disable`, `enable`.
* `dhcp_relay_type` - DHCP relay type (regular or IPsec). Valid values: `regular`, `ipsec`.
* `dhcp_smart_relay` - Enable/disable DHCP smart relay. Valid values: `disable`, `enable`.
* `dhcp_relay_agent_option` - Enable/disable DHCP relay agent option. Valid values: `enable`, `disable`.
* `dhcp_classless_route_addition` - Enable/disable addition of classless static routes retrieved from DHCP server. Valid values: `enable`, `disable`.
* `management_ip` - High Availability in-band management IP address of this interface.
* `ip` - Interface IPv4 address and subnet mask, syntax: X.X.X.X/24.
* `allowaccess` - Permitted types of management access to this interface.
* `gwdetect` - Enable/disable detect gateway alive for first. Valid values: `enable`, `disable`.
* `ping_serv_status` - PING server status.
* `detectserver` - Gateway's ping server for this IP.
* `detectprotocol` - Protocols used to detect the server. Valid values: `ping`, `tcp-echo`, `udp-echo`.
* `ha_priority` - HA election priority for the PING server.
* `fail_detect` - Enable/disable fail detection features for this interface. Valid values: `enable`, `disable`.
* `fail_detect_option` - Options for detecting that this interface has failed. Valid values: `detectserver`, `link-down`.
* `fail_alert_method` - Select link-failed-signal or link-down method to alert about a failed link. Valid values: `link-failed-signal`, `link-down`.
* `fail_action_on_extender` - Action on extender when interface fail . Valid values: `soft-restart`, `hard-restart`, `reboot`.
* `fail_alert_interfaces` - Names of the FortiGate interfaces from which the link failure alert is sent for this interface. The structure of `fail_alert_interfaces` block is documented below.
* `dhcp_client_identifier` - DHCP client identifier.
* `dhcp_renew_time` - DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.
* `ipunnumbered` - Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.
* `username` - Username of the PPPoE account, provided by your ISP.
* `pppoe_egress_cos` - CoS in VLAN tag for outgoing PPPoE/PPP packets. Valid values: `cos0`, `cos1`, `cos2`, `cos3`, `cos4`, `cos5`, `cos6`, `cos7`.
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation. Valid values: `enable`, `disable`.
* `password` - PPPoE account's password.
* `idle_timeout` - PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.
* `detected_peer_mtu` - MTU of detected peer (0 - 4294967295).
* `disc_retry_timeout` - Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.
* `padt_retry_timeout` - PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.
* `service_name` - PPPoE service name.
* `ac_name` - PPPoE server name.
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect.
* `defaultgw` - Enable to get the gateway IP from the DHCP or PPPoE server. Valid values: `enable`, `disable`.
* `dns_server_override` - Enable/disable use DNS acquired by DHCP or PPPoE. Valid values: `enable`, `disable`.
* `dns_server_protocol` - DNS transport protocols. Valid values: `cleartext`, `dot`, `doh`.
* `auth_type` - PPP authentication type to use. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
* `pptp_client` - Enable/disable PPTP client. Valid values: `enable`, `disable`.
* `pptp_user` - PPTP user name.
* `pptp_password` - PPTP password.
* `pptp_server_ip` - PPTP server IP address.
* `pptp_auth_type` - PPTP authentication type. Valid values: `auto`, `pap`, `chap`, `mschapv1`, `mschapv2`.
* `pptp_timeout` - Idle timer in minutes (0 for disabled).
* `arpforward` - Enable/disable ARP forwarding. Valid values: `enable`, `disable`.
* `ndiscforward` - Enable/disable NDISC forwarding. Valid values: `enable`, `disable`.
* `broadcast_forward` - Enable/disable broadcast forwarding. Valid values: `enable`, `disable`.
* `bfd` - Bidirectional Forwarding Detection (BFD) settings. Valid values: `global`, `enable`, `disable`.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval.
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval.
* `l2forward` - Enable/disable l2 forwarding. Valid values: `enable`, `disable`.
* `icmp_send_redirect` - Enable/disable ICMP send redirect. Valid values: `enable`, `disable`.
* `icmp_accept_redirect` - Enable/disable ICMP accept redirect. Valid values: `enable`, `disable`.
* `reachable_time` - IPv4 reachable time in milliseconds (30000 - 3600000, default = 30000).
* `vlanforward` - Enable/disable traffic forwarding between VLANs on this interface. Valid values: `enable`, `disable`.
* `stpforward` - Enable/disable STP forwarding. Valid values: `enable`, `disable`.
* `stpforward_mode` - Configure STP forwarding mode. Valid values: `rpl-all-ext-id`, `rpl-bridge-ext-id`, `rpl-nothing`.
* `ips_sniffer_mode` - Enable/disable the use of this interface as a one-armed sniffer. Valid values: `enable`, `disable`.
* `ident_accept` - Enable/disable authentication for this interface. Valid values: `enable`, `disable`.
* `ipmac` - Enable/disable IP/MAC binding. Valid values: `enable`, `disable`.
* `subst` - Enable to always send packets from this interface to a destination MAC address. Valid values: `enable`, `disable`.
* `macaddr` - Change the interface's MAC address.
* `virtual_mac` - Change the interface's virtual MAC address.
* `substitute_dst_mac` - Destination MAC address that all packets are sent to from this interface.
* `speed` - Interface speed. The default setting and the options available depend on the interface hardware.
* `status` - Bring the interface up or shut the interface down. Valid values: `up`, `down`.
* `netbios_forward` - Enable/disable NETBIOS forwarding. Valid values: `disable`, `enable`.
* `wins_ip` - WINS server IP.
* `type` - Interface type.
* `dedicated_to` - Configure interface for single purpose. Valid values: `none`, `management`.
* `trust_ip_1` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_2` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_3` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip6_1` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_2` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_3` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `mtu_override` - Enable to set a custom MTU for this interface. Valid values: `enable`, `disable`.
* `mtu` - MTU value for this interface.
* `ring_rx` - RX ring size.
* `ring_tx` - TX ring size.
* `wccp` - Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers. Valid values: `enable`, `disable`.
* `netflow_sampler` - Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both). Valid values: `disable`, `tx`, `rx`, `both`.
* `netflow_sample_rate` - NetFlow sample rate.  Sample one packet every configured number of packets
(1 - 65535, default = 1, which means standard NetFlow where all packets are sampled).
* `netflow_sampler_id` - Netflow sampler ID.
* `sflow_sampler` - Enable/disable sFlow on this interface. Valid values: `enable`, `disable`.
* `drop_overlapped_fragment` - Enable/disable drop overlapped fragment packets. Valid values: `enable`, `disable`.
* `drop_fragment` - Enable/disable drop fragment packets. Valid values: `enable`, `disable`.
* `scan_botnet_connections` - Enable monitoring or blocking connections to Botnet servers through this interface. Valid values: `disable`, `block`, `monitor`.
* `src_check` - Enable/disable source IP check. Valid values: `enable`, `disable`.
* `sample_rate` - sFlow sample rate (10 - 99999).
* `polling_interval` - sFlow polling interval in seconds (1 - 255).
* `sample_direction` - Data that NetFlow collects (rx, tx, or both). Valid values: `tx`, `rx`, `both`.
* `explicit_web_proxy` - Enable/disable the explicit web proxy on this interface. Valid values: `enable`, `disable`.
* `explicit_ftp_proxy` - Enable/disable the explicit FTP proxy on this interface. Valid values: `enable`, `disable`.
* `proxy_captive_portal` - Enable/disable proxy captive portal on this interface. Valid values: `enable`, `disable`.
* `tcp_mss` - TCP maximum segment size. 0 means do not change segment size.
* `mediatype` - Select SFP media interface type Valid values: `none`, `gmii`, `sgmii`, `sr`, `lr`, `cr`, `sr4`, `lr4`, `cr4`.
* `inbandwidth` - Bandwidth limit for incoming traffic, 0 means unlimited. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: 0 - 16776000 kbps. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.17, >= 7.2.1: 0 - 80000000 kbps.
* `outbandwidth` - Bandwidth limit for outgoing traffic, 0 means unlimited. On FortiOS versions 6.2.0-6.2.6: 0 - 16776000 kbps. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.17, >= 7.2.1: 0 - 80000000 kbps.
* `egress_shaping_profile` - Outgoing traffic shaping profile.
* `ingress_shaping_profile` - Incoming traffic shaping profile.
* `disconnect_threshold` - Time in milliseconds to wait before sending a notification that this interface is down or disconnected.
* `spillover_threshold` - Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `ingress_spillover_threshold` - Ingress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `weight` - Default weight for static routes (if route has no weight configured).
* `interface` - Interface name.
* `external` - Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet). Valid values: `enable`, `disable`.
* `vlan_protocol` - Ethernet protocol of VLAN. Valid values: `8021q`, `8021ad`.
* `vlanid` - VLAN ID (1 - 4094).
* `trunk` - Enable/disable VLAN trunk. Valid values: `enable`, `disable`.
* `forward_domain` - Transparent mode forward domain.
* `remote_ip` - Remote IP address of tunnel.
* `member` - Physical interfaces that belong to the aggregate or redundant interface. The structure of `member` block is documented below.
* `lacp_mode` - LACP mode. Valid values: `static`, `passive`, `active`.
* `lacp_ha_secondary` - LACP HA secondary member. Valid values: `enable`, `disable`.
* `lacp_ha_slave` - LACP HA slave. Valid values: `enable`, `disable`.
* `system_id_type` - Method in which system ID is generated. Valid values: `auto`, `user`.
* `system_id` - Define a system ID for the aggregate interface.
* `lacp_speed` - How often the interface sends LACP messages. Valid values: `slow`, `fast`.
* `min_links` - Minimum number of aggregated ports that must be up.
* `min_links_down` - Action to take when less than the configured minimum number of links are active. Valid values: `operational`, `administrative`.
* `algorithm` - Frame distribution algorithm.
* `link_up_delay` - Number of milliseconds to wait before considering a link is up.
* `aggregate_type` - Type of aggregation. Valid values: `physical`, `vxlan`.
* `priority_override` - Enable/disable fail back to higher priority port once recovered. Valid values: `enable`, `disable`.
* `aggregate` - Aggregate interface.
* `redundant_interface` - Redundant interface.
* `managed_device` - Available when FortiLink is enabled, used for managed devices through FortiLink interface. The structure of `managed_device` block is documented below.
* `devindex` - Device Index.
* `vindex` - Switch control interface VLAN ID.
* `switch` - Contained in switch.
* `description` - Description.
* `alias` - Alias will be displayed with the interface name to make it easier to distinguish.
* `security_mode` - Turn on captive portal authentication for this interface. Valid values: `none`, `captive-portal`, `802.1X`.
* `captive_portal` - Enable/disable captive portal.
* `security_mac_auth_bypass` - Enable/disable MAC authentication bypass. Valid values: `mac-auth-only`, `enable`, `disable`.
* `security_ip_auth_bypass` - Enable/disable IP authentication bypass. Valid values: `enable`, `disable`.
* `security_external_web` - URL of external authentication web server.
* `security_external_logout` - URL of external authentication logout server.
* `replacemsg_override_group` - Replacement message override group.
* `security_redirect_url` - URL redirection after disclaimer/authentication.
* `auth_cert` - HTTPS server certificate.
* `auth_portal_addr` - Address of captive portal.
* `security_exempt_list` - Name of security-exempt-list.
* `security_groups` - User groups that can authenticate with the captive portal. The structure of `security_groups` block is documented below.
* `ike_saml_server` - Configure IKE authentication SAML server.
* `stp` - Enable/disable STP. Valid values: `disable`, `enable`.
* `stp_ha_secondary` - Control STP behaviour on HA secondary. Valid values: `disable`, `enable`, `priority-adjust`.
* `device_identification` - Enable/disable passively gathering of device identity information about the devices on the network connected to this interface. Valid values: `enable`, `disable`.
* `exclude_signatures` - Exclude IOT or OT application signatures. Valid values: `iot`, `ot`.
* `device_user_identification` - Enable/disable passive gathering of user identity information about users on this interface. Valid values: `enable`, `disable`.
* `device_identification_active_scan` - Enable/disable active gathering of device identity information about the devices on the network connected to this interface. Valid values: `enable`, `disable`.
* `device_access_list` - Device access list.
* `device_netscan` - Enable/disable inclusion of devices detected on this interface in network vulnerability scans. Valid values: `disable`, `enable`.
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception. Valid values: `enable`, `disable`, `vdom`.
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission. Valid values: `enable`, `disable`, `vdom`.
* `lldp_network_policy` - LLDP-MED network policy profile.
* `fortiheartbeat` - Enable/disable FortiHeartBeat (FortiTelemetry on GUI). Valid values: `enable`, `disable`.
* `broadcast_forticlient_discovery` - Enable/disable broadcasting FortiClient discovery messages. Valid values: `enable`, `disable`.
* `endpoint_compliance` - Enable/disable endpoint compliance enforcement. Valid values: `enable`, `disable`.
* `estimated_upstream_bandwidth` - Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.
* `estimated_downstream_bandwidth` - Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.
* `measured_upstream_bandwidth` - Measured upstream bandwidth (kbps). 
* `measured_downstream_bandwidth` - Measured downstream bandwidth (kbps). 
* `bandwidth_measure_time` - Bandwidth measure time 
* `monitor_bandwidth` - Enable monitoring bandwidth on this interface. Valid values: `enable`, `disable`.
* `vrrp_virtual_mac` - Enable/disable use of virtual MAC for VRRP. Valid values: `enable`, `disable`.
* `vrrp` - VRRP configuration. The structure of `vrrp` block is documented below.
* `phy_setting` - PHY settings The structure of `phy_setting` block is documented below.
* `role` - Interface role. Valid values: `lan`, `wan`, `dmz`, `undefined`.
* `snmp_index` - Permanent SNMP Index of the interface.
* `secondary_ip` - Enable/disable adding a secondary IP to this interface. Valid values: `enable`, `disable`.
* `secondaryip` - Second IP address of interface. The structure of `secondaryip` block is documented below.
* `preserve_session_route` - Enable/disable preservation of session route when dirty. Valid values: `enable`, `disable`.
* `auto_auth_extension_device` - Enable/disable automatic authorization of dedicated Fortinet extension device on this interface. Valid values: `enable`, `disable`.
* `ap_discover` - Enable/disable automatic registration of unknown FortiAP devices. Valid values: `enable`, `disable`.
* `fortilink_stacking` - Enable/disable FortiLink switch-stacking on this interface. Valid values: `enable`, `disable`.
* `fortilink_neighbor_detect` - Protocol for FortiGate neighbor discovery. Valid values: `lldp`, `fortilink`.
* `ip_managed_by_fortiipam` - Enable/disable automatic IP address assignment of this interface by FortiIPAM.
* `managed_subnetwork_size` - Number of IP addresses to be allocated by FortiIPAM and used by this FortiGate unit's DHCP server settings.
* `fortilink_split_interface` - Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy. Valid values: `enable`, `disable`.
* `internal` - Implicitly created.
* `fortilink_backup_link` - fortilink split interface backup link.
* `switch_controller_access_vlan` - Block FortiSwitch port-to-port traffic. Valid values: `enable`, `disable`.
* `switch_controller_traffic_policy` - Switch controller traffic policy for the VLAN.
* `switch_controller_rspan_mode` - Stop Layer2 MAC learning and interception of BPDUs and other packets on this interface. Valid values: `disable`, `enable`.
* `switch_controller_netflow_collect` - NetFlow collection and processing. Valid values: `disable`, `enable`.
* `switch_controller_mgmt_vlan` - VLAN to use for FortiLink management purposes.
* `switch_controller_igmp_snooping` - Switch controller IGMP snooping. Valid values: `enable`, `disable`.
* `switch_controller_igmp_snooping_proxy` - Switch controller IGMP snooping proxy. Valid values: `enable`, `disable`.
* `switch_controller_igmp_snooping_fast_leave` - Switch controller IGMP snooping fast-leave. Valid values: `enable`, `disable`.
* `switch_controller_dhcp_snooping` - Switch controller DHCP snooping. Valid values: `enable`, `disable`.
* `switch_controller_dhcp_snooping_verify_mac` - Switch controller DHCP snooping verify MAC. Valid values: `enable`, `disable`.
* `switch_controller_dhcp_snooping_option82` - Switch controller DHCP snooping option82. Valid values: `enable`, `disable`.
* `dhcp_snooping_server_list` - Configure DHCP server access list. The structure of `dhcp_snooping_server_list` block is documented below.
* `switch_controller_arp_inspection` - Enable/disable FortiSwitch ARP inspection.
* `switch_controller_learning_limit` - Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).
* `switch_controller_nac` - Integrated NAC settings for managed FortiSwitch.
* `switch_controller_dynamic` - Integrated FortiLink settings for managed FortiSwitch.
* `switch_controller_feature` - Interface's purpose when assigning traffic (read only).
* `switch_controller_iot_scanning` - Enable/disable managed FortiSwitch IoT scanning. Valid values: `enable`, `disable`.
* `switch_controller_offload` - Enable/disable managed FortiSwitch routing offload. Valid values: `enable`, `disable`.
* `switch_controller_offload_ip` - IP for routing offload on FortiSwitch.
* `switch_controller_offload_gw` - Enable/disable managed FortiSwitch routing offload gateway. Valid values: `enable`, `disable`.
* `swc_vlan` - Creation status for switch-controller VLANs.
* `swc_first_create` - Initial create for switch-controller VLANs.
* `color` - Color of icon on the GUI.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `eap_supplicant` - Enable/disable EAP-Supplicant. Valid values: `enable`, `disable`.
* `eap_method` - EAP method. Valid values: `tls`, `peap`.
* `eap_identity` - EAP identity.
* `eap_password` - EAP password.
* `eap_ca_cert` - EAP CA certificate name.
* `eap_user_cert` - EAP user certificate name.
* `default_purdue_level` - default purdue level of device detected on this interface. Valid values: `1`, `1.5`, `2`, `2.5`, `3`, `3.5`, `4`, `5`, `5.5`.
* `forward_error_correction` - Configure forward error correction (FEC). Valid values: `none`, `disable`, `cl91-rs-fec`, `cl74-fc-fec`.
* `ipv6` - IPv6 of interface. The structure of `ipv6` block is documented below.
* `autogenerated` - Indicates whether the interface is automatically created by FortiGate, for example, created during the VPN creation process. If it is, set it to "auto", else keep it empty.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `client_options` block supports:

* `id` - ID.
* `code` - DHCP client option code.
* `type` - DHCP client option type. Valid values: `hex`, `string`, `ip`, `fqdn`.
* `value` - DHCP client option value.
* `ip` - DHCP option IPs.

The `fail_alert_interfaces` block supports:

* `name` - Names of the physical interfaces belonging to the aggregate or redundant interface.

The `member` block supports:

* `interface_name` - Physical interface name.

The `managed_device` block supports:

* `name` - Managed dev identifier.

The `security_groups` block supports:

* `name` - Names of user groups that can authenticate with the captive portal.

The `vrrp` block supports:

* `vrid` - Virtual router identifier (1 - 255).
* `version` - VRRP version. Valid values: `2`, `3`.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrip` - IP address of the virtual router.
* `priority` - Priority of the virtual router (1 - 255).
* `adv_interval` - Advertisement interval. On FortiOS versions 6.2.0-7.4.7: 1 - 255 seconds. On FortiOS versions >= 7.6.0: 250 - 255000 milliseconds.
* `start_time` - Startup time (1 - 255 seconds).
* `preempt` - Enable/disable preempt mode. Valid values: `enable`, `disable`.
* `accept_mode` - Enable/disable accept mode. Valid values: `enable`, `disable`.
* `vrdst` - Monitor the route to this destination.
* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination. Valid values: `enable`, `disable`.
* `status` - Enable/disable this VRRP configuration. Valid values: `enable`, `disable`.
* `proxy_arp` - VRRP Proxy ARP configuration. The structure of `proxy_arp` block is documented below.

The `proxy_arp` block supports:

* `id` - ID.
* `ip` - Set IP addresses of proxy ARP.

The `phy_setting` block supports:

* `signal_ok_threshold_value` - Signal-ok-threshold value(0 - 12).

The `secondaryip` block supports:

* `id` - ID.
* `ip` - Secondary IP address of the interface.
* `secip_relay_ip` - DHCP relay IP address.
* `allowaccess` - Management access settings for the secondary IP address.
* `gwdetect` - Enable/disable detect gateway alive for first. Valid values: `enable`, `disable`.
* `ping_serv_status` - PING server status.
* `detectserver` - Gateway's ping server for this IP.
* `detectprotocol` - Protocols used to detect the server. Valid values: `ping`, `tcp-echo`, `udp-echo`.
* `ha_priority` - HA election priority for the PING server.

The `dhcp_snooping_server_list` block supports:

* `name` - DHCP server name.
* `server_ip` - IP address for DHCP server.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block supports:

* `name` - Tag name.

The `ipv6` block supports:

* `ip6_mode` - Addressing mode (static, DHCP, delegated). Valid values: `static`, `dhcp`, `pppoe`, `delegated`.
* `client_options` - DHCP6 client options. The structure of `client_options` block is documented below.
* `nd_mode` - Neighbor discovery mode. Valid values: `basic`, `SEND-compatible`.
* `nd_cert` - Neighbor discovery certificate.
* `nd_security_level` - Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
* `nd_timestamp_delta` - Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
* `nd_timestamp_fuzz` - Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
* `nd_cga_modifier` - Neighbor discovery CGA modifier.
* `ip6_dns_server_override` - Enable/disable using the DNS server acquired by DHCP. Valid values: `enable`, `disable`.
* `ip6_address` - Primary IPv6 address prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_extra_addr` - Extra IPv6 address prefixes of interface. The structure of `ip6_extra_addr` block is documented below.
* `ip6_allowaccess` - Allow management access to the interface.
* `ip6_send_adv` - Enable/disable sending advertisements about the interface. Valid values: `enable`, `disable`.
* `icmp6_send_redirect` - Enable/disable sending of ICMPv6 redirects. Valid values: `enable`, `disable`.
* `ip6_manage_flag` - Enable/disable the managed flag. Valid values: `enable`, `disable`.
* `ip6_other_flag` - Enable/disable the other IPv6 flag. Valid values: `enable`, `disable`.
* `ip6_max_interval` - IPv6 maximum interval (4 to 1800 sec).
* `ip6_min_interval` - IPv6 minimum interval (3 to 1350 sec).
* `ip6_link_mtu` - IPv6 link MTU.
* `ra_send_mtu` - Enable/disable sending link MTU in RA packet. Valid values: `enable`, `disable`.
* `ip6_reachable_time` - IPv6 reachable time (milliseconds; 0 means unspecified).
* `ip6_retrans_time` - IPv6 retransmit time (milliseconds; 0 means unspecified).
* `ip6_default_life` - Default life (sec).
* `ip6_hop_limit` - Hop limit (0 means unspecified).
* `ip6_adv_rio` - Enable/disable sending advertisements with route information option. Valid values: `enable`, `disable`.
* `ip6_route_pref` - Set route preference to the interface (default = medium). Valid values: `medium`, `high`, `low`.
* `ip6_route_list` - Advertised route list. The structure of `ip6_route_list` block is documented below.
* `autoconf` - Enable/disable address auto config. Valid values: `enable`, `disable`.
* `unique_autoconf_addr` - Enable/disable unique auto config address. Valid values: `enable`, `disable`.
* `interface_identifier` - IPv6 interface identifier.
* `ip6_prefix_mode` - Assigning a prefix from DHCP or RA. Valid values: `dhcp6`, `ra`.
* `ip6_upstream_interface` - Interface name providing delegated information.
* `ip6_delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `ip6_subnet` -  Subnet to routing prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_prefix_list` - Advertised prefix list. The structure of `ip6_prefix_list` block is documented below.
* `ip6_rdnss_list` - Advertised IPv6 RDNSS list. The structure of `ip6_rdnss_list` block is documented below.
* `ip6_dnssl_list` - Advertised IPv6 DNSS list. The structure of `ip6_dnssl_list` block is documented below.
* `ip6_delegated_prefix_list` - Advertised IPv6 delegated prefix list. The structure of `ip6_delegated_prefix_list` block is documented below.
* `dhcp6_relay_service` - Enable/disable DHCPv6 relay. Valid values: `disable`, `enable`.
* `dhcp6_relay_type` - DHCPv6 relay type. Valid values: `regular`.
* `dhcp6_relay_source_interface` - Enable/disable use of address on this interface as the source address of the relay message. Valid values: `disable`, `enable`.
* `dhcp6_relay_ip` - DHCPv6 relay IP address.
* `dhcp6_relay_source_ip` - IPv6 address used by the DHCP6 relay as its source IP.
* `dhcp6_relay_interface_id` - DHCP6 relay interface ID.
* `dhcp6_client_options` - DHCPv6 client options. Valid values: `rapid`, `iapd`, `iana`.
* `dhcp6_prefix_delegation` - Enable/disable DHCPv6 prefix delegation. Valid values: `enable`, `disable`.
* `dhcp6_information_request` - Enable/disable DHCPv6 information request. Valid values: `enable`, `disable`.
* `dhcp6_iapd_list` - DHCPv6 IA-PD list The structure of `dhcp6_iapd_list` block is documented below.
* `dhcp6_prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `dhcp6_prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `dhcp6_prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `cli_conn6_status` - CLI IPv6 connection status.
* `vrrp_virtual_mac6` - Enable/disable virtual MAC for VRRP. Valid values: `enable`, `disable`.
* `vrip6_link_local` - Link-local IPv6 address of virtual router.
* `vrrp6` - IPv6 VRRP configuration. The structure of `vrrp6` block is documented below.

The `client_options` block supports:

* `id` - ID.
* `code` - DHCPv6 option code.
* `type` - DHCPv6 option type. Valid values: `hex`, `string`, `ip6`, `fqdn`.
* `value` - DHCPv6 option value (hexadecimal value must be even).
* `ip6` - DHCP option IP6s.

The `ip6_extra_addr` block supports:

* `prefix` - IPv6 address prefix.

The `ip6_route_list` block supports:

* `route` - IPv6 route.
* `route_pref` - Set route preference to the interface (default = medium). Valid values: `medium`, `high`, `low`.
* `route_life_time` - Route life time in seconds (0 - 65535, default = 1800).

The `ip6_prefix_list` block supports:

* `prefix` - IPv6 prefix.
* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `enable`, `disable`.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `enable`, `disable`.
* `valid_life_time` - Valid life time (sec).
* `preferred_life_time` - Preferred life time (sec).
* `rdnss` - Recursive DNS server option.
* `dnssl` - DNS search list option. The structure of `dnssl` block is documented below.

The `dnssl` block supports:

* `domain` - Domain name.

The `ip6_rdnss_list` block supports:

* `rdnss` - Recursive DNS server option.
* `rdnss_life_time` - Recursive DNS server life time in seconds (0 - 4294967295, default = 1800).

The `ip6_dnssl_list` block supports:

* `domain` - Domain name.
* `dnssl_life_time` - DNS search list time in seconds (0 - 4294967295, default = 1800).

The `ip6_delegated_prefix_list` block supports:

* `prefix_id` - Prefix ID.
* `upstream_interface` - Name of the interface that provides delegated information.
* `delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `enable`, `disable`.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `enable`, `disable`.
* `subnet` -  Add subnet ID to routing prefix.
* `rdnss_service` - Recursive DNS service option. Valid values: `delegated`, `default`, `specify`.
* `rdnss` - Recursive DNS server option.

The `dhcp6_iapd_list` block supports:

* `iaid` - Identity association identifier.
* `prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).

The `vrrp6` block supports:

* `vrid` - Virtual router identifier (1 - 255).
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrip6` - IPv6 address of the virtual router.
* `priority` - Priority of the virtual router (1 - 255).
* `adv_interval` - Advertisement interval. On FortiOS versions 6.2.0-7.4.7: 1 - 255 seconds. On FortiOS versions >= 7.6.0: 250 - 255000 milliseconds.
* `start_time` - Startup time (1 - 255 seconds).
* `preempt` - Enable/disable preempt mode. Valid values: `enable`, `disable`.
* `accept_mode` - Enable/disable accept mode. Valid values: `enable`, `disable`.
* `vrdst6` - Monitor the route to this destination.
* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination. Valid values: `enable`, `disable`.
* `status` - Enable/disable VRRP. Valid values: `enable`, `disable`.

-> If you want to modify the configuration of the interface which is automatically created by FortiGate (for example, the interface created during the VPN creation process), please set the `autogenerated` argument to "auto".


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Interface can be imported using any of these accepted formats:
```
$ terraform import fortios_system_interface.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_interface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
