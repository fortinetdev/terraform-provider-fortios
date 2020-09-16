---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_fortishield"
description: |-
  Configure FortiGuard - AntiSpam.
---

# fortios_spamfilter_fortishield
Configure FortiGuard - AntiSpam.

## Example Usage

```hcl
resource "fortios_spamfilter_fortishield" "trname" {
  spam_submit_force   = "enable"
  spam_submit_srv     = "www.nospammer.net"
  spam_submit_txt2htm = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `spam_submit_srv` - Hostname of the spam submission server.
* `spam_submit_force` - Enable/disable force insertion of a new mime entity for the submission text.
* `spam_submit_txt2htm` - Enable/disable conversion of text email to HTML email.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Spamfilter Fortishield can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_spamfilter_fortishield.labelname SpamfilterFortishield
$ unset "FORTIOS_IMPORT_TABLE"
```
