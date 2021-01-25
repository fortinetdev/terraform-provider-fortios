---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_ldap"
description: |-
  Configure LDAP server entries.
---

# fortios_user_ldap
Configure LDAP server entries.

## Example Usage

```hcl
resource "fortios_user_ldap" "trname" {
  account_key_filter      = "(&(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))"
  account_key_processing  = "same"
  cnid                    = "cn"
  dn                      = "EIWNCIEW"
  group_member_check      = "user-attr"
  group_object_filter     = "(&(objectcategory=group)(member=*))"
  member_attr             = "memberOf"
  name                    = "1"
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
```

## Argument Reference


The following arguments are supported:

* `name` - LDAP server entry name.
* `server` - (Required) LDAP server CN domain name or IP.
* `secondary_server` - Secondary LDAP server CN domain name or IP.
* `tertiary_server` - Tertiary LDAP server CN domain name or IP.
* `server_identity_check` - Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate).
* `source_ip` - Source IP for communications to LDAP server.
* `cnid` - Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
* `dn` - (Required) Distinguished name used to look up entries on the LDAP server.
* `type` - Authentication type for LDAP searches.
* `username` - Username (full DN) for initial binding.
* `password` - Password for initial binding.
* `group_member_check` - Group member checking methods.
* `group_search_base` - Search base used for group searching.
* `group_object_filter` - Filter used for group searching.
* `group_filter` - Filter used for group matching.
* `secure` - Port to be used for authentication.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `ca_cert` - CA certificate name.
* `port` - Port to be used for communication with the LDAP server (default = 389).
* `password_expiry_warning` - Enable/disable password expiry warnings.
* `password_renewal` - Enable/disable online password renewal.
* `member_attr` - Name of attribute from which to get group membership.
* `account_key_processing` - Account key processing operation, either keep or strip domain string of UPN in the token.
* `account_key_filter` - Account key filter, using the UPN as the search filter.
* `search_type` - Search type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Ldap can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_ldap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
