---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_dynamicportpolicy"
description: |-
  Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.
---

# fortios_switchcontroller_dynamicportpolicy
Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - Dynamic port policy name.
* `description` - Description for the Dynamic port policy.
* `fortilink` - FortiLink interface for which this Dynamic port policy belongs to.
* `policy` - Port policies with matching criteria and actions. The structure of `policy` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `policy` block supports:

* `name` - Policy name.
* `description` - Description for the policy.
* `status` - Enable/disable policy. Valid values: `enable`, `disable`.
* `category` - Category of Dynamic port policy. Valid values: `device`, `interface-tag`.
* `interface_tags` - Policy matching the FortiSwitch interface object tags. The structure of `interface_tags` block is documented below.
* `mac` - Policy matching MAC address.
* `hw_vendor` - Match policy based on hardware vendor.
* `type` - Policy matching type.
* `family` - Policy matching family.
* `host` - Policy matching host.
* `lldp_profile` - LLDP profile to be applied when using this policy.
* `qos_policy` - QoS policy to be applied when using this policy.
* `n802_1x` - 802.1x security policy to be applied when using this policy.
* `vlan_policy` - VLAN policy to be applied when using this policy.
* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.

The `interface_tags` block supports:

* `tag_name` - FortiSwitch port tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController DynamicPortPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_dynamicportpolicy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_dynamicportpolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
