---
subcategory: "FortiGate Automation"
layout: "fortios"
page_title: "FortiOS: fortios_automation_setting"
description: |-
  Automation setting configuration.
---

# fortios_automation_setting
Automation setting configuration. Applies to FortiOS Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `max_concurrent_stitches` - Maximum number of automation stitches that are allowed to run concurrently.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Automation Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_automation_setting.labelname AutomationSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_automation_setting.labelname AutomationSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
