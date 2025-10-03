---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_hostchecksoftware"
description: |-
  SSL-VPN host check software.
---

# fortios_vpnsslweb_hostchecksoftware
SSL-VPN host check software. Applies to FortiOS Version `<= 7.6.2`.

## Example Usage

```hcl
resource "fortios_vpnsslweb_hostchecksoftware" "trname" {
  name    = "hostchecksoftwares1"
  os_type = "windows"
  type    = "fw"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `os_type` - OS type. Valid values: `windows`, `macos`.
* `type` - Type. Valid values: `av`, `fw`.
* `version` - Version.
* `guid` - Globally unique ID.
* `check_item_list` - Check item list. The structure of `check_item_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `check_item_list` block supports:

* `id` - ID (0 - 4294967295).
* `action` - Action. Valid values: `require`, `deny`.
* `type` - Type. Valid values: `file`, `registry`, `process`.
* `target` - Target.
* `version` - Version.
* `md5s` - MD5 checksum. The structure of `md5s` block is documented below.

The `md5s` block supports:

* `id` - Hex string of MD5 checksum.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnSslWeb HostCheckSoftware can be imported using any of these accepted formats:
```
$ terraform import fortios_vpnsslweb_hostchecksoftware.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnsslweb_hostchecksoftware.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
