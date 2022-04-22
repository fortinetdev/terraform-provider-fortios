---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_rulesettings"
description: |-
  Configure application rule settings.
---

# fortios_application_rulesettings
Configure application rule settings.

## Argument Reference

The following arguments are supported:

* `fosid` - Rule ID.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Application RuleSettings can be imported using any of these accepted formats:
```
$ terraform import fortios_application_rulesettings.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_application_rulesettings.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
