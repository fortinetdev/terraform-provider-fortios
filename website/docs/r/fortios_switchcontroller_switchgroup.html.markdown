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
* `members` - FortiSwitch members belonging to this switch group. The structure of `members` block is documented below.

The `members` block supports:

* `name` - Managed device ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SwitchGroup can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_switchgroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
