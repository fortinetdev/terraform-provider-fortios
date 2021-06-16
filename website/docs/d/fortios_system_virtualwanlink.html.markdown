---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualwanlink"
description: |-
  Get information on fortios system virtualwanlink.
---

# Data Source: fortios_system_virtualwanlink
Use this data source to get information on fortios system virtualwanlink

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable SD-WAN.
* `load_balance_mode` - Algorithm or mode to use for load balancing Internet traffic to SD-WAN members.
* `neighbor_hold_down` - Enable/disable hold switching from the secondary neighbor to the primary neighbor.
* `neighbor_hold_down_time` - Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
* `neighbor_hold_boot_time` - Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
* `fail_detect` - Enable/disable SD-WAN Internet connection status checking (failure detection).
* `fail_alert_interfaces` - Physical interfaces that will be alerted. The structure of `fail_alert_interfaces` block is documented below.
* `zone` - Configure SD-WAN zones. The structure of `zone` block is documented below.
* `members` - Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
* `health_check` - SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `health_check` block is documented below.
* `neighbor` - Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
* `service` - Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.

The `fail_alert_interfaces` block contains:

* `name` - Physical interface name.

The `zone` block contains:

* `name` - Zone name.

The `members` block contains:

* `seq_num` - Sequence number(1-255).
* `interface` - Interface name.
* `gateway` - The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
* `source` - Source IP address used in the health-check packet to the server.
* `gateway6` - IPv6 gateway.
* `source6` - Source IPv6 address used in the health-check packet to the server.
* `cost` - Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
* `weight` - Weight of this interface for weighted load balancing. (0 - 255) More traffic is directed to interfaces with higher weights.
* `priority` - Priority of the interface (0 - 4294967295). Used for SD-WAN rules or priority rules.
* `spillover_threshold` - Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `ingress_spillover_threshold` - Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `volume_ratio` - Measured volume ratio (this value / sum of all values = percentage of link volume, 0 - 255).
* `status` - Enable/disable this interface in the SD-WAN.
* `comment` - Comments.

The `health_check` block contains:

* `name` - Status check or health check name.
* `probe_packets` - Enable/disable transmission of probe packets.
* `addr_mode` - Address mode (IPv4 or IPv6).
* `system_dns` - Enable/disable system DNS as the probe server.
* `server` - IP address or FQDN name of the server.
* `protocol` - Protocol used to determine if the FortiGate can communicate with the server.
* `port` - Port number used to communicate with the server over the selected protocol.
* `security_mode` - Twamp controller security mode.
* `password` - Twamp controller password in authentication mode
* `packet_size` - Packet size of a twamp test session,
* `ha_priority` - HA election priority (1 - 50).
* `http_get` - URL used to communicate with the server if the protocol if the protocol is HTTP.
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_match` - Response string expected from the server if the protocol is HTTP.
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `interval` - Status check interval, or the time between attempting to connect to the server (1 - 3600 sec, default = 5).
* `probe_timeout` - Time to wait before a probe packet is considered lost (500 - 5000 msec, default = 500).
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

The `members` block contains:

* `seq_num` - Member sequence number.

The `sla` block contains:

* `id` - SLA ID.
* `link_cost_factor` - Criteria on which to base link selection.
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).

The `neighbor` block contains:

* `ip` - IP address of neighbor.
* `member` - Member sequence number.
* `role` - Role of neighbor.
* `health_check` - SD-WAN health-check name.
* `sla_id` - SLA ID.

The `service` block contains:

* `id` - Priority rule ID (1 - 4000).
* `name` - Priority rule name.
* `addr_mode` - Address mode (IPv4 or IPv6).
* `input_device` - Source interface name. The structure of `input_device` block is documented below.
* `input_device_negate` - Enable/disable negation of input device match.
* `mode` - Control how the priority rule sets the priority of interfaces in the SD-WAN.
* `role` - Service role to work with neighbor.
* `standalone_action` - Enable/disable service when selected neighbor role is standalone while service role is not standalone.
* `quality_link` - Quality grade.
* `member` - Member sequence number.
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
* `internet_service_id` - Internet service ID list. The structure of `internet_service_id` block is documented below.
* `internet_service_group` - Internet Service group list. The structure of `internet_service_group` block is documented below.
* `internet_service_app_ctrl` - Application control based Internet Service ID list. The structure of `internet_service_app_ctrl` block is documented below.
* `internet_service_app_ctrl_group` - Application control based Internet Service group list. The structure of `internet_service_app_ctrl_group` block is documented below.
* `internet_service_ctrl` - Control-based Internet Service ID list. The structure of `internet_service_ctrl` block is documented below.
* `internet_service_ctrl_group` - Control-based Internet Service group list. The structure of `internet_service_ctrl_group` block is documented below.
* `health_check` - Health check.
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
* `sla_compare_method` - Method to compare SLA value for sla and load balance mode. 

The `input_device` block contains:

* `name` - Interface name.

The `dst` block contains:

* `name` - Address or address group name.

The `src` block contains:

* `name` - Address or address group name.

The `dst6` block contains:

* `name` - Address6 or address6 group name.

The `src6` block contains:

* `name` - Address6 or address6 group name.

The `users` block contains:

* `name` - User name.

The `groups` block contains:

* `name` - Group name.

The `internet_service_custom` block contains:

* `name` - Custom Internet service name.

The `internet_service_custom_group` block contains:

* `name` - Custom Internet Service group name.

The `internet_service_name` block contains:

* `name` - Internet service name.

The `internet_service_id` block contains:

* `id` - Internet service ID.

The `internet_service_group` block contains:

* `name` - Internet Service group name.

The `internet_service_app_ctrl` block contains:

* `id` - Application control based Internet Service ID.

The `internet_service_app_ctrl_group` block contains:

* `name` - Application control based Internet Service group name.

The `internet_service_ctrl` block contains:

* `id` - Control-based Internet Service ID.

The `internet_service_ctrl_group` block contains:

* `name` - Control-based Internet Service group name.

The `sla` block contains:

* `health_check` - Virtual WAN Link health-check.
* `id` - SLA ID.

The `priority_members` block contains:

* `seq_num` - Member sequence number.

