---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system3gmodem_custom"
description: |-
  3G MODEM custom.
---

# fortios_system3gmodem_custom
3G MODEM custom. Applies to FortiOS Version `7.0.4`.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `vendor` - MODEM vendor name.
* `model` - MODEM model name.
* `vendor_id` - USB vendor ID in hexadecimal format (0000-ffff).
* `product_id` - USB product ID in hexadecimal format (0000-ffff).
* `class_id` - USB interface class in hexadecimal format (00-ff).
* `init_string` - Init string in hexadecimal format (even length).
* `modeswitch_string` - USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System3GModem Custom can be imported using any of these accepted formats:
```
$ terraform import fortios_system3gmodem_custom.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system3gmodem_custom.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
