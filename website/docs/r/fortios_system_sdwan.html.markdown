---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan"
description: |-
  Configure redundant internet connections using SD-WAN (formerly virtual WAN link).
---

# fortios_system_sdwan
Configure redundant internet connections using SD-WAN (formerly virtual WAN link). Applies to FortiOS Version `>= 6.4.1`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable SD-WAN. Valid values: `disable`, `enable`.
* `load_balance_mode` - Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
* `speedtest_bypass_routing` - Enable/disable bypass routing when speedtest on a SD-WAN member. Valid values: `disable`, `enable`.
* `duplication_max_num` - Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
* `duplication_max_discrepancy` - Maximum discrepancy between two packets for deduplication in milliseconds (250 - 1000, default = 250).
* `neighbor_hold_down` - Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
* `neighbor_hold_down_time` - Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
* `app_perf_log_period` - Time interval in seconds that applicationperformance logs are generated (0 - 3600, default = 0).
* `neighbor_hold_boot_time` - Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
* `fail_detect` - Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
* `fail_alert_interfaces` - Physical interfaces that will be alerted. The structure of `fail_alert_interfaces` block is documented below.
* `zone` - Configure SD-WAN zones. The structure of `zone` block is documented below.
* `members` - FortiGate interfaces added to the SD-WAN. The structure of `members` block is documented below.
* `health_check` - SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `health_check` block is documented below.
* `health_check_fortiguard` - SD-WAN status checking or health checking. Identify a server predefine by FortiGuard and determine how SD-WAN verifies that FGT can communicate with it. The structure of `health_check_fortiguard` block is documented below.
* `neighbor` - Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
* `service` - Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN. The structure of `service` block is documented below.
* `duplication` - Create SD-WAN duplication rule. The structure of `duplication` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `fail_alert_interfaces` block supports:

* `name` - Physical interface name.

The `zone` block supports:

* `name` - Zone name.
* `advpn_select` - Enable/disable selection of ADVPN based on SDWAN information. Valid values: `enable`, `disable`.
* `advpn_health_check` - Health check for ADVPN local overlay link quality.
* `service_sla_tie_break` - Method of selecting member if more than one meets the SLA.
* `minimum_sla_meet_members` - Minimum number of members which meet SLA when the neighbor is preferred.

The `members` block supports:

* `seq_num` - Sequence number On FortiOS versions >= 6.4.2: 1-512. On FortiOS versions 6.4.1: 1-255.
* `interface` - Interface name.
* `zone` - Zone name.
* `gateway` - The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
* `preferred_source` - Preferred source of route for this member.
* `source` - Source IP address used in the health-check packet to the server.
* `gateway6` - IPv6 gateway.
* `source6` - Source IPv6 address used in the health-check packet to the server.
* `cost` - Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
* `weight` - Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.
* `priority` - Priority of the interface for IPv4 . Used for SD-WAN rules or priority rules. On FortiOS versions 6.4.1: 0 - 65535. On FortiOS versions >= 7.0.4: 1 - 65535, default = 1.
* `priority6` - Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.
* `priority_in_sla` - Preferred priority of routes to this member when this member is in-sla (0 - 65535, default = 0).
* `priority_out_sla` - Preferred priority of routes to this member when this member is out-of-sla (0 - 65535, default = 0).
* `spillover_threshold` - Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `ingress_spillover_threshold` - Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `volume_ratio` - Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).
* `status` - Enable/disable this interface in the SD-WAN. Valid values: `disable`, `enable`.
* `transport_group` - Measured transport group (0 - 255).
* `comment` - Comments.

The `health_check` block supports:

