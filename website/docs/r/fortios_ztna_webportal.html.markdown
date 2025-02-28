---
subcategory: "FortiGate Ztna"
layout: "fortios"
page_title: "FortiOS: fortios_ztna_webportal"
description: |-
  Configure ztna web-portal.
---

# fortios_ztna_webportal
Configure ztna web-portal. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `name` - ZTNA proxy name.
* `vip` - Virtual IP name.
* `host` - Virtual or real host name.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `log_blocked_traffic` - Enable/disable logging of blocked traffic. Valid values: `disable`, `enable`.
* `auth_portal` - Enable/disable authentication portal. Valid values: `disable`, `enable`.
* `auth_virtual_host` - Virtual host for authentication portal.
* `vip6` - Virtual IPv6 name.
* `auth_rule` - Authentication Rule.
* `display_bookmark` - Enable to display the web portal bookmark widget. Valid values: `enable`, `disable`.
* `focus_bookmark` - Enable to prioritize the placement of the bookmark section over the quick-connection section in the ztna web-portal. Valid values: `enable`, `disable`.
* `display_status` - Enable to display the web portal status widget. Valid values: `enable`, `disable`.
* `display_history` - Enable to display the web portal user login history widget. Valid values: `enable`, `disable`.
* `policy_auth_sso` - Enable policy sso authentication. Valid values: `enable`, `disable`.
* `heading` - Web portal heading message.
* `theme` - Web portal color scheme. Valid values: `jade`, `neutrino`, `mariner`, `graphite`, `melongene`, `jet-stream`, `security-fabric`, `dark-matter`, `onyx`, `eclipse`.
* `clipboard` - Enable to support RDP/VPC clipboard functionality. Valid values: `enable`, `disable`.
* `default_window_width` - Screen width (range from 0 - 65535, default = 1024).
* `default_window_height` - Screen height (range from 0 - 65535, default = 768).
* `cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.
* `forticlient_download` - Enable/disable download option for FortiClient. Valid values: `enable`, `disable`.
* `forticlient_download_method` - FortiClient download method. Valid values: `direct`, `ssl-vpn`.
* `customize_forticlient_download_url` - Enable support of customized download URL for FortiClient. Valid values: `enable`, `disable`.
* `windows_forticlient_download_url` - Download URL for Windows FortiClient.
* `macos_forticlient_download_url` - Download URL for Mac FortiClient.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ztna WebPortal can be imported using any of these accepted formats:
```
$ terraform import fortios_ztna_webportal.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ztna_webportal.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
