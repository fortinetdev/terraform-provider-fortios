---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_sensitivity"
description: |-
  Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source.
---

# fortios_dlp_sensitivity
Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source.

## Argument Reference

The following arguments are supported:

* `name` - DLP Sensitivity Levels.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Sensitivity can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_dlp_sensitivity.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
