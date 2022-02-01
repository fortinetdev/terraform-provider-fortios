---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpvenueurl"
description: |-
  Configure venue URL.
---

# fortios_wirelesscontrollerhotspot20_anqpvenueurl
Configure venue URL. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `name` - Name of venue url.
* `value_list` - URL list. The structure of `value_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `value_list` block supports:

* `index` - URL index.
* `number` - Venue number.
* `value` - Venue URL value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpVenueUrl can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_anqpvenueurl.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
