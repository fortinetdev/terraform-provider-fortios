---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_saml"
description: |-
  Get information on an fortios user saml.
---

# Data Source: fortios_user_saml
Use this data source to get information on an fortios user saml

## Argument Reference

* `name` - (Required) Specify the name of the desired user saml.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - SAML server entry name.
* `cert` - Certificate to sign SAML messages.
* `entity_id` - SP entity ID.
* `single_sign_on_url` - SP single sign-on URL.
* `single_logout_url` - SP single logout URL.
* `idp_entity_id` - IDP entity ID.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `idp_single_logout_url` - IDP single logout url.
* `idp_cert` - IDP Certificate name.
* `user_name` - User name in assertion statement.
* `group_name` - Group name in assertion statement.

