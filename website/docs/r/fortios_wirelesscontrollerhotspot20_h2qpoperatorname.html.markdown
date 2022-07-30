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
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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
$ terraform import fortios_wirelesscontrollerhotspot20_h2qpoperatorname.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_h2qpoperatorname.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
