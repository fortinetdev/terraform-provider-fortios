---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf"
description: |-
  Configure OSPF.
---

# fortios_router_ospf
Configure OSPF.

~> The provider supports the definition of Ospf-Interface in Router Ospf `fortios_router_ospf`, and also allows the definition of separate Ospf-Interface resources `fortios_routerospf_ospfinterface`, but do not use a `fortios_router_ospf` with in-line Ospf-Interface in conjunction with any `fortios_routerospf_ospfinterface` resources, otherwise conflicts and overwrite will occur.

~> The provider supports the definition of Network in Router Ospf `fortios_router_ospf`, and also allows the definition of separate Network resources `fortios_routerospf_network`, but do not use a `fortios_router_ospf` with in-line Network in conjunction with any `fortios_routerospf_network` resources, otherwise conflicts and overwrite will occur.

~> The provider supports the definition of Neighbor in Router Ospf `fortios_router_ospf`, and also allows the definition of separate Neighbor resources `fortios_routerospf_neighbor`, but do not use a `fortios_router_ospf` with in-line Neighbor in conjunction with any `fortios_routerospf_neighbor` resources, otherwise conflicts and overwrite will occur.



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

* `abr_type` - Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `distance_external` - Administrative external distance.
* `distance_inter_area` - Administrative inter-area distance.
* `distance_intra_area` - Administrative intra-area distance.
* `database_overflow` - Enable/disable database overflow. Valid values: `enable`, `disable`.
* `database_overflow_max_lsas` - Database overflow maximum LSAs.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `default_information_originate` - Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type. Valid values: `1`, `2`.
* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `distance` - Distance of the route.
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility. Valid values: `enable`, `disable`.
* `router_id` - (Required) Router ID.
* `spf_timers` - SPF calculation frequency.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `log_neighbour_changes` - Enable logging of OSPF neighbour's changes Valid values: `enable`, `disable`.
* `distribute_list_in` - Filter incoming routes.
* `distribute_route_map_in` - Filter incoming external routes by route-map.
* `restart_mode` - OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.
* `restart_period` - Graceful restart period.
* `restart_on_topology_change` - Enable/disable continuing graceful restart upon topology change. Valid values: `enable`, `disable`.
* `area` - OSPF area configuration. The structure of `area` block is documented below.
* `ospf_interface` - OSPF interface configuration. The structure of `ospf_interface` block is documented below.
* `network` - OSPF network configuration. The structure of `network` block is documented below.
* `neighbor` - OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `summary_address` - IP address summary configuration. The structure of `summary_address` block is documented below.
* `distribute_list` - Distribute list configuration. The structure of `distribute_list` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `area` block supports:

* `id` - Area entry IP address.
* `shortcut` - Enable/disable shortcut option. Valid values: `disable`, `enable`, `default`.
* `authentication` - Authentication type.
* `default_cost` - Summary default cost of stub or NSSA area.
* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate`, `never`, `always`.
* `stub_type` - Stub summary setting. Valid values: `no-summary`, `summary`.
* `type` - Area type setting. Valid values: `regular`, `nssa`, `stub`.
* `nssa_default_information_originate` - Redistribute, advertise, or do not originate Type-7 default route into NSSA area. Valid values: `enable`, `always`, `disable`.
* `nssa_default_information_originate_metric` - OSPF default metric.
* `nssa_default_information_originate_metric_type` - OSPF metric type for default routes. Valid values: `1`, `2`.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area. Valid values: `enable`, `disable`.
* `comments` - Comment.
* `range` - OSPF area range configuration. The structure of `range` block is documented below.
* `virtual_link` - OSPF virtual link configuration. The structure of `virtual_link` block is documented below.
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.

The `range` block supports:

* `id` - Range entry ID.
* `prefix` - Prefix.
* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.
* `substitute` - Substitute prefix.
* `substitute_status` - Enable/disable substitute status. Valid values: `enable`, `disable`.

The `virtual_link` block supports:

* `name` - Virtual link entry name.
* `authentication` - Authentication type.
* `authentication_key` - Authentication key.
* `keychain` - Message-digest key-chain name.
* `md5_key` - MD5 key.
* `md5_keychain` - Authentication MD5 key-chain name.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `peer` - Peer IP.
* `md5_keys` - MD5 key. The structure of `md5_keys` block is documented below.

The `md5_keys` block supports:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.

The `filter_list` block supports:

* `id` - Filter list entry ID.
* `list` - Access-list or prefix-list name.
* `direction` - Direction. Valid values: `in`, `out`.

The `ospf_interface` block supports:

* `name` - Interface entry name.
* `comments` - Comment.
* `interface` - Configuration interface name.
* `ip` - IP address.
* `authentication` - Authentication type.
* `authentication_key` - Authentication key.
* `keychain` - Message-digest key-chain name.
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
* `database_filter_out` - Enable/disable control of flooding out LSAs. Valid values: `enable`, `disable`.
* `mtu` - MTU for database description packets.
* `mtu_ignore` - Enable/disable ignore MTU. Valid values: `enable`, `disable`.
* `network_type` - Network type. Valid values: `broadcast`, `non-broadcast`, `point-to-point`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.
* `resync_timeout` - Graceful restart neighbor resynchronization timeout.
* `md5_keys` - MD5 key. The structure of `md5_keys` block is documented below.

The `md5_keys` block supports:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.

The `network` block supports:

* `id` - Network entry ID.
* `prefix` - Prefix.
* `area` - Attach the network to area.
* `comments` - Comment.

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
* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.

The `distribute_list` block supports:

* `id` - Distribute list entry ID.
* `access_list` - Access list name.
* `protocol` - Protocol type. Valid values: `connected`, `static`, `rip`.

The `redistribute` block supports:

* `name` - Redistribute name.
* `status` - status Valid values: `enable`, `disable`.
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `metric_type` - Metric type. Valid values: `1`, `2`.
* `tag` - Tag value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ospf can be imported using any of these accepted formats:
```
$ terraform import fortios_router_ospf.labelname RouterOspf

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_ospf.labelname RouterOspf
$ unset "FORTIOS_IMPORT_TABLE"
```
