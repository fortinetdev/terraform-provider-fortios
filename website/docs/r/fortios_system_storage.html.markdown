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
* `status` - Enable/disable storage.
* `media_status` - The physical status of current media.
* `order` - Set storage order.
* `partition` - Label of underlying partition.
* `device` - Partition device.
* `size` - Partition size.
* `usage` - Use hard disk for logging or WAN Optimization (default = log).
* `wanopt_mode` - WAN Optimization mode (default = mix).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Storage can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_storage.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
