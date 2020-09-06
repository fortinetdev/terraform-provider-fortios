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
* `image_type` - FortiAP region image type (png|jpeg|gif).
* `comments` - Comments.
* `grayscale` - Region image grayscale.
* `opacity` - Region image opacity (0 - 100).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Region can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_region.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
