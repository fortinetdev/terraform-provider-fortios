---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_scim"
description: |-
  Configure SCIM client entries.
---

# fortios_user_scim
Configure SCIM client entries. Applies to FortiOS Version `>= 7.6.0`.

## Argument Reference

The following arguments are supported:

* `name` - SCIM client name.
* `fosid` - SCIM client ID.
* `status` - Enable/disable System for Cross-domain Identity Management (SCIM). Valid values: `enable`, `disable`.
* `base_url` - Server URL to receive SCIM create, read, update, delete (CRUD) requests.
* `auth_method` - TLS client authentication methods (default = bearer token). Valid values: `token`, `base`.
* `token_certificate` - Certificate for token verification.
* `secret` - Secret for token verification or base authentication.
* `client_authentication_method` - TLS client authentication methods (default = bearer token). Valid values: `token`, `base`.
* `client_secret_token` - Client secret token for authentication.
* `certificate` - Certificate name.
* `client_identity_check` - Enable/disable client identity check. Valid values: `enable`, `disable`.
* `cascade` - Enable/disable to follow SCIM users/groups changes in IDP. Valid values: `disable`, `enable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Scim can be imported using any of these accepted formats:
```
$ terraform import fortios_user_scim.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_scim.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
