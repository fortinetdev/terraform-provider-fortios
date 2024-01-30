---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_standalonecluster"
description: |-
  Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.
---

# fortios_system_standalonecluster
Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `standalone_group_id` - Cluster group ID (0 - 255). Must be the same for all members.
* `group_member_id` - Cluster member ID (0 - 3).
* `layer2_connection` - Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
* `session_sync_dev` - Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
* `encryption` - Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
* `psksecret` - Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
* `asymmetric_traffic_control` - Asymmetric traffic control mode. Valid values: `cps-preferred`, `strict-anti-replay`.
* `cluster_peer` - Configure FortiGate Session Life Support Protocol (FGSP) session synchronization. The structure of `cluster_peer` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `cluster_peer` block supports:

* `sync_id` - Sync ID.
* `peervd` - VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
* `peerip` - IP address of the interface on the peer unit that is used for the session synchronization link.
* `syncvd` - Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
* `down_intfs_before_sess_sync` - List of interfaces to be turned down before session synchronization is complete. The structure of `down_intfs_before_sess_sync` block is documented below.
* `hb_interval` - Heartbeat interval (1 - 20 (100*ms). Increase to reduce false positives.
* `hb_lost_threshold` - Lost heartbeat threshold (1 - 60). Increase to reduce false positives.
* `ipsec_tunnel_sync` - Enable/disable IPsec tunnel synchronization. Valid values: `enable`, `disable`.
* `secondary_add_ipsec_routes` - Enable/disable IKE route announcement on the backup unit. Valid values: `enable`, `disable`.
* `session_sync_filter` - Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `session_sync_filter` block is documented below.

The `syncvd` block supports:

* `name` - VDOM name.

The `down_intfs_before_sess_sync` block supports:

* `name` - Interface name.

The `session_sync_filter` block supports:

* `srcintf` - Only sessions from this interface are synchronized.
* `dstintf` - Only sessions to this interface are synchronized.
* `srcaddr` - Only sessions from this IPv4 address are synchronized.
* `dstaddr` - Only sessions to this IPv4 address are synchronized.
* `srcaddr6` - Only sessions from this IPv6 address are synchronized.
* `dstaddr6` - Only sessions to this IPv6 address are synchronized.
* `custom_service` - Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services. The structure of `custom_service` block is documented below.

The `custom_service` block supports:

* `id` - Custom service ID.
* `src_port_range` - Custom service source port range.
* `dst_port_range` - Custom service destination port range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System StandaloneCluster can be imported using any of these accepted formats:
```
$ terraform import fortios_system_standalonecluster.labelname SystemStandaloneCluster

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_standalonecluster.labelname SystemStandaloneCluster
$ unset "FORTIOS_IMPORT_TABLE"
```
