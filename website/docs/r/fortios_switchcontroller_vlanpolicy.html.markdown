---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_vlanpolicy"
description: |-
  Configure VLAN policy to be applied on the managed FortiSwitch ports through port-policy.
---

# fortios_switchcontroller_vlanpolicy
Configure VLAN policy to be applied on the managed FortiSwitch ports through port-policy. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - VLAN policy name.
* `description` - Description for the VLAN policy.
* `fortilink` - FortiLink interface for which this VLAN policy belongs to.
* `vlan` - Native VLAN to be applied when using this VLAN policy.
* `allowed_vlans` - Allowed VLANs to be applied when using this VLAN policy. The structure of `allowed_vlans` block is documented below.
* `untagged_vlans` - Untagged VLANs to be applied when using this VLAN policy. The structure of `untagged_vlans` block is documented below.
* `allowed_vlans_all` - Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
* `discard_mode` - Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `allowed_vlans` block supports:

* `vlan_name` - VLAN name.

The `untagged_vlans` block supports:

* `vlan_name` - VLAN name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController VlanPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_vlanpolicy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_vlanpolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
