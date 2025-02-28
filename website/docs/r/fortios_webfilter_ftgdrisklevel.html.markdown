---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ftgdrisklevel"
description: |-
  Configure FortiGuard Web Filter risk level.
---

# fortios_webfilter_ftgdrisklevel
Configure FortiGuard Web Filter risk level. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `name` - Risk level name.
* `high` - Risk level high.
* `low` - Risk level low.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webfilter FtgdRiskLevel can be imported using any of these accepted formats:
```
$ terraform import fortios_webfilter_ftgdrisklevel.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webfilter_ftgdrisklevel.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
