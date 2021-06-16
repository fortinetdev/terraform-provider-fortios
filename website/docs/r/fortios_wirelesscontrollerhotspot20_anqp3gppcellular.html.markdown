---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqp3gppcellular"
description: |-
  Configure 3GPP public land mobile network (PLMN).
---

# fortios_wirelesscontrollerhotspot20_anqp3gppcellular
Configure 3GPP public land mobile network (PLMN).

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_anqp3gppcellular" "trname" {
  name = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - 3GPP PLMN name.
* `mcc_mnc_list` - Mobile Country Code and Mobile Network Code configuration. The structure of `mcc_mnc_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `mcc_mnc_list` block supports:

* `id` - ID.
* `mcc` - Mobile country code.
* `mnc` - Mobile network code.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 Anqp3GppCellular can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_anqp3gppcellular.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
