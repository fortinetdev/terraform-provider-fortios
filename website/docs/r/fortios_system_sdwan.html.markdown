---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan"
description: |-
  Configure redundant internet connections using SD-WAN (formerly virtual WAN link).
---

# fortios_system_sdwan
Configure redundant internet connections using SD-WAN (formerly virtual WAN link). Applies to FortiOS Version `>= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable SD-WAN.
* `load_balance_mode` - Algorithm or mode to use for load balancing Internet traffic to SD-WAN members.
* `duplication_max_num` - Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
* `neighbor_hold_down` - Enable/disable hold switching from the secondary neighbor to the primary neighbor.
* `neighbor_hold_down_time` - Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
* `neighbor_hold_boot_time` - Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
* `fail_detect` - Enable/disable SD-WAN Internet connection status checking (failure detection).
* `fail_alert_interfaces` - Physical interfaces that will be alerted. The structure of `fail_alert_interfaces` block is documented below.
* `zone` - Configure SD-WAN zones. The structure of `zone` block is documented below.
* `members` - FortiGate interfaces added to the SD-WAN. The structure of `members` block is documented below.
* `health_check` - SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `health_check` block is documented below.
* `neighbor` - Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
* `service` - Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN. The structure of `service` block is documented below.
* `duplication` - Create SD-WAN duplication rule. The structure of `duplication` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `fail_alert_interfaces` block supports:

* `name` - Physical interface name.

The `zone` block supports:

* `name` - Zone name.
* `service_sla_tie_break` - Method of selecting member if more than one meets the SLA.

The `members` block supports:

* `seq_num` - Sequence number(1-512).
* `interface` - Interface name.
* `zone` - Zone name.
* `gateway` - The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
* `source` - Source IP address used in the health-check packet to the server.
* `gateway6` - IPv6 gateway.
* `source6` - Source IPv6 address used in the health-check packet to the server.
* `cost` - Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
* `weight` - Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.
* `priority` - Priority of the interface (0 - 65535). Used for SD-WAN rules or priority rules.
* `priority6` - Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.
* `spillover_threshold` - Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `ingress_spillover_threshold` - Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `volume_ratio` - Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).
* `status` - Enable/disable this interface in the SD-WAN.
* `comment` - Comments.

The `health_check` block supports:

* `name` - Status check or health check name.
* `probe_packets` - Enable/disable transmission of probe packets.
* `addr_mode` - Address mode (IPv4 or IPv6).
* `system_dns` - Enable/disable system DNS as the probe server.
* `server` - IP address or FQDN name of the server.
* `protocol` - Protocol used to determine if the FortiGate can communicate with the server.
* `port` - Port number used to communicate with the server over the selected protocol (0-65535, default = 0, auto select. http, twamp: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21).
* `quality_measured_method` - Method to measure the quality of tcp-connect.
* `security_mode` - Twamp controller security mode.
* `user` - The user name to access probe server.
* `password` - Twamp controller password in authentication mode
* `packet_size` - Packet size of a twamp test session,
* `ha_priority` - HA election priority (1 - 50).
* `ftp_mode` - FTP mode.
* `ftp_file` - Full path and file name on the FTP server to download for FTP health-check to probe.
* `http_get` - URL used to communicate with the server if the protocol if the protocol is HTTP.
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_match` - Response string expected from the server if the protocol is HTTP.
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `dns_match_ip` - Response IP expected from DNS server if the protocol is DNS.
* `interval` - Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).
* `probe_timeout` - Time to wait before a probe packet is considered lost (500 - 3600*1000 msec, default = 500).
* `failtime` - Number of failures before server is considered lost (1 - 3600, default = 5).
* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
* `probe_count` - Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `update_cascade_interface` - Enable/disable update cascade interface.
* `update_static_route` - Enable/disable updating the static route.
* `sla_fail_log_period` - Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).
* `sla_pass_log_period` - Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).
* `threshold_warning_packetloss` - Warning threshold for packet loss (percentage, default = 0).
* `threshold_alert_packetloss` - Alert threshold for packet loss (percentage, default = 0).
* `threshold_warning_latency` - Warning threshold for latency (ms, default = 0).
* `threshold_alert_latency` - Alert threshold for latency (ms, default = 0).
* `threshold_warning_jitter` - Warning threshold for jitter (ms, default = 0).
* `threshold_alert_jitter` - Alert threshold for jitter (ms, default = 0).
* `members` - Member sequence number list. The structure of `members` block is documented below.
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.

The `members` block supports:

* `seq_num` - Member sequence number.

The `sla` block supports:

