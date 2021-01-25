---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_system"
description: |-
  Configure system-wide switch controller settings.
---

# fortios_switchcontroller_system
Configure system-wide switch controller settings.

## Argument Reference


The following arguments are supported:

* `parallel_process_override` - Enable/disable parallel process override.
* `parallel_process` - Maximum number of parallel processes (1 - 300, default = 1).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController System can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_system.labelname SwitchControllerSystem
$ unset "FORTIOS_IMPORT_TABLE"
```
