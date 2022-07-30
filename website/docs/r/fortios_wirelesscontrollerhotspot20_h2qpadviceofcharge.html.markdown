---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge"
description: |-
  Configure advice of charge.
---

# fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge
Configure advice of charge. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `name` - Plan name.
* `aoc_list` - AOC list. The structure of `aoc_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `aoc_list` block supports:

* `name` - Advice of charge ID.
* `type` - Usage charge type. Valid values: `time-based`, `volume-based`, `time-and-volume-based`, `unlimited`.
* `nai_realm_encoding` - NAI realm encoding.
* `nai_realm` - NAI realm list name.
* `plan_info` - Plan info. The structure of `plan_info` block is documented below.

The `plan_info` block supports:

* `name` - Plan name.
* `lang` - Languague code.
* `currency` - Currency code.
* `info_file` - Info file.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 H2QpAdviceOfCharge can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
