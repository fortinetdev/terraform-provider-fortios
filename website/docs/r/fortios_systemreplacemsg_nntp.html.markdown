---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemreplacemsg_nntp"
description: |-
  Replacement messages.
---

# fortios_systemreplacemsg_nntp
Replacement messages. Applies to FortiOS Version `<= 6.4.0`.

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

SystemReplacemsg Nntp can be imported using any of these accepted formats:
```
$ terraform import fortios_systemreplacemsg_nntp.labelname {{msg_type}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_systemreplacemsg_nntp.labelname {{msg_type}}
$ unset "FORTIOS_IMPORT_TABLE"
```
