---
subcategory: "FortiGate Dlp"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_fpsensitivity"
description: |-
  Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source.
---

# fortios_dlp_fpsensitivity
Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source.

## Example Usage

```hcl
resource "fortios_dlp_fpsensitivity" "trname" {
  name = "Normal"
}
```

## Argument Reference

The following arguments are supported:

* `name` - DLP Sensitivity Levels.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp FpSensitivity can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_dlp_fpsensitivity.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
