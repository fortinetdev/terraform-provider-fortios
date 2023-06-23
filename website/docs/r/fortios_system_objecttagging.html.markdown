---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_objecttagging"
description: |-
  Configure object tagging.
---

# fortios_system_objecttagging
Configure object tagging.

## Example Usage

```hcl
resource "fortios_system_objecttagging" "trname" {
  address   = "disable"
  category  = "s1"
  color     = 0
  device    = "mandatory"
  interface = "disable"
  multiple  = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `category` - Tag Category.
* `address` - Address. Valid values: `disable`, `mandatory`, `optional`.
* `device` - Device. Valid values: `disable`, `mandatory`, `optional`.
* `interface` - Interface. Valid values: `disable`, `mandatory`, `optional`.
* `multiple` - Allow multiple tag selection. Valid values: `enable`, `disable`.
* `color` - Color of icon on the GUI.
* `tags` - Tags. The structure of `tags` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `tags` block supports:

* `name` - Tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{category}}.

## Import

System ObjectTagging can be imported using any of these accepted formats:
```
$ terraform import fortios_system_objecttagging.labelname {{category}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_objecttagging.labelname {{category}}
$ unset "FORTIOS_IMPORT_TABLE"
```
