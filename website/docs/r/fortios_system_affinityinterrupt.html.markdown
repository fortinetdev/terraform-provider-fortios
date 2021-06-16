---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_affinityinterrupt"
description: |-
  Configure interrupt affinity.
---

# fortios_system_affinityinterrupt
Configure interrupt affinity.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID of the interrupt affinity setting.
* `interrupt` - (Required) Interrupt name.
* `affinity_cpumask` - (Required) Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AffinityInterrupt can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_affinityinterrupt.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
