---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ssoadmin"
description: |-
  Configure SSO admin users.
---

# fortios_system_ssoadmin
Configure SSO admin users. Applies to FortiOS Version `>= 6.2.4`.

## Example Usage

```hcl
resource "fortios_system_ssoadmin" "trname" {
  accprofile = "super_admin"
  name       = "admin1"

  vdom {
    name = "root"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - SSO admin name.
* `accprofile` - (Required) SSO admin user access profile.
* `vdom` - Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `vdom` block supports:

* `name` - Virtual domain name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SsoAdmin can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ssoadmin.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ssoadmin.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
