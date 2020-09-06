---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_vlan"
description: |-
  Configure VLANs for switch controller.
---

# fortios_switchcontroller_vlan
Configure VLANs for switch controller.

## Argument Reference

The following arguments are supported:

* `name` - Switch VLAN name.
* `vdom` - Virtual domain,
* `vlanid` - VLAN ID.
* `comments` - Comment.
* `color` - Color of icon on the GUI.
* `security` - Security.
* `auth` - Authentication.
* `radius_server` - (Required) Authentication radius server.
* `usergroup` - (Required) Authentication usergroup.
* `portal_message_override_group` - Specify captive portal replacement message override group.
* `portal_message_overrides` - Individual message overrides.
* `selected_usergroups` - (Required) Selected user group.

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
