---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ftgdlocalrisk"
description: |-
  Configure FortiGuard Web Filter local risk score.
---

# fortios_webfilter_ftgdlocalrisk
Configure FortiGuard Web Filter local risk score. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `url` - URL to rate locally.
* `status` - Enable/disable local risk score. Valid values: `enable`, `disable`.
* `comment` - Comment.
* `risk_score` - Local risk score.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{url}}.

## Import

Webfilter FtgdLocalRisk can be imported using any of these accepted formats:
```
$ terraform import fortios_webfilter_ftgdlocalrisk.labelname {{url}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webfilter_ftgdlocalrisk.labelname {{url}}
$ unset "FORTIOS_IMPORT_TABLE"
```
