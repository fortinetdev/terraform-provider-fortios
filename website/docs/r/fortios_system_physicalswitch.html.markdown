---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_physicalswitch"
description: |-
  Configure physical switches.
---

# fortios_system_physicalswitch
Configure physical switches. Applies to FortiOS Version `>= 7.0.4`.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `age_enable` - Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
* `age_val` - Layer 2 table age timer value.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System PhysicalSwitch can be imported using any of these accepted formats:
```
$ terraform import fortios_system_physicalswitch.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_physicalswitch.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
