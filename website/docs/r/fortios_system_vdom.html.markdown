---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdom"
description: |-
  Configure virtual domain.
---

# fortios_system_vdom
Configure virtual domain.

## Example Usage

```hcl
resource "fortios_system_vdom" "trname" {
  name       = "testvdom"
  short_name = "testvdom"
  temporary  = 0
}
```

## Argument Reference

The following arguments are supported:

* `name` - VDOM name.
* `short_name` - VDOM short name.
* `vcluster_id` - Virtual cluster ID (0 - 4294967295).
* `flag` - Flag.
* `temporary` - Temporary.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Vdom can be imported using any of these accepted formats:
```
$ terraform import fortios_system_vdom.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_vdom.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
