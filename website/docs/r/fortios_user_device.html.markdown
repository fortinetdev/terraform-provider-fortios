---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_device"
description: |-
  Configure devices.
---

# fortios_user_device
Configure devices. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_user_device" "trname" {
  alias    = "1"
  category = "amazon-device"
  mac      = "08:00:20:0a:8c:6d"
  type     = "unknown"
}
```

## Argument Reference

The following arguments are supported:

* `alias` - Device alias.
* `mac` - Device MAC address.
* `user` - User name.
* `master_device` - Master device (optional).
* `comment` - Comment.
* `avatar` - Image file for avatar (maximum 4K base64 encoded).
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `type` - Device type.
* `category` - Device category.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block supports:

* `name` - Tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{alias}}.

## Import

User Device can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_device.labelname {{alias}}
$ unset "FORTIOS_IMPORT_TABLE"
```
