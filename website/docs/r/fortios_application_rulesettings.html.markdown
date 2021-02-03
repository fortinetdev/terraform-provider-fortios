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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Application RuleSettings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_application_rulesettings.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
