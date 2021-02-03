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

* `default_vlan` - Default VLAN (native) assigned to all switch ports upon discovery.
* `quarantine` - VLAN for quarantined traffic.
* `rspan` - VLAN for RSPAN/ERSPAN mirrored traffic.
* `voice` - VLAN dedicated for voice devices.
* `video` - VLAN dedicated for video devices.
* `nac` - VLAN for NAC onboarding devices.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchControllerInitialConfig Vlans can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerinitialconfig_vlans.labelname SwitchControllerInitialConfigVlans
$ unset "FORTIOS_IMPORT_TABLE"
```
