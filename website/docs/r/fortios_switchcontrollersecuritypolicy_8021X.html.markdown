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

* `name` - (Required) Policy name.
* `security_mode` - Port or MAC based 802.1X security mode.
* `user_group` - Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `user_group` block is documented below.
* `mac_auth_bypass` - Enable/disable MAB for this policy.
* `open_auth` - Enable/disable open authentication for this policy.
* `eap_passthru` - Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication.
* `guest_vlan` - Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients.
* `guest_vlanid` - Guest VLAN ID.
* `guest_vlan_id` - Guest VLAN name.
* `guest_auth_delay` - Guest authentication delay (1 - 900  sec, default = 30).
* `auth_fail_vlan` - Enable to allow limited access to clients that cannot authenticate.
* `auth_fail_vlanid` - VLAN ID on which authentication failed.
* `auth_fail_vlan_id` - VLAN ID on which authentication failed.
* `framevid_apply` - Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN.
* `radius_timeout_overwrite` - Enable to override the global RADIUS session timeout.
* `policy_type` - Policy type.

The `user_group` block supports:

* `name` - Group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerSecurityPolicy 8021X can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollersecuritypolicy_8021X.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
