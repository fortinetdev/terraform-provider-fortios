---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_local"
description: |-
  Configure local users.
---

# fortios_user_local
Configure local users.

## Example Usage

```hcl
resource "fortios_user_ldap" "trname3" {
  account_key_filter      = "(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"
  account_key_processing  = "same"
  cnid                    = "cn"
  dn                      = "EIWNCIEW"
  group_member_check      = "user-attr"
  group_object_filter     = "(&(objectcategory=group)(member=*))"
  member_attr             = "memberOf"
  name                    = "s13"
  password_expiry_warning = "disable"
  password_renewal        = "disable"
  port                    = 389
  secure                  = "disable"
  server                  = "1.1.1.1"
  server_identity_check   = "disable"
  source_ip               = "0.0.0.0"
  ssl_min_proto_version   = "default"
  type                    = "simple"
}

resource "fortios_user_local" "trname" {
  auth_concurrent_override = "disable"
  auth_concurrent_value    = 0
  authtimeout              = 0
  ldap_server              = fortios_user_ldap.trname3.name
  name                     = "userlocal1"
  passwd_time              = "0000-00-00 00:00:00"
  sms_server               = "fortiguard"
  status                   = "enable"
  two_factor               = "disable"
  type                     = "ldap"
}
```

## Argument Reference

The following arguments are supported:

* `name` - User name.
* `fosid` - User ID.
* `status` - (Required) Enable/disable allowing the local user to authenticate with the FortiGate unit.
* `type` - (Required) Authentication method.
* `passwd` - User's password.
* `ldap_server` - Name of LDAP server with which the user must authenticate.
* `radius_server` - Name of RADIUS server with which the user must authenticate.
* `tacacs_server` - Name of TACACS+ server with which the user must authenticate.
* `two_factor` - Enable/disable two-factor authentication.
* `fortitoken` - Two-factor recipient's FortiToken serial number.
* `email_to` - Two-factor recipient's email address.
* `sms_server` - Send SMS through FortiGuard or other external server.
* `sms_custom_server` - Two-factor recipient's SMS server.
* `sms_phone` - Two-factor recipient's mobile phone number.
* `passwd_policy` - Password policy to apply to this user, as defined in config user password-policy.
* `passwd_time` - Time of the last password update.
* `authtimeout` - Time in minutes before the authentication timeout for a user is reached.
* `workstation` - Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
* `auth_concurrent_override` - Enable/disable overriding the policy-auth-concurrent under config system global.
* `auth_concurrent_value` - Maximum number of concurrent logins permitted from the same user.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Local can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_local.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
