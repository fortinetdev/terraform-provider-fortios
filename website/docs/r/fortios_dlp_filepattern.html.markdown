---
subcategory: "FortiGate Dlp"
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
* `entries` - Configure file patterns used by DLP blocking.

The `entries` block supports:

* `filter_type` - Filter by file name pattern or by file type.
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
