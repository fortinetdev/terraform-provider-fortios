---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_scheme"
description: |-
  Configure Authentication Schemes.
---

# fortios_authentication_scheme
Configure Authentication Schemes.

## Example Usage

```hcl
resource "fortios_user_fsso" "trname3" {
  name       = "2"
  port       = 8000
  port2      = 8000
  port3      = 8000
  port4      = 8000
  port5      = 8000
  server     = "1.1.1.1"
  source_ip  = "0.0.0.0"
  source_ip6 = "::"
}

resource "fortios_authentication_scheme" "trname" {
  fsso_agent_for_ntlm = fortios_user_fsso.trname3.name
  fsso_guest          = "disable"
  method              = "ntlm"
  name                = "1"
  negotiate_ntlm      = "enable"
  require_tfa         = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Authentication scheme name.
* `method` - (Required) Authentication methods (default = basic).
* `negotiate_ntlm` - Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable`, `disable`.
* `kerberos_keytab` - Kerberos keytab setting.
* `domain_controller` - Domain controller setting.
* `saml_server` - SAML configuration.
* `saml_timeout` - SAML authentication timeout in seconds.
* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `require_tfa` - Enable/disable two-factor authentication (default = disable). Valid values: `enable`, `disable`.
* `fsso_guest` - Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable`, `disable`.
* `user_cert` - Enable/disable authentication with user certificate (default = disable). Valid values: `enable`, `disable`.
* `user_database` - Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.
* `ssh_ca` - SSH CA name.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `user_database` block supports:

* `name` - Authentication server name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Authentication Scheme can be imported using any of these accepted formats:
```
$ terraform import fortios_authentication_scheme.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_authentication_scheme.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
