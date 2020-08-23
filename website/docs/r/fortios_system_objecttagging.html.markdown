---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_objecttagging"
description: |-
  Configure object tagging.
---

# fortios_system_objecttagging
Configure object tagging.

## Example Usage

```hcl
resource "fortios_system_objecttagging" "trname" {
  address   = "disable"
  category  = "s1"
  color     = 0
  device    = "mandatory"
  interface = "disable"
  multiple  = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `category` - Tag Category.
* `address` - Address.
* `device` - Device.
* `interface` - Interface.
* `multiple` - Allow multiple tag selection.
* `color` - Color of icon on the GUI.
* `tags` - Tags.

The `tags` block supports:

* `name` - Tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{category}}.

## Import

System ObjectTagging can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_objecttagging.labelname {{category}}
$ unset "FORTIOS_IMPORT_TABLE"
```
