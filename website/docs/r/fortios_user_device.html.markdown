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
* `type` - Device type. Valid values: `unknown`, `android-phone`, `android-tablet`, `blackberry-phone`, `blackberry-playbook`, `forticam`, `fortifone`, `fortinet-device`, `gaming-console`, `ip-phone`, `ipad`, `iphone`, `linux-pc`, `mac`, `media-streaming`, `printer`, `router-nat-device`, `windows-pc`, `windows-phone`, `windows-tablet`, `other-network-device`.
* `category` - Device category. Valid values: `none`, `amazon-device`, `android-device`, `blackberry-device`, `fortinet-device`, `ios-device`, `windows-device`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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
$ terraform import fortios_user_device.labelname {{alias}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_device.labelname {{alias}}
$ unset "FORTIOS_IMPORT_TABLE"
```
