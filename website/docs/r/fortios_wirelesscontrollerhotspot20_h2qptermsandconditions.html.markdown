---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qptermsandconditions"
description: |-
  Configure terms and conditions.
---

# fortios_wirelesscontrollerhotspot20_h2qptermsandconditions
Configure terms and conditions. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `name` - Terms and Conditions ID.
* `filename` - Filename.
* `timestamp` - Timestamp.
* `url` - URL.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 H2QpTermsAndConditions can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_h2qptermsandconditions.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_h2qptermsandconditions.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
