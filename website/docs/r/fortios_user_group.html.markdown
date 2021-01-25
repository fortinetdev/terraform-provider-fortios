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
* `group_type` - Set the group to be for firewall authentication, FSSO, RSSO, or guest users.
* `authtimeout` - Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.
* `auth_concurrent_override` - Enable/disable overriding the global number of concurrent authentication sessions for this user group.
* `auth_concurrent_value` - Maximum number of concurrent authenticated connections per user (0 - 100).
* `http_digest_realm` - Realm attribute for MD5-digest authentication.
* `sso_attribute_value` - Name of the RADIUS user group that this local user group represents.
* `member` - Names of users, peers, LDAP severs, or RADIUS servers to add to the user group. The structure of `member` block is documented below.
* `match` - Group matches. The structure of `match` block is documented below.
* `user_id` - Guest user ID type.
* `password` - Guest user password type.
* `user_name` - Enable/disable the guest user name entry.
* `sponsor` - Set the action for the sponsor guest user field.
* `company` - Set the action for the company guest user field.
* `email` - Enable/disable the guest user email address field.
* `mobile_phone` - Enable/disable the guest user mobile phone number field.
* `sms_server` - Send SMS through FortiGuard or other external server.
* `sms_custom_server` - SMS server.
* `expire_type` - Determine when the expiration countdown begins.
* `expire` - Time in seconds before guest user accounts expire. (1 - 31536000 sec)
* `max_accounts` - Maximum number of guest accounts that can be created for this group (0 means unlimited).
* `multiple_guest_add` - Enable/disable addition of multiple guests.
* `guest` - Guest User. The structure of `guest` block is documented below.

The `member` block supports:

* `name` - Group member name.

The `match` block supports:

* `id` - ID.
* `server_name` - Name of remote auth server.
* `group_name` - Name of matching group on remote auththentication server.

The `guest` block supports:

* `user_id` - Guest ID.
* `name` - Guest name.
* `password` - Guest password.
* `mobile_phone` - Mobile phone.
* `sponsor` - Set the action for the sponsor guest user field.
* `company` - Set the action for the company guest user field.
* `email` - Email.
* `expiration` - Expire time.
* `comment` - Comment.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

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
