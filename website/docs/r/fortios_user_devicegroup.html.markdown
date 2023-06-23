---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_devicegroup"
description: |-
  Configure device groups.
---

# fortios_user_devicegroup
Configure device groups. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_user_device" "trnames12" {
  alias    = "user_devices2"
  category = "amazon-device"
  mac      = "08:00:20:0a:1c:1d"
  type     = "unknown"
}

resource "fortios_user_devicegroup" "trname" {
  name = "user_devicegroupd1"

  member {
    name = fortios_user_device.trnames12.alias
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Device group name.
* `member` - Device group member. The structure of `member` block is documented below.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `comment` - Comment.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Device name.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block supports:

* `name` - Tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User DeviceGroup can be imported using any of these accepted formats:
```
$ terraform import fortios_user_devicegroup.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_devicegroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
