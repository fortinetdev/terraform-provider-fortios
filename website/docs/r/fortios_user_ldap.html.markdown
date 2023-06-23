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
* `server_identity_check` - Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
* `source_ip` - Source IP for communications to LDAP server.
* `source_port` - Source port to be used for communication with the LDAP server.
* `cnid` - Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
* `dn` - (Required) Distinguished name used to look up entries on the LDAP server.
* `type` - Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.
* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
* `two_factor_filter` - Filter used to synchronize users to FortiToken Cloud.
* `username` - Username (full DN) for initial binding.
* `password` - Password for initial binding.
* `group_member_check` - Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.
* `group_search_base` - Search base used for group searching.
* `group_object_filter` - Filter used for group searching.
* `group_filter` - Filter used for group matching.
* `secure` - Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
* `ca_cert` - CA certificate name.
* `port` - Port to be used for communication with the LDAP server (default = 389).
* `password_expiry_warning` - Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
* `password_renewal` - Enable/disable online password renewal. Valid values: `enable`, `disable`.
* `member_attr` - Name of attribute from which to get group membership.
* `account_key_processing` - Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.
* `account_key_upn_san` - Define SAN in certificate for user principle name matching. Valid values: `othername`, `rfc822name`, `dnsname`.
* `account_key_filter` - Account key filter, using the UPN as the search filter.
* `search_type` - Search type. Valid values: `recursive`.
* `client_cert_auth` - Enable/disable using client certificate for TLS authentication. Valid values: `enable`, `disable`.
* `client_cert` - Client certificate name.
* `obtain_user_info` - Enable/disable obtaining of user information. Valid values: `enable`, `disable`.
* `user_info_exchange_server` - MS Exchange server from which to fetch user information.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `antiphish` - Enable/disable AntiPhishing credential backend. Valid values: `enable`, `disable`.
* `password_attr` - Name of attribute to get password hash.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Ldap can be imported using any of these accepted formats:
```
$ terraform import fortios_user_ldap.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_ldap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
