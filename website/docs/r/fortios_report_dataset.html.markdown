---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_dataset"
description: |-
  Report dataset configuration.
---

# fortios_report_dataset
Report dataset configuration. Applies to FortiOS Version `<= 7.0.0`.

## Example Usage

```hcl
resource "fortios_report_dataset" "trname" {
  name   = "1"
  policy = 0
  query  = "select * from testdb"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `policy` - Used by monitor policy.
* `query` - SQL query statement.
* `field` - Fields. The structure of `field` block is documented below.
* `parameters` - Parameters. The structure of `parameters` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `field` block supports:

* `id` - Field ID (1 to number of columns in SQL result).
* `type` - Field type. Valid values: `text`, `integer`, `double`.
* `name` - Name.
* `displayname` - Display name.

The `parameters` block supports:

* `id` - Parameter ID (1 to number of columns in SQL result).
* `display_name` - Display name.
* `field` - SQL field name.
* `data_type` - Data type. Valid values: `text`, `integer`, `double`, `long-integer`, `date-time`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Report Dataset can be imported using any of these accepted formats:
```
$ terraform import fortios_report_dataset.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_report_dataset.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
