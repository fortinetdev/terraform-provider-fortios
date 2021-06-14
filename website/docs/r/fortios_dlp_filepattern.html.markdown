---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_filepattern"
description: |-
  Configure file patterns used by DLP blocking.
---

# fortios_dlp_filepattern
Configure file patterns used by DLP blocking.

## Example Usage

```hcl
resource "fortios_dlp_filepattern" "trname" {
  fosid = 9
  name  = "alldocs"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table containing the file pattern list.
* `comment` - Optional comments.
* `entries` - Configure file patterns used by DLP blocking. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `filter_type` - Filter by file name pattern or by file type. Valid values: `pattern`, `type`.
* `pattern` - Add a file name pattern.
* `file_type` - Select a file type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dlp Filepattern can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_dlp_filepattern.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
