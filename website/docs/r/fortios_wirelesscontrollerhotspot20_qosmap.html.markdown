---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_qosmap"
description: |-
  Configure QoS map set.
---

# fortios_wirelesscontrollerhotspot20_qosmap
Configure QoS map set.

## Argument Reference

The following arguments are supported:

* `name` - QOS-MAP name.
* `dscp_except` - Differentiated Services Code Point (DSCP) exceptions. The structure of `dscp_except` block is documented below.
* `dscp_range` - Differentiated Services Code Point (DSCP) ranges. The structure of `dscp_range` block is documented below.

The `dscp_except` block supports:

* `index` - DSCP exception index.
* `dscp` - DSCP value.
* `up` - User priority.

The `dscp_range` block supports:

* `index` - DSCP range index.
* `up` - User priority.
* `low` - DSCP low value.
* `high` - DSCP high value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 QosMap can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_qosmap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
