---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_group"
description: |-
  Configure user groups.
---

# fortios_user_group
Configure user groups.

## Example Usage

```hcl
resource "fortios_user_group" "trname" {
  company            = "optional"
  email              = "enable"
  expire             = 14400
  expire_type        = "immediately"
  group_type         = "firewall"
  max_accounts       = 0
  mobile_phone       = "disable"
  multiple_guest_add = "disable"
  name               = "s1"

  member {
    name = "guest"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Group name.
* `fosid` - Group ID.
* `group_type` - Set the group to be for firewall authentication, FSSO, RSSO, or guest users. Valid values: `firewall`, `fsso-service`, `rsso`, `guest`.
* `authtimeout` - Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.
* `auth_concurrent_override` - Enable/disable overriding the global number of concurrent authentication sessions for this user group. Valid values: `enable`, `disable`.
* `auth_concurrent_value` - Maximum number of concurrent authenticated connections per user (0 - 100).
* `http_digest_realm` - Realm attribute for MD5-digest authentication.
* `sso_attribute_value` - Name of the RADIUS user group that this local user group represents.
* `member` - Names of users, peers, LDAP severs, or RADIUS servers to add to the user group. The structure of `member` block is documented below.
* `match` - Group matches. The structure of `match` block is documented below.
* `user_id` - Guest user ID type. Valid values: `email`, `auto-generate`, `specify`.
* `password` - Guest user password type. Valid values: `auto-generate`, `specify`, `disable`.
* `user_name` - Enable/disable the guest user name entry. Valid values: `disable`, `enable`.
* `sponsor` - Set the action for the sponsor guest user field. Valid values: `optional`, `mandatory`, `disabled`.
* `company` - Set the action for the company guest user field. Valid values: `optional`, `mandatory`, `disabled`.
* `email` - Enable/disable the guest user email address field. Valid values: `disable`, `enable`.
* `mobile_phone` - Enable/disable the guest user mobile phone number field. Valid values: `disable`, `enable`.
* `sms_server` - Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.
* `sms_custom_server` - SMS server.
* `expire_type` - Determine when the expiration countdown begins. Valid values: `immediately`, `first-successful-login`.
* `expire` - Time in seconds before guest user accounts expire. (1 - 31536000 sec)
* `max_accounts` - Maximum number of guest accounts that can be created for this group (0 means unlimited).
* `multiple_guest_add` - Enable/disable addition of multiple guests. Valid values: `disable`, `enable`.
* `guest` - Guest User. The structure of `guest` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Group member name.

The `match` block supports:

* `id` - ID.
* `server_name` - Name of remote auth server.
* `group_name` - Name of matching group on remote auththentication server.

The `guest` block supports:

* `id` - Guest ID.
* `user_id` - Guest ID.
* `name` - Guest name.
* `password` - Guest password.
* `mobile_phone` - Mobile phone.
* `sponsor` - Set the action for the sponsor guest user field.
* `company` - Set the action for the company guest user field.
* `email` - Email.
* `expiration` - Expire time.
* `comment` - Comment.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Group can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_group.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
