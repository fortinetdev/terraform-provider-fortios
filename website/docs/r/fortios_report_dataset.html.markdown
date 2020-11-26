---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_dataset"
description: |-
  Report dataset configuration.
---

# fortios_report_dataset
Report dataset configuration.

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

The `field` block supports:

* `id` - Field ID (1 to number of columns in SQL result).
* `type` - Field type.
* `name` - Name.
* `displayname` - Display name.

The `parameters` block supports:

* `id` - Parameter ID (1 to number of columns in SQL result).
* `display_name` - Display name.
* `field` - SQL field name.
* `data_type` - Data type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Report Dataset can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_report_dataset.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
