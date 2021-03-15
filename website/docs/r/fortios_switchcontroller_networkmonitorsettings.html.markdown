---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_networkmonitorsettings"
description: |-
  Configure network monitor settings.
---

# fortios_switchcontroller_networkmonitorsettings
Configure network monitor settings.

## Example Usage

```hcl
resource "fortios_switchcontroller_networkmonitorsettings" "trname" {
  network_monitoring = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `network_monitoring` - Enable/disable passive gathering of information by FortiSwitch units concerning other network devices. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController NetworkMonitorSettings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_networkmonitorsettings.labelname SwitchControllerNetworkMonitorSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
