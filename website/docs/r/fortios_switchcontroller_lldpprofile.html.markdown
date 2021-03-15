---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_lldpprofile"
description: |-
  Configure FortiSwitch LLDP profiles.
---

# fortios_switchcontroller_lldpprofile
Configure FortiSwitch LLDP profiles.

## Example Usage

```hcl
resource "fortios_switchcontroller_lldpprofile" "trname" {
  auto_isl                 = "enable"
  auto_isl_hello_timer     = 3
  auto_isl_port_group      = 0
  auto_isl_receive_timeout = 60
  med_tlvs                 = "inventory-management network-policy"
  name                     = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `med_tlvs` - Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
* `n8021_tlvs` - Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
* `n8023_tlvs` - Transmitted IEEE 802.3 TLVs.
* `auto_isl` - Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
* `auto_isl_hello_timer` - Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
* `auto_isl_receive_timeout` - Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
* `auto_isl_port_group` - Auto inter-switch LAG port group ID (0 - 9).
* `auto_mclag_icl` - Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
* `med_network_policy` - Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `med_network_policy` block is documented below.
* `med_location_service` - Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `med_location_service` block is documented below.
* `custom_tlvs` - Configuration method to edit custom TLV entries. The structure of `custom_tlvs` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `med_network_policy` block supports:

* `name` - Policy type name.
* `status` - Enable or disable this TLV. Valid values: `disable`, `enable`.
* `vlan_intf` - VLAN interface to advertise; if configured on port.
* `assign_vlan` - Enable/disable VLAN assignment when this profile is applied on managed FortiSwitch port. Valid values: `disable`, `enable`.
* `vlan` - ID of VLAN to advertise, if configured on port (0 - 4094, 0 = priority tag).
* `priority` - Advertised Layer 2 priority (0 - 7; from lowest to highest priority).
* `dscp` - Advertised Differentiated Services Code Point (DSCP) value, a packet header value indicating the level of service requested for traffic, such as high priority or best effort delivery.

The `med_location_service` block supports:

* `name` - Location service type name.
* `status` - Enable or disable this TLV. Valid values: `disable`, `enable`.
* `sys_location_id` - Location service ID.

The `custom_tlvs` block supports:

* `name` - TLV name (not sent).
* `oui` - Organizationally unique identifier (OUI), a 3-byte hexadecimal number, for this TLV.
* `subtype` - Organizationally defined subtype (0 - 255).
* `information_string` - Organizationally defined information string (0 - 507 hexadecimal bytes).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController LldpProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_lldpprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
