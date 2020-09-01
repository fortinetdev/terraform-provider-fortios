---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_bword"
description: |-
  Configure AntiSpam banned word list.
---

# fortios_spamfilter_bword
Configure AntiSpam banned word list.

## Example Usage

```hcl
resource "fortios_spamfilter_bword" "trname" {
  comment = "test"
  fosid   = 1
  name    = "s1"

  entries {
    action       = "clear"
    language     = "western"
    pattern      = "test*patten"
    pattern_type = "wildcard"
    score        = 10
    status       = "enable"
    where        = "subject"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter banned word.

The `entries` block supports:

* `status` - Enable/disable status.
* `id` - Banned word entry ID.
* `pattern` - Pattern for the banned word.
* `pattern_type` - Wildcard pattern or regular expression.
* `action` - Mark spam or good.
* `where` - Component of the email to be scanned.
* `language` - Language for the banned word.
* `score` - Score value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Spamfilter Bword can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_spamfilter_bword.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
