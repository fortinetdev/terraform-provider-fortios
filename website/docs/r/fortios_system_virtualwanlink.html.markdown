---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualwanlink"
description: |-
  Configure redundant internet connections using SD-WAN (formerly virtual WAN link).
---

# fortios_system_virtualwanlink
Configure redundant internet connections using SD-WAN (formerly virtual WAN link). Applies to FortiOS Version `<= 6.4.0`.

## Example Usage

```hcl
resource "fortios_system_virtualwanlink" "trname" {
  fail_detect       = "disable"
  load_balance_mode = "source-ip-based"
  status            = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable SD-WAN. Valid values: `disable`, `enable`.
* `load_balance_mode` - Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
* `neighbor_hold_down` - Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
* `neighbor_hold_down_time` - Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
* `neighbor_hold_boot_time` - Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
* `fail_detect` - Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
* `fail_alert_interfaces` - Physical interfaces that will be alerted. The structure of `fail_alert_interfaces` block is documented below.
* `zone` - Configure SD-WAN zones. The structure of `zone` block is documented below.
* `members` - Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
* `health_check` - SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `health_check` block is documented below.
* `neighbor` - Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
* `service` - Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `fail_alert_interfaces` block supports:

* `name` - Physical interface name.

The `zone` block supports:

* `name` - Zone name.

The `members` block supports:

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
* `status` - Enable/disable this interface in the SD-WAN. Valid values: `disable`, `enable`.
* `comment` - Comments.

The `health_check` block supports:

* `name` - Status check or health check name.
* `probe_packets` - Enable/disable transmission of probe packets. Valid values: `disable`, `enable`.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
* `system_dns` - Enable/disable system DNS as the probe server. Valid values: `disable`, `enable`.
* `server` - IP address or FQDN name of the server.
* `protocol` - Protocol used to determine if the FortiGate can communicate with the server.
* `port` - Port number used to communicate with the server over the selected protocol.
* `security_mode` - Twamp controller security mode. Valid values: `none`, `authentication`.
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
* `update_cascade_interface` - Enable/disable update cascade interface. Valid values: `enable`, `disable`.
* `update_static_route` - Enable/disable updating the static route. Valid values: `enable`, `disable`.
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
* `link_cost_factor` - Criteria on which to base link selection. Valid values: `latency`, `jitter`, `packet-loss`.
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).

The `neighbor` block supports:

* `ip` - IP address of neighbor.
* `member` - Member sequence number.
* `role` - Role of neighbor. Valid values: `standalone`, `primary`, `secondary`.
* `health_check` - SD-WAN health-check name.
* `sla_id` - SLA ID.

The `service` block supports:

* `id` - Priority rule ID (1 - 4000).
* `name` - Priority rule name.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
* `input_device` - Source interface name. The structure of `input_device` block is documented below.
* `input_device_negate` - Enable/disable negation of input device match. Valid values: `enable`, `disable`.
* `mode` - Control how the priority rule sets the priority of interfaces in the SD-WAN. Valid values: `auto`, `manual`, `priority`, `sla`, `load-balance`.
* `role` - Service role to work with neighbor. Valid values: `standalone`, `primary`, `secondary`.
* `standalone_action` - Enable/disable service when selected neighbor role is standalone while service role is not standalone. Valid values: `enable`, `disable`.
* `quality_link` - Quality grade.
* `member` - Member sequence number.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `protocol` - Protocol number.
* `start_port` - Start destination port number.
* `end_port` - End destination port number.
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
* `internet_service_name` - Internet service name list. The structure of `internet_service_name` block is documented below.
* `internet_service_id` - Internet service ID list. The structure of `internet_service_id` block is documented below.
* `internet_service_group` - Internet Service group list. The structure of `internet_service_group` block is documented below.
* `internet_service_app_ctrl` - Application control based Internet Service ID list. The structure of `internet_service_app_ctrl` block is documented below.
* `internet_service_app_ctrl_group` - Application control based Internet Service group list. The structure of `internet_service_app_ctrl_group` block is documented below.
* `internet_service_ctrl` - Control-based Internet Service ID list. The structure of `internet_service_ctrl` block is documented below.
* `internet_service_ctrl_group` - Control-based Internet Service group list. The structure of `internet_service_ctrl_group` block is documented below.
* `health_check` - Health check.
* `link_cost_factor` - Link cost factor. Valid values: `latency`, `jitter`, `packet-loss`, `inbandwidth`, `outbandwidth`, `bibandwidth`, `custom-profile-1`.
* `packet_loss_weight` - Coefficient of packet-loss in the formula of custom-profile-1.
* `latency_weight` - Coefficient of latency in the formula of custom-profile-1.
* `jitter_weight` - Coefficient of jitter in the formula of custom-profile-1.
* `bandwidth_weight` - Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
* `link_cost_threshold` - Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
* `hold_down_time` - Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
* `dscp_forward` - Enable/disable forward traffic DSCP tag. Valid values: `enable`, `disable`.
* `dscp_reverse` - Enable/disable reverse traffic DSCP tag. Valid values: `enable`, `disable`.
* `dscp_forward_tag` - Forward traffic DSCP tag.
* `dscp_reverse_tag` - Reverse traffic DSCP tag.
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.
* `priority_members` - Member sequence number list. The structure of `priority_members` block is documented below.
* `status` - Enable/disable SD-WAN service. Valid values: `enable`, `disable`.
* `gateway` - Enable/disable SD-WAN service gateway. Valid values: `enable`, `disable`.
* `default` - Enable/disable use of SD-WAN as default service. Valid values: `enable`, `disable`.
* `sla_compare_method` - Method to compare SLA value for sla and load balance mode.  Valid values: `order`, `number`.

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

The `internet_service_id` block supports:

* `id` - Internet service ID.

The `internet_service_group` block supports:

* `name` - Internet Service group name.

The `internet_service_app_ctrl` block supports:

* `id` - Application control based Internet Service ID.

The `internet_service_app_ctrl_group` block supports:

* `name` - Application control based Internet Service group name.

The `internet_service_ctrl` block supports:

* `id` - Control-based Internet Service ID.

The `internet_service_ctrl_group` block supports:

* `name` - Control-based Internet Service group name.

The `sla` block supports:

* `health_check` - Virtual WAN Link health-check.
* `id` - SLA ID.

The `priority_members` block supports:

* `seq_num` - Member sequence number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VirtualWanLink can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_virtualwanlink.labelname SystemVirtualWanLink
$ unset "FORTIOS_IMPORT_TABLE"
```
