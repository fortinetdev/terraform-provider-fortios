---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_switchgroup"
description: |-
  Configure FortiSwitch switch groups.
---

# fortios_switchcontroller_switchgroup
Configure FortiSwitch switch groups.

## Argument Reference

The following arguments are supported:

* `name` - Switch group name.
* `description` - Optional switch group description.
* `fortilink` - FortiLink interface to which switch group members belong.
* `members` - FortiSwitch members belonging to this switch group. The structure of `members` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `members` block supports:

* `switch_id` - Managed device ID.
* `name` - Managed device ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SwitchGroup can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_switchgroup.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_switchgroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
