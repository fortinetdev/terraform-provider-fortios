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

* `name` - WAN metric name.
* `link_status` - Link status. Valid values: `up`, `down`, `in-test`.
* `symmetric_wan_link` - WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
* `link_at_capacity` - Link at capacity. Valid values: `enable`, `disable`.
* `uplink_speed` - Uplink speed (in kilobits/s).
* `downlink_speed` - Downlink speed (in kilobits/s).
* `uplink_load` - Uplink load.
* `downlink_load` - Downlink load.
* `load_measurement_duration` - Load measurement duration (in tenths of a second).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


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
