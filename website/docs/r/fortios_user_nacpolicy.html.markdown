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
* `severity` - NAC policy matching devices vulnerability severity lists. The structure of `severity` block is documented below.
* `switch_fortilink` - FortiLink interface for which this NAC policy belongs to.
* `switch_group` - List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switch_group` block is documented below.
* `switch_scope` - List of managed FortiSwitches on which NAC policy can be applied. The structure of `switch_scope` block is documented below.
* `switch_auto_auth` - NAC device auto authorization when discovered and nac-policy matched. Valid values: `global`, `disable`, `enable`.
* `switch_port_policy` - switch-port-policy to be applied on the matched NAC policy.
* `switch_mac_policy` - switch-mac-policy to be applied on the matched NAC policy.
* `firewall_address` - Dynamic firewall address to associate MAC which match this policy.
* `ssid_policy` - SSID policy to be applied on the matched NAC policy.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `severity` block supports:

* `severity_num` - Enter multiple severity levels, where 0 = Info, 1 = Low, ..., 4 = Critical

The `switch_group` block supports:

* `name` - Managed FortiSwitch group name from available options.

The `switch_scope` block supports:

* `switch_id` - Managed FortiSwitch name from available options.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User NacPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_user_nacpolicy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_nacpolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
