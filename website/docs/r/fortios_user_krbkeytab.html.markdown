---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_krbkeytab"
description: |-
  Configure Kerberos keytab entries.
---

# fortios_user_krbkeytab
Configure Kerberos keytab entries.

## Example Usage

```hcl
resource "fortios_user_ldap" "trname2" {
  account_key_filter      = "(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"
  account_key_processing  = "same"
  cnid                    = "cn"
  dn                      = "EIWNCIEW"
  group_member_check      = "user-attr"
  group_object_filter     = "(&(objectcategory=group)(member=*))"
  member_attr             = "memberOf"
  name                    = "s1"
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

resource "fortios_user_krbkeytab" "trname" {
  keytab      = "ZXdlY2VxcmVxd3Jld3E="
  ldap_server = fortios_user_ldap.trname2.name
  name        = "1"
  principal   = "testprin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Kerberos keytab entry name.
* `pac_data` - Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
* `principal` - (Required) Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
* `ldap_server` - (Required) LDAP server name.
* `keytab` - (Required) base64 coded keytab file containing a pre-shared key.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User KrbKeytab can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_krbkeytab.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
