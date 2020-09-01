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
* `temporary` - Temporary.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Vdom can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_vdom.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
