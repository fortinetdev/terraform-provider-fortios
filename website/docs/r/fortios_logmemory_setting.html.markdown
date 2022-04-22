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

* `status` - Enable/disable logging to the FortiGate's memory. Valid values: `enable`, `disable`.
* `diskfull` - Action to take when memory is full. Valid values: `overwrite`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogMemory Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_logmemory_setting.labelname LogMemorySetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_logmemory_setting.labelname LogMemorySetting
$ unset "FORTIOS_IMPORT_TABLE"
```
