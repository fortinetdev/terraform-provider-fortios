---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_lognulldevice_filter"
description: |-
  Filters for null device logging.
---

# fortios_lognulldevice_filter
Filters for null device logging.

## Example Usage

```hcl
resource "fortios_lognulldevice_filter" "trname" {
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

* `severity` - Lowest severity level to log.
* `forward_traffic` - Enable/disable forward traffic logging.
* `local_traffic` - Enable/disable local in or out traffic logging.
* `multicast_traffic` - Enable/disable multicast traffic logging.
* `sniffer_traffic` - Enable/disable sniffer traffic logging.
* `anomaly` - Enable/disable anomaly logging.
* `netscan_discovery` - Enable/disable netscan discovery event logging.
* `netscan_vulnerability` - Enable/disable netscan vulnerability event logging.
* `voip` - Enable/disable VoIP logging.
* `gtp` - Enable/disable GTP messages logging.
* `dns` - Enable/disable detailed DNS event logging.
* `ssh` - Enable/disable SSH logging.
* `filter` - Null-device log filter.
* `filter_type` - Include/exclude logs that match the filter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogNullDevice Filter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_lognulldevice_filter.labelname LogNullDeviceFilter
$ unset "FORTIOS_IMPORT_TABLE"
```