* `name` - Status check or health check name.
* `fortiguard` - Enable/disable use of FortiGuard predefined server. Valid values: `disable`, `enable`.
* `fortiguard_name` - Predefined health-check target name.
* `probe_packets` - Enable/disable transmission of probe packets. Valid values: `disable`, `enable`.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
* `system_dns` - Enable/disable system DNS as the probe server. Valid values: `disable`, `enable`.
* `server` - IP address or FQDN name of the server.
* `detect_mode` - The mode determining how to detect the server.
* `protocol` - Protocol used to determine if the FortiGate can communicate with the server.
* `port` - Port number used to communicate with the server over the selected protocol (0 - 65535, default = 0, auto select. http, tcp-connect: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21, twamp: 862).
* `quality_measured_method` - Method to measure the quality of tcp-connect. Valid values: `half-open`, `half-close`.
* `security_mode` - Twamp controller security mode. Valid values: `none`, `authentication`.
* `user` - The user name to access probe server.
* `password` - Twamp controller password in authentication mode
* `packet_size` - Packet size of a TWAMP test session. (124/158 - 1024)
* `ha_priority` - HA election priority (1 - 50).
* `ftp_mode` - FTP mode. Valid values: `passive`, `port`.
* `ftp_file` - Full path and file name on the FTP server to download for FTP health-check to probe.
* `http_get` - URL used to communicate with the server if the protocol if the protocol is HTTP.
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_match` - Response string expected from the server if the protocol is HTTP.
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `dns_match_ip` - Response IP expected from DNS server if the protocol is DNS.
* `interval` - Status check interval in milliseconds, or the time between attempting to connect to the server (default = 500). On FortiOS versions 6.4.1-7.0.10, 7.2.0-7.2.4: 500 - 3600*1000 msec. On FortiOS versions 7.0.11-7.0.17, >= 7.2.6: 20 - 3600*1000 msec.
* `probe_timeout` - Time to wait before a probe packet is considered lost (default = 500). On FortiOS versions 6.4.2-7.0.10, 7.2.0-7.2.4: 500 - 3600*1000 msec. On FortiOS versions 6.4.1: 500 - 5000 msec. On FortiOS versions 7.0.11-7.0.17, >= 7.2.6: 20 - 3600*1000 msec.
* `agent_probe_timeout` - Time to wait before a probe packet is considered lost when detect-mode is agent (5000 - 3600*1000 msec, default = 60000).
* `remote_probe_timeout` - Time to wait before a probe packet is considered lost when detect-mode is remote (20 - 3600*1000 msec, default = 5000).
* `failtime` - Number of failures before server is considered lost (1 - 3600, default = 5).
* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
* `probe_count` - Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `update_cascade_interface` - Enable/disable update cascade interface. Valid values: `enable`, `disable`.
* `update_static_route` - Enable/disable updating the static route. Valid values: `enable`, `disable`.
* `embed_measured_health` - Enable/disable embedding measured health information. Valid values: `enable`, `disable`.
* `sla_id_redistribute` - Select the ID from the SLA sub-table. The selected SLA's priority value will be distributed into the routing table (0 - 32, default = 0).
* `sla_fail_log_period` - Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).
* `sla_pass_log_period` - Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).
* `threshold_warning_packetloss` - Warning threshold for packet loss (percentage, default = 0).
* `threshold_alert_packetloss` - Alert threshold for packet loss (percentage, default = 0).
* `threshold_warning_latency` - Warning threshold for latency (ms, default = 0).
* `threshold_alert_latency` - Alert threshold for latency (ms, default = 0).
* `threshold_warning_jitter` - Warning threshold for jitter (ms, default = 0).
* `threshold_alert_jitter` - Alert threshold for jitter (ms, default = 0).
* `vrf` - Virtual Routing Forwarding ID.
* `source` - Source IP address used in the health-check packet to the server.
* `source6` - Source IPv6 addressused in the health-check packet to server.
* `members` - Member sequence number list. The structure of `members` block is documented below.
* `mos_codec` - Codec to use for MOS calculation (default = g711). Valid values: `g711`, `g722`, `g729`.
* `class_id` - Traffic class ID.
* `packet_loss_weight` - Coefficient of packet-loss in the formula of custom-profile-1.
* `latency_weight` - Coefficient of latency in the formula of custom-profile-1.
* `jitter_weight` - Coefficient of jitter in the formula of custom-profile-1.
* `bandwidth_weight` - Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.

The `members` block supports:

* `seq_num` - Member sequence number.

The `sla` block supports:

* `id` - SLA ID.
* `link_cost_factor` - Criteria on which to base link selection.
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).
* `mos_threshold` - Minimum Mean Opinion Score for SLA to be marked as pass. (1.0 - 5.0, default = 3.6).
* `custom_profile_threshold` - Custom profile threshold for SLA to be marked as pass(0 - 10000000, default = 0).
* `priority_in_sla` - Value to be distributed into routing table when in-sla (0 - 65535, default = 0).
* `priority_out_sla` - Value to be distributed into routing table when out-sla (0 - 65535, default = 0).

The `health_check_fortiguard` block supports:

* `target_name` - Status check or predefined health-check targets name.
* `probe_packets` - Enable/disable transmission of probe packets. Valid values: `disable`, `enable`.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
* `system_dns` - Enable/disable system DNS as the probe server. Valid values: `disable`, `enable`.
* `server` - Predefined IP address or FQDN name from FortiGuard.
* `detect_mode` - The mode determining how to detect the server. Valid values: `active`, `passive`, `prefer-passive`, `remote`, `agent-based`.
* `protocol` - Protocol used to determine if the FortiGate can communicate with the server. Valid values: `ping`, `tcp-echo`, `udp-echo`, `http`, `https`, `twamp`, `dns`, `tcp-connect`, `ftp`.
* `port` - Port number used to communicate with the server over the selected protocol (0 - 65535, default = 0, auto select. http, tcp-connect: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21, twamp: 862).
* `quality_measured_method` - Method to measure the quality of tcp-connect. Valid values: `half-open`, `half-close`.
* `security_mode` - Twamp controller security mode. Valid values: `none`, `authentication`.
* `user` - The user name to access probe server.
* `password` - TWAMP controller password in authentication mode.
* `packet_size` - Packet size of a TWAMP test session. (124/158 - 1024)
* `ha_priority` - HA election priority (1 - 50).
* `ftp_mode` - FTP mode. Valid values: `passive`, `port`.
* `ftp_file` - Full path and file name on the FTP server to download for FTP health-check to probe.
* `http_get` - URL used to communicate with the server if the protocol if the protocol is HTTP.
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_match` - Response string expected from the server if the protocol is HTTP.
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `dns_match_ip` - Response IP expected from DNS server if the protocol is DNS.
* `interval` - Status check interval in milliseconds, or the time between attempting to connect to the server (20 - 3600*1000 msec, default = 500).
* `probe_timeout` - Time to wait before a probe packet is considered lost (20 - 3600*1000 msec, default = 500).
* `failtime` - Number of failures before server is considered lost (1 - 3600, default = 5).
* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
* `probe_count` - Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `update_cascade_interface` - Enable/disable update cascade interface. Valid values: `enable`, `disable`.
* `update_static_route` - Enable/disable updating the static route. Valid values: `enable`, `disable`.
* `embed_measured_health` - Enable/disable embedding measured health information. Valid values: `enable`, `disable`.
* `sla_id_redistribute` - Select the ID from the SLA sub-table. The selected SLA's priority value will be distributed into the routing table (0 - 32, default = 0).
* `sla_fail_log_period` - Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).
* `sla_pass_log_period` - Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).
* `threshold_warning_packetloss` - Warning threshold for packet loss (percentage, default = 0).
* `threshold_alert_packetloss` - Alert threshold for packet loss (percentage, default = 0).
* `threshold_warning_latency` - Warning threshold for latency (ms, default = 0).
* `threshold_alert_latency` - Alert threshold for latency (ms, default = 0).
* `threshold_warning_jitter` - Warning threshold for jitter (ms, default = 0).
* `threshold_alert_jitter` - Alert threshold for jitter (ms, default = 0).
* `vrf` - Virtual Routing Forwarding ID.
* `source` - Source IP address used in the health-check packet to the server.
* `source6` - Source IPv6 address used in the health-check packet to server.
* `members` - Member sequence number list. The structure of `members` block is documented below.
* `mos_codec` - Codec to use for MOS calculation (default = g711). Valid values: `g711`, `g722`, `g729`.
* `class_id` - Traffic class ID.
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.

