---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_igmpsnooping"
description: |-
  Configure FortiSwitch IGMP snooping global settings.
---

# fortios_switchcontroller_igmpsnooping
Configure FortiSwitch IGMP snooping global settings.

## Example Usage

```hcl
resource "fortios_switchcontroller_igmpsnooping" "trname" {
  aging_time              = 300
  flood_unknown_multicast = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `aging_time` - Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
* `flood_unknown_multicast` - Enable/disable unknown multicast flooding.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController IgmpSnooping can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_igmpsnooping.labelname SwitchControllerIgmpSnooping
$ unset "FORTIOS_IMPORT_TABLE"
```
