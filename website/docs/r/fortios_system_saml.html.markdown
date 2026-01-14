---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_saml"
description: |-
  Global settings for SAML authentication.
---

# fortios_system_saml
Global settings for SAML authentication. Applies to FortiOS Version `>= 6.2.4`.

## Example Usage

```hcl
resource "fortios_system_saml" "trname" {
  default_login_page = "normal"
  default_profile    = "admin_no_access"
  life               = 30
  role               = "service-provider"
  status             = "disable"
  tolerance          = 5
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable SAML authentication (default = disable). Valid values: `enable`, `disable`.
* `role` - SAML role. Valid values: `identity-provider`, `service-provider`.
* `default_login_page` - Choose default login page. Valid values: `normal`, `sso`.
* `default_profile` - Default profile for new SSO admin.
* `cert` - Certificate to sign SAML messages.
* `binding_protocol` - IdP Binding protocol. Valid values: `post`, `redirect`.
* `portal_url` - SP portal URL.
* `entity_id` - SP entity ID.
* `single_sign_on_url` - SP single sign-on URL.
* `single_logout_url` - SP single logout URL.
* `idp_entity_id` - IDP entity ID.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `idp_single_logout_url` - IDP single logout URL.
* `idp_cert` - IDP certificate name.
* `server_address` - Server address.
* `require_signed_resp_and_asrt` - Require both response and assertion from IDP to be signed when FGT acts as SP (default = disable). Valid values: `enable`, `disable`.
* `tolerance` - Tolerance to the range of time when the assertion is valid (in minutes).
* `life` - Length of the range of time when the assertion is valid (in minutes).
* `service_providers` - Authorized service providers. The structure of `service_providers` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `service_providers` block supports:

* `name` - Name.
* `prefix` - Prefix.
* `sp_binding_protocol` - SP binding protocol. Valid values: `post`, `redirect`.
* `sp_cert` - SP certificate name.
* `sp_entity_id` - SP entity ID.
* `sp_single_sign_on_url` - SP single sign-on URL.
* `sp_single_logout_url` - SP single logout URL.
* `sp_portal_url` - SP portal URL.
* `idp_entity_id` - IDP entity ID.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `idp_single_logout_url` - IDP single logout URL.
* `assertion_attributes` - Customized SAML attributes to send along with assertion. The structure of `assertion_attributes` block is documented below.

The `assertion_attributes` block supports:

* `name` - Name.
* `type` - Type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Saml can be imported using any of these accepted formats:
```
$ terraform import fortios_system_saml.labelname SystemSaml

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_saml.labelname SystemSaml
$ unset "FORTIOS_IMPORT_TABLE"
```
