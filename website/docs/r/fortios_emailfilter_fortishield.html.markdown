---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_fortishield"
description: |-
  Configure FortiGuard - AntiSpam.
---

# fortios_emailfilter_fortishield
Configure FortiGuard - AntiSpam.

## Argument Reference

The following arguments are supported:

* `spam_submit_srv` - Hostname of the spam submission server.
* `spam_submit_force` - Enable/disable force insertion of a new mime entity for the submission text.
* `spam_submit_txt2htm` - Enable/disable conversion of text email to HTML email.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Emailfilter Fortishield can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_emailfilter_fortishield.labelname EmailfilterFortishield
$ unset "FORTIOS_IMPORT_TABLE"
```
