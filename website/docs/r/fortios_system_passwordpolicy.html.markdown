---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_passwordpolicy"
description: |-
  Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
---

# fortios_system_passwordpolicy
Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.

## Example Usage

```hcl
resource "fortios_system_passwordpolicy" "trname" {
  apply_to              = "admin-password"
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
* `apply_to` - Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
* `minimum_length` - Minimum password length (8 - 128, default = 8).
* `min_lower_case_letter` - Minimum number of lowercase characters in password (0 - 128, default = 0).
* `min_upper_case_letter` - Minimum number of uppercase characters in password (0 - 128, default = 0).
* `min_non_alphanumeric` - Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
* `min_number` - Minimum number of numeric characters in password (0 - 128, default = 0).
* `min_change_characters` - Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
* `change_4_characters` - Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
* `expire_status` - Enable/disable password expiration. Valid values: `enable`, `disable`.
* `expire_day` - Number of days after which passwords expire (1 - 999 days, default = 90).
* `reuse_password` - Enable/disable reuse of password. On FortiOS versions 6.2.0-7.0.0: If both reuse-password and change-4-characters are enabled, change-4-characters overrides.. On FortiOS versions >= 7.0.1: If both reuse-password and min-change-characters are enabled, min-change-characters overrides.. Valid values: `enable`, `disable`.
* `reuse_password_limit` - Number of times passwords can be reused (0 - 20, default = 0. If set to 0, can reuse password an unlimited number of times.).
* `login_lockout_upon_downgrade` - Enable/disable administrative user login lockout upon downgrade (defaut = disable). If enabled, downgrading the FortiOS firmware to a lower version where safer passwords are unsupported will lock out administrative users. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System PasswordPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_system_passwordpolicy.labelname SystemPasswordPolicy

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_passwordpolicy.labelname SystemPasswordPolicy
$ unset "FORTIOS_IMPORT_TABLE"
```
