---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemreplacemsg_admin"
description: |-
  Replacement messages.
---

# fortios_systemreplacemsg_admin
Replacement messages.

## Argument Reference

The following arguments are supported:

* `msg_type` - (Required) Message type.
* `buffer` - Message string.
* `header` - Header flag.
* `format` - Format flag.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{msg_type}}.

## Import

SystemReplacemsg Admin can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemreplacemsg_admin.labelname {{msg_type}}
$ unset "FORTIOS_IMPORT_TABLE"
```
