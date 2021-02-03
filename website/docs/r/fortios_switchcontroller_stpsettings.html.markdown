---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_stpsettings"
description: |-
  Configure FortiSwitch spanning tree protocol (STP).
---

# fortios_switchcontroller_stpsettings
Configure FortiSwitch spanning tree protocol (STP).

## Example Usage

```hcl
resource "fortios_switchcontroller_stpsettings" "trname" {
  forward_time  = 15
  hello_time    = 2
  max_age       = 20
  max_hops      = 20
  pending_timer = 4
  revision      = 0
  status        = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of global STP settings configuration.
* `status` - Enable/disable STP.
* `revision` - STP revision number (0 - 65535).
* `hello_time` - Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
* `forward_time` - Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
* `max_age` - Maximum time before a bridge port saves its configuration BPDU information (6 - 40 sec, default = 20).
* `max_hops` - Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
* `pending_timer` - Pending time (1 - 15 sec, default = 4).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController StpSettings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_stpsettings.labelname SwitchControllerStpSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
