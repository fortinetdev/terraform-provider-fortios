---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_addrgrp"
description: |-
  Configure the MAC address group.
---

# fortios_wirelesscontroller_addrgrp
Configure the MAC address group. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1,6.4.2,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6`.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `default_policy` - Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
* `addresses` - Manually selected group of addresses. The structure of `addresses` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `addresses` block supports:

* `id` - Address ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController Addrgrp can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_addrgrp.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_addrgrp.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
