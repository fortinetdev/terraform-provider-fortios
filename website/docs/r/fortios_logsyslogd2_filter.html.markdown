---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logsyslogd2_filter"
description: |-
  Filters for remote system server.
---

# fortios_logsyslogd2_filter
Filters for remote system server.

## Example Usage

```hcl
resource "fortios_logsyslogd2_filter" "trname" {
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

* `severity` - Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `forward_traffic` - Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
* `local_traffic` - Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
* `multicast_traffic` - Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
* `sniffer_traffic` - Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
* `anomaly` - Enable/disable anomaly logging. Valid values: `enable`, `disable`.
* `netscan_discovery` - Enable/disable netscan discovery event logging.
* `netscan_vulnerability` - Enable/disable netscan vulnerability event logging.
* `voip` - Enable/disable VoIP logging. Valid values: `enable`, `disable`.
* `gtp` - Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
* `free_style` - Free Style Filters The structure of `free_style` block is documented below.
* `dns` - Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
* `ssh` - Enable/disable SSH logging. Valid values: `enable`, `disable`.
* `filter` - Syslog 2 filter.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `free_style` block supports:

* `id` - Entry ID.
* `category` - Log category. Valid values: `traffic`, `event`, `virus`, `webfilter`, `attack`, `spam`, `anomaly`, `voip`, `dlp`, `app-ctrl`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `icap`.
* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include`, `exclude`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogSyslogd2 Filter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logsyslogd2_filter.labelname LogSyslogd2Filter
$ unset "FORTIOS_IMPORT_TABLE"
```
