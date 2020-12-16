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
* `default_command_log` - Enable/disable logging unmatched shell commands.
* `shell_commands` - SSH command filter. The structure of `shell_commands` block is documented below.

The `shell_commands` block supports:

* `id` - Id.
* `type` - Matching type.
* `pattern` - SSH shell command pattern.
* `action` - Action to take for URL filter matches.
* `log` - Enable/disable logging.
* `alert` - Enable/disable alert.
* `severity` - Log severity.


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
