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
* `hostname` - Hostname of the server to connect to.
* `username` - User name to sign in with. Must have proper permissions for service.
* `password` - Password for specified username.
* `ip_address` - (Required) Domain controller IP address.
* `ip6` - Domain controller IPv6 address.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_ip6` - FortiGate IPv6 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `extra_server` - extra servers. The structure of `extra_server` block is documented below.
* `domain_name` - Domain DNS name.
* `replication_port` - Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.
* `ldap_server` - (Required) LDAP server name.
* `change_detection` - Enable/disable detection of a configuration change in the Active Directory server. Valid values: `enable`, `disable`.
* `change_detection_period` - Minutes to detect a configuration change in the Active Directory server (5 - 10080 minutes (7 days), default = 60).
* `dns_srv_lookup` - Enable/disable DNS service lookup. Valid values: `enable`, `disable`.
* `ad_mode` - Set Active Directory mode. Valid values: `none`, `ds`, `lds`.
* `adlds_dn` - AD LDS distinguished name.
* `adlds_ip_address` - AD LDS IPv4 address.
* `adlds_ip6` - AD LDS IPv6 address.
* `adlds_port` - Port number of AD LDS service (default = 389).
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `extra_server` block supports:

* `id` - Server ID.
* `ip_address` - Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `source_ip_address` - FortiGate IPv4 address to be used for communication with the domain controller.
* `source_port` - Source port to be used for communication with the domain controller.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User DomainController can be imported using any of these accepted formats:
```
$ terraform import fortios_user_domaincontroller.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_domaincontroller.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
