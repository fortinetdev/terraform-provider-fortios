---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_bleprofile"
description: |-
  Configure Bluetooth Low Energy profile.
---

# fortios_wirelesscontroller_bleprofile
Configure Bluetooth Low Energy profile.

## Argument Reference

The following arguments are supported:

* `name` - Bluetooth Low Energy profile name.
* `comment` - Comment.
* `advertising` - Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
* `ibeacon_uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `major_id` - Major ID.
* `minor_id` - Minor ID.
* `eddystone_namespace` - Eddystone namespace ID.
* `eddystone_instance` - Eddystone instance ID.
* `eddystone_url` - Eddystone URL.
* `eddystone_url_encode_hex` - Eddystone encoded URL hexadecimal string
* `txpower` - Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
* `beacon_interval` - Beacon interval (default = 100 msec).
* `ble_scanning` - Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController BleProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_bleprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_bleprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