* `id` - SLA ID.
* `link_cost_factor` - Criteria on which to base link selection.
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).

The `neighbor` block supports:

* `ip` - IP/IPv6 address of neighbor.
* `member` - Member sequence number.
* `role` - Role of neighbor.
* `health_check` - SD-WAN health-check name.
* `sla_id` - SLA ID.

The `service` block supports:

* `id` - SD-WAN rule ID (1 - 4000).
* `name` - SD-WAN rule name.
* `addr_mode` - Address mode (IPv4 or IPv6).
* `input_device` - Source interface name. The structure of `input_device` block is documented below.
* `input_device_negate` - Enable/disable negation of input device match.
* `mode` - Control how the SD-WAN rule sets the priority of interfaces in the SD-WAN.
* `minimum_sla_meet_members` - Minimum number of members which meet SLA.
* `hash_mode` - Hash algorithm for selected priority members for load balance mode.
* `role` - Service role to work with neighbor.
* `standalone_action` - Enable/disable service when selected neighbor role is standalone while service role is not standalone.
* `quality_link` - Quality grade.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `protocol` - Protocol number.
* `start_port` - Start destination port number.
* `end_port` - End destination port number.
* `route_tag` - IPv4 route map route-tag.
* `dst` - Destination address name. The structure of `dst` block is documented below.
* `dst_negate` - Enable/disable negation of destination address match.
* `src` - Source address name. The structure of `src` block is documented below.
* `dst6` - Destination address6 name. The structure of `dst6` block is documented below.
* `src6` - Source address6 name. The structure of `src6` block is documented below.
* `src_negate` - Enable/disable negation of source address match.
* `users` - User name. The structure of `users` block is documented below.
* `groups` - User groups. The structure of `groups` block is documented below.
* `internet_service` - Enable/disable use of Internet service for application-based load balancing.
* `internet_service_custom` - Custom Internet service name list. The structure of `internet_service_custom` block is documented below.
* `internet_service_custom_group` - Custom Internet Service group list. The structure of `internet_service_custom_group` block is documented below.
* `internet_service_name` - Internet service name list. The structure of `internet_service_name` block is documented below.
* `internet_service_group` - Internet Service group list. The structure of `internet_service_group` block is documented below.
* `internet_service_app_ctrl` - Application control based Internet Service ID list. The structure of `internet_service_app_ctrl` block is documented below.
* `internet_service_app_ctrl_group` - Application control based Internet Service group list. The structure of `internet_service_app_ctrl_group` block is documented below.
* `health_check` - Health check list. The structure of `health_check` block is documented below.
* `link_cost_factor` - Link cost factor.
* `packet_loss_weight` - Coefficient of packet-loss in the formula of custom-profile-1.
* `latency_weight` - Coefficient of latency in the formula of custom-profile-1.
* `jitter_weight` - Coefficient of jitter in the formula of custom-profile-1.
* `bandwidth_weight` - Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
* `link_cost_threshold` - Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
* `hold_down_time` - Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
* `dscp_forward` - Enable/disable forward traffic DSCP tag.
* `dscp_reverse` - Enable/disable reverse traffic DSCP tag.
* `dscp_forward_tag` - Forward traffic DSCP tag.
* `dscp_reverse_tag` - Reverse traffic DSCP tag.
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.
* `priority_members` - Member sequence number list. The structure of `priority_members` block is documented below.
* `status` - Enable/disable SD-WAN service.
* `gateway` - Enable/disable SD-WAN service gateway.
* `default` - Enable/disable use of SD-WAN as default service.
* `sla_compare_method` - Method to compare SLA value for SLA mode.
* `tie_break` - Method of selecting member if more than one meets the SLA.
* `use_shortcut_sla` - Enable/disable use of ADVPN shortcut for quality comparison.

The `input_device` block supports:

* `name` - Interface name.

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

The `internet_service_name` block supports:

* `name` - Internet service name.

The `internet_service_group` block supports:

* `name` - Internet Service group name.

The `internet_service_app_ctrl` block supports:

* `id` - Application control based Internet Service ID.

The `internet_service_app_ctrl_group` block supports:

* `name` - Application control based Internet Service group name.

The `health_check` block supports:

* `name` - Health check name.

The `sla` block supports:

* `health_check` - SD-WAN health-check.
* `id` - SLA ID.

The `priority_members` block supports:

* `seq_num` - Member sequence number.

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
* `packet_duplication` - Configure packet duplication method.
* `packet_de_duplication` - Enable/disable discarding of packets that have been duplicated.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_sdwan.labelname SystemSdwan
$ unset "FORTIOS_IMPORT_TABLE"
```
