---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomlink"
description: |-
  Configure VDOM links.
---

# fortios_system_vdomlink
Configure VDOM links.

## Argument Reference

The following arguments are supported:

* `name` - VDOM link name (maximum = 8 characters).
* `vcluster` - Virtual cluster. Valid values: `vcluster1`, `vcluster2`.
* `type` - VDOM link type: PPP or Ethernet. Valid values: `ppp`, `ethernet`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VdomLink can be imported using any of these accepted formats:
```
$ terraform import fortios_system_vdomlink.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_vdomlink.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
