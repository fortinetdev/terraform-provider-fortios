---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_classificationsettings"
description: |-
  Configure app classification setting.
---

# fortios_application_classificationsettings
Configure app classification setting. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `default_app_classification` - Default application classification. Valid values: `unclassified`, `unsanctioned`, `sanctioned`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Application ClassificationSettings can be imported using any of these accepted formats:
```
$ terraform import fortios_application_classificationsettings.labelname ApplicationClassificationSettings

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_application_classificationsettings.labelname ApplicationClassificationSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
