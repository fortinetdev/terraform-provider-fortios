---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf6"
description: |-
  Configure IPv6 OSPF.
---

# fortios_router_ospf6
Configure IPv6 OSPF.

## Example Usage

```hcl
resource "fortios_router_ospf6" "trname" {
  abr_type                        = "standard"
  auto_cost_ref_bandwidth         = 1000
  bfd                             = "disable"
  default_information_metric      = 10
  default_information_metric_type = "2"
  default_information_originate   = "disable"
  default_metric                  = 10
  log_neighbour_changes           = "enable"
  router_id                       = "0.0.0.0"
  spf_timers                      = "5 10"

  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "connected"
    status      = "disable"
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "static"
    status      = "disable"
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "rip"
    status      = "disable"
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "bgp"
    status      = "disable"
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "isis"
    status      = "disable"
  }
}
```

## Argument Reference


The following arguments are supported:

* `abr_type` - Area border router type.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `default_information_originate` - Enable/disable generation of default route.
* `log_neighbour_changes` - Enable logging of OSPFv3 neighbour's changes
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type.
* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `router_id` - (Required) A.B.C.D, in IPv4 address format.
* `spf_timers` - SPF calculation frequency.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).
* `area` - OSPF6 area configuration. The structure of `area` block is documented below.
* `ospf6_interface` - OSPF6 interface configuration. The structure of `ospf6_interface` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `summary_address` - IPv6 address summary configuration. The structure of `summary_address` block is documented below.

The `area` block supports:

* `id` - Area entry IP address.
* `default_cost` - Summary default cost of stub or NSSA area.
* `nssa_translator_role` - NSSA translator role type.
* `stub_type` - Stub summary setting.
* `type` - Area type setting.
* `nssa_default_information_originate` - Enable/disable originate type 7 default into NSSA area.
* `nssa_default_information_originate_metric` - OSPFv3 default metric.
* `nssa_default_information_originate_metric_type` - OSPFv3 metric type for default routes.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area.
* `range` - OSPF6 area range configuration. The structure of `range` block is documented below.
* `virtual_link` - OSPF6 virtual link configuration. The structure of `virtual_link` block is documented below.

The `range` block supports:

* `id` - Range entry ID.
* `prefix6` - IPv6 prefix.
* `advertise` - Enable/disable advertise status.

The `virtual_link` block supports:

* `name` - Virtual link entry name.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `peer` - A.B.C.D, peer router ID.

The `ospf6_interface` block supports:

* `name` - Interface entry name.
* `area_id` - A.B.C.D, in IPv4 address format.
* `interface` - Configuration interface name.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - priority
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `status` - Enable/disable OSPF6 routing on this interface.
* `network_type` - Network type.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).
* `mtu` - MTU for OSPFv3 packets.
* `mtu_ignore` - Enable/disable ignoring MTU field in DBD packets.
* `neighbor` - OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.

The `neighbor` block supports:

* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - priority

The `redistribute` block supports:

* `name` - Redistribute name.
* `status` - status
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `metric_type` - Metric type.

The `passive_interface` block supports:

* `name` - Passive interface name.

The `summary_address` block supports:

* `id` - Summary address entry ID.
* `prefix6` - IPv6 prefix.
* `advertise` - Enable/disable advertise status.
* `tag` - Tag value.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ospf6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_ospf6.labelname RouterOspf6
$ unset "FORTIOS_IMPORT_TABLE"
```
