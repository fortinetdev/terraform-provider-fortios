---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_devicegroup"
description: |-
  Configure device groups.
---

# fortios_user_devicegroup
Configure device groups.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_devicegroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