The `members` block supports:

* `seq_num` - Member sequence number.

The `sla` block supports:

* `id` - SLA ID.
* `link_cost_factor` - Criteria on which to base link selection. Valid values: `latency`, `jitter`, `packet-loss`, `mos`, `remote`.
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).
* `mos_threshold` - Minimum Mean Opinion Score for SLA to be marked as pass. (1.0 - 5.0, default = 3.6).
* `priority_in_sla` - Value to be distributed into routing table when in-sla (0 - 65535, default = 0).
* `priority_out_sla` - Value to be distributed into routing table when out-sla (0 - 65535, default = 0).

The `neighbor` block supports:

* `ip` - IP/IPv6 address of neighbor.
* `member_block` - Member sequence number list. *Due to the data type change of API, for other versions of FortiOS, please check variable `member`.* The structure of `member_block` block is documented below.
* `minimum_sla_meet_members` - Minimum number of members which meet SLA when the neighbor is preferred.
* `member` - Member sequence number. *Due to the data type change of API, for other versions of FortiOS, please check variable `member_block`.*
* `service_id` - SD-WAN service ID to work with the neighbor.
* `mode` - What metric to select the neighbor. Valid values: `sla`, `speedtest`.
* `role` - Role of neighbor. Valid values: `standalone`, `primary`, `secondary`.
* `route_metric` - Route-metric of neighbor. Valid values: `preferable`, `priority`.
* `health_check` - SD-WAN health-check name.
* `sla_id` - SLA ID.

The `member_block` block supports:

