---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_admin"
description: |-
  Configure admin users.
---

# fortios_system_admin
Configure admin users.

## Example Usage

```hcl
resource "fortios_system_admin" "trname" {
  accprofile                 = "super_admin"
  accprofile_override        = "disable"
  allow_remove_admin_session = "enable"
  force_password_change      = "disable"
  guest_auth                 = "disable"
  hidden                     = 0
  name                       = "ss"
  password                   = "fdafdace"
  password_expire            = "0000-00-00 00:00:00"
  peer_auth                  = "disable"
  radius_vdom_override       = "disable"
  remote_auth                = "disable"
  two_factor                 = "disable"
  wildcard                   = "disable"
  vdom {
    name = "root"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - User name.
* `wildcard` - Enable/disable wildcard RADIUS authentication. Valid values: `enable`, `disable`.
* `remote_auth` - Enable/disable authentication using a remote RADIUS, LDAP, or TACACS+ server. Valid values: `enable`, `disable`.
* `remote_group` - User group name used for remote auth.
* `password` - Admin user password.
* `peer_auth` - Set to enable peer certificate authentication (for HTTPS admin access). Valid values: `enable`, `disable`.
* `peer_group` - Name of peer group defined under config user group which has PKI members. Used for peer certificate authentication (for HTTPS admin access).
* `trusthost1` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost2` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost3` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost4` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost5` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost6` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost7` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost8` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost9` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost10` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `ip6_trusthost1` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost2` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost3` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost4` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost5` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost6` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost7` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost8` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost9` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost10` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `accprofile` - Access profile for this administrator. Access profiles control administrator access to FortiGate features.
* `allow_remove_admin_session` - Enable/disable allow admin session to be removed by privileged admin users. Valid values: `enable`, `disable`.
* `comments` - Comment.
* `hidden` - Admin user hidden attribute.
* `vdom` - Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
* `ssh_public_key1` - Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
* `ssh_public_key2` - Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
* `ssh_public_key3` - Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
* `ssh_certificate` - Select the certificate to be used by the FortiGate for authentication with an SSH client.
* `schedule` - Firewall schedule used to restrict when the administrator can log in. No schedule means no restrictions.
* `accprofile_override` - Enable to use the name of an access profile provided by the remote authentication server to control the FortiGate features that this administrator can access. Valid values: `enable`, `disable`.
* `radius_vdom_override` - Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access. Valid values: `enable`, `disable`.
* `password_expire` - Password expire time.
* `force_password_change` - Enable/disable force password change on next login. Valid values: `enable`, `disable`.
* `gui_dashboard` - GUI dashboards. The structure of `gui_dashboard` block is documented below.
* `two_factor` - Enable/disable two-factor authentication.
* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
* `fortitoken` - This administrator's FortiToken serial number.
* `email_to` - This administrator's email address.
* `sms_server` - Send SMS messages using the FortiGuard SMS server or a custom server. Valid values: `fortiguard`, `custom`.
* `sms_custom_server` - Custom SMS server to send SMS messages to.
* `sms_phone` - Phone number on which the administrator receives SMS messages.
* `guest_auth` - Enable/disable guest authentication. Valid values: `disable`, `enable`.
* `guest_usergroups` - Select guest user groups. The structure of `guest_usergroups` block is documented below.
* `guest_lang` - Guest management portal language.
* `history0` - history0
* `history1` - history1
* `login_time` - Record user login time. The structure of `login_time` block is documented below.
* `gui_global_menu_favorites` - Favorite GUI menu IDs for the global VDOM. The structure of `gui_global_menu_favorites` block is documented below.
* `gui_vdom_menu_favorites` - Favorite GUI menu IDs for VDOMs. The structure of `gui_vdom_menu_favorites` block is documented below.
* `gui_new_feature_acknowledge` - Acknowledgement of new features. The structure of `gui_new_feature_acknowledge` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `vdom` block supports:

* `name` - Virtual domain name.

The `gui_dashboard` block supports:

* `id` - Dashboard ID.
* `name` - Dashboard name.
* `scope` - Dashboard scope. Valid values: `global`, `vdom`.
* `layout_type` - Layout type. Valid values: `responsive`, `fixed`.
* `columns` - Number of columns.
* `widget` - Dashboard widgets. The structure of `widget` block is documented below.

The `widget` block supports:

* `id` - Widget ID.
* `type` - Widget type. Valid values: `sysinfo`, `licinfo`, `vminfo`, `forticloud`, `cpu-usage`, `memory-usage`, `disk-usage`, `log-rate`, `sessions`, `session-rate`, `tr-history`, `analytics`, `usb-modem`, `admins`, `security-fabric`, `security-fabric-ranking`, `ha-status`, `vulnerability-summary`, `host-scan-summary`, `fortiview`, `botnet-activity`, `fortimail`.
* `x_pos` - X position.
* `y_pos` - Y position.
* `width` - Width.
* `height` - Height.
* `interface` - Interface to monitor.
* `region` - Security Audit Rating region. Valid values: `default`, `custom`.
* `industry` - Security Audit Rating industry. Valid values: `default`, `custom`.
* `fabric_device` - Fabric device to monitor.
* `title` - Widget title.
* `report_by` - Field to aggregate the data by. Valid values: `source`, `destination`, `country`, `intfpair`, `srcintf`, `dstintf`, `policy`, `wificlient`, `shaper`, `endpoint-vulnerability`, `endpoint-device`, `application`, `cloud-app`, `cloud-user`, `web-domain`, `web-category`, `web-search-phrase`, `threat`, `system`, `unauth`, `admin`, `vpn`.
* `timeframe` - Timeframe period of reported data. Valid values: `realtime`, `5min`, `hour`, `day`, `week`.
* `sort_by` - Field to sort the data by.
* `visualization` - Visualization to use. Valid values: `table`, `bubble`, `country`, `chord`.
* `filters` - FortiView filters. The structure of `filters` block is documented below.

The `filters` block supports:

* `id` - FortiView Filter ID.
* `key` - Filter key.
* `value` - Filter value.

The `guest_usergroups` block supports:

* `name` - Select guest user groups.

The `login_time` block supports:

* `usr_name` - User name.
* `last_login` - Last successful login time.
* `last_failed_login` - Last failed login time.

The `gui_global_menu_favorites` block supports:

* `id` - Select menu ID.

The `gui_vdom_menu_favorites` block supports:

* `id` - Select menu ID.

The `gui_new_feature_acknowledge` block supports:

* `id` - Select menu ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Admin can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_admin.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
