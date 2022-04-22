---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_apstatus"
description: |-
  Configure access point status (rogue | accepted | suppressed).
---

# fortios_wirelesscontroller_apstatus
Configure access point status (rogue | accepted | suppressed).

## Argument Reference

The following arguments are supported:

* `fosid` - AP ID.
* `bssid` - Access Point's (AP's) BSSID.
* `ssid` - Access Point's (AP's) SSID.
* `status` - Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController ApStatus can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_apstatus.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_apstatus.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
