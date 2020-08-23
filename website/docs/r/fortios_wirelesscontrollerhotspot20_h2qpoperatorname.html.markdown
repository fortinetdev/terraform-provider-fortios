---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpoperatorname"
description: |-
  Configure operator friendly name.
---

# fortios_wirelesscontrollerhotspot20_h2qpoperatorname
Configure operator friendly name.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_h2qpoperatorname" "trname" {
  name = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Friendly name ID.
* `value_list` - Name list.

The `value_list` block supports:

* `index` - Value index.
* `lang` - Language code.
* `value` - Friendly name value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 H2QpOperatorName can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_h2qpoperatorname.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
