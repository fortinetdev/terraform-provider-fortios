---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_replacemsgimage"
description: |-
  Configure replacement message images.
---

# fortios_system_replacemsgimage
Configure replacement message images.

## Example Usage

```hcl
resource "fortios_system_replacemsgimage" "trname" {
  image_type   = "png"
  name         = "logo_S1"
  image_base64 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEWAAABFgAVshLGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII="
}
```

## Argument Reference

The following arguments are supported:

* `name` - Image name.
* `image_type` - Image type. Valid values: `gif`, `jpg`, `tiff`, `png`.
* `image_base64` - Image data.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ReplacemsgImage can be imported using any of these accepted formats:
```
$ terraform import fortios_system_replacemsgimage.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_replacemsgimage.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
