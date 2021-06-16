---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_fortishield"
description: |-
  Configure FortiGuard - AntiSpam.
---

# fortios_spamfilter_fortishield
Configure FortiGuard - AntiSpam. Applies to FortiOS Version `<= 6.2.0`.

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
* `spam_submit_force` - Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
* `spam_submit_txt2htm` - Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


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
