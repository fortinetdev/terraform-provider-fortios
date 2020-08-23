---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzer2_overridefilter"
description: |-
  Override filters for FortiAnalyzer.
---

# fortios_logfortianalyzer2_overridefilter
Override filters for FortiAnalyzer.

## Example Usage

```hcl
resource "fortios_logfortianalyzer2_overridefilter" "trname" {
  anomaly           = "enable"
  dlp_archive       = "enable"
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
* `dlp_archive` - Enable/disable DLP archive logging.
* `gtp` - Enable/disable GTP messages logging.
* `dns` - Enable/disable detailed DNS event logging.
* `ssh` - Enable/disable SSH logging.
* `filter` - FortiAnalyzer 2 log filter.
* `filter_type` - Include/exclude logs that match the filter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortianalyzer2 OverrideFilter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logfortianalyzer2_overridefilter.labelname LogFortianalyzer2OverrideFilter
$ unset "FORTIOS_IMPORT_TABLE"
```
