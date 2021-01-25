---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_deviceaccesslist"
description: |-
  Configure device access control lists.
---

# fortios_user_deviceaccesslist
Configure device access control lists.

## Example Usage

```hcl
resource "fortios_user_deviceaccesslist" "trname" {
  default_action = "accept"
  name           = "1"
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) Device access list name.
* `default_action` - Accept or deny unknown/unspecified devices.
* `device_list` - Device list. The structure of `device_list` block is documented below.

The `device_list` block supports:

* `id` - Entry ID.
* `device` - Firewall device or device group.
* `action` - Allow or block device.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User DeviceAccessList can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_deviceaccesslist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
