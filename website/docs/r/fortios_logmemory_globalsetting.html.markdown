---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logmemory_globalsetting"
description: |-
  Global settings for memory logging.
---

# fortios_logmemory_globalsetting
Global settings for memory logging.

## Example Usage

```hcl
resource "fortios_logmemory_globalsetting" "trname" {
  full_final_warning_threshold  = 95
  full_first_warning_threshold  = 75
  full_second_warning_threshold = 90
  max_size                      = 163840
}
```

## Argument Reference


The following arguments are supported:

* `max_size` - Maximum amount of memory that can be used for memory logging in bytes.
* `full_first_warning_threshold` - Log full first warning threshold as a percent (1 - 98, default = 75).
* `full_second_warning_threshold` - Log full second warning threshold as a percent (2 - 99, default = 90).
* `full_final_warning_threshold` - Log full final warning threshold as a percent (3 - 100, default = 95).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogMemory GlobalSetting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logmemory_globalsetting.labelname LogMemoryGlobalSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
