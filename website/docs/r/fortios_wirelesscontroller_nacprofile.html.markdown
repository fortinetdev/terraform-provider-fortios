---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_nacprofile"
description: |-
  Configure WiFi network access control (NAC) profiles.
---

# fortios_wirelesscontroller_nacprofile
Configure WiFi network access control (NAC) profiles. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `comment` - Comment.
* `onboarding_vlan` - VLAN interface name.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController NacProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_nacprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
