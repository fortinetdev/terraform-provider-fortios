---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_bonjourprofile"
description: |-
  Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.
---

# fortios_wirelesscontroller_bonjourprofile
Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.

## Argument Reference

The following arguments are supported:

* `name` - Bonjour profile name.
* `comment` - Comment.
* `micro_location` - Enable/disable Micro location for Bonjour profile (default = disable). Valid values: `enable`, `disable`.
* `policy_list` - Bonjour policy list. The structure of `policy_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `policy_list` block supports:

* `policy_id` - Policy ID.
* `description` - Description.
* `from_vlan` - VLAN ID from which the Bonjour service is advertised (0 - 4094, default = 0).
* `to_vlan` - VLAN ID to which the Bonjour service is made available (0 - 4094, default = all).
* `services` - Bonjour services for the VLAN connecting to the Bonjour network.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController BonjourProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_bonjourprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_bonjourprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
