---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_domaincontroller"
description: |-
  Configure domain controller entries.
---

# fortios_user_domaincontroller
Configure domain controller entries.

## Example Usage

```hcl
resource "fortios_user_ldap" "trname1" {
  account_key_filter      = "(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"
  account_key_processing  = "same"
  cnid                    = "cn"
  dn                      = "EIWNCIEW"
  group_member_check      = "user-attr"
  group_object_filter     = "(&(objectcategory=group)(member=*))"
  member_attr             = "memberOf"
  name                    = "s2"
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

resource "fortios_user_domaincontroller" "trname" {
  domain_name = "s.com"
  ip_address  = "1.1.1.1"
  ldap_server = fortios_user_ldap.trname1.name
  name        = "1"
  port        = 445
}
```

## Argument Reference

The following arguments are supported:

* `name` - Domain controller entry name.
* `ip_address` - (Required) Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `extra_server` - extra servers. The structure of `extra_server` block is documented below.
* `domain_name` - Domain DNS name.
* `ldap_server` - (Required) LDAP server name.

The `extra_server` block supports:

* `id` - Server ID.
* `ip_address` - Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User DomainController can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_domaincontroller.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
