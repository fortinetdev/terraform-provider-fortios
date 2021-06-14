---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_devicecategory"
description: |-
  Configure device categories.
---

# fortios_user_devicecategory
Configure device categories. Applies to FortiOS Version `<= 6.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Device category name.
* `desc` - Device category description.
* `comment` - Comment.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User DeviceCategory can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_devicecategory.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
