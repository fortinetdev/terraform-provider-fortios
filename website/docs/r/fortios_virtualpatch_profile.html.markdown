---
subcategory: "FortiGate Virtual-Patch"
layout: "fortios"
page_title: "FortiOS: fortios_virtualpatch_profile"
description: |-
  Configure virtual-patch profile.
---

# fortios_virtualpatch_profile
Configure virtual-patch profile. Applies to FortiOS Version `>= 7.4.1`.

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `comment` - Comment.
* `severity` - Relative severity of the signature (low, medium, high, critical). Valid values: `low`, `medium`, `high`, `critical`.
* `action` - Action (pass/block). Valid values: `pass`, `block`.
* `log` - Enable/disable logging of detection. Valid values: `enable`, `disable`.
* `exemption` - Exempt devices or rules. The structure of `exemption` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `exemption` block supports:

* `id` - IDs.
* `status` - Enable/disable exemption. Valid values: `enable`, `disable`.
* `rule` - Patch signature rule IDs. The structure of `rule` block is documented below.
* `device` - Device MAC addresses. The structure of `device` block is documented below.

The `rule` block supports:

* `id` - Rule IDs.

The `device` block supports:

* `mac` - Device MAC address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VirtualPatch Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_virtualpatch_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_virtualpatch_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
