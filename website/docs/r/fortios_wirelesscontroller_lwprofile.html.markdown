---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_lwprofile"
description: |-
  Configure LoRaWAN profile.
---

# fortios_wirelesscontroller_lwprofile
Configure LoRaWAN profile. Applies to FortiOS Version `>= 7.6.5`.

## Argument Reference

The following arguments are supported:

* `name` - LoRaWAN profile name.
* `comment` - Comment.
* `lw_protocol` - Configure LoRaWAN protocol (default = basics-station) Valid values: `basics-station`, `packet-forwarder`.
* `cups_server` - CUPS (Configuration and Update Server) domain name or IP address of LoRaWAN device.
* `cups_server_port` - CUPS Port value of LoRaWAN device.
* `cups_api_key` - CUPS API key of LoRaWAN device.
* `tc_server` - TC (Traffic Controller) domain name or IP address of LoRaWAN device.
* `tc_server_port` - TC Port value of LoRaWAN device.
* `tc_api_key` - TC API key of LoRaWAN device.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController LwProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_lwprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_lwprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
