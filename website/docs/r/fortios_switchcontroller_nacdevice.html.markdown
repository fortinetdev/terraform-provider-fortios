---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_nacdevice"
description: |-
  Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.
---

# fortios_switchcontroller_nacdevice
Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - Device ID.
* `description` - Description for the learned NAC device.
* `status` - Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
* `mac` - MAC address of the learned NAC device.
* `last_known_switch` - Managed FortiSwitch where NAC device is last learned.
* `last_known_port` - Managed FortiSwitch port where NAC device is last learned.
* `matched_nac_policy` - Matched NAC policy for the learned NAC device.
* `port_policy` - Port policy to be applied on this learned NAC device.
* `mac_policy` - MAC policy to be applied on this learned NAC device.
* `last_seen` - Device last seen.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController NacDevice can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_nacdevice.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
