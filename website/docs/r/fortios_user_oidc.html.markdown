---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_oidc"
description: |-
  OpenID Connect server entry configuration.
---

# fortios_user_oidc
OpenID Connect server entry configuration. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - OpenID Connect server entry name.
* `type` - Type of OpenID Connect config. Valid values: `discovery`, `manual`.
* `client_id` - OpenID Connect server client ID.
* `client_secret` - OpenID Connect server client secret.
* `discovery_url` - OpenID Connect server discovery URL.
* `authorization_url` - OpenID Connect server authorization URL.
* `token_url` - OpenID Connect server token URL.
* `jwks_uri` - URL of the OP's JWK Set document.
* `domain_hint` - Domain Hint.
* `issuer` - OpenID Connect server issuer.
* `verify_issuer` - Verify issuer in ID token (default = enable). Valid values: `enable`, `disable`.
* `user_attr_name` - Which field in ID token is username. Valid values: `email`, `sub`, `preferred_username`.
* `user_regex` - username must match this regex (case insensitive).
* `ldap_server` - LDAP server name(s). The structure of `ldap_server` block is documented below.
* `clock_tolerance` - Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `ldap_server` block supports:

* `name` - LDAP server name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Oidc can be imported using any of these accepted formats:
```
$ terraform import fortios_user_oidc.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_oidc.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
