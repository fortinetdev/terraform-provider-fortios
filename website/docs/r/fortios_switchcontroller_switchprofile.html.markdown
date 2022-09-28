---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_switchprofile"
description: |-
  Configure FortiSwitch switch profile.
---

# fortios_switchcontroller_switchprofile
Configure FortiSwitch switch profile.

## Example Usage

```hcl
resource "fortios_switchcontroller_switchprofile" "trname" {
  login_passwd_override = "enable"
  name                  = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - FortiSwitch Profile name.
* `login_passwd_override` - Enable/disable overriding the admin administrator password for a managed FortiSwitch with the FortiGate admin administrator account password. Valid values: `enable`, `disable`.
* `login_passwd` - Login password of managed FortiSwitch.
* `login` - Enable/disable FortiSwitch serial console. Valid values: `enable`, `disable`.
* `revision_backup_on_logout` - Enable/disable automatic revision backup upon logout from FortiSwitch. Valid values: `enable`, `disable`.
* `revision_backup_on_upgrade` - Enable/disable automatic revision backup upon FortiSwitch image upgrade. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SwitchProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_switchprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_switchprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
