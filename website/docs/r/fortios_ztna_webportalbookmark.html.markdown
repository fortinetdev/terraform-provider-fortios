---
subcategory: "FortiGate Ztna"
layout: "fortios"
page_title: "FortiOS: fortios_ztna_webportalbookmark"
description: |-
  Configure ztna web-portal bookmark.
---

# fortios_ztna_webportalbookmark
Configure ztna web-portal bookmark. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `name` - Bookmark name.
* `users` - User name. The structure of `users` block is documented below.
* `groups` - User groups. The structure of `groups` block is documented below.
* `bookmarks` - Bookmark table. The structure of `bookmarks` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `users` block supports:

* `name` - User name.

The `groups` block supports:

* `name` - Group name.

The `bookmarks` block supports:

* `name` - Bookmark name.
* `apptype` - Application type. Valid values: `ftp`, `rdp`, `sftp`, `smb`, `ssh`, `telnet`, `vnc`, `web`.
* `url` - URL parameter.
* `host` - Host name/IP parameter.
* `folder` - Network shared file folder parameter.
* `domain` - Login domain.
* `description` - Description.
* `keyboard_layout` - Keyboard layout. Valid values: `ar-101`, `ar-102`, `ar-102-azerty`, `can-mul`, `cz`, `cz-qwerty`, `cz-pr`, `da`, `nl`, `de`, `de-ch`, `de-ibm`, `en-uk`, `en-uk-ext`, `en-us`, `en-us-dvorak`, `es`, `es-var`, `fi`, `fi-sami`, `fr`, `fr-apple`, `fr-ca`, `fr-ch`, `fr-be`, `hr`, `hu`, `hu-101`, `it`, `it-142`, `ja`, `ja-106`, `ko`, `la-am`, `lt`, `lt-ibm`, `lt-std`, `lav-std`, `lav-leg`, `mk`, `mk-std`, `no`, `no-sami`, `pol-214`, `pol-pr`, `pt`, `pt-br`, `pt-br-abnt2`, `ru`, `ru-mne`, `ru-t`, `sl`, `sv`, `sv-sami`, `tuk`, `tur-f`, `tur-q`, `zh-sym-sg-us`, `zh-sym-us`, `zh-tr-hk`, `zh-tr-mo`, `zh-tr-us`.
* `security` - Security mode for RDP connection (default = any). Valid values: `any`, `rdp`, `nla`, `tls`.
* `send_preconnection_id` - Enable/disable sending of preconnection ID. Valid values: `enable`, `disable`.
* `preconnection_id` - The numeric ID of the RDP source (0-4294967295).
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `restricted_admin` - Enable/disable restricted admin mode for RDP. Valid values: `enable`, `disable`.
* `port` - Remote port.
* `logon_user` - Logon user.
* `logon_password` - Logon password.
* `color_depth` - Color depth per pixel. Valid values: `32`, `16`, `8`.
* `sso` - Single sign-on. Valid values: `disable`, `enable`.
* `width` - Screen width (range from 0 - 65535, default = 0).
* `height` - Screen height (range from 0 - 65535, default = 0).
* `vnc_keyboard_layout` - Keyboard layout. Valid values: `default`, `da`, `nl`, `en-uk`, `en-uk-ext`, `fi`, `fr`, `fr-be`, `fr-ca-mul`, `de`, `de-ch`, `it`, `it-142`, `pt`, `pt-br-abnt2`, `no`, `gd`, `es`, `sv`, `us-intl`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna WebPortalBookmark can be imported using any of these accepted formats:
```
$ terraform import fortios_ztna_webportalbookmark.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ztna_webportalbookmark.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