* `seq_num` - Member sequence number.

The `service` block supports:

* `id` - SD-WAN rule ID (1 - 4000).
* `name` - SD-WAN rule name.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
* `load_balance` - Enable/disable load-balance. Valid values: `enable`, `disable`.
* `shortcut_stickiness` - Enable/disable shortcut-stickiness of ADVPN. Valid values: `enable`, `disable`.
* `input_device` - Source interface name. The structure of `input_device` block is documented below.
* `input_device_negate` - Enable/disable negation of input device match. Valid values: `enable`, `disable`.
* `input_zone` - Source input-zone name. The structure of `input_zone` block is documented below.
* `mode` - Control how the SD-WAN rule sets the priority of interfaces in the SD-WAN.
* `zone_mode` - Enable/disable zone mode. Valid values: `enable`, `disable`.
* `minimum_sla_meet_members` - Minimum number of members which meet SLA.
* `hash_mode` - Hash algorithm for selected priority members for load balance mode. Valid values: `round-robin`, `source-ip-based`, `source-dest-ip-based`, `inbandwidth`, `outbandwidth`, `bibandwidth`.
* `shortcut_priority` - High priority of ADVPN shortcut for this service. Valid values: `enable`, `disable`, `auto`.
* `role` - Service role to work with neighbor. Valid values: `standalone`, `primary`, `secondary`.
* `standalone_action` - Enable/disable service when selected neighbor role is standalone while service role is not standalone. Valid values: `enable`, `disable`.
* `quality_link` - Quality grade.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `protocol` - Protocol number.
* `start_port` - Start destination port number.
* `end_port` - End destination port number.
* `start_src_port` - Start source port number.
* `end_src_port` - End source port number.
* `route_tag` - IPv4 route map route-tag.
* `dst` - Destination address name. The structure of `dst` block is documented below.
* `dst_negate` - Enable/disable negation of destination address match. Valid values: `enable`, `disable`.
* `src` - Source address name. The structure of `src` block is documented below.
* `dst6` - Destination address6 name. The structure of `dst6` block is documented below.
* `src6` - Source address6 name. The structure of `src6` block is documented below.
* `src_negate` - Enable/disable negation of source address match. Valid values: `enable`, `disable`.
* `users` - User name. The structure of `users` block is documented below.
* `groups` - User groups. The structure of `groups` block is documented below.
* `internet_service` - Enable/disable use of Internet service for application-based load balancing. Valid values: `enable`, `disable`.
* `internet_service_custom` - Custom Internet service name list. The structure of `internet_service_custom` block is documented below.
* `internet_service_custom_group` - Custom Internet Service group list. The structure of `internet_service_custom_group` block is documented below.
* `internet_service_fortiguard` - FortiGuard Internet service name list. The structure of `internet_service_fortiguard` block is documented below.
* `internet_service_name` - Internet service name list. The structure of `internet_service_name` block is documented below.
* `internet_service_group` - Internet Service group list. The structure of `internet_service_group` block is documented below.
* `internet_service_app_ctrl` - Application control based Internet Service ID list. The structure of `internet_service_app_ctrl` block is documented below.
* `internet_service_app_ctrl_group` - Application control based Internet Service group list. The structure of `internet_service_app_ctrl_group` block is documented below.
* `internet_service_app_ctrl_category` - IDs of one or more application control categories. The structure of `internet_service_app_ctrl_category` block is documented below.
* `health_check` - Health check list. The structure of `health_check` block is documented below.
* `link_cost_factor` - Link cost factor. Valid values: `latency`, `jitter`, `packet-loss`, `inbandwidth`, `outbandwidth`, `bibandwidth`, `custom-profile-1`.
* `packet_loss_weight` - Coefficient of packet-loss in the formula of custom-profile-1.
* `latency_weight` - Coefficient of latency in the formula of custom-profile-1.
* `jitter_weight` - Coefficient of jitter in the formula of custom-profile-1.
* `bandwidth_weight` - Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
* `link_cost_threshold` - Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
* `hold_down_time` - Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
* `sla_stickiness` - Enable/disable SLA stickiness (default = disable). Valid values: `enable`, `disable`.
* `dscp_forward` - Enable/disable forward traffic DSCP tag. Valid values: `enable`, `disable`.
* `dscp_reverse` - Enable/disable reverse traffic DSCP tag. Valid values: `enable`, `disable`.
* `dscp_forward_tag` - Forward traffic DSCP tag.
* `dscp_reverse_tag` - Reverse traffic DSCP tag.
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.
* `priority_members` - Member sequence number list. The structure of `priority_members` block is documented below.
* `priority_zone` - Priority zone name list. The structure of `priority_zone` block is documented below.
* `status` - Enable/disable SD-WAN service. Valid values: `enable`, `disable`.
* `gateway` - Enable/disable SD-WAN service gateway. Valid values: `enable`, `disable`.
* `default` - Enable/disable use of SD-WAN as default service. Valid values: `enable`, `disable`.
* `sla_compare_method` - Method to compare SLA value for SLA mode. Valid values: `order`, `number`.
* `fib_best_match_force` - Enable/disable force using fib-best-match oif as outgoing interface. Valid values: `disable`, `enable`.
* `tie_break` - Method of selecting member if more than one meets the SLA.
* `use_shortcut_sla` - Enable/disable use of ADVPN shortcut for quality comparison. Valid values: `enable`, `disable`.
* `passive_measurement` - Enable/disable passive measurement based on the service criteria. Valid values: `enable`, `disable`.
* `agent_exclusive` - Set/unset the service as agent use exclusively. Valid values: `enable`, `disable`.
* `shortcut` - Enable/disable shortcut for this service. Valid values: `enable`, `disable`.
* `comment` - Comments.

