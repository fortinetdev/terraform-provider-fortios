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
* `fortilink` - Enable FortiLink to dedicate this interface to manage other Fortinet devices.
* `mode` - Addressing mode (static, DHCP, PPPoE).
* `distance` - Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
* `priority` - Priority of learned routes.
* `dhcp_relay_service` - Enable/disable allowing this interface to act as a DHCP relay.
* `dhcp_relay_ip` - DHCP relay IP address.
* `dhcp_relay_type` - DHCP relay type (regular or IPsec).
* `dhcp_relay_agent_option` - Enable/disable DHCP relay agent option.
* `management_ip` - High Availability in-band management IP address of this interface.
* `ip` - Interface IPv4 address and subnet mask, syntax: X.X.X.X/24.
* `allowaccess` - Permitted types of management access to this interface.
* `gwdetect` - Enable/disable detect gateway alive for first.
* `ping_serv_status` - PING server status.
* `detectserver` - Gateway's ping server for this IP.
* `detectprotocol` - Protocols used to detect the server.
* `ha_priority` - HA election priority for the PING server.
* `fail_detect` - Enable/disable fail detection features for this interface.
* `fail_detect_option` - Options for detecting that this interface has failed.
* `fail_alert_method` - Select link-failed-signal or link-down method to alert about a failed link.
* `fail_action_on_extender` - Action on extender when interface fail .
* `fail_alert_interfaces` - Names of the FortiGate interfaces from which the link failure alert is sent for this interface.
* `dhcp_client_identifier` - DHCP client identifier.
* `dhcp_renew_time` - DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.
* `ipunnumbered` - Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.
* `username` - Username of the PPPoE account, provided by your ISP.
* `pppoe_unnumbered_negotiate` - Enable/disable PPPoE unnumbered negotiation.
* `password` - PPPoE account's password.
* `idle_timeout` - PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.
* `detected_peer_mtu` - MTU of detected peer (0 - 4294967295).
* `disc_retry_timeout` - Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.
* `padt_retry_timeout` - PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.
* `service_name` - PPPoE service name.
* `ac_name` - PPPoE server name.
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum missed LCP echo messages before disconnect.
* `defaultgw` - Enable to get the gateway IP from the DHCP or PPPoE server.
* `dns_server_override` - Enable/disable use DNS acquired by DHCP or PPPoE.
* `auth_type` - PPP authentication type to use.
* `pptp_client` - Enable/disable PPTP client.
* `pptp_user` - PPTP user name.
* `pptp_password` - PPTP password.
* `pptp_server_ip` - PPTP server IP address.
* `pptp_auth_type` - PPTP authentication type.
* `pptp_timeout` - Idle timer in minutes (0 for disabled).
* `arpforward` - Enable/disable ARP forwarding.
* `ndiscforward` - Enable/disable NDISC forwarding.
* `broadcast_forward` - Enable/disable broadcast forwarding.
* `bfd` - Bidirectional Forwarding Detection (BFD) settings.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval.
* `bfd_detect_mult` - BFD detection multiplier.
* `bfd_required_min_rx` - BFD required minimal receive interval.
* `l2forward` - Enable/disable l2 forwarding.
* `icmp_send_redirect` - Enable/disable ICMP send redirect.
* `icmp_accept_redirect` - Enable/disable ICMP accept redirect.
* `vlanforward` - Enable/disable traffic forwarding between VLANs on this interface.
* `stpforward` - Enable/disable STP forwarding.
* `stpforward_mode` - Configure STP forwarding mode.
* `ips_sniffer_mode` - Enable/disable the use of this interface as a one-armed sniffer.
* `ident_accept` - Enable/disable authentication for this interface.
* `ipmac` - Enable/disable IP/MAC binding.
* `subst` - Enable to always send packets from this interface to a destination MAC address.
* `macaddr` - Change the interface's MAC address.
* `substitute_dst_mac` - Destination MAC address that all packets are sent to from this interface.
* `speed` - Interface speed. The default setting and the options available depend on the interface hardware.
* `status` - Bring the interface up or shut the interface down.
* `netbios_forward` - Enable/disable NETBIOS forwarding.
* `wins_ip` - WINS server IP.
* `type` - Interface type.
* `dedicated_to` - Configure interface for single purpose.
* `trust_ip_1` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_2` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip_3` - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
* `trust_ip6_1` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_2` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `trust_ip6_3` - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
* `mtu_override` - Enable to set a custom MTU for this interface.
* `mtu` - MTU value for this interface.
* `wccp` - Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers.
* `netflow_sampler` - Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both).
* `sflow_sampler` - Enable/disable sFlow on this interface.
* `drop_overlapped_fragment` - Enable/disable drop overlapped fragment packets.
* `drop_fragment` - Enable/disable drop fragment packets.
* `scan_botnet_connections` - Enable monitoring or blocking connections to Botnet servers through this interface.
* `src_check` - Enable/disable source IP check.
* `sample_rate` - sFlow sample rate (10 - 99999).
* `polling_interval` - sFlow polling interval (1 - 255 sec).
* `sample_direction` - Data that NetFlow collects (rx, tx, or both).
* `explicit_web_proxy` - Enable/disable the explicit web proxy on this interface.
* `explicit_ftp_proxy` - Enable/disable the explicit FTP proxy on this interface.
* `proxy_captive_portal` - Enable/disable proxy captive portal on this interface.
* `tcp_mss` - TCP maximum segment size. 0 means do not change segment size.
* `inbandwidth` - Bandwidth limit for incoming traffic (0 - 16776000 kbps), 0 means unlimited.
* `outbandwidth` - Bandwidth limit for outgoing traffic (0 - 16776000 kbps).
* `egress_shaping_profile` - Outgoing traffic shaping profile.
* `disconnect_threshold` - Time in milliseconds to wait before sending a notification that this interface is down or disconnected.
* `spillover_threshold` - Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
* `ingress_spillover_threshold` - Ingress Spillover threshold (0 - 16776000 kbps).
* `weight` - Default weight for static routes (if route has no weight configured).
* `interface` - Interface name.
* `external` - Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet).
* `vlanid` - VLAN ID (1 - 4094).
* `forward_domain` - Transparent mode forward domain.
* `remote_ip` - Remote IP address of tunnel.
* `member` - Physical interfaces that belong to the aggregate or redundant interface.
* `lacp_mode` - LACP mode.
* `lacp_ha_slave` - LACP HA slave.
* `lacp_speed` - How often the interface sends LACP messages.
* `min_links` - Minimum number of aggregated ports that must be up.
* `min_links_down` - Action to take when less than the configured minimum number of links are active.
* `algorithm` - Frame distribution algorithm.
* `link_up_delay` - Number of milliseconds to wait before considering a link is up.
* `priority_override` - Enable/disable fail back to higher priority port once recovered.
* `aggregate` - Aggregate interface.
* `redundant_interface` - Redundant interface.
* `managed_device` - Available when FortiLink is enabled, used for managed devices through FortiLink interface.
* `devindex` - Device Index.
* `vindex` - Switch control interface VLAN ID.
* `switch` - Contained in switch.
* `description` - Description.
* `alias` - Alias will be displayed with the interface name to make it easier to distinguish.
* `security_mode` - Turn on captive portal authentication for this interface.
* `captive_portal` - Enable/disable captive portal.
* `security_mac_auth_bypass` - Enable/disable MAC authentication bypass.
* `security_external_web` - URL of external authentication web server.
* `security_external_logout` - URL of external authentication logout server.
* `replacemsg_override_group` - Replacement message override group.
* `security_redirect_url` - URL redirection after disclaimer/authentication.
* `security_exempt_list` - Name of security-exempt-list.
* `security_groups` - User groups that can authenticate with the captive portal.
* `device_identification` - Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
* `device_user_identification` - Enable/disable passive gathering of user identity information about users on this interface.
* `device_identification_active_scan` - Enable/disable active gathering of device identity information about the devices on the network connected to this interface.
* `device_access_list` - Device access list.
* `device_netscan` - Enable/disable inclusion of devices detected on this interface in network vulnerability scans.
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception.
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission.
* `lldp_network_policy` - LLDP-MED network policy profile.
* `fortiheartbeat` - Enable/disable FortiHeartBeat (FortiTelemetry on GUI).
* `broadcast_forticlient_discovery` - Enable/disable broadcasting FortiClient discovery messages.
* `endpoint_compliance` - Enable/disable endpoint compliance enforcement.
* `estimated_upstream_bandwidth` - Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.
* `estimated_downstream_bandwidth` - Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.
* `vrrp_virtual_mac` - Enable/disable use of virtual MAC for VRRP.
* `vrrp` - VRRP configuration.
* `role` - Interface role.
* `snmp_index` - Permanent SNMP Index of the interface.
* `secondary_IP` - Enable/disable adding a secondary IP to this interface.
* `secondaryip` - Second IP address of interface.
* `preserve_session_route` - Enable/disable preservation of session route when dirty.
* `auto_auth_extension_device` - Enable/disable automatic authorization of dedicated Fortinet extension device on this interface.
* `ap_discover` - Enable/disable automatic registration of unknown FortiAP devices.
* `fortilink_stacking` - Enable/disable FortiLink switch-stacking on this interface.
* `fortilink_split_interface` - Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy.
* `internal` - Implicitly created.
* `fortilink_backup_link` - fortilink split interface backup link.
* `switch_controller_access_vlan` - Block FortiSwitch port-to-port traffic.
* `switch_controller_traffic_policy` - Switch controller traffic policy for the VLAN.
* `switch_controller_igmp_snooping` - Switch controller IGMP snooping.
* `switch_controller_dhcp_snooping` - Switch controller DHCP snooping.
* `switch_controller_dhcp_snooping_verify_mac` - Switch controller DHCP snooping verify MAC.
* `switch_controller_dhcp_snooping_option82` - Switch controller DHCP snooping option82.
* `switch_controller_arp_inspection` - Enable/disable FortiSwitch ARP inspection.
* `switch_controller_learning_limit` - Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).
* `color` - Color of icon on the GUI.
* `tagging` - Config object tagging.
* `ipv6` - IPv6 of interface.

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
* `version` - VRRP version.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrip` - IP address of the virtual router.
* `priority` - Priority of the virtual router (1 - 255).
* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `start_time` - Startup time (1 - 255 seconds).
* `preempt` - Enable/disable preempt mode.
* `accept_mode` - Enable/disable accept mode.
* `vrdst` - Monitor the route to this destination.
* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `status` - Enable/disable this VRRP configuration.
* `proxy_arp` - VRRP Proxy ARP configuration.

