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
* `value_list` - Name list. The structure of `value_list` block is documented below.

The `value_list` block supports:

* `index` - Value index.
* `lang` - Language code.
* `value` - Friendly name value.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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
