---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_saml"
description: |-
  SAML server entry configuration.
---

# fortios_user_saml
SAML server entry configuration.

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
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Saml can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_saml.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
