---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzercloud_overridesetting"
description: |-
  Override FortiAnalyzer Cloud settings.
---

# fortios_logfortianalyzercloud_overridesetting
Override FortiAnalyzer Cloud settings.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable logging to FortiAnalyzer.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortianalyzerCloud OverrideSetting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logfortianalyzercloud_overridesetting.labelname LogFortianalyzerCloudOverrideSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
