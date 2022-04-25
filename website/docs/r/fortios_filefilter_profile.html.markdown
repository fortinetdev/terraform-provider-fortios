---
subcategory: "FortiGate File-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_filefilter_profile"
description: |-
  Configure file-filter profiles.
---

# fortios_filefilter_profile
Configure file-filter profiles. Applies to FortiOS Version `>= 6.4.1`.

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `comment` - Comment.
* `feature_set` - Flow/proxy feature set. Valid values: `flow`, `proxy`.
* `replacemsg_group` - Replacement message group
* `log` - Enable/disable file-filter logging. Valid values: `disable`, `enable`.
* `extended_log` - Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.
* `scan_archive_contents` - Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.
* `rules` - File filter rules. The structure of `rules` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `rules` block supports:

* `name` - File-filter rule name.
* `comment` - Comment.
* `protocol` - Protocols to apply rule to. Valid values: `http`, `ftp`, `smtp`, `imap`, `pop3`, `mapi`, `cifs`, `ssh`.
* `action` - Action taken for matched file. Valid values: `log-only`, `block`.
* `direction` - Traffic direction. (HTTP, FTP, SSH, CIFS only) Valid values: `incoming`, `outgoing`, `any`.
* `password_protected` - Match password-protected files. Valid values: `yes`, `any`.
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block supports:

* `name` - File type name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FileFilter Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_filefilter_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_filefilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
