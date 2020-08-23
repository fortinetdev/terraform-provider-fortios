---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf"
description: |-
  Configure OSPF.
---

# fortios_router_ospf
Configure OSPF.

## Example Usage

```hcl
resource "fortios_router_ospf" "trname" {
  abr_type                          = "standard"
  auto_cost_ref_bandwidth           = 1000
  bfd                               = "disable"
  database_overflow                 = "disable"
  database_overflow_max_lsas        = 10000
  database_overflow_time_to_recover = 300
  default_information_metric        = 10
  default_information_metric_type   = "2"
  default_information_originate     = "disable"
  default_metric                    = 10
  distance                          = 110
  distance_external                 = 110
  distance_inter_area               = 110
  distance_intra_area               = 110
  log_neighbour_changes             = "enable"
  restart_mode                      = "none"
  restart_period                    = 120
  rfc1583_compatible                = "disable"
  router_id                         = "0.0.0.0"
  spf_timers                        = "5 10"

  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "connected"
    status      = "disable"
    tag         = 0
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "static"
    status      = "disable"
    tag         = 0
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "rip"
    status      = "disable"
    tag         = 0
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "bgp"
    status      = "disable"
    tag         = 0
  }
  redistribute {
    metric      = 0
    metric_type = "2"
    name        = "isis"
    status      = "disable"
    tag         = 0
  }
}
```

## Argument Reference

The following arguments are supported:

* `abr_type` - Area border router type.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `distance_external` - Administrative external distance.
* `distance_inter_area` - Administrative inter-area distance.
* `distance_intra_area` - Administrative intra-area distance.
* `database_overflow` - Enable/disable database overflow.
* `database_overflow_max_lsas` - Database overflow maximum LSAs.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `default_information_originate` - Enable/disable generation of default route.
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type.
* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `distance` - Distance of the route.
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility.
* `router_id` - (Required) Router ID.
* `spf_timers` - SPF calculation frequency.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `log_neighbour_changes` - Enable logging of OSPF neighbour's changes
* `distribute_list_in` - Filter incoming routes.
* `distribute_route_map_in` - Filter incoming external routes by route-map.
* `restart_mode` - OSPF restart mode (graceful or LLS).
* `restart_period` - Graceful restart period.
* `area` - OSPF area configuration.
* `ospf_interface` - OSPF interface configuration.
* `network` - OSPF network configuration.
* `neighbor` - OSPF neighbor configuration are used when OSPF runs on non-broadcast media
* `passive_interface` - Passive interface configuration.
* `summary_address` - IP address summary configuration.
* `distribute_list` - Distribute list configuration.
* `redistribute` - Redistribute configuration.

The `area` block supports:

* `id` - Area entry IP address.
* `shortcut` - Enable/disable shortcut option.
* `authentication` - Authentication type.
* `default_cost` - Summary default cost of stub or NSSA area.
* `nssa_translator_role` - NSSA translator role type.
* `stub_type` - Stub summary setting.
* `type` - Area type setting.
* `nssa_default_information_originate` - Redistribute, advertise, or do not originate Type-7 default route into NSSA area.
* `nssa_default_information_originate_metric` - OSPF default metric.
* `nssa_default_information_originate_metric_type` - OSPF metric type for default routes.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area.
* `range` - OSPF area range configuration.
* `virtual_link` - OSPF virtual link configuration.
* `filter_list` - OSPF area filter-list configuration.

The `range` block supports:

* `id` - Range entry ID.
* `prefix` - Prefix.
* `advertise` - Enable/disable advertise status.
* `substitute` - Substitute prefix.
* `substitute_status` - Enable/disable substitute status.

The `virtual_link` block supports:

* `name` - Virtual link entry name.
* `authentication` - Authentication type.
* `authentication_key` - Authentication key.
* `md5_key` - MD5 key.
* `md5_keychain` - Authentication MD5 key-chain name.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `peer` - Peer IP.

The `filter_list` block supports:

* `id` - Filter list entry ID.
* `list` - Access-list or prefix-list name.
* `direction` - Direction.

The `ospf_interface` block supports:

* `name` - Interface entry name.
* `interface` - Configuration interface name.
* `ip` - IP address.
* `authentication` - Authentication type.
* `authentication_key` - Authentication key.
* `md5_key` - MD5 key.
* `md5_keychain` - Authentication MD5 key-chain name.
* `prefix_length` - Prefix length.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - Priority.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `hello_multiplier` - Number of hello packets within dead interval.
* `database_filter_out` - Enable/disable control of flooding out LSAs.
* `mtu` - MTU for database description packets.
* `mtu_ignore` - Enable/disable ignore MTU.
* `network_type` - Network type.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `status` - Enable/disable status.
* `resync_timeout` - Graceful restart neighbor resynchronization timeout.

The `network` block supports:

* `id` - Network entry ID.
* `prefix` - Prefix.
* `area` - Attach the network to area.

The `neighbor` block supports:

* `id` - Neighbor entry ID.
* `ip` - Interface IP address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - Priority.

The `passive_interface` block supports:

* `name` - Passive interface name.

The `summary_address` block supports:

* `id` - Summary address entry ID.
* `prefix` - Prefix.
* `tag` - Tag value.
* `advertise` - Enable/disable advertise status.

The `distribute_list` block supports:

* `id` - Distribute list entry ID.
* `access_list` - Access list name.
* `protocol` - Protocol type.

The `redistribute` block supports:

* `name` - Redistribute name.
* `status` - status
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `metric_type` - Metric type.
* `tag` - Tag value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ospf can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_ospf.labelname RouterOspf
$ unset "FORTIOS_IMPORT_TABLE"
```
