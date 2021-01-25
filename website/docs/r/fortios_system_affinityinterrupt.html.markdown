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
