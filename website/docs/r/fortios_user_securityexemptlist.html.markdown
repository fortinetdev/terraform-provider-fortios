---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_securityexemptlist"
description: |-
  Configure security exemption list.
---

# fortios_user_securityexemptlist
Configure security exemption list.

## Argument Reference

The following arguments are supported:

* `name` - Name of the exempt list.
* `description` - Description.
* `rule` - Configure rules for exempting users from captive portal authentication. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `rule` block supports:

* `id` - ID.
* `srcaddr` - Source addresses or address groups. The structure of `srcaddr` block is documented below.
* `devices` - Devices or device groups. The structure of `devices` block is documented below.
* `dstaddr` - Destination addresses or address groups. The structure of `dstaddr` block is documented below.
* `service` - Destination services. The structure of `service` block is documented below.

The `srcaddr` block supports:

* `name` - Address or group name.

The `devices` block supports:

* `name` - Device or group name.

The `dstaddr` block supports:

* `name` - Address or group name.

The `service` block supports:

* `name` - Service name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User SecurityExemptList can be imported using any of these accepted formats:
```
$ terraform import fortios_user_securityexemptlist.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_securityexemptlist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
