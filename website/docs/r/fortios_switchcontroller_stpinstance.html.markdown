---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_stpinstance"
description: |-
  Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.
---

# fortios_switchcontroller_stpinstance
Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.

## Argument Reference

The following arguments are supported:

* `fosid` - Instance ID.
* `vlan_range` - Configure VLAN range for STP instance. The structure of `vlan_range` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `vlan_range` block supports:

* `vlan_name` - VLAN name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController StpInstance can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_stpinstance.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