The `proxy_arp` block supports:

* `id` - ID.
* `ip` - Set IP addresses of proxy ARP.

The `secondaryip` block supports:

* `id` - ID.
* `ip` - Secondary IP address of the interface.
* `allowaccess` - Management access settings for the secondary IP address.
* `gwdetect` - Enable/disable detect gateway alive for first.
* `ping_serv_status` - PING server status.
* `detectserver` - Gateway's ping server for this IP.
* `detectprotocol` - Protocols used to detect the server.
* `ha_priority` - HA election priority for the PING server.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags.

The `tags` block supports:

* `name` - Tag name.

The `ipv6` block supports:

* `ip6_mode` - Addressing mode (static, DHCP, delegated).
* `nd_mode` - Neighbor discovery mode.
* `nd_cert` - Neighbor discovery certificate.
* `nd_security_level` - Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
* `nd_timestamp_delta` - Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
* `nd_timestamp_fuzz` - Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
* `nd_cga_modifier` - Neighbor discovery CGA modifier.
* `ip6_dns_server_override` - Enable/disable using the DNS server acquired by DHCP.
* `ip6_address` - Primary IPv6 address prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_extra_addr` - Extra IPv6 address prefixes of interface.
* `ip6_allowaccess` - Allow management access to the interface.
* `ip6_send_adv` - Enable/disable sending advertisements about the interface.
* `ip6_manage_flag` - Enable/disable the managed flag.
* `ip6_other_flag` - Enable/disable the other IPv6 flag.
* `ip6_max_interval` - IPv6 maximum interval (4 to 1800 sec).
* `ip6_min_interval` - IPv6 minimum interval (3 to 1350 sec).
* `ip6_link_mtu` - IPv6 link MTU.
* `ip6_reachable_time` - IPv6 reachable time (milliseconds; 0 means unspecified).
* `ip6_retrans_time` - IPv6 retransmit time (milliseconds; 0 means unspecified).
* `ip6_default_life` - Default life (sec).
* `ip6_hop_limit` - Hop limit (0 means unspecified).
* `autoconf` - Enable/disable address auto config.
* `ip6_upstream_interface` - Interface name providing delegated information.
* `ip6_subnet` -  Subnet to routing prefix, syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx
* `ip6_prefix_list` - Advertised prefix list.
* `ip6_delegated_prefix_list` - Advertised IPv6 delegated prefix list.
* `dhcp6_relay_service` - Enable/disable DHCPv6 relay.
* `dhcp6_relay_type` - DHCPv6 relay type.
* `dhcp6_relay_ip` - DHCPv6 relay IP address.
* `dhcp6_client_options` - DHCPv6 client options.
* `dhcp6_prefix_delegation` - Enable/disable DHCPv6 prefix delegation.
* `dhcp6_information_request` - Enable/disable DHCPv6 information request.
* `dhcp6_prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `dhcp6_prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `dhcp6_prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).
* `vrrp_virtual_mac6` - Enable/disable virtual MAC for VRRP.
* `vrip6_link_local` - Link-local IPv6 address of virtual router.
* `vrrp6` - IPv6 VRRP configuration.

