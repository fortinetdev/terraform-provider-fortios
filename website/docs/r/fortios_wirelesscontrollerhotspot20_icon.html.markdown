---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_icon"
description: |-
  Configure OSU provider icon.
---

# fortios_wirelesscontrollerhotspot20_icon
Configure OSU provider icon.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_icon" "trname" {
  name = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Icon list ID.
* `icon_list` - Icon list. The structure of `icon_list` block is documented below.

The `icon_list` block supports:

* `name` - Icon name.
* `lang` - Language code.
* `file` - Icon file.
* `type` - Icon type.
* `width` - Icon width.
* `height` - Icon height.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 Icon can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_icon.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
