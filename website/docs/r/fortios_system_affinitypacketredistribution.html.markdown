---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_affinitypacketredistribution"
description: |-
  Configure packet redistribution.
---

# fortios_system_affinitypacketredistribution
Configure packet redistribution. Applies to FortiOS Version `6.2.0,6.2.4,6.2.6,6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.0.14,7.0.15,7.0.16,7.0.17,7.0.18,7.0.19,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.2.7,7.2.8,7.2.9,7.2.10,7.2.11,7.2.12,7.4.0,7.4.1,7.4.2,7.4.3,7.4.4,7.4.5,7.4.6,7.4.7,7.4.8,7.4.9,7.4.11,7.4.12,7.6.0,7.6.1,7.6.2,7.6.3,7.6.4,7.6.5,7.6.6,7.6.7,8.0.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID of the packet redistribution setting.
* `interface` - (Required) Physical interface name on which to perform packet redistribution.
* `rxqid` - (Required) ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution (255 = all queues).
* `round_robin` - Enable/disable round-robin redistribution to multiple CPUs. Valid values: `enable`, `disable`.
* `affinity_cpumask` - (Required) Hexadecimal cpumask, empty value means all CPUs.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


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
