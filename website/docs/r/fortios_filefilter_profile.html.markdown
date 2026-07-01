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
* `scan_archive_contents` - Enable/disable archive contents scan. Valid values: `disable`, `enable`.
* `rules` - File filter rules. The structure of `rules` block is documented below.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `enable`, `disable`.
* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `rules` block supports:

* `name` - File-filter rule name.
* `comment` - Comment.
* `protocol` - Protocols to apply rule to.
* `action` - Action taken for matched file.
* `direction` - Traffic direction. On FortiOS versions 6.4.1-7.4.1: HTTP, FTP, SSH, CIFS only. On FortiOS versions 7.4.2-7.6.6: HTTP, FTP, SSH, CIFS, and MAPI only. On FortiOS versions >= 7.6.7: HTTP, FTP, SSH, WEBSOCKET, CIFS, and MAPI only. Valid values: `incoming`, `outgoing`, `any`.
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
