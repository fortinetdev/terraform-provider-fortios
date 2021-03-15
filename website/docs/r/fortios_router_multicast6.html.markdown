---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicast6"
description: |-
  Configure IPv6 multicast.
---

# fortios_router_multicast6
Configure IPv6 multicast.

## Example Usage

```hcl
resource "fortios_router_multicast6" "trname" {
  multicast_pmtu    = "disable"
  multicast_routing = "disable"

  pim_sm_global {
    register_rate_limit = 0
  }
}
```

## Argument Reference

The following arguments are supported:

* `multicast_routing` - Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
* `multicast_pmtu` - Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
* `interface` - Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
* `pim_sm_global` - PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `interface` block supports:

* `name` - Interface name.
* `hello_interval` - Interval between sending PIM hello messages  (1 - 65535 sec, default = 30)..
* `hello_holdtime` - Time before old neighbour information expires (1 - 65535 sec, default = 105).

The `pim_sm_global` block supports:

* `register_rate_limit` - Limit of packets/sec per source registered through this RP (0 means unlimited).
* `rp_address` - Statically configured RP addresses. The structure of `rp_address` block is documented below.

The `rp_address` block supports:

* `id` - ID of the entry.
* `ip6_address` - RP router IPv6 address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Multicast6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_multicast6.labelname RouterMulticast6
$ unset "FORTIOS_IMPORT_TABLE"
```
