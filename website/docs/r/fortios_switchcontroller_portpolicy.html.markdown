---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_portpolicy"
description: |-
  Configure port policy to be applied on the managed FortiSwitch ports through NAC device.
---

# fortios_switchcontroller_portpolicy
Configure port policy to be applied on the managed FortiSwitch ports through NAC device. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - Port policy name.
* `description` - Description for the port policy.
* `fortilink` - FortiLink interface for which this port policy belongs to.
* `lldp_profile` - LLDP profile to be applied when using this port-policy.
* `qos_policy` - QoS policy to be applied when using this port-policy.
* `n802_1x` - 802.1x security policy to be applied when using this port-policy.
* `vlan_policy` - VLAN policy to be applied when using this port-policy.
* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController PortPolicy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_portpolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
