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

* `prefer_scenario` - Set the preferred cache behavior towards the balance between latency and hit-ratio. Valid values: `balance`, `prefer-speed`, `prefer-cache`.
* `collaboration` - Enable/disable cache-collaboration between cache-service clusters. Valid values: `enable`, `disable`.
* `device_id` - Set identifier for this cache device.
* `acceptable_connections` - Set strategy when accepting cache collaboration connection. Valid values: `any`, `peers`.
* `dst_peer` - Modify cache-service destination peer list. The structure of `dst_peer` block is documented below.
* `src_peer` - Modify cache-service source peer list. The structure of `src_peer` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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
$ terraform import fortios_wanopt_cacheservice.labelname WanoptCacheService

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wanopt_cacheservice.labelname WanoptCacheService
$ unset "FORTIOS_IMPORT_TABLE"
```
