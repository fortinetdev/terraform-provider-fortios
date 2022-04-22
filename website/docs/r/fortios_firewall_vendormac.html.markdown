---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vendormac"
description: |-
  Show vendor and the MAC address they have.
---

# fortios_firewall_vendormac
Show vendor and the MAC address they have. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - Vendor ID.
* `name` - Vendor name.
* `mac_number` - Total number of MAC addresses.
* `obsolete` - Indicates whether the Vendor ID can be used.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall VendorMac can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_vendormac.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_vendormac.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