The `ip6_extra_addr` block supports:

* `prefix` - IPv6 address prefix.

The `ip6_prefix_list` block supports:

* `prefix` - IPv6 prefix.
* `autonomous_flag` - Enable/disable the autonomous flag.
* `onlink_flag` - Enable/disable the onlink flag.
* `valid_life_time` - Valid life time (sec).
* `preferred_life_time` - Preferred life time (sec).
* `rdnss` - Recursive DNS server option.
* `dnssl` - DNS search list option.

The `dnssl` block supports:

* `domain` - Domain name.

The `ip6_delegated_prefix_list` block supports:

* `prefix_id` - Prefix ID.
* `upstream_interface` - Name of the interface that provides delegated information.
* `autonomous_flag` - Enable/disable the autonomous flag.
* `onlink_flag` - Enable/disable the onlink flag.
* `subnet` -  Add subnet ID to routing prefix.
* `rdnss_service` - Recursive DNS service option.
* `rdnss` - Recursive DNS server option.

The `vrrp6` block supports:

* `vrid` - Virtual router identifier (1 - 255).
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrip6` - IPv6 address of the virtual router.
* `priority` - Priority of the virtual router (1 - 255).
* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `start_time` - Startup time (1 - 255 seconds).
* `preempt` - Enable/disable preempt mode.
* `accept_mode` - Enable/disable accept mode.
* `vrdst6` - Monitor the route to this destination.
* `status` - Enable/disable VRRP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Interface can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_interface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
