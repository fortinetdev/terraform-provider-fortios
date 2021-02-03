---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_mheader"
description: |-
  Configure AntiSpam MIME header.
---

# fortios_spamfilter_mheader
Configure AntiSpam MIME header. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_spamfilter_mheader" "trname" {
  comment = "test"
  fosid   = 1
  name    = "s1"

  entries {
    action       = "spam"
    fieldbody    = "scstest"
    fieldname    = "EIWEtest"
    id           = 1
    pattern_type = "wildcard"
    status       = "enable"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter mime header content. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `status` - Enable/disable status.
* `id` - Mime header entry ID.
* `fieldname` - Pattern for header field name.
* `fieldbody` - Pattern for the header field body.
* `pattern_type` - Wildcard pattern or regular expression.
* `action` - Mark spam or good.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Spamfilter Mheader can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_spamfilter_mheader.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
