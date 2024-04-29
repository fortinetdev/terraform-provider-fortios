---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_apcfgprofile"
description: |-
  Configure AP local configuration profiles.
---

# fortios_wirelesscontroller_apcfgprofile
Configure AP local configuration profiles. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - AP local configuration profile name.
* `ap_family` - FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.
* `comment` - Comment.
* `ac_type` - Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.
* `ac_timer` - Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
* `ac_ip` - IP address of the validation controller that AP must be able to join after applying AP local configuration.
* `ac_port` - Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
* `command_list` - AP local configuration command list. The structure of `command_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `command_list` block supports:

* `id` - Command ID.
* `type` - The command type (default = non-password). Valid values: `non-password`, `password`.
* `name` - AP local configuration command name.
* `value` - AP local configuration command value.
* `passwd_value` - AP local configuration command password value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController ApcfgProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_apcfgprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_apcfgprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
