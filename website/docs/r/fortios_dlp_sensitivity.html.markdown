---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_sensitivity"
description: |-
  Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source.
---

# fortios_dlp_sensitivity
Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `name` - DLP Sensitivity Levels.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Sensitivity can be imported using any of these accepted formats:
```
$ terraform import fortios_dlp_sensitivity.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_dlp_sensitivity.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
