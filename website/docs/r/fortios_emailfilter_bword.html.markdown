---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_bword"
description: |-
  Configure AntiSpam banned word list.
---

# fortios_emailfilter_bword
Configure AntiSpam banned word list. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `name` - Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter banned word. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `status` - Enable/disable status. Valid values: `enable`, `disable`.
* `id` - Banned word entry ID.
* `pattern` - Pattern for the banned word.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
* `action` - Mark spam or good. Valid values: `spam`, `clear`.
* `where` - Component of the email to be scanned. Valid values: `subject`, `body`, `all`.
* `language` - Language for the banned word. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`.
* `score` - Score value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter Bword can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_emailfilter_bword.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
