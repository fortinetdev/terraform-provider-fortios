---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_region"
description: |-
  Configure FortiAP regions (for floor plans and maps).
---

# fortios_wirelesscontroller_region
Configure FortiAP regions (for floor plans and maps).

## Argument Reference

The following arguments are supported:

* `name` - FortiAP region name.
* `image_type` - FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
* `comments` - Comments.
* `grayscale` - Region image grayscale. Valid values: `enable`, `disable`.
* `opacity` - Region image opacity (0 - 100).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Region can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_region.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_region.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
