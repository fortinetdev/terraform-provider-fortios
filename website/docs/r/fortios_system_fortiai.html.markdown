---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortiai"
description: |-
  Configure FortiAI.
---

# fortios_system_fortiai
Configure FortiAI. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FortiAI. Valid values: `disable`, `enable`.
* `source_ip` - Source IP address for communications to FortiAI.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortiai can be imported using any of these accepted formats:
```
$ terraform import fortios_system_fortiai.labelname SystemFortiai

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_fortiai.labelname SystemFortiai
$ unset "FORTIOS_IMPORT_TABLE"
```
