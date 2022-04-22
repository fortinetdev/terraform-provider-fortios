---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerqos_dot1pmap"
description: |-
  Configure FortiSwitch QoS 802.1p.
---

# fortios_switchcontrollerqos_dot1pmap
Configure FortiSwitch QoS 802.1p.

## Example Usage

```hcl
resource "fortios_switchcontrollerqos_dot1pmap" "trname" {
  name       = "1"
  priority_0 = "queue-0"
  priority_1 = "queue-0"
  priority_2 = "queue-0"
  priority_3 = "queue-0"
  priority_4 = "queue-0"
  priority_5 = "queue-0"
  priority_6 = "queue-0"
  priority_7 = "queue-0"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Dot1p map name.
* `description` - Description of the 802.1p name.
* `egress_pri_tagging` - Enable/disable egress priority-tag frame. Valid values: `disable`, `enable`.
* `priority_0` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_1` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_2` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_3` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_4` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_5` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_6` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `priority_7` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerQos Dot1PMap can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrollerqos_dot1pmap.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollerqos_dot1pmap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
