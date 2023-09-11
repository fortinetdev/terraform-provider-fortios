---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_nacsettings"
description: |-
  Configure integrated NAC settings for FortiSwitch.
---

# fortios_switchcontroller_nacsettings
Configure integrated NAC settings for FortiSwitch. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,7.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - NAC settings name.
* `mode` - Set NAC mode to be used on the FortiSwitch ports. Valid values: `local`, `global`.
* `inactive_timer` - Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered.
* `auto_auth` - Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable`, `enable`.
* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
* `link_down_flush` - Clear NAC devices on switch ports on link down event. Valid values: `disable`, `enable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController NacSettings can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_nacsettings.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_nacsettings.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
