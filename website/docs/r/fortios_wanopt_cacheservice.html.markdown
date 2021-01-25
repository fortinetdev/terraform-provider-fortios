---
subcategory: "FortiGate WANOpt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_cacheservice"
description: |-
  Designate cache-service for wan-optimization and webcache.
---

# fortios_wanopt_cacheservice
Designate cache-service for wan-optimization and webcache.

## Example Usage

```hcl
resource "fortios_wanopt_cacheservice" "trname" {
  acceptable_connections = "any"
  collaboration          = "disable"
  device_id              = "default_dev_id"
  prefer_scenario        = "balance"
}
```

## Argument Reference


The following arguments are supported:

* `prefer_scenario` - Set the preferred cache behavior towards the balance between latency and hit-ratio.
* `collaboration` - Enable/disable cache-collaboration between cache-service clusters.
* `device_id` - Set identifier for this cache device.
* `acceptable_connections` - Set strategy when accepting cache collaboration connection.
* `dst_peer` - Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
* `src_peer` - Modify cache-service source peer list. The structure of `src_peer` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dst_peer` block supports:

* `device_id` - Device ID of this peer.
* `auth_type` - Set authentication type for this peer.
* `encode_type` - Set encode type for this peer.
* `priority` - Set priority for this peer.
* `ip` - Set cluster IP address of this peer.

The `src_peer` block supports:

* `device_id` - Device ID of this peer.
* `auth_type` - Set authentication type for this peer.
* `encode_type` - Set encode type for this peer.
* `priority` - Set priority for this peer.
* `ip` - Set cluster IP address of this peer.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wanopt CacheService can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wanopt_cacheservice.labelname WanoptCacheService
$ unset "FORTIOS_IMPORT_TABLE"
```
