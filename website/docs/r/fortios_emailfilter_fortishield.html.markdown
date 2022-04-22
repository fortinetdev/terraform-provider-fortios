---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_fortishield"
description: |-
  Configure FortiGuard - AntiSpam.
---

# fortios_emailfilter_fortishield
Configure FortiGuard - AntiSpam. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `spam_submit_srv` - Hostname of the spam submission server.
* `spam_submit_force` - Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
* `spam_submit_txt2htm` - Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Emailfilter Fortishield can be imported using any of these accepted formats:
```
$ terraform import fortios_emailfilter_fortishield.labelname EmailfilterFortishield

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_emailfilter_fortishield.labelname EmailfilterFortishield
$ unset "FORTIOS_IMPORT_TABLE"
```
