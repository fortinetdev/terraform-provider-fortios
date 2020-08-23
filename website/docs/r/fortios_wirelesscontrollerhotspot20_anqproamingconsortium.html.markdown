---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqproamingconsortium"
description: |-
  Configure roaming consortium.
---

# fortios_wirelesscontrollerhotspot20_anqproamingconsortium
Configure roaming consortium.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_anqproamingconsortium" "trname" {
  name = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Roaming consortium name.
* `oi_list` - Organization identifier list.

The `oi_list` block supports:

* `index` - OI index.
* `oi` - Organization identifier.
* `comment` - Comment.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpRoamingConsortium can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_anqproamingconsortium.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
