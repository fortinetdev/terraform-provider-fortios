---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemreplacemsg_nacquar"
description: |-
  Replacement messages.
---

# fortios_systemreplacemsg_nacquar
Replacement messages.

## Argument Reference

The following arguments are supported:

* `msg_type` - (Required) Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{msg_type}}.

## Import

SystemReplacemsg NacQuar can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemreplacemsg_nacquar.labelname {{msg_type}}
$ unset "FORTIOS_IMPORT_TABLE"
```
