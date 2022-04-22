---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_saml"
description: |-
  SAML server entry configuration.
---

# fortios_user_saml
SAML server entry configuration. Applies to FortiOS Version `>= 6.2.4`.

## Example Usage

```hcl
resource "fortios_user_saml" "tr3" {
  cert                   = "Fortinet_Factory"
  entity_id              = "https://1.1.1.1"
  idp_cert               = "cer11"
  idp_entity_id          = "https://1.1.1.1/acc"
  idp_single_logout_url  = "https://1.1.1.1/lo"
  idp_single_sign_on_url = "https://1.1.1.1/sou"
  name                   = "testwebvpn"
  single_logout_url      = "https://1.1.1.1/logout"
  single_sign_on_url     = "https://1.1.1.1/sign"
  user_name              = "ad111"
}
```

## Argument Reference

The following arguments are supported:

* `name` - SAML server entry name.
* `cert` - Certificate to sign SAML messages.
* `entity_id` - (Required) SP entity ID.
* `single_sign_on_url` - (Required) SP single sign-on URL.
* `single_logout_url` - SP single logout URL.
* `idp_entity_id` - (Required) IDP entity ID.
* `idp_single_sign_on_url` - (Required) IDP single sign-on URL.
* `idp_single_logout_url` - IDP single logout url.
* `idp_cert` - (Required) IDP Certificate name.
* `user_name` - User name in assertion statement.
* `group_name` - Group name in assertion statement.
* `digest_method` - Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
* `limit_relaystate` - Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
* `clock_tolerance` - Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
* `adfs_claim` - Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
* `user_claim_type` - User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
* `group_claim_type` - Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Saml can be imported using any of these accepted formats:
```
$ terraform import fortios_user_saml.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_saml.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
