---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzercloud_filter"
description: |-
  Filters for FortiAnalyzer Cloud.
---

# fortios_logfortianalyzercloud_filter
Filters for FortiAnalyzer Cloud.

## Argument Reference

The following arguments are supported:

* `severity` - Lowest severity level to log.
* `forward_traffic` - Enable/disable forward traffic logging.
* `local_traffic` - Enable/disable local in or out traffic logging.
* `multicast_traffic` - Enable/disable multicast traffic logging.
* `sniffer_traffic` - Enable/disable sniffer traffic logging.
* `anomaly` - Enable/disable anomaly logging.
* `voip` - Enable/disable VoIP logging.
* `dlp_archive` - Enable/disable DLP archive logging.
* `gtp` - Enable/disable GTP messages logging.
* `free_style` - Free Style Filters The structure of `free_style` block is documented below.
* `filter` - FortiAnalyzer Cloud log filter.
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

LogFortianalyzerCloud Filter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logfortianalyzercloud_filter.labelname LogFortianalyzerCloudFilter
$ unset "FORTIOS_IMPORT_TABLE"
```
