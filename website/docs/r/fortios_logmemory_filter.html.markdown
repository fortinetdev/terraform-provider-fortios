---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logmemory_filter"
description: |-
  Filters for memory buffer.
---

# fortios_logmemory_filter
Filters for memory buffer.

## Example Usage

```hcl
resource "fortios_logmemory_filter" "trname" {
  anomaly           = "enable"
  dns               = "enable"
  filter_type       = "include"
  forward_traffic   = "enable"
  gtp               = "enable"
  local_traffic     = "enable"
  multicast_traffic = "enable"
  severity          = "information"
  sniffer_traffic   = "enable"
  ssh               = "enable"
  voip              = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `severity` - Log every message above and including this severity level.
* `forward_traffic` - Enable/disable forward traffic logging.
* `local_traffic` - Enable/disable local in or out traffic logging.
* `multicast_traffic` - Enable/disable multicast traffic logging.
* `sniffer_traffic` - Enable/disable sniffer traffic logging.
* `anomaly` - Enable/disable anomaly logging.
* `netscan_discovery` - Enable/disable netscan discovery event logging.
* `netscan_vulnerability` - Enable/disable netscan vulnerability event logging.
* `voip` - Enable/disable VoIP logging.
* `gtp` - Enable/disable GTP messages logging.
* `free_style` - Free Style Filters The structure of `free_style` block is documented below.
* `dns` - Enable/disable detailed DNS event logging.
* `ssh` - Enable/disable SSH logging.
* `event` - Enable/disable event logging.
* `system` - Enable/disable system activity logging.
* `radius` - Enable/disable RADIUS messages logging.
* `ipsec` - Enable/disable IPsec negotiation messages logging.
* `dhcp` - Enable/disable DHCP service messages logging.
* `ppp` - Enable/disable L2TP/PPTP/PPPoE logging.
* `admin` - Enable/disable admin login/logout logging.
* `ha` - Enable/disable HA logging.
* `auth` - Enable/disable firewall authentication logging.
* `pattern` - Enable/disable pattern update logging.
* `sslvpn_log_auth` - Enable/disable SSL user authentication logging.
* `sslvpn_log_adm` - Enable/disable SSL administrator login logging.
* `sslvpn_log_session` - Enable/disable SSL session logging.
* `vip_ssl` - Enable/disable VIP SSL logging.
* `ldb_monitor` - Enable/disable VIP real server health monitoring logging.
* `wan_opt` - Enable/disable WAN optimization event logging.
* `wireless_activity` - Enable/disable wireless activity event logging.
* `cpu_memory_usage` - Enable/disable CPU & memory usage logging every 5 minutes.
* `filter` - Memory log filter.
* `filter_type` - Include/exclude logs that match the filter.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `free_style` block supports:

* `id` - Entry ID.
* `category` - Log category.
* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogMemory Filter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logmemory_filter.labelname LogMemoryFilter
$ unset "FORTIOS_IMPORT_TABLE"
```
