---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logmemory_setting"
description: |-
  Settings for memory buffer.
---

# fortios_logmemory_setting
Settings for memory buffer.

## Example Usage

```hcl
resource "fortios_logmemory_setting" "trname" {
  diskfull = "overwrite"
  status   = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `status` - Enable/disable logging to the FortiGate's memory.
* `diskfull` - Action to take when memory is full.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogMemory Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logmemory_setting.labelname LogMemorySetting
$ unset "FORTIOS_IMPORT_TABLE"
```
