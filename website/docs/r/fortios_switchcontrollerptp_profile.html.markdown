---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerptp_profile"
description: |-
  Global PTP profile.
---

# fortios_switchcontrollerptp_profile
Global PTP profile. Applies to FortiOS Version `>= 7.4.1`.

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `description` - Description.
* `mode` - Select PTP mode. Valid values: `transparent-e2e`, `transparent-p2p`.
* `ptp_profile` - Configure PTP power profile. Valid values: `C37.238-2017`.
* `transport` - Configure PTP transport mode. Valid values: `l2-mcast`.
* `domain` - Configure PTP domain value (0 - 255, default = 254).
* `pdelay_req_interval` - Configure PTP peer delay request interval. Valid values: `1sec`, `2sec`, `4sec`, `8sec`, `16sec`, `32sec`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerPtp Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrollerptp_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollerptp_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
