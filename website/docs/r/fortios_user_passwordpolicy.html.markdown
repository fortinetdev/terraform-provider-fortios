---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_passwordpolicy"
description: |-
  Configure user password policy.
---

# fortios_user_passwordpolicy
Configure user password policy.

## Example Usage

```hcl
resource "fortios_user_passwordpolicy" "trname" {
  expire_days = 22
  name        = "1"
  warn_days   = 13
}
```

## Argument Reference

The following arguments are supported:

* `name` - Password policy name.
* `expire_days` - Time in days before the user's password expires.
* `warn_days` - Time in days before a password expiration warning message is displayed to the user upon login.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User PasswordPolicy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_passwordpolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
