---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ssoadmin"
description: |-
  Configure SSO admin users.
---

# fortios_system_ssoadmin
Configure SSO admin users.

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

The `vdom` block supports:

* `name` - Virtual domain name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SsoAdmin can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ssoadmin.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
