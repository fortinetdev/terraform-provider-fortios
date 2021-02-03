---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_address"
description: |-
  Configure the client with its MAC address.
---

# fortios_wirelesscontroller_address
Configure the client with its MAC address.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `mac` - MAC address.
* `policy` - Allow or block the client with this MAC address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController Address can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_address.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
