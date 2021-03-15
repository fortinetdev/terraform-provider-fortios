---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemreplacemsg_automation"
description: |-
  Replacement messages.
---

# fortios_systemreplacemsg_automation
Replacement messages. Applies to FortiOS Version `>= 7.0.0`.

## Argument Reference

The following arguments are supported:

* `msg_type` - Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag. Valid values: `none`, `text`, `html`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{msg_type}}.

## Import

SystemReplacemsg Automation can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemreplacemsg_automation.labelname {{msg_type}}
$ unset "FORTIOS_IMPORT_TABLE"
```
