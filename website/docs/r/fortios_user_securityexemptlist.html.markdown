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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_securityexemptlist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
