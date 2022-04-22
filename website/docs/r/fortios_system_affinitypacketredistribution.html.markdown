---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_affinitypacketredistribution"
description: |-
  Configure packet redistribution.
---

# fortios_system_affinitypacketredistribution
Configure packet redistribution.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID of the packet redistribution setting.
* `interface` - (Required) Physical interface name on which to perform packet redistribution.
* `rxqid` - (Required) ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution.
* `affinity_cpumask` - (Required) Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AffinityPacketRedistribution can be imported using any of these accepted formats:
```
$ terraform import fortios_system_affinitypacketredistribution.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_affinitypacketredistribution.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
