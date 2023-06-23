---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_fortilinksettings"
description: |-
  Configure integrated FortiLink settings for FortiSwitch.
---

# fortios_switchcontroller_fortilinksettings
Configure integrated FortiLink settings for FortiSwitch. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - FortiLink settings name.
* `fortilink` - FortiLink interface to which this fortilink-setting belongs.
* `inactive_timer` - Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
* `link_down_flush` - Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.
* `nac_ports` - NAC specific configuration. The structure of `nac_ports` block is documented below.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `nac_ports` block supports:

* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered.
* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
* `lan_segment` - Enable/disable LAN segment feature on the FortiLink interface. Valid values: `enabled`, `disabled`.
* `nac_lan_interface` - Configure NAC LAN interface.
* `nac_segment_vlans` - Configure NAC segment VLANs. The structure of `nac_segment_vlans` block is documented below.
* `parent_key` - Parent key name.
* `member_change` - Member change flag.

The `nac_segment_vlans` block supports:

* `vlan_name` - VLAN interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController FortilinkSettings can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_fortilinksettings.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_fortilinksettings.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
