---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ftgdlocalcat"
description: |-
  Configure FortiGuard Web Filter local categories.
---

# fortios_webfilter_ftgdlocalcat
Configure FortiGuard Web Filter local categories.

## Example Usage

```hcl
resource "fortios_webfilter_ftgdlocalcat" "trname" {
  desc   = "s1"
  fosid  = 188
  status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable the local category. Valid values: `enable`, `disable`.
* `fosid` - Local category ID.
* `desc` - Local category description.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{desc}}.

## Import

Webfilter FtgdLocalCat can be imported using any of these accepted formats:
```
$ terraform import fortios_webfilter_ftgdlocalcat.labelname {{desc}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webfilter_ftgdlocalcat.labelname {{desc}}
$ unset "FORTIOS_IMPORT_TABLE"
```
