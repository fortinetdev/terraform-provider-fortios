---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollersecuritypolicy_8021X"
description: |-
  Configure 802.1x MAC Authentication Bypass (MAB) policies.
---

# fortios_switchcontrollersecuritypolicy_8021X
Configure 802.1x MAC Authentication Bypass (MAB) policies.

## Example Usage

```hcl
resource "fortios_switchcontrollersecuritypolicy_8021X" "trname" {
  auth_fail_vlan           = "disable"
  auth_fail_vlanid         = 0
  eap_passthru             = "disable"
  framevid_apply           = "enable"
  guest_auth_delay         = 30
  guest_vlan               = "disable"
  guest_vlanid             = 100
  mac_auth_bypass          = "disable"
  name                     = "1"
  open_auth                = "disable"
  policy_type              = "802.1X"
  radius_timeout_overwrite = "disable"
  security_mode            = "802.1X"
  user_group {
    name = "Guest-group"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Policy name.
* `security_mode` - Port or MAC based 802.1X security mode. Valid values: `802.1X`, `802.1X-mac-based`.
* `user_group` - Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `user_group` block is documented below.
* `mac_auth_bypass` - Enable/disable MAB for this policy. Valid values: `disable`, `enable`.
* `open_auth` - Enable/disable open authentication for this policy. Valid values: `disable`, `enable`.
* `eap_passthru` - Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable`, `enable`.
* `eap_auto_untagged_vlans` - Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable`, `enable`.
* `guest_vlan` - Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable`, `enable`.
* `guest_vlanid` - Guest VLAN ID.
* `guest_vlan_id` - Guest VLAN name.
* `guest_auth_delay` - Guest authentication delay (1 - 900  sec, default = 30).
* `auth_fail_vlan` - Enable to allow limited access to clients that cannot authenticate. Valid values: `disable`, `enable`.
* `auth_fail_vlanid` - VLAN ID on which authentication failed.
* `auth_fail_vlan_id` - VLAN ID on which authentication failed.
* `framevid_apply` - Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable`, `enable`.
* `radius_timeout_overwrite` - Enable to override the global RADIUS session timeout. Valid values: `disable`, `enable`.
* `policy_type` - Policy type. Valid values: `802.1X`.
* `authserver_timeout_period` - Authentication server timeout period (3 - 15 sec, default = 3).
* `authserver_timeout_vlan` - Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable.  Valid values: `disable`, `enable`.
* `authserver_timeout_vlanid` - Authentication server timeout VLAN name.
* `authserver_timeout_tagged` - Configure timeout option for the tagged VLAN which allows limited access when the authentication server is unavailable. Valid values: `disable`, `lldp-voice`, `static`.
* `authserver_timeout_tagged_vlanid` - Tagged VLAN name for which the timeout option is applied to (only one VLAN ID).
* `dacl` - Enable/disable dynamic access control list on this interface. Valid values: `disable`, `enable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `user_group` block supports:

* `name` - Group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerSecurityPolicy 8021X can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrollersecuritypolicy_8021X.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollersecuritypolicy_8021X.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
