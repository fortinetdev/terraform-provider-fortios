---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicecustomgroup"
description: |-
  Configure custom Internet Service group.
---

# fortios_firewall_internetservicecustomgroup
Configure custom Internet Service group.

## Argument Reference

The following arguments are supported:

* `name` - Custom Internet Service group name.
* `comment` - Comment.
* `member` - Custom Internet Service group members. The structure of `member` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Group member name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall InternetServiceCustomGroup can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetservicecustomgroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
