---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_content"
description: |-
  Configure Web filter banned word table.
---

# fortios_webfilter_content
Configure Web filter banned word table.

## Example Usage

```hcl
resource "fortios_webfilter_content" "trname" {
  fosid = 1
  name  = "chaeei"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - Configure banned word entries.

The `entries` block supports:

* `name` - Banned word.
* `pattern_type` - Banned word pattern type: wildcard pattern or Perl regular expression.
* `status` - Enable/disable banned word.
* `lang` - Language of banned word.
* `score` - Score, to be applied every time the word appears on a web page (0 - 4294967295, default = 10).
* `action` - Block or exempt word when a match is found.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter Content can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webfilter_content.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
