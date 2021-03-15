---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_passwordpolicyguestadmin"
description: |-
  Configure the password policy for guest administrators.
---

# fortios_system_passwordpolicyguestadmin
Configure the password policy for guest administrators.

## Example Usage

```hcl
resource "fortios_system_passwordpolicyguestadmin" "trname" {
  apply_to              = "guest-admin-password"
  change_4_characters   = "disable"
  expire_day            = 90
  expire_status         = "disable"
  min_lower_case_letter = 0
  min_non_alphanumeric  = 0
  min_number            = 0
  min_upper_case_letter = 0
  minimum_length        = 8
  reuse_password        = "enable"
  status                = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
* `apply_to` - Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
* `minimum_length` - Minimum password length (8 - 128, default = 8).
* `min_lower_case_letter` - Minimum number of lowercase characters in password (0 - 128, default = 0).
* `min_upper_case_letter` - Minimum number of uppercase characters in password (0 - 128, default = 0).
* `min_non_alphanumeric` - Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
* `min_number` - Minimum number of numeric characters in password (0 - 128, default = 0).
* `change_4_characters` - Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
* `expire_status` - Enable/disable password expiration. Valid values: `enable`, `disable`.
* `expire_day` - Number of days after which passwords expire (1 - 999 days, default = 90).
* `reuse_password` - Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System PasswordPolicyGuestAdmin can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_passwordpolicyguestadmin.labelname SystemPasswordPolicyGuestAdmin
$ unset "FORTIOS_IMPORT_TABLE"
```
