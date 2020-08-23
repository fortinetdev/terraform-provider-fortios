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
* `image_type` - Image type.
* `image_base64` - Image data.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ReplacemsgImage can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_replacemsgimage.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
