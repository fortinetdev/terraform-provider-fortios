---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpwanmetric"
description: |-
  Configure WAN metrics.
---

# fortios_wirelesscontrollerhotspot20_h2qpwanmetric
Configure WAN metrics.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_h2qpwanmetric" "trname" {
  downlink_load             = 0
  downlink_speed            = 2400
  link_at_capacity          = "disable"
  link_status               = "up"
  load_measurement_duration = 0
  name                      = "1"
  symmetric_wan_link        = "symmetric"
  uplink_load               = 0
  uplink_speed              = 2400
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) WAN metric name.
* `link_status` - Link status.
* `symmetric_wan_link` - WAN link symmetry.
* `link_at_capacity` - Link at capacity.
* `uplink_speed` - Uplink speed (in kilobits/s).
* `downlink_speed` - Downlink speed (in kilobits/s).
* `uplink_load` - Uplink load.
* `downlink_load` - Downlink load.
* `load_measurement_duration` - Load measurement duration (in tenths of a second).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 H2QpWanMetric can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_h2qpwanmetric.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
