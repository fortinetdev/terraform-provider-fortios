---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsecurityrating_controls"
description: |-
  Settings for individual Security Rating controls.
---

# fortios_systemsecurityrating_controls
Settings for individual Security Rating controls. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `name` - Security Rating control name.
* `display_report` - Enable/disable displaying the Security Rating control in the default report. Valid values: `enable`, `disable`.
* `display_insight` - Enable/disable displaying the Security Rating control as an insight across the GUI. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemSecurityRating Controls can be imported using any of these accepted formats:
```
$ terraform import fortios_systemsecurityrating_controls.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_systemsecurityrating_controls.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
