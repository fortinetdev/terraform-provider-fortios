---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_profile"
description: |-
  Configure DLP profiles.
---

# fortios_dlp_profile
Configure DLP profiles. Applies to FortiOS Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - Name of the DLP profile.
* `comment` - Comment.
* `feature_set` - Flow/proxy feature set. Valid values: `flow`, `proxy`.
* `replacemsg_group` - Replacement message group used by this DLP profile.
* `rule` - Set up DLP rules for this profile. The structure of `rule` block is documented below.
* `dlp_log` - Enable/disable DLP logging. Valid values: `enable`, `disable`.
* `extended_log` - Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
* `nac_quar_log` - Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
* `full_archive_proto` - Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
* `summary_proto` - Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `rule` block supports:

* `id` - ID.
* `name` - Filter name.
* `severity` - Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.
* `type` - Select whether to check the content of messages (an email message) or files (downloaded files or email attachments). Valid values: `file`, `message`.
* `proto` - Check messages or files over one or more of these protocols. Valid values: `smtp`, `pop3`, `imap`, `http-get`, `http-post`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.
* `filter_by` - Select the type of content to match. Valid values: `sensor`, `mip`, `fingerprint`, `encrypted`, `none`.
* `file_size` - Match files this size or larger (0 - 4294967295 kbytes).
* `sensitivity` - Select a DLP file pattern sensitivity to match. The structure of `sensitivity` block is documented below.
* `match_percentage` - Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
* `file_type` - Select the number of a DLP file pattern table to match.
* `sensor` - Select DLP sensors. The structure of `sensor` block is documented below.
* `label` - MIP label dictionary.
* `archive` - Enable/disable DLP archiving. Valid values: `disable`, `enable`.
* `action` - Action to take with content that this DLP profile matches. Valid values: `allow`, `log-only`, `block`, `quarantine-ip`.
* `expiry` - Quarantine duration in days, hours, minutes (format = dddhhmm).

The `sensitivity` block supports:

* `name` - Select a DLP sensitivity.

The `sensor` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_dlp_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_dlp_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
