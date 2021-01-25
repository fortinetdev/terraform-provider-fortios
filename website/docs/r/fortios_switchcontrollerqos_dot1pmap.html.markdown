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
* `priority_0` - COS queue mapped to dot1p priority number.
* `priority_1` - COS queue mapped to dot1p priority number.
* `priority_2` - COS queue mapped to dot1p priority number.
* `priority_3` - COS queue mapped to dot1p priority number.
* `priority_4` - COS queue mapped to dot1p priority number.
* `priority_5` - COS queue mapped to dot1p priority number.
* `priority_6` - COS queue mapped to dot1p priority number.
* `priority_7` - COS queue mapped to dot1p priority number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerQos Dot1PMap can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerqos_dot1pmap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
