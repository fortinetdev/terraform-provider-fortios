---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationstitch"
description: |-
  Automation stitches.
---

# fortios_system_automationstitch
Automation stitches.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `description` - Description.
* `status` - (Required) Enable/disable this stitch. Valid values: `enable`, `disable`.
* `trigger` - (Required) Trigger name.
* `condition` - Automation conditions. The structure of `condition` block is documented below.
* `condition_logic` - Apply AND/OR logic to the specified automation conditions. Valid values: `and`, `or`.
* `actions` - Configure stitch actions. *Due to the data type change of API, for other versions of FortiOS, please check variable `action`.* The structure of `actions` block is documented below.
* `action` - Action names. *Due to the data type change of API, for other versions of FortiOS, please check variable `actions`.* The structure of `action` block is documented below.
* `destination` - Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `condition` block supports:

* `name` - Condition name.

The `actions` block supports:

* `id` - Entry ID.
* `action` - Action name.
* `delay` - Delay before execution (in seconds).
* `required` - Required in action chain. Valid values: `enable`, `disable`.

The `action` block supports:

* `name` - Action name.

The `destination` block supports:

* `name` - Destination name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationStitch can be imported using any of these accepted formats:
```
$ terraform import fortios_system_automationstitch.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_automationstitch.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
