---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerinitialconfig_vlans"
description: |-
  Configure initial template for auto-generated VLAN interfaces.
---

# fortios_switchcontrollerinitialconfig_vlans
Configure initial template for auto-generated VLAN interfaces. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `optional_vlans` - Auto-generate pre-configured VLANs upon switch discovery. Valid values: `enable`, `disable`.
* `default_vlan` - Default VLAN (native) assigned to all switch ports upon discovery.
* `quarantine` - VLAN for quarantined traffic.
* `rspan` - VLAN for RSPAN/ERSPAN mirrored traffic.
* `voice` - VLAN dedicated for voice devices.
* `video` - VLAN dedicated for video devices.
* `nac` - VLAN for NAC onboarding devices.
* `nac_segment` - VLAN for NAC segemnt primary interface.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchControllerInitialConfig Vlans can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrollerinitialconfig_vlans.labelname SwitchControllerInitialConfigVlans

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollerinitialconfig_vlans.labelname SwitchControllerInitialConfigVlans
$ unset "FORTIOS_IMPORT_TABLE"
```
