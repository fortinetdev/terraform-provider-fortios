---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ha"
description: |-
  Get information on fortios system ha.
---

# Data Source: fortios_system_ha
Use this data source to get information on fortios system ha

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `group_id` - Cluster group ID  (0 - 255). Must be the same for all members.
* `group_name` - Cluster group name. Must be the same for all members.
* `mode` - HA mode. Must be the same for all members. FGSP requires standalone.
* `sync_packet_balance` - Enable/disable HA packet distribution to multiple CPUs.
* `password` - Cluster password. Must be the same for all members.
* `key` - key
* `hbdev` - Heartbeat interfaces. Must be the same for all members.
* `unicast_hb` - Enable/disable unicast heartbeat.
* `unicast_hb_peerip` - Unicast heartbeat peer IP.
* `unicast_hb_netmask` - Unicast heartbeat netmask.
* `session_sync_dev` - Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
* `route_ttl` - TTL for primary unit routes (5 - 3600 sec). Increase to maintain active routes during failover.
* `route_wait` - Time to wait before sending new routes to the cluster (0 - 3600 sec).
* `route_hold` - Time to wait between routing table updates to the cluster (0 - 3600 sec).
* `multicast_ttl` - HA multicast TTL on master (5 - 3600 sec).
* `evpn_ttl` - HA EVPN FDB TTL on primary box (5 - 3600 sec).
* `load_balance_all` - Enable to load balance TCP sessions. Disable to load balance proxy sessions only.
* `sync_config` - Enable/disable configuration synchronization.
* `encryption` - Enable/disable heartbeat message encryption.
* `authentication` - Enable/disable heartbeat message authentication.
* `hb_interval` - Time between sending heartbeat packets (1 - 20 (100*ms)). Increase to reduce false positives.
* `hb_interval_in_milliseconds` - Number of milliseconds for each heartbeat interval: 100ms or 10ms.
* `hb_lost_threshold` - Number of lost heartbeats to signal a failure (1 - 60). Increase to reduce false positives.
* `hello_holddown` - Time to wait before changing from hello to work state (5 - 300 sec).
* `gratuitous_arps` - Enable/disable gratuitous ARPs. Disable if link-failed-signal enabled.
* `arps` - Number of gratuitous ARPs (1 - 60). Lower to reduce traffic. Higher to reduce failover time.
* `arps_interval` - Time between gratuitous ARPs  (1 - 20 sec). Lower to reduce failover time. Higher to reduce traffic.
* `session_pickup` - Enable/disable session pickup. Enabling it can reduce session down time when fail over happens.
* `session_pickup_connectionless` - Enable/disable UDP and ICMP session sync.
* `session_pickup_expectation` - Enable/disable session helper expectation session sync for FGSP.
* `session_pickup_nat` - Enable/disable NAT session sync for FGSP.
* `session_pickup_delay` - Enable to sync sessions longer than 30 sec. Only longer lived sessions need to be synced.
* `link_failed_signal` - Enable to shut down all interfaces for 1 sec after a failover. Use if gratuitous ARPs do not update network.
* `uninterruptible_upgrade` - Enable to upgrade a cluster without blocking network traffic.
* `uninterruptible_primary_wait` - Number of minutes the primary HA unit waits before the secondary HA unit is considered upgraded and the system is started before starting its own upgrade (1 - 300, default = 30).
* `standalone_mgmt_vdom` - Enable/disable standalone management VDOM.
* `ha_mgmt_status` - Enable to reserve interfaces to manage individual cluster units.
* `ha_mgmt_interfaces` - Reserve interfaces to manage individual cluster units. The structure of `ha_mgmt_interfaces` block is documented below.
* `ha_eth_type` - HA heartbeat packet Ethertype (4-digit hex).
* `hc_eth_type` - Transparent mode HA heartbeat packet Ethertype (4-digit hex).
* `l2ep_eth_type` - Telnet session HA heartbeat packet Ethertype (4-digit hex).
* `ha_uptime_diff_margin` - Normally you would only reduce this value for failover testing.
* `standalone_config_sync` - Enable/disable FGSP configuration synchronization.
* `unicast_status` - Enable/disable unicast connection.
* `unicast_gateway` - Default route gateway for unicast interface.
* `unicast_peers` - Number of unicast peers. The structure of `unicast_peers` block is documented below.
* `logical_sn` - Enable/disable usage of the logical serial number.
* `vcluster2` - Enable/disable virtual cluster 2 for virtual clustering.
* `vcluster_id` - Cluster ID.
* `override` - Enable and increase the priority of the unit that should always be primary (master).
* `priority` - Increase the priority to select the primary unit (0 - 255).
* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `schedule` - Type of A-A load balancing. Use none if you have external load balancers.
* `weight` - Weight-round-robin weight for each cluster unit. Syntax <priority> <weight>.
* `cpu_threshold` - Dynamic weighted load balancing CPU usage weight and high and low thresholds.
* `memory_threshold` - Dynamic weighted load balancing memory usage weight and high and low thresholds.
* `http_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of HTTP proxy sessions.
* `ftp_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of FTP proxy sessions.
* `imap_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of IMAP proxy sessions.
* `nntp_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of NNTP proxy sessions.
* `pop3_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of POP3 proxy sessions.
* `smtp_proxy_threshold` - Dynamic weighted load balancing weight and high and low number of SMTP proxy sessions.
* `monitor` - Interfaces to check for port monitoring (or link failure).
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover.
* `pingserver_slave_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover.
* `pingserver_flip_timeout` - Time to wait in minutes before renegotiating after a remote IP monitoring failover.
* `vcluster_status` - Enable/disable virtual cluster for virtual clustering.
* `vcluster` - Virtual cluster table. The structure of `vcluster` block is documented below.
* `vdom` - VDOMs in virtual cluster 1.
* `secondary_vcluster` - Configure virtual cluster 2. The structure of `secondary_vcluster` block is documented below.
* `ha_direct` - Enable/disable using ha-mgmt interface for syslog, SNMP, remote authentication (RADIUS), FortiAnalyzer, and FortiSandbox.
* `ssd_failover` - Enable/disable automatic HA failover on SSD disk failure.
* `memory_compatible_mode` - Enable/disable memory compatible mode.
* `memory_based_failover` - Enable/disable memory based failover.
* `memory_failover_threshold` - Memory usage threshold to trigger memory based failover (0 means using conserve mode threshold in system.global).
* `memory_failover_monitor_period` - Duration of high memory usage before memory based failover is triggered in seconds (1 - 300, default = 60).
* `memory_failover_sample_rate` - Rate at which memory usage is sampled in order to measure memory usage in seconds (1 - 60, default = 1).
* `memory_failover_flip_timeout` - Time to wait between subsequent memory based failovers in minutes (6 - 2147483647, default = 6).
* `failover_hold_time` - Time to wait before failover (0 - 300 sec, default = 0), to avoid flip.
* `inter_cluster_session_sync` - Enable/disable synchronization of sessions among HA clusters.

The `ha_mgmt_interfaces` block contains:

* `id` - Table ID.
* `interface` - Interface to reserve for HA management.
* `dst` - Default route destination for reserved HA management interface.
* `gateway` - Default route gateway for reserved HA management interface.
* `gateway6` - Default IPv6 gateway for reserved HA management interface.

The `unicast_peers` block contains:

* `id` - Table ID.
* `peer_ip` - Unicast peer IP.

The `vcluster` block contains:

* `vcluster_id` - ID.
* `override` - Enable and increase the priority of the unit that should always be primary (master).
* `priority` - Increase the priority to select the primary unit (0 - 255).
* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `monitor` - Interfaces to check for port monitoring (or link failure).
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover.
* `pingserver_slave_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover.
* `vdom` - Virtual domain(s) in the virtual cluster. The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - Virtual domain name.

The `secondary_vcluster` block contains:

* `vcluster_id` - Cluster ID.
* `override` - Enable and increase the priority of the unit that should always be primary (master).
* `priority` - Increase the priority to select the primary unit (0 - 255).
* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `monitor` - Interfaces to check for port monitoring (or link failure).
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover.
* `pingserver_slave_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover.
* `vdom` - VDOMs in virtual cluster 2.

