---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_hamonitor"
description: |-
  Configure HA monitor.
---

# fortios_system_hamonitor
Configure HA monitor.

## Example Usage

```hcl
resource "fortios_system_hamonitor" "trname" {
  monitor_vlan           = "disable"
  vlan_hb_interval       = 5
  vlan_hb_lost_threshold = 3
}
```

## Argument Reference

The following arguments are supported:

* `monitor_vlan` - Enable/disable monitor VLAN interfaces. Valid values: `enable`, `disable`.
* `vlan_hb_interval` - Configure heartbeat interval (seconds).
* `vlan_hb_lost_threshold` - VLAN lost heartbeat threshold (1 - 60).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System HaMonitor can be imported using any of these accepted formats:
```
$ terraform import fortios_system_hamonitor.labelname SystemHaMonitor

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_hamonitor.labelname SystemHaMonitor
$ unset "FORTIOS_IMPORT_TABLE"
```
