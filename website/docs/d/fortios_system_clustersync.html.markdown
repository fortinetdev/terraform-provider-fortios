---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_clustersync"
description: |-
  Get information on an fortios system clustersync.
---

# Data Source: fortios_system_clustersync
Use this data source to get information on an fortios system clustersync

## Argument Reference

* `sync_id` - (Required) Specify the sync_id of the desired system clustersync.

## Attribute Reference

The following attributes are exported:

* `sync_id` - Sync ID.
* `peervd` - VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
* `peerip` - IP address of the interface on the peer unit that is used for the session synchronization link.
* `syncvd` - Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
* `down_intfs_before_sess_sync` - List of interfaces to be turned down before session synchronization is complete. The structure of `down_intfs_before_sess_sync` block is documented below.
* `hb_interval` - Heartbeat interval (1 - 10 sec).
* `hb_lost_threshold` - Lost heartbeat threshold (1 - 10).
* `ipsec_tunnel_sync` - Enable/disable IPsec tunnel synchronization.
* `slave_add_ike_routes` - Enable/disable IKE route announcement on the backup unit.
* `session_sync_filter` - Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `session_sync_filter` block is documented below.

The `syncvd` block contains:

* `name` - VDOM name.

The `down_intfs_before_sess_sync` block contains:

* `name` - Interface name.

The `session_sync_filter` block contains:

* `srcintf` - Only sessions from this interface are synchronized. You can only enter one interface name. To synchronize sessions for multiple source interfaces, add multiple filters.
* `dstintf` - Only sessions to this interface are synchronized. You can only enter one interface name. To synchronize sessions to multiple destination interfaces, add multiple filters.
* `srcaddr` - Only sessions from this IPv4 address are synchronized. You can only enter one address. To synchronize sessions from multiple source addresses, add multiple filters.
* `dstaddr` - Only sessions to this IPv4 address are synchronized. You can only enter one address. To synchronize sessions for multiple destination addresses, add multiple filters.
* `srcaddr6` - Only sessions from this IPv6 address are synchronized. You can only enter one address. To synchronize sessions from multiple source addresses, add multiple filters.
* `dstaddr6` - Only sessions to this IPv6 address are synchronized. You can only enter one address. To synchronize sessions for multiple destination addresses, add multiple filters.
* `custom_service` - Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custome services. The structure of `custom_service` block is documented below.

The `custom_service` block contains:

* `id` - Custom service ID.
* `src_port_range` - Custom service source port range.
* `dst_port_range` - Custom service destination port range.

