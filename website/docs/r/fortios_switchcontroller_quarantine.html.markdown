---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_quarantine"
description: |-
  Configure FortiSwitch quarantine support.
---

# fortios_switchcontroller_quarantine
Configure FortiSwitch quarantine support.

## Argument Reference

The following arguments are supported:

* `quarantine` - Enable/disable quarantine. Valid values: `enable`, `disable`.
* `targets` - Quarantine MACs. The structure of `targets` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `targets` block supports:

* `mac` - Quarantine MAC.
* `entry_id` - FSW entry id for the quarantine MAC.
* `description` - Description for the quarantine MAC.
* `tag` - Tags for the quarantine MAC. The structure of `tag` block is documented below.

The `tag` block supports:

* `tags` - Tag string(eg. string1 string2 string3).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController Quarantine can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_quarantine.labelname SwitchControllerQuarantine

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_quarantine.labelname SwitchControllerQuarantine
$ unset "FORTIOS_IMPORT_TABLE"
```
