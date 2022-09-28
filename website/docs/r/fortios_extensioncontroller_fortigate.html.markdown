---
subcategory: "FortiGate Extension-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extensioncontroller_fortigate"
description: |-
  FortiGate controller configuration.
---

# fortios_extensioncontroller_fortigate
FortiGate controller configuration. Applies to FortiOS Version `>= 7.2.1`.

## Argument Reference

The following arguments are supported:

* `name` - FortiGate entry name.
* `fosid` - FortiGate serial number.
* `authorized` - Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
* `hostname` - FortiGate hostname.
* `description` - Description.
* `vdom` - VDOM.
* `device_id` - device-id
* `profile` - FortiGate profile configuration.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController Fortigate can be imported using any of these accepted formats:
```
$ terraform import fortios_extensioncontroller_fortigate.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_extensioncontroller_fortigate.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
