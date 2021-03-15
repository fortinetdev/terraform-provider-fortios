---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_vlan"
description: |-
  Configure VLANs for switch controller.
---

# fortios_switchcontroller_vlan
Configure VLANs for switch controller. Applies to FortiOS Version `<= 6.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - Switch VLAN name.
* `vdom` - Virtual domain,
* `vlanid` - VLAN ID.
* `comments` - Comment.
* `color` - Color of icon on the GUI.
* `security` - Security. Valid values: `open`, `captive-portal`, `8021x`.
* `auth` - Authentication. Valid values: `radius`, `usergroup`.
* `radius_server` - Authentication radius server.
* `usergroup` - Authentication usergroup.
* `portal_message_override_group` - Specify captive portal replacement message override group.
* `portal_message_overrides` - Individual message overrides. The structure of `portal_message_overrides` block is documented below.
* `selected_usergroups` - Selected user group. The structure of `selected_usergroups` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `portal_message_overrides` block supports:

* `auth_disclaimer_page` - Override auth-disclaimer-page message with message from portal-message-overrides group.
* `auth_reject_page` - Override auth-reject-page message with message from portal-message-overrides group.
* `auth_login_page` - Override auth-login-page message with message from portal-message-overrides group.
* `auth_login_failed_page` - Override auth-login-failed-page message with message from portal-message-overrides group.

The `selected_usergroups` block supports:

* `name` - User group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController Vlan can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_vlan.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
