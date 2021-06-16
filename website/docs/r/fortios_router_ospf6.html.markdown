---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf6"
description: |-
  Configure IPv6 OSPF.
---

# fortios_router_ospf6
Configure IPv6 OSPF.

~> The provider supports the definition of Ospf6-Interface in Router Ospf6 `fortios_router_ospf6`, and also allows the definition of separate Ospf6-Interface resources `fortios_routerospf6_ospf6interface`, but do not use a `fortios_router_ospf6` with in-line Ospf6-Interface in conjunction with any `fortios_routerospf6_ospf6interface` resources, otherwise conflicts and overwrite will occur.



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

* `abr_type` - Area border router type. Valid values: `cisco`, `ibm`, `standard`.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `default_information_originate` - Enable/disable generation of default route. Valid values: `enable`, `always`, `disable`.
* `log_neighbour_changes` - Enable logging of OSPFv3 neighbour's changes Valid values: `enable`, `disable`.
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type. Valid values: `1`, `2`.
* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `router_id` - (Required) A.B.C.D, in IPv4 address format.
* `spf_timers` - SPF calculation frequency.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable`, `disable`.
* `area` - OSPF6 area configuration. The structure of `area` block is documented below.
* `ospf6_interface` - OSPF6 interface configuration. The structure of `ospf6_interface` block is documented below.
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.
* `summary_address` - IPv6 address summary configuration. The structure of `summary_address` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `area` block supports:

* `id` - Area entry IP address.
* `default_cost` - Summary default cost of stub or NSSA area.
* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate`, `never`, `always`.
* `stub_type` - Stub summary setting. Valid values: `no-summary`, `summary`.
* `type` - Area type setting. Valid values: `regular`, `nssa`, `stub`.
* `nssa_default_information_originate` - Enable/disable originate type 7 default into NSSA area. Valid values: `enable`, `disable`.
* `nssa_default_information_originate_metric` - OSPFv3 default metric.
* `nssa_default_information_originate_metric_type` - OSPFv3 metric type for default routes. Valid values: `1`, `2`.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area. Valid values: `enable`, `disable`.
* `authentication` - Authentication mode. Valid values: `none`, `ah`, `esp`.
* `key_rollover_interval` - Key roll-over interval.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
* `range` - OSPF6 area range configuration. The structure of `range` block is documented below.
* `virtual_link` - OSPF6 virtual link configuration. The structure of `virtual_link` block is documented below.

The `ipsec_keys` block supports:

* `spi` - Security Parameters Index.
* `auth_key` - Authentication key.
* `enc_key` - Encryption key.

The `range` block supports:

* `id` - Range entry ID.
* `prefix6` - IPv6 prefix.
* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.

The `virtual_link` block supports:

* `name` - Virtual link entry name.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `peer` - A.B.C.D, peer router ID.
* `authentication` - Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
* `key_rollover_interval` - Key roll-over interval.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block supports:

* `spi` - Security Parameters Index.
* `auth_key` - Authentication key.
* `enc_key` - Encryption key.

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
* `status` - Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.
* `network_type` - Network type. Valid values: `broadcast`, `point-to-point`, `non-broadcast`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.
* `mtu` - MTU for OSPFv3 packets.
* `mtu_ignore` - Enable/disable ignoring MTU field in DBD packets. Valid values: `enable`, `disable`.
* `authentication` - Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.
* `key_rollover_interval` - Key roll-over interval.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.
* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.
* `neighbor` - OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media The structure of `neighbor` block is documented below.

The `ipsec_keys` block supports:

* `spi` - Security Parameters Index.
* `auth_key` - Authentication key.
* `enc_key` - Encryption key.

The `neighbor` block supports:

* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `priority` - priority

The `redistribute` block supports:

* `name` - Redistribute name.
* `status` - status Valid values: `enable`, `disable`.
* `metric` - Redistribute metric setting.
* `routemap` - Route map name.
* `metric_type` - Metric type. Valid values: `1`, `2`.

The `passive_interface` block supports:

* `name` - Passive interface name.

The `summary_address` block supports:

* `id` - Summary address entry ID.
* `prefix6` - IPv6 prefix.
* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.
* `tag` - Tag value.


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
