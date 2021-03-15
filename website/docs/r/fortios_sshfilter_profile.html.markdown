---
subcategory: "FortiGate SSHFilter"
layout: "fortios"
page_title: "FortiOS: fortios_sshfilter_profile"
description: |-
  SSH filter profile.
---

# fortios_sshfilter_profile
SSH filter profile.

## Example Usage

```hcl
resource "fortios_sshfilter_profile" "trname" {
  block               = "x11"
  default_command_log = "enable"
  log                 = "x11"
  name                = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - SSH filter profile name.
* `block` - SSH blocking options.
* `log` - SSH logging options.
* `default_command_log` - Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
* `shell_commands` - SSH command filter. The structure of `shell_commands` block is documented below.
* `file_filter` - File filter. The structure of `file_filter` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `shell_commands` block supports:

* `id` - Id.
* `type` - Matching type. Valid values: `simple`, `regex`.
* `pattern` - SSH shell command pattern.
* `action` - Action to take for URL filter matches. Valid values: `block`, `allow`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `alert` - Enable/disable alert. Valid values: `enable`, `disable`.
* `severity` - Log severity. Valid values: `low`, `medium`, `high`, `critical`.

The `file_filter` block supports:

* `status` - Enable/disable file filter. Valid values: `enable`, `disable`.
* `log` - Enable/disable file filter logging. Valid values: `enable`, `disable`.
* `scan_archive_contents` - Enable/disable file filter archive contents scan. Valid values: `enable`, `disable`.
* `entries` - File filter entries. The structure of `entries` block is documented below.

The `entries` block supports:

* `filter` - Add a file filter.
* `comment` - Comment.
* `action` - Action taken for matched file. Valid values: `log`, `block`.
* `direction` - Match files transmitted in the session's originating or reply direction. Valid values: `incoming`, `outgoing`, `any`.
* `password_protected` - Match password-protected files. Valid values: `yes`, `any`.
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block supports:

* `name` - File type name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SshFilter Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_sshfilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
