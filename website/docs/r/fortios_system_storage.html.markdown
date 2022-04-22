---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_storage"
description: |-
  Configure logical storage.
---

# fortios_system_storage
Configure logical storage.

## Argument Reference

The following arguments are supported:

* `name` - Storage name.
* `status` - Enable/disable storage. Valid values: `enable`, `disable`.
* `media_status` - The physical status of current media. Valid values: `enable`, `disable`, `fail`.
* `order` - Set storage order.
* `partition` - Label of underlying partition.
* `device` - Partition device.
* `size` - Partition size.
* `usage` - Use hard disk for logging or WAN Optimization (default = log).
* `wanopt_mode` - WAN Optimization mode (default = mix). Valid values: `mix`, `wanopt`, `webcache`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Storage can be imported using any of these accepted formats:
```
$ terraform import fortios_system_storage.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_storage.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
