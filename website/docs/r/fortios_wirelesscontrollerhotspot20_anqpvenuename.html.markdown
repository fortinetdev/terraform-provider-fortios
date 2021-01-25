---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpvenuename"
description: |-
  Configure venue name duple.
---

# fortios_wirelesscontrollerhotspot20_anqpvenuename
Configure venue name duple.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_anqpvenuename" "trname" {
  name = "1"
  value_list {
    index = 1
    lang  = "CN"
    value = "3"
  }
}
```

## Argument Reference


The following arguments are supported:

* `name` - Name of venue name duple.
* `value_list` - Name list. The structure of `value_list` block is documented below.

The `value_list` block supports:

* `index` - Value index.
* `lang` - Language code.
* `value` - Venue name value.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpVenueName can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_anqpvenuename.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
