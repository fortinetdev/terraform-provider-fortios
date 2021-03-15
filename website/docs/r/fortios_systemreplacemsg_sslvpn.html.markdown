---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemreplacemsg_sslvpn"
description: |-
  Replacement messages.
---

# fortios_systemreplacemsg_sslvpn
Replacement messages.

## Argument Reference

The following arguments are supported:

* `msg_type` - (Required) Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{msg_type}}.

## Import

SystemReplacemsg Sslvpn can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemreplacemsg_sslvpn.labelname {{msg_type}}
$ unset "FORTIOS_IMPORT_TABLE"
```
