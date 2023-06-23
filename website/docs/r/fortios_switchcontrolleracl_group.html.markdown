---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrolleracl_group"
description: |-
  Configure ACL groups to be applied on managed FortiSwitch ports.
---

# fortios_switchcontrolleracl_group
Configure ACL groups to be applied on managed FortiSwitch ports. Applies to FortiOS Version `>= 7.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - Group name.
* `ingress` - Configure ingress ACL policies in group. The structure of `ingress` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ingress` block supports:

* `id` - ACL ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerAcl Group can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrolleracl_group.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrolleracl_group.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
