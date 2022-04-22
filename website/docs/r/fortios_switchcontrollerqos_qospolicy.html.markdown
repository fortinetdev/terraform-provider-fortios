---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerqos_qospolicy"
description: |-
  Configure FortiSwitch QoS policy.
---

# fortios_switchcontrollerqos_qospolicy
Configure FortiSwitch QoS policy.

## Argument Reference

The following arguments are supported:

* `name` - (Required) QoS policy name.
* `default_cos` - (Required) Default cos queue for untagged packets.
* `trust_dot1p_map` - QoS trust 802.1p map.
* `trust_ip_dscp_map` - QoS trust ip dscp map.
* `queue_policy` - QoS egress queue policy.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerQos QosPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrollerqos_qospolicy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollerqos_qospolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
