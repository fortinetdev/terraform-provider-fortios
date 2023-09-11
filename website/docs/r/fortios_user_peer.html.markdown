---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_peer"
description: |-
  Configure peer users.
---

# fortios_user_peer
Configure peer users.

## Example Usage

```hcl
resource "fortios_user_peer" "trname1" {
  ca                  = "EC-ACC"
  cn_type             = "string"
  ldap_mode           = "password"
  mandatory_ca_verify = "enable"
  name                = "u1"
  two_factor          = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Peer name.
* `mandatory_ca_verify` - Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
* `ca` - Name of the CA certificate as returned by the execute vpn certificate ca list command.
* `subject` - Peer certificate name constraints.
* `cn` - Peer certificate common name.
* `cn_type` - Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
* `mfa_mode` - MFA mode for remote peer authentication/authorization. Valid values: `none`, `password`, `subject-identity`.
* `mfa_server` - Name of a remote authenticator. Performs client access right check.
* `mfa_username` - Unified username for remote authentication.
* `mfa_password` - Unified password for remote authentication. This field may be left empty when RADIUS authentication is used, in which case the FortiGate will use the RADIUS username as a password. 
* `ldap_server` - Name of an LDAP server defined under the user ldap command. Performs client access rights check.
* `ldap_username` - Username for LDAP server bind.
* `ldap_password` - Password for LDAP server bind.
* `ldap_mode` - Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
* `ocsp_override_server` - Online Certificate Status Protocol (OCSP) server for certificate retrieval.
* `two_factor` - Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
* `passwd` - Peer's password used for two-factor authentication.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Peer can be imported using any of these accepted formats:
```
$ terraform import fortios_user_peer.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_peer.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
