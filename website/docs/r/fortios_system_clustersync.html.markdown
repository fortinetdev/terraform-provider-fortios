---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_clustersync"
description: |-
  Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.
---

# fortios_system_clustersync
Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.

## Example Usage

```hcl
resource "fortios_system_clustersync" "trname" {
  hb_interval          = 3
  hb_lost_threshold    = 3
  peerip               = "1.1.1.1"
  peervd               = "root"
  slave_add_ike_routes = "enable"
  sync_id              = 1
}
```

## Argument Reference

The following arguments are supported:

* `sync_id` - Sync ID.
* `peervd` - VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
* `peerip` - IP address of the interface on the peer unit that is used for the session synchronization link.
* `syncvd` - Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.
* `down_intfs_before_sess_sync` - List of interfaces to be turned down before session synchronization is complete. The structure of `down_intfs_before_sess_sync` block is documented below.
* `hb_interval` - Heartbeat interval (1 - 10 sec).
* `hb_lost_threshold` - Lost heartbeat threshold (1 - 10).
* `slave_add_ike_routes` - Enable/disable IKE route announcement on the backup unit.
* `session_sync_filter` - Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `session_sync_filter` block is documented below.

The `syncvd` block supports:

* `name` - VDOM name.

The `down_intfs_before_sess_sync` block supports:

* `name` - Interface name.

The `session_sync_filter` block supports:

* `srcintf` - Only sessions from this interface are synchronized. You can only enter one interface name. To synchronize sessions for multiple source interfaces, add multiple filters.
* `dstintf` - Only sessions to this interface are synchronized. You can only enter one interface name. To synchronize sessions to multiple destination interfaces, add multiple filters.
* `srcaddr` - Only sessions from this IPv4 address are synchronized. You can only enter one address. To synchronize sessions from multiple source addresses, add multiple filters.
* `dstaddr` - Only sessions to this IPv4 address are synchronized. You can only enter one address. To synchronize sessions for multiple destination addresses, add multiple filters.
* `srcaddr6` - Only sessions from this IPv6 address are synchronized. You can only enter one address. To synchronize sessions from multiple source addresses, add multiple filters.
* `dstaddr6` - Only sessions to this IPv6 address are synchronized. You can only enter one address. To synchronize sessions for multiple destination addresses, add multiple filters.
* `custom_service` - Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custome services. The structure of `custom_service` block is documented below.

The `custom_service` block supports:

* `id` - Custom service ID.
* `src_port_range` - Custom service source port range.
* `dst_port_range` - Custom service destination port range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{sync_id}}.

## Import

System ClusterSync can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_clustersync.labelname {{sync_id}}
$ unset "FORTIOS_IMPORT_TABLE"
```