The `input_device` block supports:

* `name` - Interface name.

The `input_zone` block supports:

* `name` - Zone.

The `dst` block supports:

* `name` - Address or address group name.

The `src` block supports:

* `name` - Address or address group name.

The `dst6` block supports:

* `name` - Address6 or address6 group name.

The `src6` block supports:

* `name` - Address6 or address6 group name.

The `users` block supports:

* `name` - User name.

The `groups` block supports:

* `name` - Group name.

The `internet_service_custom` block supports:

* `name` - Custom Internet service name.

The `internet_service_custom_group` block supports:

* `name` - Custom Internet Service group name.

The `internet_service_fortiguard` block supports:

* `name` - FortiGuard Internet service name.

The `internet_service_name` block supports:

* `name` - Internet service name.

The `internet_service_group` block supports:

* `name` - Internet Service group name.

The `internet_service_app_ctrl` block supports:

* `id` - Application control based Internet Service ID.

The `internet_service_app_ctrl_group` block supports:

* `name` - Application control based Internet Service group name.

The `internet_service_app_ctrl_category` block supports:

* `id` - Application control category ID.

The `health_check` block supports:

* `name` - Health check name.

The `sla` block supports:

* `health_check` - SD-WAN health-check.
* `id` - SLA ID.

The `priority_members` block supports:

* `seq_num` - Member sequence number.

The `priority_zone` block supports:

* `name` - Priority zone name.

The `duplication` block supports:

* `id` - Duplication rule ID (1 - 255).
* `service_id` - SD-WAN service rule ID list. The structure of `service_id` block is documented below.
* `srcaddr` - Source address or address group names. The structure of `srcaddr` block is documented below.
* `dstaddr` - Destination address or address group names. The structure of `dstaddr` block is documented below.
* `srcaddr6` - Source address6 or address6 group names. The structure of `srcaddr6` block is documented below.
* `dstaddr6` - Destination address6 or address6 group names. The structure of `dstaddr6` block is documented below.
* `srcintf` - Incoming (ingress) interfaces or zones. The structure of `srcintf` block is documented below.
* `dstintf` - Outgoing (egress) interfaces or zones. The structure of `dstintf` block is documented below.
* `service` - Service and service group name. The structure of `service` block is documented below.
* `packet_duplication` - Configure packet duplication method. Valid values: `disable`, `force`, `on-demand`.
* `sla_match_service` - Enable/disable packet duplication matching health-check SLAs in service rule. Valid values: `enable`, `disable`.
* `packet_de_duplication` - Enable/disable discarding of packets that have been duplicated. Valid values: `enable`, `disable`.

The `service_id` block supports:

* `id` - SD-WAN service rule ID.

The `srcaddr` block supports:

* `name` - Address or address group name.

The `dstaddr` block supports:

* `name` - Address or address group name.

The `srcaddr6` block supports:

* `name` - Address6 or address6 group name.

The `dstaddr6` block supports:

* `name` - Address6 or address6 group name.

The `srcintf` block supports:

* `name` - Interface, zone or SDWAN zone name.

The `dstintf` block supports:

* `name` - Interface, zone or SDWAN zone name.

The `service` block supports:

* `name` - Service and service group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Sdwan can be imported using any of these accepted formats:
```
$ terraform import fortios_system_sdwan.labelname SystemSdwan

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_sdwan.labelname SystemSdwan
$ unset "FORTIOS_IMPORT_TABLE"
```
