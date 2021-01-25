---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_hostchecksoftware"
description: |-
  SSL-VPN host check software.
---

# fortios_vpnsslweb_hostchecksoftware
SSL-VPN host check software.

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
* `os_type` - OS type.
* `type` - Type.
* `version` - Version.
* `guid` - Globally unique ID.
* `check_item_list` - Check item list. The structure of `check_item_list` block is documented below.

The `check_item_list` block supports:

* `id` - ID (0 - 4294967295).
* `action` - Action.
* `type` - Type.
* `target` - Target.
* `version` - Version.
* `md5s` - MD5 checksum. The structure of `md5s` block is documented below.

The `md5s` block supports:

* `id` - Hex string of MD5 checksum.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnSslWeb HostCheckSoftware can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnsslweb_hostchecksoftware.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
