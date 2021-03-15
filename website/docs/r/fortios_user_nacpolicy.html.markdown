---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_nacpolicy"
description: |-
  Configure NAC policy matching pattern to identify matching NAC devices.
---

# fortios_user_nacpolicy
Configure NAC policy matching pattern to identify matching NAC devices. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - NAC policy name.
* `description` - Description for the NAC policy matching pattern.
* `category` - Category of NAC policy.
* `status` - Enable/disable NAC policy. Valid values: `enable`, `disable`.
* `mac` - NAC policy matching MAC address.
* `hw_vendor` - NAC policy matching hardware vendor.
* `type` - NAC policy matching type.
* `family` - NAC policy matching family.
* `os` - NAC policy matching operating system.
* `hw_version` - NAC policy matching hardware version.
* `sw_version` - NAC policy matching software version.
* `host` - NAC policy matching host.
* `user` - NAC policy matching user.
* `src` - NAC policy matching source.
* `user_group` - NAC policy matching user group.
* `ems_tag` - NAC policy matching EMS tag.
* `switch_fortilink` - FortiLink interface for which this NAC policy belongs to.
* `switch_scope` - List of managed FortiSwitches on which NAC policy can be applied. The structure of `switch_scope` block is documented below.
* `switch_auto_auth` - NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
* `switch_port_policy` - switch-port-policy to be applied on the matched NAC policy.
* `switch_mac_policy` - switch-mac-policy to be applied on the matched NAC policy.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `switch_scope` block supports:

* `switch_id` - Managed FortiSwitch name from available options.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User NacPolicy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_nacpolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
