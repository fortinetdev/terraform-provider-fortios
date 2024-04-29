---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_usergroupbookmark"
description: |-
  Configure SSL VPN user group bookmark.
---

# fortios_vpnsslweb_usergroupbookmark
Configure SSL VPN user group bookmark.

## Example Usage

```hcl
resource "fortios_vpnsslweb_usergroupbookmark" "trname" {
  name = "Guest-group"

  bookmarks {
    apptype                  = "citrix"
    listening_port           = 0
    name                     = "b1"
    port                     = 0
    remote_port              = 0
    security                 = "rdp"
    server_layout            = "en-us-qwerty"
    show_status_window       = "disable"
    sso                      = "disable"
    sso_credential           = "sslvpn-login"
    sso_credential_sent_once = "disable"
    url                      = "www.aaa.com"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Group name.
* `bookmarks` - Bookmark table. The structure of `bookmarks` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `bookmarks` block supports:

* `name` - Bookmark name.
* `apptype` - Application type.
* `url` - URL parameter.
* `host` - Host name/IP parameter.
* `folder` - Network shared file folder parameter.
* `domain` - Login domain.
* `additional_params` - Additional parameters.
* `listening_port` - Listening port (0 - 65535).
* `remote_port` - Remote port (0 - 65535).
* `show_status_window` - Enable/disable showing of status window. Valid values: `enable`, `disable`.
* `description` - Description.
* `keyboard_layout` - Keyboard layout.
* `server_layout` - Server side keyboard layout.
* `security` - Security mode for RDP connection (default = any). Valid values: `rdp`, `nla`, `tls`, `any`.
* `send_preconnection_id` - Enable/disable sending of preconnection ID. Valid values: `enable`, `disable`.
* `preconnection_id` - The numeric ID of the RDP source. On FortiOS versions 6.2.0-6.4.2, 7.0.0: 0-2147483648. On FortiOS versions 6.4.10-6.4.15, >= 7.0.1: 0-4294967295.
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `restricted_admin` - Enable/disable restricted admin mode for RDP. Valid values: `enable`, `disable`.
* `port` - Remote port.
* `logon_user` - Logon user.
* `logon_password` - Logon password.
* `color_depth` - Color depth per pixel. Valid values: `32`, `16`, `8`.
* `sso` - Single Sign-On. Valid values: `disable`, `static`, `auto`.
* `form_data` - Form data. The structure of `form_data` block is documented below.
* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.
* `sso_username` - SSO user name.
* `sso_password` - SSO password.
* `sso_credential_sent_once` - Single sign-on credentials are only sent once to remote server. Valid values: `enable`, `disable`.
* `width` - Screen width. On FortiOS versions 7.0.4-7.0.5: range from 640 - 65535, default = 1024. On FortiOS versions >= 7.0.6: range from 0 - 65535, default = 0.
* `height` - Screen height. On FortiOS versions 7.0.4-7.0.5: range from 480 - 65535, default = 768. On FortiOS versions >= 7.0.6: range from 0 - 65535, default = 0.
* `vnc_keyboard_layout` - Keyboard layout. Valid values: `default`, `da`, `nl`, `en-uk`, `en-uk-ext`, `fi`, `fr`, `fr-be`, `fr-ca-mul`, `de`, `de-ch`, `it`, `it-142`, `pt`, `pt-br-abnt2`, `no`, `gd`, `es`, `sv`, `us-intl`.

The `form_data` block supports:

* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnSslWeb UserGroupBookmark can be imported using any of these accepted formats:
```
$ terraform import fortios_vpnsslweb_usergroupbookmark.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnsslweb_usergroupbookmark.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
