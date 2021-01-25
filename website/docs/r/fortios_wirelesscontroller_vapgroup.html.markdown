---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_vapgroup"
description: |-
  Configure virtual Access Point (VAP) groups.
---

# fortios_wirelesscontroller_vapgroup
Configure virtual Access Point (VAP) groups.

## Argument Reference


The following arguments are supported:

* `name` - (Required) Group Name
* `comment` - Comment.
* `vaps` - List of SSIDs to be included in the VAP group. The structure of `vaps` block is documented below.

The `vaps` block supports:

* `name` - vap name
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController VapGroup can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_vapgroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
