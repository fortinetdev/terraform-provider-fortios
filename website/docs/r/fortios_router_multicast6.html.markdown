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
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `interface` block supports:

* `name` - Interface name.
* `hello_interval` - Interval between sending PIM hello messages in seconds (1 - 65535, default = 30).
* `hello_holdtime` - Time before old neighbor information expires in seconds (1 - 65535, default = 105).

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
$ terraform import fortios_router_multicast6.labelname RouterMulticast6

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_multicast6.labelname RouterMulticast6
$ unset "FORTIOS_IMPORT_TABLE"
```
